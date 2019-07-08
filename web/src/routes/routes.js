// import DashboardLayout from "@/pages/Dashboard/Layout/DashboardLayout.vue";
// import AuthLayout from "@/pages/Dashboard/Pages/AuthLayout.vue";

import DashboardLayout from "@/views/Layout/DashboardLayout.vue";
import AuthLayout from "@/views/Layout/AuthLayout.vue";
// Dashboard pages
import Dashboard from "@/views/Dashboard/Dashboard.vue";
// User pages
import Users from "@/views/Users/Users.vue";
import AddUser from "@/views/Users/Add.vue";
import DetailUser from "@/views/Users/Detail.vue";
// Device pages
import Devices from "@/views/Devices/Devices.vue";
import AddDevice from "@/views/Devices/Add.vue";
import DetailDevice from "@/views/Devices/Detail.vue";
// Login page
import Login from "@/views/Login.vue";
import Logout from "@/views/Logout.vue";

// 404
import Page404 from "@/views/404.vue";

// Pages
import User from "@/pages/Dashboard/Pages/UserProfile.vue";
import Pricing from "@/pages/Dashboard/Pages/Pricing.vue";
import TimeLine from "@/pages/Dashboard/Pages/TimeLinePage.vue";
// import RtlSupport from "@/pages/Dashboard/Pages/RtlSupport.vue";
// import Login from "@/views/Login.vue";
import Register from "@/pages/Dashboard/Pages/Register.vue";
import Lock from "@/pages/Dashboard/Pages/Lock.vue";

// Components pages
import Buttons from "@/pages/Dashboard/Components/Buttons.vue";
import GridSystem from "@/pages/Dashboard/Components/GridSystem.vue";
import Panels from "@/pages/Dashboard/Components/Panels.vue";
// import SweetAlert from "@/pages/Dashboard/Components/SweetAlert.vue";
import Notifications from "@/pages/Dashboard/Components/Notifications.vue";
import Icons from "@/pages/Dashboard/Components/Icons.vue";
import Typography from "@/pages/Dashboard/Components/Typography.vue";

// Forms pages
import RegularForms from "@/pages/Dashboard/Forms/RegularForms.vue";
import ExtendedForms from "@/pages/Dashboard/Forms/ExtendedForms.vue";
import ValidationForms from "@/pages/Dashboard/Forms/ValidationForms.vue";
import Wizard from "@/pages/Dashboard/Forms/Wizard.vue";

// TableList pages
import RegularTables from "@/pages/Dashboard/Tables/RegularTables.vue";
import ExtendedTables from "@/pages/Dashboard/Tables/ExtendedTables.vue";
import PaginatedTables from "@/pages/Dashboard/Tables/PaginatedTables.vue";

// Maps pages
import GoogleMaps from "@/pages/Dashboard/Maps/GoogleMaps.vue";
import FullScreenMap from "@/pages/Dashboard/Maps/FullScreenMap.vue";
// import VectorMaps from "@/pages/Dashboard/Maps/VectorMaps.vue";

// Calendar
// import Calendar from "@/pages/Dashboard/Calendar.vue";
// Charts
import Charts from "@/pages/Dashboard/Charts.vue";
import Widgets from "@/pages/Dashboard/Widgets.vue";

const UsersMenu = {
  path: "/users",
  component: DashboardLayout,
  redirect: "/users/",
  name: "用户管理",
  children: [
    {
      path: "/",
      name: "用户列表",
      components: { default: Users }
    },
    {
      path: "add",
      name: "添加用户",
      components: { default: AddUser }
    },
    {
      path: ":id([a-z0-9]{24})",
      name: "用户详情",
      components: { default: DetailUser }
    },
    {
      path: ":id([a-z0-9]{24})/add_device",
      name: "添加设备(给用户)",
      components: { default: AddDevice }
    }
  ]
};

const DeviceMenu = {
  path: "/devices",
  component: DashboardLayout,
  redirect: "/devices/",
  name: "设备管理",
  children: [
    {
      path: "/",
      name: "设备列表",
      components: { default: Devices }
    },
    {
      path: "add",
      name: "添加设备",
      components: { default: AddDevice }
    },
    {
      path: ":id([a-z0-9]{24})",
      name: "设备详情",
      components: { default: DetailDevice }
    }
  ]
};

const componentsMenu = {
  path: "/components",
  component: DashboardLayout,
  redirect: "/components/buttons",
  name: "Components",
  children: [
    {
      path: "buttons",
      name: "Buttons",
      components: { default: Buttons }
    },
    {
      path: "grid-system",
      name: "Grid System",
      components: { default: GridSystem }
    },
    {
      path: "panels",
      name: "Panels",
      components: { default: Panels }
    },
    // {
    //   path: "sweet-alert",
    //   name: "Sweet Alert",
    //   components: { default: SweetAlert }
    // },
    {
      path: "notifications",
      name: "Notifications",
      components: { default: Notifications }
    },
    {
      path: "icons",
      name: "Icons",
      components: { default: Icons }
    },
    {
      path: "typography",
      name: "Typography",
      components: { default: Typography }
    }
  ]
};
const formsMenu = {
  path: "/forms",
  component: DashboardLayout,
  redirect: "/forms/regular",
  name: "Forms",
  children: [
    {
      path: "regular",
      name: "Regular Forms",
      components: { default: RegularForms }
    },
    {
      path: "extended",
      name: "Extended Forms",
      components: { default: ExtendedForms }
    },
    {
      path: "validation",
      name: "Validation Forms",
      components: { default: ValidationForms }
    },
    {
      path: "wizard",
      name: "Wizard",
      components: { default: Wizard }
    }
  ]
};

const tablesMenu = {
  path: "/table-list",
  component: DashboardLayout,
  redirect: "/table-list/regular",
  name: "Tables",
  children: [
    {
      path: "regular",
      name: "Regular Tables",
      components: { default: RegularTables }
    },
    {
      path: "extended",
      name: "Extended Tables",
      components: { default: ExtendedTables }
    },
    {
      path: "paginated",
      name: "Pagianted Tables",
      components: { default: PaginatedTables }
    }
  ]
};

const mapsMenu = {
  path: "/maps",
  component: DashboardLayout,
  name: "Maps",
  redirect: "/maps/google",
  children: [
    {
      path: "google",
      name: "Google Maps",
      components: { default: GoogleMaps }
    },
    {
      path: "full-screen",
      name: "Full Screen Map",
      meta: {
        hideContent: true,
        hideFooter: true,
        navbarAbsolute: true
      },
      components: { default: FullScreenMap }
    }
    // {
    //   path: "vector-map",
    //   name: "Vector Map",
    //   components: { default: VectorMaps }
    // }
  ]
};

const pagesMenu = {
  path: "/pages",
  component: DashboardLayout,
  name: "Pages",
  redirect: "/pages/user",
  children: [
    {
      path: "user",
      name: "User Page",
      components: { default: User }
    },
    {
      path: "timeline",
      name: "Timeline Page",
      components: { default: TimeLine }
    }
  ]
};

const authPages = {
  path: "/",
  component: AuthLayout,
  name: "Authentication",
  children: [
    {
      path: "/login",
      name: "登录",
      component: Login
    },
    {
      path: "/logout",
      name: "注销",
      component: Logout
    },
    {
      path: "/register",
      name: "Register",
      component: Register
    },
    {
      path: "/pricing",
      name: "Pricing",
      component: Pricing
    },
    {
      path: "/lock",
      name: "Lock",
      component: Lock
    }
  ]
};

const routes = [
  {
    path: "/",
    redirect: "/dashboard",
    name: "Home"
  },
  UsersMenu,
  DeviceMenu,
  // componentsMenu,
  // formsMenu,
  // tablesMenu,
  // mapsMenu,
  // pagesMenu,
  authPages,
  {
    path: "/",
    component: DashboardLayout,
    children: [
      {
        path: "dashboard",
        name: "控制台",
        components: { default: Dashboard },
        meta: {
          requiresAuth: true
        }
      },
      // {
      //   path: "charts",
      //   name: "Charts",
      //   components: { default: Charts }
      // },
      // {
      //   path: "widgets",
      //   name: "Widgets",
      //   components: { default: Widgets }
      // },
      {
        path: "*",
        name: "404",
        components: { default: Page404 }
      }
    ]
  }
];

export default routes;
