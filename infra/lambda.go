package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"os"
	"strings"
	"time"
)

func getLambdas(stack constructs.Construct, stage string) []*apiFunctionResource {
	dirs, err := os.ReadDir("src/app/lambda")
	if err != nil {
		panic(err)
	}

	var apiMeta []*apiFunctionResource

	for _, dir := range dirs {
		dataStrings := strings.Split(dir.Name(), "_")

		name := dataStrings[0]
		method := ""

		if len(dataStrings) == 2 {
			method = dataStrings[1]
		}

		function, functionName := buildLambda(stack, dir.Name(), stage)
		version := buildLambdaVersion(stack, function, functionName)

		md := &apiFunctionResource{
			apiPath:            name,
			apiMethod:          strings.ToUpper(method),
			apiFunctionVersion: version,
		}

		apiMeta = append(apiMeta, md)
	}
	return apiMeta
}

func buildLambda(stack constructs.Construct, path, stage string) (awslambda.IFunction, string) {
	name := path + "-" + stage
	functionProps := &awslambda.FunctionProps{
		FunctionName: &name,
		Code:         awslambda.AssetCode_FromAsset(jsii.String("src/app/lambda/"+path+"/bootstrap.zip"), nil),
		Handler:      jsii.String("bootstrap.zip"),
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Architecture: awslambda.Architecture_ARM_64(),
		CurrentVersionOptions: &awslambda.VersionOptions{
			RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
		},
		Environment: &map[string]*string{
			"STAGE":      s(stage),
			"STACK_NAME": s(appName),
		},
	}

	return awslambda.NewFunction(stack, jsii.String(path+"-lambda-"+stage), functionProps), path
}

func buildLambdaVersion(stack constructs.Construct, function awslambda.IFunction, name string) awslambda.IVersion {

	props := &awslambda.VersionProps{
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
		Lambda:        function,
	}

	return awslambda.NewVersion(stack, s(name+"-version-"+time.Now().String()), props)
}
