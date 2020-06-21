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
	flags.BoolVarP(&verbose, "verbose", "v", false, "Show more logs.")
	flags.StringVar(&configFile, "config", "", fmt.Sprintf("The configuration file to use. By default it is ~/.%s/.%s.yaml",
		appName, appName))
	flags.StringVarP(&indexSite, "indexer", "x", "ihcph", "The ihcph site to use.")
	flags.BoolVarP(&singleRun, "single_run", "s", false, "Only checks for results once.")
	viper.SetDefault("verbose", false)
	_ = viper.BindPFlag("verbose", flags.Lookup("verbose"))
	_ = viper.BindEnv("verbose")
	viper.SetDefault("indexer", "ihcph")
	_ = viper.BindPFlag("indexer", flags.Lookup("indexer"))
	_ = viper.BindEnv("indexer")
	viper.SetDefault("telegram_token", "")
	_ = viper.BindEnv("telegram_token")
	//Single run config.
	_ = viper.BindEnv("single_run")
	_ = viper.BindPFlag("single_run", flags.Lookup("single_run"))
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
