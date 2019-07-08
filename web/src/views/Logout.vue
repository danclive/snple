<template>
  <div class="md-layout text-center">
    <div class="md-layout-item md-size-50 md-small-size-100">
      <md-progress-bar md-mode="indeterminate"></md-progress-bar>
      <!-- <md-progress-bar class="md-accent" md-mode="indeterminate"></md-progress-bar> -->
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      username: "admin",
      password: "public"
    };
  },
  created() {
    setTimeout(() => {
      this.$store.dispatch("Logout").then(() => {
        this.$router.push({ name: "登录" });
      });
    }, 1000);
  }
};
</script>

<style></style>
