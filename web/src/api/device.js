import request from "./util.js";

export function deviceList(params) {
  return request({
    url: "/v1/device/",
    method: "get",
    params
  });
}

export function deviceDetail(id) {
  return request({
    url: "/v1/device/" + id,
    method: "get"
  });
}

export function deviceAdd(data) {
  return request({
    url: "/v1/device/",
    method: "post",
    data
  });
}

export function deviceUpdate(id, data) {
  return request({
    url: "/v1/device/" + id,
    method: "patch",
    data
  });
}

export function deviceDelete(id) {
  return request({
    url: "/v1/device/" + id,
    method: "delete"
  });
}

export function genid() {
  return request({
    url: "/v1/util/genid",
    method: "get"
  });
}

export function userDeviceList(id, params) {
  return request({
    url: "/v1/user/" + id + "/device",
    method: "get",
    params
  });
}

export function userDeviceAdd(id, data) {
  return request({
    url: "/v1/user/" + id + "/device",
    method: "post",
    data
  });
}
