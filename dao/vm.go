package dao

import "github.com/hwb2017/CMDBDemo/model"

var vmCollection = model.VMCollection{}

func (d *Dao) ListVMBasicView() (interface{}, error){
	return vmCollection.ListBasicView(d.client)
}