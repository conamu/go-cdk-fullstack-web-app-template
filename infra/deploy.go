package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/jsii-runtime-go"
	"github.com/conamu/go-cdk-fullstack-web-app-template/src/pkg/config"
	"github.com/spf13/viper"
	"os"
)

// this is where the infra is built and deployed

func main() {
	defer jsii.Close()

	env := os.Getenv("ENV")
	config.Init(env)
	projectId := viper.GetString("project-id-slug")

	app := awscdk.NewApp(nil)

	stack := awscdk.NewStack(app, s(projectId), &awscdk.StackProps{
		Env:         awsEnv(),
		Description: s(projectId),
	})

	// This is necessary to be able to use git branch names in cloudformation stacks
	stage = removeNumbersAndSpecialChars(stage)

	StackName = buildApplicationName()

	requireApiKey := true

	if stage != "production" && stage != "staging" {
		requireApiKey = false
	}

	lambdaApiMeta := getLambdas(stack, stage)

	// Grant permissions to api gateway to invoke functions
	for _, meta := range lambdaApiMeta {
		meta.apiFunctionVersion.GrantInvoke(awsiam.NewServicePrincipal(s("apigateway.amazonaws.com"), &awsiam.ServicePrincipalOpts{}))
	}
	ApiGatewayRoot := buildApiGateway(stack, StackName)

	buildApiResources(stack, ApiGatewayRoot, lambdaApiMeta, requireApiKey, stage)

	awscdk.NewCfnOutput(stack, s("api-url"), &awscdk.CfnOutputProps{
		Value: ApiGatewayRoot.Url(),
	})

	networkingStack(stack, projectId+"NetStack", &awscdk.NestedStackProps{
		Description: s(projectId + " Networking Stack"),
	})

	appStack(stack, projectId+"AppStack", &awscdk.NestedStackProps{
		Description: s(projectId + " Application stack"),
	})

	app.Synth(nil)
}
