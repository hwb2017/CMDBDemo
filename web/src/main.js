import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import { Layout, Menu, Icon, Button, Table, Tag, Divider } from "ant-design-vue";

Vue.config.productionTip = false;
Vue.use(Layout);
Vue.use(Menu);
Vue.use(Icon);
Vue.use(Button);
Vue.use(Table);
Vue.use(Tag);
Vue.use(Divider);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");