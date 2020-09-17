import request from "@/utils/request";

const state = {
  virtualmachines: [],
  vmLifecycles: []
};

const actions = {
  async getVirtualMachines({ commit }) {
    const response = await request({
      url: "/baseApi/virtualMachine/ListVMBasicView",
      method: "get"
    });
    commit('saveVirtualMachines', response);
  },
  async getVMLifecycles({ commit }) {
    const response = await request({
      url: "/openApi/vmLifecycle/ListVMLifecycle",
      method: "get"
    });
    commit('saveVMLifecycles', response);
  }
};

const mutations = {
  saveVirtualMachines(state, { data }) {
    state.virtualmachines = data["data"];
  },
  saveVMLifecycles(state, { data }) {
    state.vmLifecycles = data["data"];
  }  
};

export default {
  namespaced: true,
  state,
  actions,
  mutations
};