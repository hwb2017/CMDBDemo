package dao

import "github.com/hwb2017/CMDBDemo/model"

func (d *Dao) CreateVMLifecycleAssociation(vmLifecycleID, vmID string) error {
	vmLifecycleAssociation := model.VMLifecycleAssociation{
		VMLifecycleID: vmLifecycleID,
		VMID: vmID,
	}
	return vmLifecycleAssociation.Create(d.client)
}