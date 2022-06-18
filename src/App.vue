<template>
  <v-app>
    <v-main>
     <v-parallax src="https://cdn.vuetifyjs.com/images/parallax/material.jpg"></v-parallax>

      <router-view/>

    </v-main>
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'App',

  data () {
    return {
      //
    }
  },
})
</script>
<style lang="scss">
body{
  overflow: none;
}
.v-parallax {
  height: 100vh;
}
</style>