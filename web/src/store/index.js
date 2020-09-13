import Vue from "vue";
import Vuex from "vuex";
import infra from "./modules/infrastructure";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    infra
  }
});
