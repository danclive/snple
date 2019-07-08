import Vue from "vue";
import DashboardPlugin from "./material-dashboard";

// Plugins
import App from "./App.vue";
import Chartist from "chartist";

// plugin setup
// Vue.use(VueRouter);
Vue.use(DashboardPlugin);

// global library setup
Object.defineProperty(Vue.prototype, "$Chartist", {
  get() {
    return this.$root.Chartist;
  }
});

import "material-design-icons-iconfont/dist/material-design-icons.css";

import store from "./store";
import router from "./router";

/* eslint-disable no-new */
const vue = new Vue({
  el: "#app",
  render: h => h(App),
  router,
  store,
  data: {
    Chartist: Chartist
  }
});

window.$notify = vue.$notify;
