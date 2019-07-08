<template>
  <div class="md-layout md-alignment-top-center">
    <div class="md-layout-item md-medium-size-100 md-size-50">
      <add-form header-color="green" @on-submit="onSubmit"></add-form>
    </div>
  </div>
</template>

<script>
import AddForm from "./_AddForm.vue";
import { userAdd } from "@/api/user";
export default {
  components: {
    AddForm
  },
  methods: {
    onSubmit(user) {
      userAdd(user).then(res => {
        // console.log(res);
        this.$router.push({ name: "用户列表" });
      });
    }
  }
};
</script>
