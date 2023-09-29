/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"

	"crownest/config"
	"crownest/github"

	"github.com/argoproj/argo-cd/v2/pkg/apiclient"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crownest",
	Short: "Send ArgoCD diffs to your PRs",
	Long:  "Send ArgoCD diffs to your PRs before merging them",

	Run: func(cmd *cobra.Command, args []string) {
		_, err := apiclient.NewClient(&apiclient.ClientOptions{})
		if err != nil {
			fmt.Println(err)
		}
		eventService := github.NewGithubEventService()
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		r.Any("/hook", func(c *gin.Context) {
			eventService.Handler.HandleEventRequest(c.Request)
		})
		r.Run("0.0.0.0:30000")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&config.CfgFile, "config", "", "config file (default is $HOME/.crownest.yaml)")
	cobra.OnInitialize(config.LoadConfig)
}
