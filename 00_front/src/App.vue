<template>
    <router-view></router-view>
</template>

<script>
export default {
  name: 'App',
  // 起動時にDBの値をストア
  created() {
    this.$store.dispatch('fetchDB')
  }
}
</script>