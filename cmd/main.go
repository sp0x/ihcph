package main

import (
	"fmt"
	"github.com/sp0x/torrentd/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var appName = "ihcph"
var rootCmd = &cobra.Command{
	Use:   appName,
	Short: "International house appointment tracking service.",
	Run:   runWatcher,
}
var configFile = ""
var appConfig config.ViperConfig
var indexSite string

func init() {
	cobra.OnInitialize(func() {
		appConfig = initConfig(configFile, appName)
	})
	flags := rootCmd.PersistentFlags()
	var verbose bool
	var singleRun bool
	storage := ""
	chatDb := ""
	flags.BoolVarP(&verbose, "verbose", "v", false, "Show more logs.")
	flags.StringVar(&configFile, "config", "", fmt.Sprintf("The configuration file to use. By default it is ~/.%s/.%s.yaml",
		appName, appName))
	flags.StringVarP(&indexSite, "indexer", "x", "ihcph", "The ihcph site to use.")
	flags.BoolVarP(&singleRun, "single_run", "s", false, "Only checks for results once.")
	//Chat
	flags.StringVarP(&chatDb, "chat_db", "c", "./db/chats.db", "The database where chats would be stored")
	//Storage
	flags.StringVarP(&storage, "storage", "o", "boltdb", `The storage backing to use.
Currently supported storage backings: boltdb, firebase, sqlite`)
	firebaseProject := ""
	firebaseCredentials := ""
	flags.StringVarP(&firebaseCredentials, "firebase_project", "", "", "The project id for firebase")
	flags.StringVarP(&firebaseProject, "firebase_credentials_file", "", "", "The service credentials for firebase")

	viper.SetDefault("verbose", false)
	_ = viper.BindPFlag("verbose", flags.Lookup("verbose"))
	_ = viper.BindEnv("verbose")
	viper.SetDefault("indexer", "ihcph")
	_ = viper.BindPFlag("indexer", flags.Lookup("indexer"))
	_ = viper.BindEnv("indexer")
	//Single run config.
	_ = viper.BindEnv("single_run")
	_ = viper.BindPFlag("single_run", flags.Lookup("single_run"))
	//Storage config
	_ = viper.BindPFlag("storage", flags.Lookup("storage"))
	_ = viper.BindEnv("storage")
	//Firestore related
	_ = viper.BindPFlag("firebase_project", flags.Lookup("firebase_project"))
	_ = viper.BindEnv("firebase_project")
	_ = viper.BindPFlag("firebase_credentials_file", flags.Lookup("firebase_credentials_file"))
	_ = viper.BindEnv("firebase_credentials_file")
	//Chat
	viper.SetDefault("telegram_token", "")
	_ = viper.BindEnv("telegram_token")
	viper.SetDefault("chat_db", "./db/chats.db")
	_ = viper.BindEnv("chat_db")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
