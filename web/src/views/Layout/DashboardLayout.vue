<template>
  <div
    class="wrapper"
    :class="[
      { 'nav-open': $sidebar.showSidebar }
    ]"
  >
    <notifications></notifications>
    <side-bar
      :active-color="sidebarBackground"
      :background-image="sidebarBackgroundImage"
      :data-background-color="sidebarBackgroundColor"
    >
      <!-- <user-menu></user-menu> -->
      <!-- <mobile-menu></mobile-menu> -->
      <template slot="links">
        <sidebar-item
          :link="{ name: '控制台', icon: 'dashboard', path: '/dashboard' }"
        >
        </sidebar-item>
        <sidebar-item
          :link="{ name: '用户管理', icon: 'people', path: '/users' }"
        >
        </sidebar-item>
        <sidebar-item
          :link="{ name: '设备管理', icon: 'devices', path: '/devices' }"
        >
        </sidebar-item>
        <sidebar-item
          :link="{ name: '设置', icon: 'settings', path: '/settings' }"
        >
        </sidebar-item>
        <sidebar-item
          :link="{ name: '注销', icon: 'exit_to_app', path: '/logout' }"
        >
        </sidebar-item>
      </template>
    </side-bar>
    <div class="main-panel">
      <top-navbar></top-navbar>
      <div
        :class="{ content: !$route.meta.hideContent }"
        @click="toggleSidebar"
      >
        <!-- <fade-transition :duration="200"> -->
          <!-- your content here -->
          <router-view></router-view>
        <!-- </fade-transition> -->
      </div>
      <content-footer v-if="!$route.meta.hideFooter"></content-footer>
    </div>
  </div>
</template>
<script>
/* eslint-disable no-new */
import PerfectScrollbar from "perfect-scrollbar";
import "perfect-scrollbar/css/perfect-scrollbar.css";

function hasElement(className) {
  return document.getElementsByClassName(className).length > 0;
}

function initScrollbar(className) {
  if (hasElement(className)) {
    new PerfectScrollbar(`.${className}`);
  } else {
    // try to init it later in case this component is loaded async
    setTimeout(() => {
      initScrollbar(className);
    }, 100);
  }
}

import TopNavbar from "./TopNavbar.vue";
import ContentFooter from "./ContentFooter.vue";
// import MobileMenu from "./Extra/MobileMenu.vue";
// import FixedPlugin from "../../FixedPlugin.vue";
// import UserMenu from "./Extra/UserMenu.vue";
// import { FadeTransition } from "vue2-transitions";

export default {
  components: {
    TopNavbar,
    ContentFooter
    // MobileMenu,
    // FixedPlugin,
    // UserMenu,
    // FadeTransition
  },
  data() {
    return {
      sidebarBackgroundColor: "black",
      sidebarBackground: "green",
      sidebarBackgroundImage: "./img/sidebar-2.jpg",
      sidebarMini: false,
      sidebarImg: true
    };
  },
  created() {
    this.$store.dispatch("Mine").then(() => {

    }).catch(err => {
      if (err.response) {
        if (err.response.status === 404) {
          this.$router.push({ name: "登录" });
        }
      }
    });
  },
  methods: {
    toggleSidebar() {
      if (this.$sidebar.showSidebar) {
        this.$sidebar.displaySidebar(false);
      }
    },
    minimizeSidebar() {
      if (this.$sidebar) {
        this.$sidebar.toggleMinimize();
      }
    }
  },
  mounted() {
    const docClasses = document.body.classList;
    const isWindows = navigator.platform.startsWith("Win");
    if (isWindows) {
      // if we are on windows OS we activate the perfectScrollbar function
      initScrollbar("sidebar");
      initScrollbar("sidebar-wrapper");
      initScrollbar("main-panel");

      docClasses.add("perfect-scrollbar-on");
    } else {
      docClasses.add("perfect-scrollbar-off");
    }
  },
  watch: {
    sidebarMini() {
      this.minimizeSidebar();
    }
  }
};
</script>
<style lang="scss">
$scaleSize: 0.95;
@keyframes zoomIn95 {
  from {
    opacity: 0;
    transform: scale3d($scaleSize, $scaleSize, $scaleSize);
  }
  to {
    opacity: 1;
  }
}
.main-panel .zoomIn {
  animation-name: zoomIn95;
}
@keyframes zoomOut95 {
  from {
    opacity: 1;
  }
  to {
    opacity: 0;
    transform: scale3d($scaleSize, $scaleSize, $scaleSize);
  }
}
.main-panel .zoomOut {
  animation-name: zoomOut95;
}
</style>
