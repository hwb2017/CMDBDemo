package errcode

var (
	ErrorListVMBasicView = NewError(500000, "列出虚拟机基本视图失败")
	ErrorCreateVMLifecycle = NewError(500100, "创建虚拟机生命周期失败")
	ErrorListVMLifecycle = NewError(500101, "列出虚拟机生命周期失败")
	ErrorGetVMLifecycle = NewError(500102, "获取虚拟机生命周期失败")
	ErrorDeleteVMLifecycle = NewError(500103, "删除虚拟机生命周期失败")
	ErrorUpdateVMLifecycle = NewError(500104,"更新虚拟机生命周期失败")
)
