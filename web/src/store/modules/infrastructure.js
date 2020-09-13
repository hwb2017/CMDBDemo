import request from "@/utils/request";

const state = {
  virtualmachines = [];
};

const actions = {
  async getVirtualMachines({ commit }) {
    const response = await request({
      url: "baseApi/virtualMachine/ListVMBasicView",
      method: "get"
    });
    commit('saveVirtualMachines', response);
  }
};

const mutations = {
  saveVirtualMachines(state, { data }) {
    state.virtualmachines = data;
  }
};

export default {
  namespaced: true,
  state,
  actions,
  mutations
};