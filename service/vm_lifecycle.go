package service

import (
	"github.com/hwb2017/CMDBDemo/model"
)

type CreateVMLifecycleRequest struct {
	Maintainer string `json:"maintainer"`
	Applicant string `json:"applicant"`
	VMLifecycleRules []model.VMLifecycleRule `json:"vm_lifecycle_rules"`
	VMIDs []string `json:"vm_ids"`
}

func (s *Service) CreateVMLifecycle(param *CreateVMLifecycleRequest) error{
	vmLifecycleID, err := s.dao.CreateVMLifecycle(
    	param.Applicant,
    	param.Maintainer,
    	param.VMIDs,
    	param.VMLifecycleRules)
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