package cmd

import (
	"fmt"

	 
	"github.com/abhijit1859/eternal_grimoire/internal/config"
	 
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	cfg     config.Config
)

var rootCmd = &cobra.Command{
	Use:   "grimoire",
	Short: "A database backup and recovery CLI",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		cfg, err = config.LoadConfig(cfgFile)
		if err != nil {
			return fmt.Errorf("failed to load configuration (%s): %w", cfgFile, err)
		}

		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {

	rootCmd.PersistentFlags().StringVarP(&cfgFile,
		"config",
		"c",
		"config.yml",
		"Path to YAML configuration file",
	)

}
