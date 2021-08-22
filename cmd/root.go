package cmd

import (
	"context"
	"docker-pull-push/pkg/aws"
	"docker-pull-push/pkg/docker"
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"strings"
)

var sourceImage string
var targetImage string
var registryType string
var awsRegion string
var awsProfile string

var rootCmd = &cobra.Command{
	Use:   "docker-pull-push",
	Short: "Cli to migrate docker images from dockerhub to private docker repository",
	Long: `For example:

docker-pull-push -s docker.elastic.co/elasticsearch/elasticsearch:7.13 -t AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com/elasticsearch:7.13.3`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Initializing..")
		serverAddress := strings.Split(targetImage, "/")[0]
		ctx := context.Background()
		cli := docker.GetLogin()
		docker.PullImage(ctx, cli, sourceImage)
		docker.TagImage(ctx, cli, sourceImage, targetImage)
		log.Info("tag docker image completed")
		svc := aws.Client(awsRegion, awsProfile)
		token := aws.GetAuthToken(svc)
		docker.PushImage(ctx, cli, serverAddress, targetImage, "AWS", token)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&sourceImage, "sourceImage", "s", "", "docker.elastic.co/elasticsearch/elasticsearch:7.13")
	rootCmd.PersistentFlags().StringVarP(&targetImage, "targetImage", "t", "", "AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com/elasticsearch:7.13.3")
	rootCmd.PersistentFlags().StringVarP(&registryType, "registryType", "r", "docker", "ECR")
	rootCmd.PersistentFlags().StringVarP(&awsRegion, "awsRegion", "l", "us-east-1", "AWS Region")
	rootCmd.PersistentFlags().StringVarP(&awsProfile, "awsProfile", "p", "default", "AWS profile")
	err := rootCmd.MarkPersistentFlagRequired("sourceImage")
	if err != nil {
		return
	}
	err = rootCmd.MarkPersistentFlagRequired("targetImage")
	if err != nil {
		return
	}
}

func initConfig() {
}
