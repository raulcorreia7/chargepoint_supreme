package main

import (
	"backend/config"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	lambda "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	lambdago "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type BackendStackProps struct {
	awscdk.StackProps
}

func ChargepointStack(scope constructs.Construct, id string, props *BackendStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here
	lambdago.NewGoFunction(stack, jsii.String(config.CityHandler), &lambdago.GoFunctionProps{
		Entry: jsii.String("handlers/city/main.go"),
		Bundling: &lambdago.BundlingOptions{
			GoBuildFlags: jsii.Strings("-ldflags \"-s -w\""),
			GoProxies: &[]*string{
				lambdago.GoFunction_GOOGLE_GOPROXY(),
				jsii.String("direct"),
			},
		},

		Runtime:      lambda.Runtime_PROVIDED_AL2023(),
		Architecture: lambda.Architecture_ARM_64(),
	})
	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	ChargepointStack(app, config.StackName, &BackendStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {

	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
