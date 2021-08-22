package aws

import (
	"context"
	"docker-pull-push/pkg/utils"
	b64 "encoding/base64"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"strings"
)


func Client(region string, profile string) *ecr.Client {
	awsConfig := aws.Config{}
	var err error
	if region == "" {
		awsConfig, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	} else if profile == "" {
		awsConfig, err = config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))
	} else if (region == "") && (profile == "") {
		awsConfig, err = config.LoadDefaultConfig(context.TODO())
	} else {
		awsConfig, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion(region), config.WithSharedConfigProfile(profile))
	}
	utils.Check(err)
	client := ecr.NewFromConfig(awsConfig)
	return client
}

func GetAuthToken(svc *ecr.Client) string {
	token := ""
	ecrSvc, err := svc.GetAuthorizationToken(context.TODO(), &ecr.GetAuthorizationTokenInput{RegistryIds: []string{"884749261746"}})
	utils.Check(err)
	for _, x := range ecrSvc.AuthorizationData {
		sDec, _ := b64.StdEncoding.DecodeString(aws.ToString(x.AuthorizationToken))
		sDecs := utils.BytesToString(sDec)
		s := strings.Split(sDecs, ":")
		token = s[1]

	}
	return token
}