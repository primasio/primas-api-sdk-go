package config

const (
	CONST_Keystore_Passphrase = "Test123:::"

	CONST_Server_Version   = "v3"
	CONST_Testnet_Server   = "https://staging.primas.io"
	CONST_Mainnet_Server   = "https://rigel-a.primas.network"
	CONST_Localhost_Server = "http://10.0.0.63:8080"

	CONST_Server = CONST_Localhost_Server + "/" + CONST_Server_Version
)
