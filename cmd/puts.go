package cmd

import (
	"novem/cmd/argpkgs"
	"novem/cmd/fsys"
	"novem/cmd/index"
	"novem/cmd/put"
	"os"

	"github.com/nultero/tics"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var putsCmd = &cobra.Command{
	Use:     "puts [FILES]",
	Short:   argpkgs.PutsDesc,
	Aliases: argpkgs.PutsAliases,
	Args:    cobra.MinimumNArgs(1),
	Run:     puts,
}

// TODOOOO puts should be able to work on a directory; whether prompt or not idk

func puts(cmd *cobra.Command, args []string) {

	nvDir, ihndl := viper.GetString(dataDir), viper.GetString(idxFile)
	ihndl = fsys.MeshPaths(nvDir, ihndl)
	idx := index.Init(ihndl)
	nvDir = fsys.AppendSlash(nvDir)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		tics.ThrowSys(puts, err)
	}

	for _, arg := range args {
		put.Cmd(nvDir, arg, homeDir, RecurseFlag, &idx)
	}

	idx.Write(ihndl)
}

func init() {
	rootCmd.AddCommand(putsCmd)
}
