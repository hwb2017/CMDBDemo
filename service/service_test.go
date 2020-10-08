package service

import (
	"context"
	"github.com/hwb2017/CMDBDemo/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

type MySuite struct{
	svc *Service
	client *mongo.Client
}

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&MySuite{})

var (
	createVMLifecycleReq = &CreateVMLifecycleRequest{
		Applicant: "unitTest",
		Maintainer: "unitTest",
		VMLifecycleRules: []model.VMLifecycleRule{
			{Operation: model.StopOperation,
			 ActionTime: time.Now()},
		},
		VMIDs: []string{"unitTest"},
	}
)

func (s *MySuite) SetUpSuite(c *C) {
	mongodbClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic("failed to create mongodb client")
	}
	mongodbClient.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		panic("failed to ping mongodb")
	}
	s.svc = New(mongodbClient)
	s.client = mongodbClient
}

func (s *MySuite) TearDownSuite(c *C) {
	s.client.Disconnect(context.TODO())
}

func (s *MySuite) TestListBasic(c *C) {
	_, err := s.svc.ListVMBasicView(&model.QueryOptions{})
	c.Assert(err, IsNil)
}

func (s *MySuite) TestListVMLifecycle(c *C) {
	_, err := s.svc.ListVMLifecycle()
	c.Assert(err, IsNil)
}

func (s *MySuite) TestCreateVMLifecycle(c *C) {
	err := s.svc.CreateVMLifecycle(createVMLifecycleReq)
	c.Assert(err, IsNil)
}