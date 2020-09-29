package service

import (
	"encoding/json"
	"fmt"
	"github.com/hwb2017/CMDBDemo/lib/cloudapi"
	"github.com/hwb2017/CMDBDemo/model"
	"github.com/hwb2017/CMDBDemo/utils"
	"time"
)

func (s *Service) ListVMBasicView() (interface{}, error) {
	return s.dao.ListVMBasicView()
}

func (s *Service) SyncInstances() error {
	return nil
}

func (s *Service) SyncAlicloudInstances() error {
	rawInstances, err := cloudapi.DescribeAlicloudInstances()
	if err != nil {
		return err
	}
	instancesMapping := make(map[string]interface{},len(rawInstances))
	instanceIds := make([]string, 0, len(rawInstances))
    for _, v := range rawInstances {
		instanceAttr := make(map[string]interface{})
		j, _ := json.Marshal(v)
		json.Unmarshal(j, &instanceAttr)
		delete(instanceAttr, "Cpu")
		instanceId := fmt.Sprintf("%v", instanceAttr["InstanceId"])
		instanceAttr["_id"] = instanceId
		instanceAttr["_syncTime"] = time.Now().Format("2006-01-02 15:04:05")
		instancesMapping[instanceId] = instanceAttr
		instanceIds = append(instanceIds, instanceId)
	}

	storedInstanceIds := make([]string, 0, len(instancesMapping))
	res, err := s.dao.FindWithProjection("infrastructure","alicloud_instance", "_id")
	if err != nil {
		return err
	}
	for _, v := range res {
		storedInstanceIds = append(storedInstanceIds, fmt.Sprintf("%v", v["_id"]))
	}

	insertInstanceIds := utils.StrSliceDiff(instanceIds, storedInstanceIds)
	deleteInstanceIds := utils.StrSliceDiff(storedInstanceIds, instanceIds)
	updateInstanceIds := utils.StrSliceIntersection(storedInstanceIds, instanceIds)

	bulkSyncModels := model.BulkSyncModels{
		ModelMapping: instancesMapping,
		InsertIDs: insertInstanceIds,
		DeleteIDs: deleteInstanceIds,
		UpdateIDs: updateInstanceIds,
	}
	err = s.dao.BulkSync("infrastructure","alicloud_instance", bulkSyncModels)
	if err != nil {
		return err
	}
	return nil
}