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
	flags.BoolVarP(&verbose, "verbose", "v", false, "Show more logs")
	flags.StringVar(&configFile, "config", "", fmt.Sprintf("The configuration file to use. By default it is ~/.%s/.%s.yaml",
		appName, appName))
	flags.StringVarP(&indexSite, "indexer", "x", "ihcph", "The ihcph site to use.")
	viper.SetDefault("verbose", false)
	_ = viper.BindPFlag("verbose", flags.Lookup("verbose"))
	_ = viper.BindEnv("verbose")
	viper.SetDefault("indexer", "ihcph")
	_ = viper.BindPFlag("indexer", flags.Lookup("indexer"))
	_ = viper.BindEnv("indexer")
	viper.SetDefault("telegram_token", "")
	_ = viper.BindEnv("telegram_token")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
