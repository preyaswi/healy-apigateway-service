package config

import "github.com/spf13/viper"

type Config struct {
	Port    string `mapstructure:"PORT"`
	PatientSvc string `mapstructure:"PATIENT_SVC"`
}

var envs = []string{"PORT", "PATIENT_SVC"}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil

}
