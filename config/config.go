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
	"log"

	"github.com/spf13/viper"
)

var (
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
	environment := flag.String("c", "config", "")
	flag.Parse()

	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(*environment)
	config.AddConfigPath("./conf/")
	config.ReadInConfig()

	paramServerVersion := config.GetString("server_version")
	if paramServerVersion != "" {
		Gogal_Server_Version = paramServerVersion
	}

	paramKeystorePassphrase := config.GetString("keystore_passphrase")
	if paramKeystorePassphrase != "" {
		Gogal_Keystore_Passphrase = paramKeystorePassphrase
	}

	paramServerUrl := config.GetString("server_url")
	if paramServerUrl != "" {
		Gogal_Server = paramServerUrl + "/" + Gogal_Server_Version
	}
}
