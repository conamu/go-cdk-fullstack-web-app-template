package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func networkingStack(scope constructs.Construct, id string, props *awscdk.NestedStackProps) awscdk.Stack {
	netStack := awscdk.NewNestedStack(scope, &id, props)

	vpc = awsec2.NewVpc(netStack, jsii.String("cdk-template-vpc"), &awsec2.VpcProps{
		MaxAzs: jsii.Number(3),
	})

	return netStack
}
