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

package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	CONST_Config_Dir = "/Users/kevinchen/Documents/Golang/primas/primas-api-sdk-go/src/github.com/primasio/primas-api-sdk-go/conf"
)

var (
	Gogal_Keystore_Dir        string = `/Users/kevinchen/Documents/Golang/primas/primas-api-sdk-go/src/github.com/primasio/primas-api-sdk-go/keystore`
	Gogal_Keystore_Passphrase string = "Test123:::"

	Gogal_Server_Version   string = "v3"
	Gogal_Testnet_Server   string = "https://staging.primas.io"
	Gogal_Mainnet_Server   string = "https://rigel-a.primas.network"
	Gogal_Localhost_Server string = "http://10.0.0.5:8080"

	Gogal_Server string = Gogal_Testnet_Server + "/" + Gogal_Server_Version
)

var config *viper.Viper

func init() {
	log.Println("read config init... ")
	isTesting := false
	for _, item := range os.Args {
		if strings.ToLower(item) == "-test.v" {
			isTesting = true
			break
		}
		if strings.ToLower(item) == "-test.v=true" {
			isTesting = true
			break
		}
		if strings.ToLower(item) == "-test.run" {
			isTesting = true
			break
		}
	}

	environment := flag.String("c", "config", "")
	flag.Parse()

	config := viper.New()
	if isTesting {
		log.Println("testing env... ")

		config.AddConfigPath(CONST_Config_Dir)
	} else {
		config.AddConfigPath("./conf/")
	}
	config.SetConfigType("yaml")
	config.SetConfigName(*environment)

	config.ReadInConfig()
	fmt.Println("config filename:", config.ConfigFileUsed())

	paramServerVersion := config.GetString("server_version")
	if paramServerVersion != "" {
		Gogal_Server_Version = paramServerVersion
	}

	paramKeystoreDir := config.GetString("keystore_dir")
	if paramKeystoreDir != "" {
		Gogal_Keystore_Dir = paramKeystoreDir
	}

	paramKeystorePassphrase := config.GetString("keystore_passphrase")
	if paramKeystorePassphrase != "" {
		Gogal_Keystore_Passphrase = paramKeystorePassphrase
	}

	paramServerUrl := config.GetString("server_url")
	if paramServerUrl != "" {
		Gogal_Server = paramServerUrl + "/" + Gogal_Server_Version
	}

	if isTesting {
		paramLocaltestServerUrl := config.GetString("localtest_server_url")
		if paramLocaltestServerUrl != "" {
			Gogal_Server = paramLocaltestServerUrl + "/" + Gogal_Server_Version
		}

		paramLocaltestKeystoreDir := config.GetString("Localtest_keystore_dir")
		if paramKeystoreDir != "" {
			Gogal_Keystore_Dir = paramLocaltestKeystoreDir
		}

		fmt.Println("Testing Gogal_Server:", Gogal_Server)
	} else {
		fmt.Println("Gogal_Server:", Gogal_Server)
	}
}
