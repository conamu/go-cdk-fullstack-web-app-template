package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

var vpc awsec2.Vpc

func appStack(scope constructs.Construct, id string, props *awscdk.NestedStackProps) awscdk.Stack {

	stack := awscdk.NewNestedStack(scope, &id, props)

	appCluster := awsecs.NewCluster(stack, jsii.String("template-app-cluster"), &awsecs.ClusterProps{
		Vpc: vpc,
	})

	// Backend task
	awsecspatterns.NewApplicationLoadBalancedFargateService(
		stack,
		jsii.String("template-fargate-service"),
		&awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
			Cluster: appCluster,
			RuntimePlatform: &awsecs.RuntimePlatform{
				CpuArchitecture:       awsecs.CpuArchitecture_ARM64(),
				OperatingSystemFamily: awsecs.OperatingSystemFamily_LINUX(),
			},
			Cpu:          jsii.Number(256),
			DesiredCount: jsii.Number(1),
			TaskImageOptions: &awsecspatterns.ApplicationLoadBalancedTaskImageOptions{
				Image: awsecs.ContainerImage_FromAsset(jsii.String("."), &awsecs.AssetImageProps{
					File:     jsii.String("./dockerfiles/be.Dockerfile"),
					Platform: awsecrassets.Platform_LINUX_ARM64(),
				}),
				ContainerPort: jsii.Number(8080),
			},
			MemoryLimitMiB:     jsii.Number(512),
			PublicLoadBalancer: jsii.Bool(true),
			ListenerPort:       jsii.Number(8080),
		},
	)

	// Frontend Task
	awsecspatterns.NewApplicationLoadBalancedFargateService(
		stack,
		jsii.String("template-fargate-service-frontend"),
		&awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
			Cluster: appCluster,
			RuntimePlatform: &awsecs.RuntimePlatform{
				CpuArchitecture:       awsecs.CpuArchitecture_ARM64(),
				OperatingSystemFamily: awsecs.OperatingSystemFamily_LINUX(),
			},
			Cpu:          jsii.Number(256),
			DesiredCount: jsii.Number(1),
			TaskImageOptions: &awsecspatterns.ApplicationLoadBalancedTaskImageOptions{
				Image: awsecs.ContainerImage_FromAsset(jsii.String("."), &awsecs.AssetImageProps{
					File:     jsii.String("./dockerfiles/fe.Dockerfile"),
					Platform: awsecrassets.Platform_LINUX_ARM64(),
				}),
				ContainerPort: jsii.Number(80),
			},
			MemoryLimitMiB:     jsii.Number(512),
			PublicLoadBalancer: jsii.Bool(true),
			ListenerPort:       jsii.Number(80),
		},
	)
	return stack
}
