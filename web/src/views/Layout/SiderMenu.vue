<template>
  <a-menu 
    theme="dark" 
    :defaultSelectedKeys="['1']" 
    :defaultOpenKeys="['sub1']"
    mode="inline"
  >
    <template v-for="item in menuData">
      <a-menu-item 
        v-if="!item.children"
        :key="item.name"
        @click="() => $router.push({ path: item.path, query: $route.query })"
      >
        {{ item.meta.title }}
      </a-menu-item>
      <sub-menu
        v-else
        :key="item.name"
        :menu-info="item"
      >
      </sub-menu>
    </template>
    <a-sub-menu key="sub1">
      <span slot="title"><a-icon type="database" /><span>基础设施管理</span></span>
      <a-menu-item key="1" @click="routeToVirtualMachines">
        虚拟机
      </a-menu-item>
      <a-menu-item key="10" @click="routeToVMLifecycle">
        虚拟机生命周期
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
import SubMenu from "./SubMenu";
export default {
  components: {
    "sub-menu": SubMenu
  }, 
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
            const menuitem = { ...item };
            delete menuitem.children;
            menuData.push(item);
            if (item.children) {
              menuData.push(...this.getMenuData(item.children));
            }
        } else if (
          item.children &&
          !item.hideInMenu
        ) {
          menuData.push(...this.getMenuData(item.children));
        }
      }
      return menuData;
    },
    routeToVirtualMachines() {
        this.$router.push({path:"/infrastructure/virtualmachines"})
        console.log(this.menuData);
    },
    routeToVMLifecycle() {
        this.$router.push({path:"/infrastructure/vmlifecycles"})
        console.log(this.menuData);
    }
  }
}
</script>

<style>
</style>