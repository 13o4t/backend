package main

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"backend/internal/myapp/app"
	"backend/internal/myapp/ports"
	"backend/internal/pkg/config"
	"backend/internal/pkg/server"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	appName = strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(os.Args[0]))

	rootCmd = &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: "Myapp",
		Long:  "Myapp",
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
	rootCmd.PersistentFlags().StringP("mysvr_grpc_addr", "", "", "mysvr GRPC address")

	_ = viper.BindPFlag("configFile", rootCmd.PersistentFlags().Lookup("config"))
	_ = viper.BindPFlag("server.port", rootCmd.PersistentFlags().Lookup("port"))
	_ = viper.BindPFlag("grpc.mysvr_addr", rootCmd.PersistentFlags().Lookup("mysvr_grpc_addr"))

	envPrefix := strings.ToUpper(appName)
	_ = viper.BindEnv("server.port", envPrefix+"_PORT")
	_ = viper.BindEnv("grpc.mysvr_addr", envPrefix+"_GRPC_MYSVR_ADDR")
}

func run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	applcation, cleanup, err := app.NewApplication(ctx, app.Options{
		GRPCUseTLS:    false,
		MysvrGRPCAddr: viper.GetString("grpc.mysvr_addr"),
	})
	if err != nil {
		return err
	}
	defer cleanup()

	return server.RunHTTP(viper.GetInt("server.port"), func(e *gin.Engine) {
		ports.RegisterHandlers(e, ports.NewHTTPServer(applcation))
	})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
