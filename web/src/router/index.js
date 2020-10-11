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
          },
          {
            path: "/infrastructure/vmlifecycles",
            name: "vmlifecycles",
            meta: { title: "虚拟机申请计划" },
            component: () => 
              import("@/views/Infrastructure/VMLifecycles")
          },          
        ]
      },
      {
        path: "/setting",
        name: "setting",
        meta: { icon: "setting", title: "系统设置" },
        component: { render: h => h("router-view") },
        children: [
          {
            path: "/setting/permission",
            name: "permission",
            meta: { title: "权限管理" },
            component: () =>
              import("@/views/Setting/Permission")
          },
          {
            path: "/setting/cloudApi",
            name: "cloudapi",
            meta: { title: "云API管理" },
            component: () =>
              import("@/views/Setting/CloudApi")
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