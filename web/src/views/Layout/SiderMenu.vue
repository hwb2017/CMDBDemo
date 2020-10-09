<template>
  <a-menu 
    theme="dark" 
    :defaultSelectedKeys="['virtualmachine']" 
    :defaultOpenKeys="['infrastructure']"
    mode="inline"
  >
    <template v-for="item in menuData">
      <a-menu-item 
        v-if="item.standalone"
        :key="item.name"
        @click="() => $router.push({ path: item.path, query: $route.query })"
      >
        {{ item.meta.title }}
      </a-menu-item>
      <sub-menu
        v-if="item.children"
        :key="item.name"
        :menu-info="item"
      >
      </sub-menu>
    </template>
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
    getMenuData(routes = [], depth = 0) {
      const menuData = [];
      for (let item of routes) {
        if (item.name && !item.hideInMenu) {
            item.depth = depth;
            if (item.depth==1 && !item.children) {
              item.standalone=true;
            }
            menuData.push(item);
            if (item.children) {
              menuData.push(...this.getMenuData(item.children, depth+1));
            }
        } else if (
          item.children &&
          !item.hideInMenu
        ) {
          menuData.push(...this.getMenuData(item.children, depth+1));
        }
      }
      return menuData;
    },
  }
}
</script>

<style>
</style>