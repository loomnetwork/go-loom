package crypto

import (
	"bytes"
	"encoding/asn1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"time"

	yubihsm "github.com/certusone/yubihsm-go"
	"github.com/certusone/yubihsm-go/commands"
	"github.com/certusone/yubihsm-go/connector"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	loom "github.com/loomnetwork/go-loom"
)

const (
	YubiDefConnURL       = "127.0.0.1:12345"
	YubiDefAuthKeyID     = 1
	YubiDefPassword      = "password"
	YubiDefPrivKeyDomain = 1
	YubiDefPrivKeyType   = PrivateKeyTypeEd25519

	YubiSecp256k1PubKeySize  = 33
	YubiSecp256k1SignDataLen = 65

	YubiEd25519PubKeySize  = 32
	YubiEd25519SignDataLen = 64
)

type YubiHsmConfig struct {
	HsmConnURL    string `json:"YubiHsmConnURL"`
	AuthKeyID     uint16 `json:"AuthKeyID"`
	AuthPasswd    string `json:"Password"`
	PrivKeyID     uint16 `json:"PrivKeyID"`
	PrivKeyDomain uint16 `json:"PrivKeyDomain"`
	PrivKeyType   string `json:"PrivKeyType"`
}

type YubiHsmPrivateKey struct {
	sessionMgr         *yubihsm.SessionManager
	privKeyType        string
	privKeyID          uint16
	pubKeyBytes        []byte
	pubKeyUncompressed []byte // this field is available only for secp256k1
}

func InitYubiHsmPrivKey(hsmConfig *YubiHsmConfig) (*YubiHsmPrivateKey, error) {
	privKey := &YubiHsmPrivateKey{}

	if hsmConfig.PrivKeyType != PrivateKeyTypeEd25519 && hsmConfig.PrivKeyType != PrivateKeyTypeSecp256k1 {
		return nil, fmt.Errorf("Invalid YubiHSM private key type '%s'", hsmConfig.PrivKeyType)
	}

	httpConnector := connector.NewHTTPConnector(hsmConfig.HsmConnURL)
	sessionMgr, err := yubihsm.NewSessionManager(httpConnector, hsmConfig.AuthKeyID, hsmConfig.AuthPasswd)
	if err != nil {
		return nil, err
	}

	privKey.sessionMgr = sessionMgr
	privKey.privKeyID = hsmConfig.PrivKeyID
	privKey.privKeyType = hsmConfig.PrivKeyType

	return privKey, nil
}

func loadYubiHsmPrivKey(filePath string) (*YubiHsmPrivateKey, error) {
	yubiHsmConfig := &YubiHsmConfig{
		HsmConnURL:    YubiDefConnURL,
		AuthKeyID:     YubiDefAuthKeyID,
		AuthPasswd:    YubiDefPassword,
		PrivKeyID:     0,
		PrivKeyDomain: YubiDefPrivKeyDomain,
		PrivKeyType:   YubiDefPrivKeyType,
	}

	jsonParams, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonParams, yubiHsmConfig)
	if err != nil {
		return nil, err
	}

	return InitYubiHsmPrivKey(yubiHsmConfig)
}

func GenYubiHsmPrivKey(filePath string) (*YubiHsmPrivateKey, error) {
	// init YubiHSM session
	yubiHsmPrivKey, err := loadYubiHsmPrivKey(filePath)
	if err != nil {
		return nil, err
	}

	// generate private key
	err = yubiHsmPrivKey.GenPrivKey()
	if err != nil {
		yubiHsmPrivKey.UnloadYubiHsmPrivKey()
		return nil, err
	}

	// export pubkey
	err = yubiHsmPrivKey.ExportPubKey()
	if err != nil {
		yubiHsmPrivKey.deletePrivKey()
		yubiHsmPrivKey.UnloadYubiHsmPrivKey()
		return nil, err
	}

	return yubiHsmPrivKey, nil
}

func (privKey *YubiHsmPrivateKey) GenPrivKey() error {
	var cap uint64
	var algo commands.Algorithm

	// generate keyID
	rand.Seed(time.Now().UnixNano())
	keyID := uint16(rand.Intn(0xFFFF))

	if privKey.privKeyType == PrivateKeyTypeEd25519 {
		cap = commands.CapabilityAsymmetricSignEddsa
		algo = commands.AlgorighmED25519
	} else {
		cap = commands.CapabilityAsymmetricSignEcdsa
		algo = commands.AlgorithmSecp256k1
	}

	// create command to generate secp256k1 key
	command, err := commands.CreateGenerateAsymmetricKeyCommand(keyID, nil, commands.Domain1, cap, algo)
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

func (privKey *YubiHsmPrivateKey) GetPrivKeyID() uint16 {
	return privKey.privKeyID
}

func (privKey *YubiHsmPrivateKey) deletePrivKey() {
	// create command to delete secp256k1 key
	command, err := commands.CreateDeleteObjectCommand(privKey.privKeyID, commands.ObjectTypeAsymmetricKey)
	if err != nil {
		return
	}

	// send command
	_, err = privKey.sessionMgr.SendEncryptedCommand(command)
	if err != nil {
		return
	}

	privKey.privKeyID = 0
}

func LoadYubiHsmPrivKey(filePath string) (*YubiHsmPrivateKey, error) {
	// init YubiHSM session
	yubiHsmPrivKey, err := loadYubiHsmPrivKey(filePath)
	if err != nil {
		return nil, err
	}
	if yubiHsmPrivKey.privKeyID == 0 {
		yubiHsmPrivKey.UnloadYubiHsmPrivKey()
		return nil, errors.New("Missing private key ID")
	}

	// try to export secp256k1 public key
	err = yubiHsmPrivKey.ExportPubKey()
	if err != nil {
		yubiHsmPrivKey.UnloadYubiHsmPrivKey()
		return nil, err
	}

	return yubiHsmPrivKey, nil
}

// unload YubiHsm private key
func (privKey *YubiHsmPrivateKey) UnloadYubiHsmPrivKey() {
	if privKey.sessionMgr == nil {
		return
	}
	privKey.sessionMgr.Destroy()
}

// export secp256k1 public key
func (privKey *YubiHsmPrivateKey) exportSecp256k1Pubkey(keyData []byte) error {
	if len(keyData) != 64 {
		return errors.New("Invalid Secp256k1 public key data size")
	}

	x := new(big.Int)
	y := new(big.Int)

	x.SetBytes(keyData[0:32])
	y.SetBytes(keyData[31:])

	privKey.pubKeyUncompressed = make([]byte, 65)
	privKey.pubKeyUncompressed[0] = 0x04
	copy(privKey.pubKeyUncompressed[1:], keyData[:])

	privKey.pubKeyBytes = make([]byte, YubiSecp256k1PubKeySize)
	copy(privKey.pubKeyBytes[:], secp256k1.CompressPubkey(x, y))
	if len(privKey.pubKeyBytes) != YubiSecp256k1PubKeySize {
		return errors.New("Invalid Secp256k1 public key size")
	}

	return nil
}

// export ed25519 public key
func (privKey *YubiHsmPrivateKey) exportEd25519Pubkey(keyData []byte) error {
	if len(keyData) != YubiEd25519PubKeySize {
		return errors.New("Invalid ed25519 public key data size")
	}

	privKey.pubKeyBytes = make([]byte, YubiEd25519PubKeySize)
	copy(privKey.pubKeyBytes[:], keyData[:])

	return nil
}

// export YubiHsm public key by private key ID
func (privKey *YubiHsmPrivateKey) ExportPubKey() error {
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

	if privKey.privKeyType == PrivateKeyTypeEd25519 {
		err = privKey.exportEd25519Pubkey(parsedResp.KeyData)
	} else {
		err = privKey.exportSecp256k1Pubkey(parsedResp.KeyData)
	}

	return err
}

func (privKey *YubiHsmPrivateKey) getSigRecID(hash []byte, sig []byte) (byte, error) {
	recIds := []byte{0x00, 0x01, 0x02, 0x03}

	for i := 0; i < len(recIds); i++ {
		tmpSig := make([]byte, 65)
		copy(tmpSig[:], sig)
		tmpSig[64] = recIds[i]
		pubKeyBytes, err := secp256k1.RecoverPubkey(hash, tmpSig[:])
		if err == nil && bytes.Equal(pubKeyBytes, privKey.pubKeyUncompressed) {
			return recIds[i], nil
		}
	}

	return 0x04, fmt.Errorf("Unable to get recovery public key from signature")
}

func (privKey *YubiHsmPrivateKey) yubiHsmSecp256k1Sign(hash []byte) ([]byte, error) {
	var sig [YubiSecp256k1SignDataLen]byte
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

	copy(sig[:], ecdsaSig.R.Bytes())
	copy(sig[32:], ecdsaSig.S.Bytes())

	normSig, err := secp256k1.NormalizeLaxDERSignature(parsedResp.Signature)
	if err != nil {
		return nil, err
	}

	recID, err := privKey.getSigRecID(hash, normSig)
	if err != nil {
		return nil, err
	}
	copy(sig[:], normSig)
	sig[64] = recID

	return sig[:], nil
}

func (privKey *YubiHsmPrivateKey) yubiHsmEd25519Sign(msg []byte) (sig []byte, err error) {
	// send command to sign data
	command, err := commands.CreateSignDataEddsaCommand(privKey.privKeyID, msg)
	if err != nil {
		return nil, err
	}
	resp, err := privKey.sessionMgr.SendEncryptedCommand(command)
	if err != nil {
		return nil, err
	}

	// parse response
	parsedResp, matched := resp.(*commands.SignDataEddsaResponse)
	if !matched {
		return nil, errors.New("Invalid response type for sign command")
	}
	if len(parsedResp.Signature) != YubiEd25519SignDataLen {
		return nil, errors.New("Invalid sign data len for ed25519")
	}

	return parsedResp.Signature, nil
}

// YubiHsmSign signs using private key in YubiHSM token
func YubiHsmSign(msg []byte, privKey *YubiHsmPrivateKey) (sig []byte, err error) {
	if privKey.privKeyType == PrivateKeyTypeEd25519 {
		sig, err = privKey.yubiHsmEd25519Sign(msg)
	} else {
		// check if msg is sha hash
		if len(msg) != 32 {
			return nil, fmt.Errorf("hash is required to be exactly 32 bytes (%d)", len(msg))
		}
		sig, err = privKey.yubiHsmSecp256k1Sign(msg)
	}

	return sig, err
}

// get pubkey bytes
func (privKey *YubiHsmPrivateKey) GetPubKeyBytes() []byte {
	return privKey.pubKeyBytes[:]
}

// get pubkey address
func (privKey *YubiHsmPrivateKey) GetPubKeyAddr() string {
	if privKey.privKeyType == PrivateKeyTypeSecp256k1 {
		ecdsaPubKey, err := crypto.UnmarshalPubkey(privKey.pubKeyUncompressed)
		if err != nil {
			privKey.deletePrivKey()
			panic(err)
		}
		pubKeyAddr := crypto.PubkeyToAddress(*ecdsaPubKey)
		return pubKeyAddr.Hex()
	}
	return loom.LocalAddressFromPublicKey(privKey.pubKeyBytes).String()
}

// get base64 encoded pubkey address
func (privKey *YubiHsmPrivateKey) GetPubKeyAddrB64Encoded() (string, error) {
	if privKey.privKeyType == PrivateKeyTypeEd25519 {
		return base64.StdEncoding.EncodeToString(loom.LocalAddressFromPublicKey(privKey.pubKeyBytes)), nil
	}

	return "", fmt.Errorf("Not suported")
}

// get key type
func (privKey *YubiHsmPrivateKey) GetKeyType() string {
	return privKey.privKeyType
}
