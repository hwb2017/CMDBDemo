import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    component: () => 
      import("@/views/Layout"),
    children: [
      {
        path: "/",
        redirect: "/infrastructure/virtualmachines"
      },
      {
        path: "/infrastructure",
        name: "infrastructure",
        meta: { icon: "database", title: "基础设施" },
        component: { render: h => h("router-view") },
        children: [
          {
            path: "/infrastructure/virtualmachines",
            name: "virtualmachines",
            meta: { title: "虚拟机" },
            component: () => 
              import("@/views/Infrastructure/VirtualMachines")
          }
        ]
      }
    ]
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;