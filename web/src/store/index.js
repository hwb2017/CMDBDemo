import Vue from "vue";
import Vuex from "vuex";
import infra from "./modules/infrastructure";

Vue.use(Vuex);

const state = {
  menuCollapsed: false
}

const mutations = {
  toggleMenuCollapsed(state) {
    state.menuCollapsed = !state.menuCollapsed;
  }
}

export default new Vuex.Store({
  state,
  mutations,
  actions: {},
  modules: {
    infra
  }
});
