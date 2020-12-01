package main

import (
	"fmt"
	"log"
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

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	// creates a config directory
	dir, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cfgPath := dir + "/.rally"
	fileInfo, err := os.Stat(cfgPath)

	if os.IsNotExist(err) {
		if err := os.Mkdir(cfgPath, 0777); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fileInfo, err = os.Stat(cfgPath)
	}

	if !fileInfo.IsDir() {
		log.Fatal("a file named .rally exists in your home folder - please remove it")
		os.Exit(1)
	}

	viper.AddConfigPath(cfgPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			// create the file?
		} else {
			// Config file was found but another error was produced
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
