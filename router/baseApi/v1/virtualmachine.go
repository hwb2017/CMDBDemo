package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hwb2017/CMDBDemo/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type vmBasicView struct  {
	InstanceType string `json:instance_type`
	HostName string `json:hostname`
	OSName string `json:osname`
	PublicIPAddresses []string `json:public_ip_addresses`
	PrivateIPAddresses []string `json:private_ip_addresses`
}

func ListVMBasicView(c * gin.Context) {
	collection := global.MongodbClient.Database("infrastructure").Collection("alicloud_instance")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// EipAddress.IpAddress PublicIpAddress.IpAddress[0] HostName OSName InstanceType NetworkInterfaces.NetworkInterface[0].PrimaryIpAddress
	findOpts := options.Find().SetProjection(bson.D{{"_id", 0}, {"EipAddress.IpAddress", 1}, {"HostName", 1},
		{"OSName", 1}, {"InstanceType", 1}, {"PublicIpAddress.IpAddress", 1}, {"NetworkInterfaces.NetworkInterface.PrimaryIpAddress", 1}})
	cursor, err := collection.Find(ctx, bson.D{}, findOpts)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to insert data to mongodb: %v\n", err)
		fmt.Fprint(os.Stderr, errMsg)
		panic(errMsg)
	}
	defer cursor.Close(ctx)

	//virtualMachines := make([]vmBasicView, 0)
	for cursor.Next(ctx) {
		var result bson.M
		//var publicIpAddresses []string
		//var privateIpAddresses []string
		err := cursor.Decode(&result)
		//result["EipAddress"].(map[string]interface{})["IpAddress"]
		//vm := vmBasicView{
		//	InstanceType: result["InstanceType"].(string),
		//	HostName: result["HostName"].(string),
		//	OSName: result["OSName"].(string),
		//}
        //virtualMachines = append(virtualMachines, vm)
        fmt.Println(result)
		if err != nil {fmt.Fprintf(os.Stderr, "Failed to parse mongodb document: %v\n", err)}
	}
	if err := cursor.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse mongodb document: %v\n", err)
	}

}