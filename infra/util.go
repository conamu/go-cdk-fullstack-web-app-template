package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"regexp"
)

type stackProps struct {
	awscdk.StackProps
}

const appName = "AWS-CDK-Template"

var stage string
var StackName string

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
