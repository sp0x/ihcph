package common

import (
	"errors"
	"github.com/sp0x/torrentd/config"
	"github.com/spf13/viper"
)

func BindConfig() {
	viper.AutomaticEnv()
	_ = viper.BindEnv("firebase_project")
	_ = viper.BindEnv("firebase_credentials_file")
}

type FirebaseConfig struct {
	Project     string
	Credentials string
}

func GetFirebaseConfig() (*FirebaseConfig, error) {
	project := viper.GetString("firebase_project")
	creds := viper.GetString("firebase_credentials_file")
	if creds == "" {
		return nil, errors.New("no firebase credentials found")
	}
	return &FirebaseConfig{
		Project:     project,
		Credentials: creds,
	}, nil
}

func GetConfig() config.Config {
	c := &config.ViperConfig{}
	_ = c.Set("storage", "firebase")
	//_ = c.Set("firebase_project", "")
	//_ = c.Set("firebase_credentials_file", "")
	return c
}
