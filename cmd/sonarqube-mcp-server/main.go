package main

import (
	"context"
	"errors"
	"log"

	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/9506hqwy/sonarqube-mcp-server/pkg/sonarqube"
)

func fromArgument(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, sonarqube.UrlKey{}, viper.GetString("url"))
	ctx = context.WithValue(ctx, sonarqube.UserKey{}, viper.GetString("user"))
	ctx = context.WithValue(ctx, sonarqube.PasswordKey{}, viper.GetString("password"))
	ctx = context.WithValue(ctx, sonarqube.UserTokenKey{}, viper.GetString("token"))
	return ctx
}

var rootCmd = &cobra.Command{
	Use:   "sonarqube-mcp-server",
	Short: "SonarQube MCP Server",
	Long:  "SonarQube MCP Server",
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewMCPServer(
			"SonarQube MCP Server",
			"0.1.0",
			server.WithToolCapabilities(false),
		)

		sonarqube.RegisterTools(s, viper.GetBool("readonly"))

		if err := server.ServeStdio(s, server.WithStdioContextFunc(fromArgument)); err != nil {
			if !errors.Is(err, context.Canceled) {
				log.Fatalf("Server error: %v", err)
			}
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().String("url", "http://127.0.0.1:9000", "SonarQube server URL.")
	rootCmd.PersistentFlags().String("user", "", "SonarQube server username.")
	rootCmd.PersistentFlags().String("password", "", "SonarQube server password.")
	rootCmd.PersistentFlags().String("token", "", "SonarQube server user type token.")
	rootCmd.PersistentFlags().Bool("readonly", true, "HTTP GET method only.")

	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("user", rootCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("readonly", rootCmd.PersistentFlags().Lookup("readonly"))
}

func initConfig() {
	viper.SetEnvPrefix("sonarqube")
	viper.AutomaticEnv()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
