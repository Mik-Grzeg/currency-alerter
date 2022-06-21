<template>
  <div class="notification">
    <v-alert
        dense
        text
        :type="typeOfAlert"
    >{{alertValue}}</v-alert>
  </div>
</template>

<script>
export default {
  props:['typeOfAlert','alertValue'],
  name: "currency-notify"
}
</script>

<style scoped>
.notification {
  position: absolute;
  bottom: 0px;
  right: 20px;
  width: 400px;
  height: 133px;
}
</style>
