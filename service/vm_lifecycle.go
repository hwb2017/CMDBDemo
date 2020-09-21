package service

import "github.com/hwb2017/CMDBDemo/model"

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
	for _, vmID := range param.VMIDs {
		err = s.dao.CreateVMLifecycleAssociation(vmLifecycleID, vmID)
		if err != nil {
			return err
		}
	}
    return nil
}