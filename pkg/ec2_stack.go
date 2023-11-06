package pkg

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type Ec2WebServiceGolangStackProps struct {
	awscdk.StackProps
}

type ec2Props struct {
	role           awsiam.IRole
	vpc            awsec2.IVpc
	security_group awsec2.ISecurityGroup
	key_pair       *awsec2.CfnKeyPair
}

type vpcProps struct {
	role awsiam.IRole
}

type sgProps struct {
	role awsiam.IRole
	vpc  awsec2.IVpc
}

func NewEC2Stack(scope constructs.Construct, id string, props *Ec2WebServiceGolangStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("Ec2WebServiceGolangQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })

	lab_role := awsiam.Role_FromRoleArn(stack, jsii.String("LabRole"), jsii.String("arn:aws:iam::866956573632:role/LabRole"), &awsiam.FromRoleArnOptions{})

	vpcProps := &vpcProps{
		role: lab_role,
	}
	vpc := createNewVPC(stack, "my_vpc", vpcProps)

	sgProps := &sgProps{
		role: lab_role,
		vpc:  vpc,
	}

	sg := createSecurityGroup(stack, "my_sg", sgProps)

	createNewEC2(stack, "FirstVM", &ec2Props{
		role:           lab_role,
		vpc:            vpc,
		security_group: sg,
	})

	return stack
}

func createNewKeyPair(stack awscdk.Stack, name string) awsec2.CfnKeyPair {
	keys := awsec2.NewCfnKeyPair(stack, jsii.String(name), &awsec2.CfnKeyPairProps{
		KeyName:   &name,
		KeyFormat: jsii.String("pem"),
		KeyType:   jsii.String("rsa"),
	})
	return keys
}

func createSecurityGroup(stack awscdk.Stack, name string, props *sgProps) awsec2.SecurityGroup {
	sg := awsec2.NewSecurityGroup(stack, jsii.String(name), &awsec2.SecurityGroupProps{
		Vpc:               props.vpc,
		SecurityGroupName: jsii.String("The awesome security group i created"),
	})
	sg.AddIngressRule(awsec2.Peer_Ipv4(jsii.String("186.154.35.0/24")), awsec2.Port_Tcp(jsii.Number(22)), jsii.String("My ssh rule"), jsii.Bool(false))
	sg.AddIngressRule(awsec2.Peer_Ipv4(jsii.String("186.154.35.0/24")), awsec2.Port_Tcp(jsii.Number(80)), jsii.String("My http rule"), jsii.Bool(false))

	return sg
}

func createNewVPC(stack awscdk.Stack, name string, props *vpcProps) awsec2.Vpc {
	vpc := awsec2.NewVpc(stack, &name, &awsec2.VpcProps{
		VpcName:                      jsii.String("My-name"),
		IpAddresses:                  awsec2.IpAddresses_Cidr(jsii.String("10.0.0.0/18")),
		RestrictDefaultSecurityGroup: jsii.Bool(false),
	})

	return vpc
}

func createNewEC2(stack awscdk.Stack, name string, props *ec2Props) awsec2.Instance {
	instance := awsec2.NewInstance(stack, &name, &awsec2.InstanceProps{
		Vpc:          props.vpc,
		VpcSubnets:   &awsec2.SubnetSelection{SubnetType: awsec2.SubnetType_PUBLIC},
		InstanceType: awsec2.InstanceType_Of(awsec2.InstanceClass_BURSTABLE2, awsec2.InstanceSize_MICRO),
		MachineImage: awsec2.NewGenericLinuxImage(&map[string]*string{
			"us-east-1": jsii.String("ami-05c13eab67c5d8861"),
		}, &awsec2.GenericLinuxImageProps{}),
		SecurityGroup: props.security_group,
		Role:          props.role,
		UserData: awsec2.UserData_ForLinux(&awsec2.LinuxUserDataOptions{
			Shebang: jsii.String("#!/bin/bash\n" +
				"sudo yum update -y\n" +
				"sudo yum install -y golang git\n" +
				"git clone https://github.com/alejovasquero/ec2-web-service-golang.git\n" +
				"cd 'ec2-web-service-golang'\n" +
				"go run cmd/app.go -v"),
		}),
		KeyName: jsii.String("Alejo_key_pair"),
	})

	return instance
}
