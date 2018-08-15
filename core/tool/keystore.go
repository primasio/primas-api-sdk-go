package tool

import (
	"encoding/hex"
	"errors"
	"io/ioutil"

	"github.com/primasio/go-ethereum/accounts"
	"github.com/primasio/go-ethereum/accounts/keystore"
	"github.com/primasio/primas-api-sdk-go/config"
)

var clientAccount *accounts.Account
var clientKeystore *keystore.KeyStore
var clientAddress string
var clientPrivateKey string

func init() {
	fileDir := "../../keystore"
	clientKeystore = keystore.NewKeyStore(fileDir, keystore.LightScryptN, keystore.LightScryptP)

	if len(clientKeystore.Accounts()) == 0 {
		panic(errors.New("client account not found"))
	}

	clientAccount = &clientKeystore.Accounts()[0]

	files, err := ioutil.ReadDir(fileDir)
	if err != nil {
		panic(err.Error())
	}

	if len(files) == 0 {
		panic("keystore is nil")
	}

	keyJson, err := ioutil.ReadFile(fileDir + "/" + files[0].Name())
	if err != nil {
		panic(err.Error())
	}
	clientAddress = clientAccount.Address.Hex()

	privateKey, err := keystore.DecryptKey(keyJson, config.CONST_Keystore_Passphrase)
	if err != nil {
		panic(err.Error())
	}

	clientPrivateKey = hex.EncodeToString(privateKey.PrivateKey.D.Bytes())

	err = clientKeystore.Unlock(*clientAccount, config.CONST_Keystore_Passphrase)
	if err != nil {
		panic(err.Error())
	}
}

func GetClientAccount() *accounts.Account {
	return clientAccount
}

func GetClientKeystore() *keystore.KeyStore {
	return clientKeystore
}

func GetClientAddress() string {
	return clientAddress
}

func GetClientPrivateKey() string {
	return clientPrivateKey
}
