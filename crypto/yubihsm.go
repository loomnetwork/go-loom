// +build evm

package crypto

import (
	"encoding/asn1"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/big"
	"math/rand"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"

	"github.com/loomnetwork/yubihsm-go"
	"github.com/loomnetwork/yubihsm-go/commands"
	"github.com/loomnetwork/yubihsm-go/connector"
)

const (
	YubiDefConnURL   = "127.0.0.1:12345"
	YubiDefAuthKeyID = 1
	YubiDefPassword  = "password"

	YubiSecp256k1PubKeySize  = 33
	YubiSecp256k1SignDataLen = 64
)

type YubiHsmParams struct {
	HsmConnURL string `json:"YubiHsmConnURL"`
	AuthKeyID  uint16 `json:"AuthKeyID"`
	AuthPasswd string `json:"Password"`
	PrivKeyID  uint16 `json:"PrivKeyID"`
}

type YubiHsmPrivateKey struct {
	yubiHsmParams *YubiHsmParams
	sessionMgr    *yubihsm.SessionManager
	privKeyID     uint16
	pubKeyBytes   [YubiSecp256k1PubKeySize]byte
}

func (privKey *YubiHsmPrivateKey) initYubiHsmSession(filePath string) error {
	yubiHsmParams := &YubiHsmParams{
		HsmConnURL: YubiDefConnURL,
		AuthKeyID:  YubiDefAuthKeyID,
		AuthPasswd: YubiDefPassword,
	}

	jsonParams, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonParams, yubiHsmParams)
	if err != nil {
		return err
	}

	httpConnector := connector.NewHTTPConnector(yubiHsmParams.HsmConnURL)
	sessionMgr, err := yubihsm.NewSessionManager(httpConnector, yubiHsmParams.AuthKeyID, yubiHsmParams.AuthPasswd)
	if err != nil {
		return err
	}

	privKey.yubiHsmParams = yubiHsmParams
	privKey.sessionMgr = sessionMgr
	privKey.privKeyID = yubiHsmParams.PrivKeyID

	return nil
}

func GenYubiHsmPrivKey(filePath string) (*YubiHsmPrivateKey, error) {
	var err error

	yubiHsmPrivKey := &YubiHsmPrivateKey{}

	// init YubiHSM session
	err = yubiHsmPrivKey.initYubiHsmSession(filePath)
	if err != nil {
		return nil, err
	}

	// generate private key
	err = yubiHsmPrivKey.genPrivKey()
	if err != nil {
		yubiHsmPrivKey.UnloadYubiHsmPrivKey()
		return nil, err
	}

	// export pubkey
	err = yubiHsmPrivKey.exportPubKey()
	if err != nil {
		yubiHsmPrivKey.deletePrivKey()
		yubiHsmPrivKey.UnloadYubiHsmPrivKey()
		return nil, err
	}

	return yubiHsmPrivKey, nil
}

func (privKey *YubiHsmPrivateKey) genPrivKey() error {
	// generate keyID
	rand.Seed(time.Now().UnixNano())
	keyID := uint16(rand.Intn(0xFFFF))

	// create command to generate secp256k1 key
	command, err := commands.CreateGenerateAsymmetricKeyCommand(keyID, nil, commands.Domain1,
		commands.CapabilityAsymmetricSignEcdsa, commands.AlgorithmSecp256k1)
	if err != nil {
		return err
	}

	// send command to YubiHSM
	_, err = privKey.sessionMgr.SendEncryptedCommand(command)
	if err != nil {
		return err
	}
	privKey.privKeyID = keyID

	return nil
}

func (privKey *YubiHsmPrivateKey) deletePrivKey() {
	// create command to delete secp256k1 key
	command, err := commands.CreateDeleteObjectCommand(privKey.privKeyID, commands.ObjectTypeAsymmetricKey)
	if err != nil {
		return
	}

	// send command
	privKey.sessionMgr.SendEncryptedCommand(command)

	privKey.privKeyID = 0
}

func LoadYubiHsmPrivKey(filePath string) (*YubiHsmPrivateKey, error) {
	var err error

	yubiHsmPrivKey := &YubiHsmPrivateKey{}

	// init YubiHSM session
	err = yubiHsmPrivKey.initYubiHsmSession(filePath)
	if err != nil {
		return nil, err
	}
	if yubiHsmPrivKey.privKeyID == 0 {
		yubiHsmPrivKey.UnloadYubiHsmPrivKey()
		return nil, errors.New("Missing private key ID")
	}

	// try to export secp256k1 public key
	err = yubiHsmPrivKey.exportPubKey()
	if err != nil {
		yubiHsmPrivKey.deletePrivKey()
		yubiHsmPrivKey.UnloadYubiHsmPrivKey()
		return nil, err
	}

	return yubiHsmPrivKey, nil
}

func (privKey *YubiHsmPrivateKey) SaveYubiHsmPrivKey(filePath string) {
	privKey.yubiHsmParams.PrivKeyID = privKey.privKeyID

	// convert to json
	jsonBytes, err := json.Marshal(privKey.yubiHsmParams)
	if err != nil {
		panic(err)
	}

	// create file
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(jsonBytes)
	if err != nil {
		panic(err)
	}
}

// unload YubiHsm private key
func (privKey *YubiHsmPrivateKey) UnloadYubiHsmPrivKey() {
	if privKey.sessionMgr == nil {
		return
	}
	privKey.sessionMgr.Destroy()
}

// export YubiHsm public key by private key ID
func (privKey *YubiHsmPrivateKey) exportPubKey() error {
	// send getpubkey command
	cmd, err := commands.CreateGetPubKeyCommand(privKey.privKeyID)
	if err != nil {
		return err
	}

	resp, err := privKey.sessionMgr.SendEncryptedCommand(cmd)
	if err != nil {
		return nil
	}

	// parse public key from response
	parsedResp, matched := resp.(*commands.GetPubKeyResponse)
	if !matched {
		return errors.New("Invalid response for exporting public key")
	}

	if parsedResp.Algorithm != commands.AlgorithmSecp256k1 {
		return errors.New("Invalid Secp256k1 key type")
	}
	if len(parsedResp.KeyData) != 64 {
		return errors.New("Invalid Secp256k1 public key data size")
	}

	x := new(big.Int)
	y := new(big.Int)

	x.SetBytes(parsedResp.KeyData[0:32])
	y.SetBytes(parsedResp.KeyData[31:])

	copy(privKey.pubKeyBytes[:], secp256k1.CompressPubkey(x, y))

	if len(privKey.pubKeyBytes) != YubiSecp256k1PubKeySize {
		return errors.New("Invalid Secp256k1 public key size")
	}

	return nil
}

// YubiHsmSign signs using private key in YubiHSM token
func YubiHsmSign(hash []byte, privKey *YubiHsmPrivateKey) (sig []byte, err error) {
	var ecdsaSig struct {
		R, S *big.Int
	}

	// send command to sign data
	command, err := commands.CreateSignDataEcdsaCommand(privKey.privKeyID, hash)
	if err != nil {
		return nil, err
	}
	resp, err := privKey.sessionMgr.SendEncryptedCommand(command)
	if err != nil {
		return nil, err
	}

	// parse response
	parsedResp, matched := resp.(*commands.SignDataEcdsaResponse)
	if !matched {
		return nil, errors.New("Invalid response type for sign command")
	}

	_, err = asn1.Unmarshal(parsedResp.Signature, &ecdsaSig)
	if err != nil {
		return nil, err
	}

	sig = ecdsaSig.R.Bytes()
	sig = append(sig, ecdsaSig.S.Bytes()...)

	if len(sig) != YubiSecp256k1SignDataLen {
		return nil, errors.New("Invalid signature length")
	}

	return sig, nil
}

// get pubkey bytes
func (privKey *YubiHsmPrivateKey) GetPubKeyBytes() []byte {
	return privKey.pubKeyBytes[:]
}
