import axios from "axios";
import LocalStore from "store";
import router from "@/router";

const API_URL = window.location.protocol + "//" + window.location.hostname + ":" + "8765/";

const instance = axios.create({
  baseURL: API_URL,
  timeout: 15000
});

instance.interceptors.request.use(function(config) {
  // Do something before request is sent
  if (LocalStore.get("token")) {
    config.headers["Authorization"] = "Bearer " + LocalStore.get("token");
  }

  // Spin.show()

  return config;
}, function(error) {
  // Do something with request error
  return Promise.reject(error);
});

instance.interceptors.response.use(function(response) {
  // Do something with response data
  // if (response.status !== 200) {
  //   if (response.status === 401) {
  //     router.push({ name: "登录" });
  //   }

  //   console.log("error");
  // }

  // 404
  if (response.data.message && response.data.message.code === 404) {
    return Promise.reject(response.data.message);
  }

  if (!response.data.success) {
    // console.log(response.data.message.info);
    // window.$notify({
    //   type: "warn",
    //   title: response.data.message.info
    // });
    // console.log(Vue.$notify);
    window.$notify({
      type: "warning",
      icon: "add_alert",
      message: response.data.message.info,
      horizontalAlign: "center",
      verticalAlign: "top"
    });

    return Promise.reject(response.data.message);
  }

  return response.data;
}, function(error) {
  // Do something with response error
  if (error.response) {
    if (error.response.status === 401) {
      router.push({ name: "登录" });
    } else if (error.response.status === 400) {
      if (error.response.data.message && error.response.data.message.code === 500) {
        window.$notify({
          type: "danger",
          icon: "add_alert",
          message: error.response.data.message.info,
          horizontalAlign: "center",
          verticalAlign: "top"
        });

        return Promise.reject(error.response.data.message);
      }
    }
  }

  return Promise.reject(error);
});

export default instance;
