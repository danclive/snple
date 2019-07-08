<template>
  <div class="md-layout md-alignment-top-center">
    <div class="md-layout-item md-medium-size-100 md-size-50">
      <add-form header-color="green" @on-submit="onSubmit"></add-form>
    </div>
  </div>
</template>

<script>
import AddForm from "./_AddForm.vue";
import { deviceAdd, userDeviceAdd } from "@/api/device";
export default {
  components: {
    AddForm
  },
  data() {
    return {
      userId: ""
    };
  },
  created() {
    const id = this.$route.params.id;
    if (id) {
      this.userId = id;
    }
  },
  methods: {
    onSubmit(device) {
      if (this.userId !== "") {
        userDeviceAdd(this.userId, device).then(res => {
          this.$router.go(-1);
        });
      } else {
        deviceAdd(device).then(res => {
          this.$router.go(-1);
        });
      }
    }
  }
};
</script>
