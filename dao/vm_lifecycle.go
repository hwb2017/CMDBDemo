package dao

import (
	"github.com/hwb2017/CMDBDemo/model"
	"time"
)

func (d *Dao) CreateVMLifecycle(applicant, maintainer string, vmIds []string, rules []model.VMLifecycleRule) (resultID string, err error) {
	vmLifecycle := model.VMLifecycle{
		Applicant: applicant,
		Maintainer: maintainer,
		VMIDs: vmIds,
		VMLifecycleRules: rules,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	return vmLifecycle.Create(d.client)
}