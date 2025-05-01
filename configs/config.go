package configs

import "github.com/spf13/viper"

type DBConfig struct {
	Container string `mapstructure:"container"`
	Port      string `mapstructure:"port"`
	Database  string `mapstructure:"database"`
	User      string `mapstructure:"user"`
	Pass      string `mapstructure:"pass"`
}

func init() {
	viper.SetDefault("db.port", "5432")
	viper.SetDefault("db.container", "localhost")
	viper.SetDefault("db.database", "livreria")
	viper.SetDefault("db.user", "postgres")
	viper.SetDefault("db.pass", "postgres")
}
func LoadDBConfig() (*DBConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg DBConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
