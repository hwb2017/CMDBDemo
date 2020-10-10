<template>
  <a-menu 
    theme="dark" 
    mode="inline"
    :selectedKeys="selectedKeys"
    :openKeys.sync="openKeys"
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
import { mapState } from "vuex";
export default {
  components: {
    "sub-menu": SubMenu
  }, 
  watch: {
    "$route.path": function(val) {
      this.selectedKeys = this.selectedKeysMap[val];
      this.openKeys = this.collapsed ? [] : this.openKeysMap[val];
    },
    "collapsed": function(val) {
      this.openKeys = val ? [] : this.openKeysMap[this.$route.path];
    },
  },
  data() {
    this.selectedKeysMap = {};
    this.openKeysMap = {};
    const menuData = this.getMenuData(this.$router.options.routes);
    return {    
      menuData,
      selectedKeys: this.selectedKeysMap[this.$route.path],
      openKeys: this.collapsed ? [] : this.openKeysMap[this.$route.path]
    };
  },
  computed: {
    ...mapState({
      collapsed: state => state.menuCollapsed
    })
  },
  methods: {
    getMenuData(routes = [], depth = 0, parentKeys = []) {
      const menuData = [];
      for (let item of routes) {
        if (item.name && !item.hideInMenu) {
            item.depth = depth;
            if (item.depth==1 && !item.children) {
              item.standalone=true;
            }
            if (item.depth==1 && item.children) {
              parentKeys = [item.name]
            }
            if (item.depth==2) {
              this.selectedKeysMap[item.path] = [item.name]
              this.openKeysMap[item.path] = parentKeys
            }
            menuData.push(item);
            if (item.children) {
              menuData.push(...this.getMenuData(item.children, depth+1, parentKeys));
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