package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/spf13/viper"
	"regexp"
)

type stackProps struct {
	awscdk.StackProps
}

var appName = viper.GetString("project-name")

var stage string
var stackName string

func s(s string) *string {
	return jsii.String(s)
}

func newStack(scope constructs.Construct, id string, props *stackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	return stack
}

func removeNumbersAndSpecialChars(input string) string {
	reg := regexp.MustCompile("[^a-zA-Z]+")
	return reg.ReplaceAllString(input, "")
}

func buildApplicationName() string {
	return appName + "-" + stage
}

func awsEnv() *awscdk.Environment {
	return &awscdk.Environment{
		Account: s(viper.GetString("aws-account")),
		Region:  s(viper.GetString("aws-region")),
	}
}
