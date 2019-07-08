import request from "./util.js";

export function listStorage(params) {
  return request({
    url: "/storage",
    method: "get",
    params
  });
}
