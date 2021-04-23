package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config ...
type Config struct {
	DBConfig    DBConfig
	PolyConfig  PolyConfig
	CurveConfig EthConfig
	BSCConfig   EthConfig
	HecoConfig  EthConfig
}

// DBConfig ...
type DBConfig struct {
	DB string
}

// PolyConfig ...
type PolyConfig struct {
	RestURL                 string
	EntranceContractAddress string
	WalletFile              string
	WalletPwd               string
}

// EthConfig ...
type EthConfig struct {
	SideChainId         uint64
	RestURL             []string
	ECCMContractAddress string
	ECCDContractAddress string
	KeyStorePath        string
	KeyStorePwdSet      map[string]string
	BlockConfig         uint64
}

// LoadConfig ...
func LoadConfig(confFile string) (config *Config, err error) {
	jsonBytes, err := ioutil.ReadFile(confFile)
	if err != nil {
		return
	}

	config = &Config{}
	err = json.Unmarshal(jsonBytes, config)
	return
}
