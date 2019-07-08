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
      <mobile-menu></mobile-menu>
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
        <sidebar-item :link="{ name: 'Pages', icon: 'image' }">
          <sidebar-item
            :link="{ name: 'Pricing', path: '/pricing' }"
          ></sidebar-item>
          <sidebar-item
            :link="{ name: 'Timeline', path: '/pages/timeline' }"
          ></sidebar-item>
          <sidebar-item
            :link="{ name: 'Login', path: '/login' }"
          ></sidebar-item>
          <sidebar-item
            :link="{ name: 'Register', path: '/register' }"
          ></sidebar-item>
          <sidebar-item
            :link="{ name: 'Lock Screen', path: '/lock' }"
          ></sidebar-item>
          <sidebar-item
            :link="{ name: 'User Profile', path: '/pages/user' }"
          ></sidebar-item>
        </sidebar-item>
        <sidebar-item :link="{ name: 'Components', icon: 'apps' }">
          <sidebar-item
            :link="{ name: 'Buttons', path: '/components/buttons' }"
          ></sidebar-item>
          <sidebar-item
            :link="{ name: 'Grid System', path: '/components/grid-system' }"
          ></sidebar-item>
          <sidebar-item
            :link="{ name: 'Panels', path: '/components/panels' }"
          ></sidebar-item>
          <sidebar-item
            :link="{ name: 'Notifications', path: '/components/notifications' }"
          ></sidebar-item>
          <sidebar-item
            :link="{ name: 'Icons', path: '/components/icons' }"
          ></sidebar-item>
          <sidebar-item
            :link="{ name: 'Typography', path: '/components/typography' }"
          ></sidebar-item>
        </sidebar-item>
        <sidebar-item :link="{ name: 'Forms', icon: 'content_paste' }">
          <sidebar-item
            :link="{ name: 'Regular Forms', path: '/forms/regular' }"
          ></sidebar-item>
          <sidebar-item
            :link="{ name: 'Extended Forms', path: '/forms/extended' }"
          ></sidebar-item>
          <sidebar-item
            :link="{ name: 'Validation Forms', path: '/forms/validation' }"
          ></sidebar-item>
          <sidebar-item
            :link="{ name: 'Wizard', path: '/forms/wizard' }"
          ></sidebar-item>
        </sidebar-item>
        <sidebar-item :link="{ name: 'Tables', icon: 'grid_on' }">
          <sidebar-item
            :link="{ name: 'Regular Tables', path: '/table-list/regular' }"
          ></sidebar-item>
          <sidebar-item
            :link="{ name: 'Extended Tables', path: '/table-list/extended' }"
          ></sidebar-item>
          <sidebar-item
            :link="{ name: 'Paginated Tables', path: '/table-list/paginated' }"
          ></sidebar-item>
        </sidebar-item>
        <sidebar-item :link="{ name: 'Maps', icon: 'place' }">
          <sidebar-item
            :link="{ name: 'Google Maps', path: '/maps/google' }"
          ></sidebar-item>
        </sidebar-item>
        <sidebar-item
          :link="{ name: 'Widgets', icon: 'widgets', path: '/widgets' }"
        ></sidebar-item>
        <sidebar-item
          :link="{ name: 'Charts', icon: 'timeline', path: '/charts' }"
        ></sidebar-item>
      </template>
    </side-bar>
    <div class="main-panel">
      <top-navbar></top-navbar>
      <div
        :class="{ content: !$route.meta.hideContent }"
        @click="toggleSidebar"
      >
        <zoom-center-transition :duration="200" mode="out-in">
          <!-- your content here -->
          <router-view></router-view>
        </zoom-center-transition>
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
import MobileMenu from "./Extra/MobileMenu.vue";
// import FixedPlugin from "../../FixedPlugin.vue";
import UserMenu from "./Extra/UserMenu.vue";
import { ZoomCenterTransition } from "vue2-transitions";

export default {
  components: {
    TopNavbar,
    ContentFooter,
    MobileMenu,
    // FixedPlugin,
    UserMenu,
    ZoomCenterTransition
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
