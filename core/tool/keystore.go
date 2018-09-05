/*
 * Copyright 2018 Primas Lab Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tool

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/primasio/go-ethereum/accounts"
	"github.com/primasio/go-ethereum/accounts/keystore"
	"github.com/primasio/primas-api-sdk-go/config"
)

var clientAccount *accounts.Account
var clientKeystore *keystore.KeyStore
var clientAddress string
var clientPrivateKey string

func LoadKeystore() {
	fmt.Println("Gogal_Keystore_Dir:", config.Gogal_Keystore_Dir)

	clientKeystore = keystore.NewKeyStore(config.Gogal_Keystore_Dir, keystore.LightScryptN, keystore.LightScryptP)

	if len(clientKeystore.Accounts()) == 0 {
		panic(errors.New("client account not found"))
	}

	clientAccount = &clientKeystore.Accounts()[0]

	files, err := ioutil.ReadDir(config.Gogal_Keystore_Dir)
	if err != nil {
		panic(err.Error())
	}

	if len(files) == 0 {
		panic("keystore is nil")
	}

	keyJson, err := ioutil.ReadFile(config.Gogal_Keystore_Dir + "/" + files[0].Name())
	if err != nil {
		panic(err.Error())
	}
	clientAddress = clientAccount.Address.Hex()

	privateKey, err := keystore.DecryptKey(keyJson, config.Gogal_Keystore_Passphrase)
	if err != nil {
		panic(err.Error())
	}

	clientPrivateKey = hex.EncodeToString(privateKey.PrivateKey.D.Bytes())

	err = clientKeystore.Unlock(*clientAccount, config.Gogal_Keystore_Passphrase)
	if err != nil {
		panic(err.Error())
	}
}

func GetClientAccount() *accounts.Account {
	if clientAccount == nil {
		LoadKeystore()
	}
	return clientAccount
}

func GetClientKeystore() *keystore.KeyStore {
	if clientKeystore == nil {
		LoadKeystore()
	}
	return clientKeystore
}

func GetClientAddress() string {
	if clientAddress == "" {
		LoadKeystore()
	}
	return clientAddress
}

func GetClientPrivateKey() string {
	if clientPrivateKey == "" {
		LoadKeystore()
	}
	return clientPrivateKey
}
