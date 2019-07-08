import request from "./util.js";

export function login(data) {
  return request({
    url: "/v1/login",
    method: "post",
    data
  });
}

export function mine() {
  return request({
    url: "/v1/mine/",
    method: "get"
  });
}

export function userList(params) {
  return request({
    url: "/v1/user/",
    method: "get",
    params
  });
}

export function userDetail(id) {
  return request({
    url: "/v1/user/" + id,
    method: "get"
  });
}

export function userAdd(data) {
  return request({
    url: "/v1/user/",
    method: "post",
    data
  });
}

export function userUpdate(id, data) {
  return request({
    url: "/v1/user/" + id,
    method: "patch",
    data
  });
}

export function userDelete(id) {
  return request({
    url: "/v1/user/" + id,
    method: "delete"
  });
}
