package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	dataFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tri",
	Short: "Tri is a todo application",
	Long:  `Keep track of your tasks with Tri.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	home, err := homedir.Dir()
	if err != nil {
		log.Println(
			"Error: Unable to detect Home Directory\nPlease set data file using --datafile.",
		)
	}

	rootCmd.PersistentFlags().
		StringVar(&dataFile, "datafile", home+string(os.PathSeparator)+".tridos.json", "data file to store tasks")

	rootCmd.PersistentFlags().
		StringVar(&cfgFile, "config", "", "config file (default \""+home+string(os.PathSeparator)+".tri.yaml)\"")
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		log.Println(
			"Error: Unable to detect Home Directory\nPlease set data file using --datafile.",
		)
	}
	// defaults in case a config file is not present
	viper.SetDefault("datafile", home+string(os.PathSeparator)+".tridos.json")

	viper.SetConfigName(".tri")
	viper.AddConfigPath(home)
	viper.AutomaticEnv()
	viper.SetEnvPrefix("tri")

	// non default config file used
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Config file:", viper.ConfigFileUsed(), "not found.")
		fmt.Println("Reverting to default configuration.")
		cfgFilePath := home + string(os.PathSeparator) + ".tri.yaml"

		f, err := os.Create(cfgFilePath)
		if err != nil {
			fmt.Println("Error: Unable to create file.")
		}
		defer f.Close()

		f.WriteString("datafile: " + home + string(os.PathSeparator) + ".tridos.json")

		f.Sync()
		fmt.Println("Created default configuration file:", cfgFilePath)
	}
}
