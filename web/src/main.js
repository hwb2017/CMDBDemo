import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import { Layout, Menu, Icon, Button, Table, Card, Descriptions, Input, Select, Tag, Pagination } from "ant-design-vue";

Vue.config.productionTip = false;
Vue.use(Layout);
Vue.use(Menu);
Vue.use(Icon);
Vue.use(Button);
Vue.use(Table);
Vue.use(Card);
Vue.use(Descriptions);
Vue.use(Input);
Vue.use(Select);
Vue.use(Tag);
Vue.use(Pagination);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");