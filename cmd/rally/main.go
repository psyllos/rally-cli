package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/psyllos/rally-cli/internal/build"
	"github.com/psyllos/rally-cli/pkg/client"
	"github.com/psyllos/rally-cli/pkg/cmd/root"
	"github.com/psyllos/rally-cli/pkg/context"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func main() {
	buildDate := build.Date
	buildVersion := build.Version

	cobra.OnInitialize(initConfig)

	rallyAPI := client.NewClient(&http.Client{})
	cmdContext := context.NewCmdContext(rallyAPI, buildVersion, buildDate)

	rootCmd := root.NewRootCmd(cmdContext)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rally.yaml)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".rally" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".rally")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
