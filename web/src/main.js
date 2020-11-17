import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import { Layout, Menu, Icon, Button, Table, Card, Descriptions, Input, Select, Tag, Pagination, Modal, Form, DatePicker, Spin } from "ant-design-vue";
import moment from "moment";

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
Vue.use(Modal);
Vue.use(Form);
Vue.use(DatePicker);
Vue.use(Spin);

Vue.filter('datefmt', function(input,fmtstr) {
  return moment.unix(input).format(fmtstr);
})

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");