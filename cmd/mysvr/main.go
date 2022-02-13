package main

import (
	"backend/internal/mysvr/app"
	"backend/internal/mysvr/ports"
	"backend/internal/pkg/config"
	"backend/internal/pkg/genprotobuf/mysvr"
	"backend/internal/pkg/server"
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var (
	appName = strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(os.Args[0]))

	rootCmd = &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: "Mysvr",
		Long:  "Mysvr",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		RunE: run,
	}
)

func init() {
	cobra.OnInitialize(func() {
		configFile := viper.GetString("configFile")
		if configFile != "" {
			_ = config.LocalSpecified(configFile)
		} else {
			_, _ = config.Local(appName, "config")
		}
	})

	rootCmd.PersistentFlags().StringP("config", "c", "", "config file")
	rootCmd.PersistentFlags().IntP("port", "", 80, "http port")

	_ = viper.BindPFlag("configFile", rootCmd.PersistentFlags().Lookup("config"))
	_ = viper.BindPFlag("server.port", rootCmd.PersistentFlags().Lookup("port"))

	envPrefix := strings.ToUpper(appName)
	_ = viper.BindEnv("server.port", envPrefix+"_PORT")
}

func run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	application := app.NewApplication(ctx)

	return server.RunGRPC(viper.GetInt("server.port"), func(server *grpc.Server) {
		mysvr.RegisterMysvrServiceServer(server, ports.NewGRPCServer(application))
	})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
