package cmd

import (
	"github.com/rms1000watt/ecs-task-router/router"

	"github.com/spf13/cobra"
)

var routerCmd = &cobra.Command{
	Use:   "router",
	Short: "This starts the router",
	Run:   runRouter,
}

var routerCfg router.Config

func init() {
	rootCmd.AddCommand(routerCmd)

	// routerCmd.Flags().StringVarP(&routerCfg.File, "file", "f", "abr.yml", "aiWARE Black Runner config file")

	setFlagsFromEnv(routerCmd)
}

func runRouter(cmd *cobra.Command, args []string) {
	configureLogging()

	router.Router(routerCfg)
}
