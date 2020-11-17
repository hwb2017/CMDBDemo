package dao

import "github.com/hwb2017/CMDBDemo/model"

var vmLifecycleAssociationCollection = &model.VMLifecycleAssociationCollection{}

func (d *Dao) CreateVMLifecycleAssociations(vs []model.VMLifecycleAssociation) error {
	return vmLifecycleAssociationCollection.BulkCreate(d.client, vs)
}

func (d *Dao) UpdateVMLifecycleAssociations(id string, vmIDs []string) error {
	return vmLifecycleAssociationCollection.BulkUpdate(d.client, id, vmIDs)
}

func (d *Dao) DeleteVMLifecycleAssociations(id string) error {
	return vmLifecycleAssociationCollection.DeleteByVMLifecycleID(d.client, id)
}