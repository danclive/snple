import Vue from "vue";
import VueRouter from "vue-router";
import routes from "./routes/routes";
import store from "./store";

Vue.use(VueRouter);

// configure router
const router = new VueRouter({
  routes, // short for routes: routes
  linkExactActiveClass: "nav-item active"
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (store.getters.isLoggedIn) {
      next();
      return;
    } else {
      next({ name: "登录" });
    }
  } else {
    next();
  }
});

export default router;
