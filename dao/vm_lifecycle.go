package dao

import (
	"github.com/hwb2017/CMDBDemo/model"
	"time"
)

var vmLifecycleCollection = &model.VMLifecycleCollection{}

func (d *Dao) CreateVMLifecycle(applicant, maintainer string, vmIds []string, rules []model.VMLifecycleRule) (resultID string, err error) {
	vmLifecycle := model.VMLifecycle{
		Applicant: applicant,
		Maintainer: maintainer,
		VMIDs: vmIds,
		VMLifecycleRules: rules,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	return vmLifecycleCollection.Create(d.client, vmLifecycle)
}

func (d *Dao) UpdateVMLifecycle(id, applicant, maintainer string, vmIds []string, rules []model.VMLifecycleRule) error {
	vmLifecycle := model.VMLifecycle{
		Applicant: applicant,
		Maintainer: maintainer,
		VMIDs: vmIds,
		VMLifecycleRules: rules,
		UpdateTime: time.Now().Unix(),
	}
	return vmLifecycleCollection.Update(d.client, id, vmLifecycle)
}

func (d *Dao) ListVMLifecycle(queryOptions *model.QueryOptions) (interface{}, error) {
	return vmLifecycleCollection.ListWithAssociation(d.client, queryOptions)
}

func (d *Dao) GetVMLifecycle(id string) (interface{}, error) {
	return vmLifecycleCollection.GetWithAssociation(d.client, id)
}

func (d *Dao) CountVMLifecycle(queryOptions *model.QueryOptions) (int, error) {
	return vmLifecycleCollection.Count(d.client, queryOptions)
}

func (d *Dao) DeleteVMLifecycle(id string) (int, error) {
	return vmLifecycleCollection.Delete(d.client, id)
}