package main

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sythe21/go-gin-template"
)

type rootOpts struct {
	env  string
	port int
}

func (opts *rootOpts) Command() *cobra.Command {
	viper.SetEnvPrefix("SERVER")

	cmd := &cobra.Command{
		Use:   "server",
		Short: "Gin HTTP server",
		RunE:  opts.RunE,
	}
	cmd.Flags().IntVarP(&opts.port, "port", "p", 8888, "Port to run the HTTP server. SERVER_PORT also supported")
	viper.BindPFlag("port", cmd.PersistentFlags().Lookup("port"))
	cmd.Flags().StringVarP(&opts.env, "env", "e", "production", "Environment server is running. SERVER_ENV also supported")
	viper.BindPFlag("env", cmd.PersistentFlags().Lookup("env"))

	return cmd
}

func (opts *rootOpts) RunE(cmd *cobra.Command, args []string) error {
	return server.AppInit(opts.env).Run(fmt.Sprintf(":%d", opts.port))
}


func main() {
	root := &rootOpts{}

	if err := root.Command().Execute(); err != nil {
		os.Exit(1)
	}
}
