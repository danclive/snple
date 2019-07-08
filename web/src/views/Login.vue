<template>
  <div class="md-layout text-center">
    <div
      class="md-layout-item md-size-33 md-medium-size-50 md-small-size-70 md-xsmall-size-100"
    >
      <login-card header-color="green">
        <h4 slot="title" class="title">登录</h4>
        <!-- <md-button
          slot="buttons"
          href="#facebook"
          class="md-just-icon md-simple md-white"
        >
          <i class="fab fa-facebook-square"></i>
        </md-button>
        <md-button
          slot="buttons"
          href="#twitter"
          class="md-just-icon md-simple md-white"
        >
          <i class="fab fa-twitter"></i>
        </md-button>
        <md-button
          slot="buttons"
          href="#google"
          class="md-just-icon md-simple md-white"
        >
          <i class="fab fa-google-plus-g"></i>
        </md-button> -->
        <!-- <p slot="description" class="description">Or Be Classical</p> -->
        <md-field class="md-form-group" slot="inputs">
          <md-icon>face</md-icon>
          <label>用户名...</label>
          <md-input v-model="username"></md-input>
        </md-field>
        <!-- <md-field class="md-form-group" slot="inputs">
          <md-icon>email</md-icon>
          <label>Email...</label>
          <md-input v-model="email" type="email"></md-input>
        </md-field> -->
        <md-field class="md-form-group" slot="inputs">
          <md-icon>lock_outline</md-icon>
          <label>密码...</label>
          <md-input v-model="password" type="password"></md-input>
        </md-field>
        <md-button slot="footer" class="md-simple md-success md-lg" @click="login">
          Lets Go
        </md-button>
      </login-card>
    </div>
  </div>
</template>
<script>
import { LoginCard } from "@/components";
export default {
  components: {
    LoginCard
  },
  data() {
    return {
      username: "admin",
      password: "public"
    };
  },
  methods: {
    login() {
      if (this.username !== "" && this.password !== "") {
        this.$store.dispatch("Login", { username: this.username, password: this.password }).then(() => {
          this.$router.push({ name: "Home" });
        }).catch(err => {
          if (err.response) {
            if (err.response.status === 401) {
              this.$notify({
                type: "warning",
                icon: "add_alert",
                message: err.response.data.message.info,
                horizontalAlign: "center",
                verticalAlign: "top"
              });
            }
          }
        });
      }
    }
  }
};
</script>

<style></style>
