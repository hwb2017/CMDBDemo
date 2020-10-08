package cloudapi

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hwb2017/CMDBDemo/utils"
)

func (a *AWS) DescribeInstances() ([]interface{}, error){
	mySession := a.Sess.Copy(&aws.Config{Region: aws.String("us-east-1")})
	svc := ec2.New(mySession)
	input := &ec2.DescribeInstancesInput{}
	result, err := svc.DescribeInstances(input)
	if err != nil {
        return nil, err
	}
	rawInstances := make([]*ec2.Instance, 0)
	for _, reservation := range result.Reservations{
		rawInstances = append(rawInstances, reservation.Instances...)
	}
	return utils.InterfaceSlice(rawInstances), nil
}