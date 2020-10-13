package service

import (
	"github.com/hwb2017/CMDBDemo/model"
	"time"
)

type CreateVMLifecycleRequest struct {
	Maintainer string `json:"maintainer"`
	Applicant string `json:"applicant" binding:"required"`
	VMLifecycleRules []VMLifecycleRule `json:"vm_lifecycle_rules" binding:"required,dive"`
	VMIDs []string `json:"vm_ids" binding:"required,dive"`
}

type VMLifecycleRule struct {
	Operation string `json:"operation" binding:"required,oneof=stop destroy"`
	ActionTime time.Time `json:"action_time" binding:"required,datetime=2006-01-02"`
}

func (s *Service) CreateVMLifecycle(param *CreateVMLifecycleRequest) error{
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

func (s *Service) ListVMLifecycle(queryOptions *model.QueryOptions) (interface{}, error) {
	return s.dao.ListVMLifecycle(queryOptions)
}

func (s *Service) CountVMLifecycle(queryOptions *model.QueryOptions) (int, error) {
	return s.dao.CountVMLifecycle(queryOptions)
}