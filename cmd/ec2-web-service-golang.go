package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"ec2-web-service-golang/pkg"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type Ec2WebServiceGolangStackProps struct {
	awscdk.StackProps
}

func NewEc2WebServiceGolangStack(scope constructs.Construct, id string, props *Ec2WebServiceGolangStackProps) awscdk.Stack {
	stack := pkg.NewEC2Stack(scope, "EC2Stack", &pkg.Ec2WebServiceGolangStackProps{})

	// The code that defines your stack goes here

	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("Ec2WebServiceGolangQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewEc2WebServiceGolangStack(app, "Ec2WebServiceGolangStack", &Ec2WebServiceGolangStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
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
