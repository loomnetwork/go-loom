package crypto

import (
	"crypto/sha256"
	"encoding/asn1"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"

	"github.com/loomnetwork/yubihsm-go"
	"github.com/loomnetwork/yubihsm-go/commands"
	"github.com/loomnetwork/yubihsm-go/connector"
)

const (
	YubiDefConnURL       = "127.0.0.1:12345"
	YubiDefAuthKeyID     = 1
	YubiDefPassword      = "password"
	YubiDefPrivKeyDomain = 1
	YubiDefPrivKeyType   = PrivateKeyTypeEd25519

	YubiSecp256k1PubKeySize  = 33
	YubiSecp256k1SignDataLen = 64

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
	pubKeyUncompressed []byte
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
	privKey.sessionMgr.SendEncryptedCommand(command)

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

	privKey.pubKeyUncompressed = make([]byte, 64)
	copy(privKey.pubKeyUncompressed[:], keyData[:])

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

func (privKey *YubiHsmPrivateKey) yubiHsmSecp256k1Sign(hash []byte) (sig []byte, err error) {
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
		return nil, errors.New("Invalid signature YubiSecp256k1SignDataLen length")
	}

	return sig, nil
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
		hash := sha256.Sum256(msg)
		sig, err = privKey.yubiHsmSecp256k1Sign(hash[:])
	}

	return sig, err
}

// get pubkey bytes
func (privKey *YubiHsmPrivateKey) GetPubKeyBytes() []byte {
	return privKey.pubKeyBytes[:]
}

// get pubkey address
func (privKey *YubiHsmPrivateKey) GetPubKeyAddr() string {
	b := append([]byte{0x4}, privKey.pubKeyUncompressed...)

	ecdsaPubKey, err := crypto.UnmarshalPubkey(b)
	if err != nil {
		privKey.deletePrivKey()
		panic(err)
	}
	pubKeyAddr := crypto.PubkeyToAddress(*ecdsaPubKey)

	return pubKeyAddr.Hex()
}
