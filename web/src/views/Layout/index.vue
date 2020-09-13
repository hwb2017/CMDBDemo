<template>
  <a-layout id="components-layout-side" style="min-height: 100vh">
    <a-layout-sider 
      v-model="collapsed" 
      collapsible
      :trigger="null"
      width="256px"
    >
      <div class="logo">CMDB Demo</div>
      <SiderMenu></SiderMenu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header style="background: #fff; padding: 0">
        <a-icon :type="collapsed ? 'menu-unfold' : 'menu-fold'" class="trigger" @click="collapsed = !collapsed"></a-icon>
      </a-layout-header>
      <a-layout-content style="margin: 0 16px">
        <router-view></router-view>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>
<script>
import SiderMenu from "./SiderMenu";
export default {
  components: {
    SiderMenu
  },
  data() {
    return {
      collapsed: false,
    };
  },
};
</script>

<style scoped>
.trigger {
  padding: 0 20px;
  line-height: 64px; 
  font-size: 20px;
}
.trigger:hover {
  background-color: #EEEEEE;
}  
.logo {
    height: 64px;
    text-align: center;
    line-height: 64px;
    overflow: hidden;
    color: #ffffff;
}
</style>
