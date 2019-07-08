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
        <div class="md-layout">
          <div class="md-layout-item md-medium-size-100 md-size-33">
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
                v-validate="modelValidations.required"
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
          </div>
          <div class="md-layout-item md-medium-size-100 md-size-33">
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
          </div>
          <div class="md-layout-item md-medium-size-100 md-size-33">
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
          </div>

          <div class="md-layout-item md-medium-size-100 md-size-100">
            <md-field style="margin-top: 0">
              <label>备注</label>
              <md-textarea v-model="desc" maxlength="500"></md-textarea>
            </md-field>
            <md-checkbox v-model="issuper">管理员</md-checkbox>
            <div class="form-category">* 必要字段</div>
          </div>
        </div>
      </md-card-content>

      <md-card-actions>
        <md-button
          native-type="submit"
          @click.native.prevent="validate"
          class="md-success"
          >更新</md-button
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
  props: {
    user: Object
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
        required: {
          required: true
        },
        password: {
          required: false,
          min: 5
        },
        confirmPassword: {
          required: false,
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
    },
    "user.name"() {
      this.username = this.user.name;
    },
    "user.desc"() {
      this.desc = this.user.desc;
    },
    "user.super"() {
      this.issuper = this.user.super;
    }
  }
};
</script>
<style lang="scss" scoped>
.md-card .md-card-actions {
  border: none;
}
</style>
