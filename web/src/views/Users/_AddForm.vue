<template>
  <form>
    <md-card>
      <md-card-header class="md-card-header-icon md-card-header-green">
        <div class="card-icon">
          <md-icon>perm_identity</md-icon>
        </div>
        <h4 class="title">用户资料</h4>
      </md-card-header>

      <md-card-content>
        <md-field
          :class="[
            { 'md-error': errors.has('username') },
            { 'md-valid': !errors.has('username') && touched.username }
          ]"
        >
          <label>用户名</label>
          <md-input
            v-model="username"
            data-vv-name="username"
            type="text"
            required
            v-validate="modelValidations.username"
          >
          </md-input>
          <slide-y-down-transition>
            <md-icon class="error" v-show="errors.has('username')">close</md-icon>
          </slide-y-down-transition>
          <slide-y-down-transition>
            <md-icon
              class="success"
              v-show="!errors.has('username') && touched.username"
              >done</md-icon
            >
          </slide-y-down-transition>
        </md-field>
        <md-field
          :class="[
            { 'md-valid': !errors.has('password') && touched.password },
            { 'md-error': errors.has('password') }
          ]"
        >
          <label>密码</label>
          <md-input
            v-model="password"
            data-vv-name="password"
            type="password"
            ref="password"
            required
            v-validate="modelValidations.password"
          >
          </md-input>
          <slide-y-down-transition>
            <md-icon
class="error"
v-show="errors.has('password')"
              >close</md-icon
            >
          </slide-y-down-transition>
          <slide-y-down-transition>
            <md-icon
              class="success"
              v-show="!errors.has('password') && touched.password"
              >done</md-icon
            >
          </slide-y-down-transition>
        </md-field>
        <md-field
          :class="[
            {
              'md-valid': !errors.has('confirmPassword') && touched.confirmPass
            },
            { 'md-error': errors.has('confirmPassword') }
          ]"
        >
          <label>确认密码</label>
          <md-input
            v-model="confirmPassword"
            data-vv-name="confirmPassword"
            data-vv-as="password"
            type="password"
            required
            v-validate="modelValidations.confirmPassword"
          >
          </md-input>
          <slide-y-down-transition>
            <md-icon
class="error"
v-show="errors.has('confirmPassword')"
              >close</md-icon
            >
          </slide-y-down-transition>
          <slide-y-down-transition>
            <md-icon
              class="success"
              v-show="!errors.has('confirmPassword') && touched.confirmPassword"
              >done</md-icon
            >
          </slide-y-down-transition>
        </md-field>
        <md-field style="margin-top: 0">
          <label>备注</label>
          <md-textarea v-model="desc" maxlength="500"></md-textarea>
        </md-field>
        <md-checkbox v-model="issuper">管理员</md-checkbox>
        <div class="form-category">* 必要字段</div>
      </md-card-content>

      <md-card-actions>
        <md-button
          native-type="submit"
          @click.native.prevent="validate"
          class="md-success"
          >确定</md-button
        >
      </md-card-actions>
    </md-card>
  </form>
</template>
<script>
import { SlideYDownTransition } from "vue2-transitions";
export default {
  components: {
    SlideYDownTransition
  },
  data() {
    return {
      username: "",
      password: "",
      confirmPassword: "",
      desc: "",
      issuper: false,
      touched: {
        username: false,
        password: false,
        confirmPass: false
      },
      modelValidations: {
        username: {
          required: true
        },
        password: {
          required: true,
          min: 5
        },
        confirmPassword: {
          required: true,
          confirmed: "password"
        }
      }
    };
  },
  methods: {
    validate() {
      this.$validator.validateAll().then(isValid => {
        if (isValid) {
          this.$emit("on-submit", {
            name: this.username,
            pass: this.password,
            desc: this.desc,
            super: this.issuper
          });
        }
      });
    }
  },
  watch: {
    username() {
      this.touched.username = true;
    },
    password() {
      this.touched.password = true;
    },
    confirmPassword() {
      this.touched.confirmPass = true;
    }
  }
};
</script>
<style lang="scss" scoped>
.md-card .md-card-actions {
  border: none;
}
</style>
