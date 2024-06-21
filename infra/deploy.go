package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

// this is where the infra is built and deployed

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	stack := awscdk.NewStack(app, jsii.String("templateStack"), &awscdk.StackProps{
		Env:         env(),
		Description: jsii.String("Template Stack"),
	})

	networkingStack(stack, "templateNetStack", &awscdk.NestedStackProps{
		Description: jsii.String("Networking Stack"),
	})

	appStack(stack, "templateAppStack", &awscdk.NestedStackProps{
		Description: jsii.String("application stack"),
	})

	app.Synth(nil)
}
