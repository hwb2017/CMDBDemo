<template>
  <a-menu 
    theme="dark" 
    :default-selected-keys="['1']" 
    mode="inline"
  >
    <a-sub-menu key="sub1">
      <span slot="title"><a-icon type="database" /><span>基础设施管理</span></span>
      <a-menu-item key="1">
        虚拟机
      </a-menu-item>
      <a-menu-item key="2">
        对象存储
      </a-menu-item>
    </a-sub-menu>
    <a-sub-menu key="sub2">
      <span slot="title"><a-icon type="appstore" /><span>应用管理</span></span>
      <a-menu-item key="3">
        概览
      </a-menu-item>
      <a-menu-item key="4">
        服务树
      </a-menu-item>
    </a-sub-menu>
    <a-sub-menu key="sub3">
      <span slot="title"><a-icon type="api" /><span>OpenAPI列表</span></span>
      <a-menu-item key="5">
        监控告警域
      </a-menu-item>
      <a-menu-item key="6">
        资源管理域
      </a-menu-item>
    </a-sub-menu>
    <a-sub-menu key="sub4">
      <span slot="title"><a-icon type="setting" /><span>系统设置</span></span>
      <a-menu-item key="7">
        权限管理
      </a-menu-item>
      <a-menu-item key="8">
        云API管理
      </a-menu-item>
      <a-menu-item key="9">
        LDAP设置
      </a-menu-item>      
    </a-sub-menu>    
  </a-menu>  
</template>

<script>
export default {
  data() {
    const menuData = this.getMenuData(this.$router.options.routes); 
    return {
      menuData
    };
  },
  methods: {
    getMenuData(routes = []) {
      const menuData = [];
      for (let item of routes) {
        if (item.name && !item.hideInMenu) {
          console.log(item)
        }
      }
      return menuData;
    }
  }
}
</script>

<style>
</style>