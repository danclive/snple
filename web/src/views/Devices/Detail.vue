<template>
  <div>
    <div class="md-layout md-alignment-top-center">
      <div class="md-layout-item md-medium-size-100 md-size-100">
        <detail-form header-color="green" :device="device" @on-submit="onSubmit"></detail-form>
      </div>
    </div>
  </div>
</template>
<script>
import DetailForm from "./_DetailForm.vue";
import { deviceDetail, deviceUpdate } from "@/api/device";
export default {
  components: {
    DetailForm
  },
  data() {
    return {
      device: {
        id: "",
        device_id: "",
        desc: "",
        time: ""
      }
    };
  },
  created() {
    this.fecthDate();
  },
  methods: {
    onSubmit(device) {
      deviceUpdate(this.device.id, device).then(res => {
        this.$notify({
          type: "success",
          icon: "add_alert",
          message: "更新成功",
          horizontalAlign: "center",
          verticalAlign: "top"
        });
      });
    },
    fecthDate() {
      const id = this.$route.params.id;

      deviceDetail(id).then(res => {
        this.device = Object.assign(this.device, res.data);
      }).catch(err => {
        if (err.code === 404) {
          this.$router.replace({ name: "404" });
        }
      });
    }
  }
};
</script>
