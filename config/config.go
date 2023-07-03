package config

import "github.com/spf13/viper"

type Config struct {
	DBDriver         string `mapstructure:"DB_DRIVER"`
	DBSoure          string `mapstructure:"DB_SOURCE"`
	DBSoureTest      string `mapstructure:"DB_SOURCE_TEST"`
	ServerAddress    string `mapstructure:"SERVER_ADDRESS"`
	CoinGeckoBaseUrl string `mapstructure:"COIN_GECKO_BASE_URL"`
	EthScanApiKey    string `mapstructure:"ETHSCAN_API_KEY"`
	EthScanUrl       string `mapstructure:"ETHSCAN_BASE_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return

}
