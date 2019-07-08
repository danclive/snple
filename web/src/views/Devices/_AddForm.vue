<template>
  <form>
    <md-card>
      <md-card-header class="md-card-header-icon md-card-header-green">
        <div class="card-icon">
          <md-icon>perm_identity</md-icon>
        </div>
        <h4 class="title">设备资料</h4>
      </md-card-header>

      <md-card-content>
        <md-field
          :class="[
            { 'md-error': errors.has('deviceId') },
            { 'md-valid': !errors.has('deviceId') && touched.deviceId }
          ]"
        >
          <label>设备ID</label>
          <md-input
            v-model="deviceId"
            data-vv-name="deviceId"
            type="text"
            required
            v-validate="modelValidations.deviceId"
          >
          </md-input>
          <slide-y-down-transition>
            <md-icon class="error" v-show="errors.has('deviceId')">close</md-icon>
          </slide-y-down-transition>
          <slide-y-down-transition>
            <md-icon
              class="success"
              v-show="!errors.has('deviceId') && touched.deviceId"
              >done</md-icon
            >
          </slide-y-down-transition>
        </md-field>
        <md-field style="margin-top: 0">
          <label>备注</label>
          <md-textarea v-model="desc" maxlength="500"></md-textarea>
        </md-field>
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
import { genid } from "@/api/device";
export default {
  components: {
    SlideYDownTransition
  },
  data() {
    return {
      deviceId: "",
      desc: "",
      touched: {
        deviceId: false,
        confirmPass: false
      },
      modelValidations: {
        deviceId: {
          required: true
        }
      }
    };
  },
  created() {
    genid().then(res => {
      this.deviceId = res.data.id;
    });
  },
  methods: {
    validate() {
      this.$validator.validateAll().then(isValid => {
        if (isValid) {
          this.$emit("on-submit", {
            device_id: this.deviceId,
            desc: this.desc
          });
        }
      });
    }
  },
  watch: {
    deviceId() {
      this.touched.deviceId = true;
    }
  }
};
</script>
<style lang="scss" scoped>
.md-card .md-card-actions {
  border: none;
}
</style>
