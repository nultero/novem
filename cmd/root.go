package cmd

import (
	"fmt"

	"github.com/nultero/tics"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

// TODOO make the fields $USER agnostic to distribute across my machines easier

// When called with no args, novem will stat the current
// directory for any files it is tracking and compare against
// its index, without delving into any subdirs(at least not without `-r` flag).
var rootCmd = &cobra.Command{
	Use:   "novem",
	Short: "Dead simple wrapper for managing dotfiles. \n" + flavorText,
	Args:  cobra.NoArgs,
	Run:   getAnyFromIndex,
}

// TODO quick list index from cwd if called with no arguments

func getAnyFromIndex(cmd *cobra.Command, args []string) {
	fmt.Println("yup")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolVarP(&RecurseFlag, "recursive", "r", false, "traverses subdirectories wherever novem was called")
}

func initConfig() {

	confMap = tics.CobraRootInitBoilerPlate(confMap, true)
	confPath := confMap[confFile]
	viper.SetConfigFile(confPath)
	viper.AutomaticEnv()

	// If a config file is found, read it in, else make one with prompt.
	err := viper.ReadInConfig()
	if err != nil {
		tics.RunConfPrompts("novem", confMap, defaultSettings)
		tics.ThrowQuiet("")
	}
}
