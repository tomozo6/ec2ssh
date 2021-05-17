package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tomozo6/ec2ssh/pkg"
)

var (
	Version = "unset"
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "ec2ssh",
	Short:   "ec2ssh is a tool that can easily ssh login to AWS EC2.",
	Long:    "ec2ssh is a tool that can easily ssh login to AWS EC2.",
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		app, _ := pkg.NewApp()
		app.GetSSMinstancesInfo()
		app.GetEC2instancesInfo()
		app.Ssh()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ec2ssh.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("session-manager", "s", false, "use SSM SessionManager. (use the InstanceID instead of IpAddress.)")
	rootCmd.Flags().StringP("ssh-user", "u", "", "ssh user")

	viper.BindPFlag("session-manager", rootCmd.Flags().Lookup("session-manager"))
	viper.BindPFlag("ssh-user", rootCmd.Flags().Lookup("ssh-user"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ec2ssh" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ec2ssh")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
