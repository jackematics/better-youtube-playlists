package config

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type EnvConfig struct {
	YoutubeApiKey string
}

func fetchYoutubeApiKey() string {
	secretName := "YOUTUBE_API_KEY"
	region := "eu-west-2"

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	
	svc := secretsmanager.NewFromConfig(config)
	
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}
	
	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Fatal(err.Error())
	}

	var secretMap map[string]string
	if result.SecretString != nil {
		secret := *result.SecretString 

		err = json.Unmarshal([]byte(secret), &secretMap)
		if err != nil {
			log.Fatal("Error fetching Youtube data api key: ", err)
		}
	} 

	return secretMap["YOUTUBE_API_KEY"]
}

var Config = EnvConfig {
	YoutubeApiKey: fetchYoutubeApiKey(),
}