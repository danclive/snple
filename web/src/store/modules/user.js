import LocalStore from "store";
import { login, mine } from "@/api/user";

const user = {
  state: {
    token: LocalStore.get("token") || "",
    user: {}
  },
  mutations: {
    SET_TOKEN(state, token) {
      state.token = token;
      LocalStore.set("token", token);
    },
    SET_USER(state, user) {
      state.user = user;
    },
    DESTORY(state) {
      LocalStore.remove("token");
      state.token = "";
      state.user = {};
    }
  },
  actions: {
    Login({ commit }, { username, password }) {
      return new Promise((resolve, reject) => {
        login({ username, password }).then(res => {
          commit("SET_TOKEN", res.data.token);
          resolve();
        }).catch(err => {
          reject(err);
        });
      });
    },

    Logout({ commit }) {
      commit("DESTORY");
    },

    Mine({ commit }) {
      return new Promise((resolve, reject) => {
        mine().then(res => {
          commit("SET_USER", res.data);
          resolve();
        }).catch(err => {
          reject(err);
        });
      });
    }
  }
};

export default user;
