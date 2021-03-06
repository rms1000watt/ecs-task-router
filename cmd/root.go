package cmd

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use: "ecs-task-router",
}

var logLevel string

func init() {
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "info", "Set log level (debug, info, warn, error, fatal)")

	setPFlagsFromEnv(rootCmd)
}

// Execute is the entrypoint into cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func configureLogging() {
	if level, err := log.ParseLevel(logLevel); err != nil {
		log.Error("log-level argument malformed: ", logLevel, ": ", err)
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(level)
	}
}

func setPFlagsFromEnv(cmd *cobra.Command) {
	// Courtesy of https://github.com/coreos/pkg/blob/master/flagutil/env.go
	cmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		key := strings.ToUpper(strings.Replace(f.Name, "-", "_", -1))
		if val := os.Getenv(key); val != "" {
			if err := cmd.PersistentFlags().Set(f.Name, val); err != nil {
				fmt.Println("Failed setting flag from env:", err)
			}
		}
	})
}

func setFlagsFromEnv(cmd *cobra.Command) {
	// Courtesy of https://github.com/coreos/pkg/blob/master/flagutil/env.go
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		key := strings.ToUpper(strings.Replace(f.Name, "-", "_", -1))
		if val := os.Getenv(key); val != "" {
			if err := cmd.Flags().Set(f.Name, val); err != nil {
				fmt.Println("Failed setting flag from env:", err)
			}
		}
	})
}
