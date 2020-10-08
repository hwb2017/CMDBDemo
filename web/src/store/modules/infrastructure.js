import request from "@/utils/request";

const state = {
  virtualmachines: [],
  vmTotal: 0,
  vmLifecycles: []
};

const actions = {
  async getVirtualMachines({ commit }, queryOptions) {
    var urlPath = "/baseApi/virtualMachine/ListVMBasicView"

    if (queryOptions) {
        var queryArray = []
        for (let key in queryOptions) {
          var keyStr = JSON.stringify(queryOptions[key])
          if (keyStr!="" && keyStr!="[]" && keyStr!="{}") {
              if (key != "searchItemKey" && key != "searchItemValue") {
                queryArray.push(key+"="+queryOptions[key])
              }
          }
        }
        if (queryOptions["searchItemKey"] && queryOptions["searchItemValue"]) {
          console.log(queryOptions["searchItemKey"],queryOptions["searchItemValue"])
          queryArray.push(queryOptions["searchItemKey"]+"="+queryOptions["searchItemValue"])
        }
        var queryString = queryArray.join("&")
        if (queryString!="") {
          queryString = "?"+queryString
        }
        urlPath = urlPath+queryString
    }
    const response = await request({
      url: urlPath,
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
    state.vmTotal = data["pagination"]["TotalRows"];
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