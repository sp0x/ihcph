package function

import "github.com/spf13/viper"

func initConfig() {
	_ = viper.BindEnv("firebase_project")
	_ = viper.BindEnv("firebase_credentials")
	//Chat
	viper.SetDefault("telegram_token", "")
	_ = viper.BindEnv("telegram_token")
}