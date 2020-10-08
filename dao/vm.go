package dao

import (
	"github.com/hwb2017/CMDBDemo/model"
)

var vmCollection = &model.VMCollection{}

func (d *Dao) ListVMBasicView(queryOptions *model.QueryOptions) (interface{}, error){
	return vmCollection.ListBasicView(d.client, queryOptions)
}

func (d *Dao) CountVM(queryOptions *model.QueryOptions) (int, error) {
	return vmCollection.Count(d.client, queryOptions)
}