package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aws/aws-sdk-go/aws/session"
)

type AliCloud struct {
	Client *ecs.Client
}

type AWS struct {
	Sess *session.Session
}

func NewAliCloudClient(accessKey, accessSecret string) (*AliCloud, error) {
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", accessKey, accessSecret)
	if err != nil {
		return nil, err
	}
	return &AliCloud{Client: client}, nil
}

// NewAWSSession read aws credentials from environment: AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY
func NewAWSSession() (*AWS, error){
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	return &AWS{Sess: sess}, nil
}