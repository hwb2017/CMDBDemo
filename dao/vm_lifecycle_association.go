package dao

import "github.com/hwb2017/CMDBDemo/model"

func (d *Dao) CreateVMLifecycleAssociations(vs []model.VMLifecycleAssociation) error {
	vmLifecycleAssociationCollection := model.VMLifecycleAssociationCollection{}
	return vmLifecycleAssociationCollection.BulkCreate(d.client, vs)
}