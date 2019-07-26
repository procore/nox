package main

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/procore/nox/internal/gaia"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	client   *gaia.Client
	cfgFile  string
	override bool
	body     string
	version  string
)

var rootCmd = &cobra.Command{
	Use:     "nox",
	Short:   "Elasticsearch infrastructure management tool",
	Long:    `A grand unified elasticsearch cli`,
	Version: version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		client = configESClient()
	},
}

// Execute is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/nox.yaml)")
	rootCmd.PersistentFlags().StringP("host", "H", "localhost", "host of your elasticsearch cluster")
	rootCmd.PersistentFlags().StringP("port", "p", "9200", "port for communication with your elasticsearch cluster")
	rootCmd.PersistentFlags().StringP("username", "u", "", "username for authentication with the cluster")
	rootCmd.PersistentFlags().StringP("password", "W", "", "password for authentication with the cluster")
	rootCmd.PersistentFlags().BoolP("tls", "t", false, "use TLS for cluster connections")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "toggle debug setting")
	rootCmd.PersistentFlags().Bool("pretty", true, "toggle pretty printing of returned json")
	rootCmd.PersistentFlags().Bool("silent", false, "toggle silent output")

	viper.BindPFlag("tls", rootCmd.PersistentFlags().Lookup("tls"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("pretty", rootCmd.PersistentFlags().Lookup("pretty"))
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("silent", rootCmd.PersistentFlags().Lookup("silent"))

}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/elasticsearch/")
		viper.SetConfigName(".nox")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
	}

}

func configESClient() *gaia.Client {
	config := gaia.NewConfig()
	config.Net.TLS.Enable = viper.GetBool("tls")
	config.User.Name = viper.GetString("username")
	config.User.Password = viper.GetString("password")
	config.Debug = viper.GetBool("debug")
	config.Net.Host = viper.GetString("host")
	config.Net.Port = viper.GetString("port")
	config.Pretty = viper.GetBool("pretty")
	return gaia.NewClient(config)
}
