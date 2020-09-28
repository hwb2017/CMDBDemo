package cloudapi

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func SyncAWSInstances() {
	mySession, _ := session.NewSession()
	svc := ec2.New(mySession)
	input := &ec2.DescribeInstancesInput{}
	result, err := svc.DescribeInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				panic(aerr.Error())
			}
		} else {
			panic(err.Error())
		}
	}
	fmt.Println(result)
}