<template>
  <div>
    <div class="md-layout md-alignment-top-center">
      <div class="md-layout-item md-medium-size-100 md-size-100">
        <detail-form header-color="green" :user="user" @on-submit="onSubmit"></detail-form>
      </div>
    </div>
    <device-list inUserDetail :userId="user.id"></device-list>
  </div>
</template>

<script>
import DetailForm from "./_DetailForm.vue";
import DeviceList from "@/views/Devices/_DeviceList.vue";
import { userDetail, userUpdate } from "@/api/user";
export default {
  components: {
    DetailForm,
    DeviceList
  },
  data() {
    return {
      user: {
        id: "",
        name: "",
        pass: "",
        desc: "",
        super: false,
        time: ""
      }
    };
  },
  created() {
    this.fecthDate();
  },
  methods: {
    onSubmit(user) {
      const user2 = Object.assign({}, user);
      user2.super = user.super ? "true" : "false";

      userUpdate(this.user.id, user2).then(res => {
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

      userDetail(id).then(res => {
        this.user = Object.assign(this.user, res.data);
      }).catch(err => {
        if (err.code === 404) {
          this.$router.replace({ name: "404" });
        }
      });
    }
  }
};
</script>
