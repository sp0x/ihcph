package common

import (
	"github.com/sp0x/torrentd/config"
	"github.com/spf13/viper"
)

func BindConfig() {
	_ = viper.BindEnv("firebase_project")
	_ = viper.BindEnv("firebase_credentials_file")
}

type FirebaseConfig struct {
	Project     string
	Credentials string
}

func GetFirebaseConfig() *FirebaseConfig {
	project := viper.GetString("firebase_project")
	creds := viper.GetString("firebase_credentials_file")
	if creds == "" {
		panic("no firebase credentials found")
	}
	return &FirebaseConfig{
		Project:     project,
		Credentials: creds,
	}
}

func GetConfig() config.Config {
	c := &config.ViperConfig{}
	_ = c.Set("storage", "firebase")
	//_ = c.Set("firebase_project", "")
	//_ = c.Set("firebase_credentials_file", "")
	return c
}
