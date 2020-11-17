package service

import (
	"fmt"
	"github.com/hwb2017/CMDBDemo/model"
	"github.com/hwb2017/CMDBDemo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type CreateVMLifecycleRequest struct {
	Maintainer string `json:"maintainer"`
	Applicant string `json:"applicant" binding:"required"`
	VMLifecycleRules []VMLifecycleRule `json:"vm_lifecycle_rules" binding:"required,dive"`
	VMIDs []string `json:"vm_ids"`
}

type UpdateVMLifecycleRequest struct {
	VMLifecycleID string `json:"id" binding:"required"`
	Maintainer string `json:"maintainer"`
	Applicant string `json:"applicant" binding:"required"`
	VMLifecycleRules []VMLifecycleRule `json:"vm_lifecycle_rules" binding:"required,dive"`
	VMIDs []string `json:"vm_ids"`
}

type DeleteVMLifecycleRequest struct {
	VMLifecycleID string `form:"id" binding:"required"`
}

type GetVMLifecycleRequest struct {
	VMLifecycleID string `form:"id" binding:"required"`
}

type VMLifecycleRule struct {
	Operation string `json:"operation" binding:"required,oneof=stop destroy"`
	ActionTime int64 `json:"action_time" binding:"required,3yr-range"`
}

func (s *Service) CreateVMLifecycle(param *CreateVMLifecycleRequest) error{
	queryOptions := &model.QueryOptions{}
	queryOptions.WithSimpleProjection("_id")
	rawVMIDs, err := s.dao.ListVMBasicView(queryOptions)
	if err != nil {
		return err
	}
	validVMIDs := make([]string, len(utils.InterfaceSlice(rawVMIDs)))
	for _, v := range utils.InterfaceSlice(rawVMIDs) {
		row := v.(bson.M)
		validVMIDs = append(validVMIDs, row["_id"].(string))
	}
	intersectionVMIDs := utils.StrSliceIntersection(param.VMIDs, validVMIDs)
	if len(intersectionVMIDs) < len(param.VMIDs) {
        return fmt.Errorf("invalid param, vm ids not exist: %s",
        	strings.Join(utils.StrSliceDiff(param.VMIDs, intersectionVMIDs),","))
	}

	vmLifecycleRules := make([]model.VMLifecycleRule,0)
	for _, rule := range param.VMLifecycleRules {
		vmOperation, err := model.ParseVMOperation(rule.Operation)
		if err != nil {
			return err
		}
		vmLifecycleRules = append(vmLifecycleRules, model.VMLifecycleRule{
			Operation: vmOperation,
			ActionTime: rule.ActionTime,
		})
	}
	vmLifecycleID, err := s.dao.CreateVMLifecycle(
    	param.Applicant,
    	param.Maintainer,
    	param.VMIDs,
    	vmLifecycleRules)
    if err != nil {
    	return err
	}
	vmLifecycleAssociations := make([]model.VMLifecycleAssociation ,len(param.VMIDs))
	for i := 0; i < len(param.VMIDs); i++ {
		vmLifecycleAssociations[i] = model.VMLifecycleAssociation{
			VMLifecycleID: vmLifecycleID,
			VMID: param.VMIDs[i],
			}
	}
	err = s.dao.CreateVMLifecycleAssociations(vmLifecycleAssociations)
	if err != nil {
		return err
	}
    return nil
}

func (s *Service) UpdateVMLifecycle(param *UpdateVMLifecycleRequest) error {
	queryOptions := &model.QueryOptions{}
	queryOptions.WithSimpleProjection("_id")
	rawVMIDs, err := s.dao.ListVMBasicView(queryOptions)
	if err != nil {
		return err
	}
	validVMIDs := make([]string, len(utils.InterfaceSlice(rawVMIDs)))
	for _, v := range utils.InterfaceSlice(rawVMIDs) {
		row := v.(bson.M)
		validVMIDs = append(validVMIDs, row["_id"].(string))
	}
	intersectionVMIDs := utils.StrSliceIntersection(param.VMIDs, validVMIDs)
	if len(intersectionVMIDs) < len(param.VMIDs) {
		return fmt.Errorf("invalid param, vm ids not exist: %s",
			strings.Join(utils.StrSliceDiff(param.VMIDs, intersectionVMIDs),","))
	}

	vmLifecycleRules := make([]model.VMLifecycleRule,0)
	for _, rule := range param.VMLifecycleRules {
		vmOperation, err := model.ParseVMOperation(rule.Operation)
		if err != nil {
			return err
		}
		vmLifecycleRules = append(vmLifecycleRules, model.VMLifecycleRule{
			Operation: vmOperation,
			ActionTime: rule.ActionTime,
		})
	}
	err = s.dao.UpdateVMLifecycle(
		param.VMLifecycleID,
		param.Applicant,
		param.Maintainer,
		param.VMIDs,
		vmLifecycleRules)
	if err != nil {
		return err
	}
	return s.dao.UpdateVMLifecycleAssociations(param.VMLifecycleID, param.VMIDs)
}

func (s *Service) ListVMLifecycle(queryOptions *model.QueryOptions) (interface{}, error) {
	return s.dao.ListVMLifecycle(queryOptions)
}

func (s *Service) GetVMLifecycle(id string) (interface{}, error) {
	return s.dao.GetVMLifecycle(id)
}

func (s *Service) DeleteVMLifecycle(id string) (int, error) {
    deletedCount, err := s.dao.DeleteVMLifecycle(id)
    if err != nil {
    	return 0, err
	}
	err = s.dao.DeleteVMLifecycleAssociations(id)
	if err != nil {
		return 0, err
	}
	return deletedCount, nil
}

func (s *Service) CountVMLifecycle(queryOptions *model.QueryOptions) (int, error) {
	return s.dao.CountVMLifecycle(queryOptions)
}