package dao

import (
	"github.com/hwb2017/CMDBDemo/model"
	"time"
)

var vmLifecycleCollection = model.VMLifecycleCollection{}

func (d *Dao) CreateVMLifecycle(applicant, maintainer string, vmIds []string, rules []model.VMLifecycleRule) (resultID string, err error) {
	vmLifecycle := model.VMLifecycle{
		Applicant: applicant,
		Maintainer: maintainer,
		VMIDs: vmIds,
		VMLifecycleRules: rules,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	return vmLifecycleCollection.Create(d.client, vmLifecycle)
}

func (d *Dao) ListVMLifecycle() (interface{}, error) {
	return vmLifecycleCollection.ListWithAssociation(d.client)
}