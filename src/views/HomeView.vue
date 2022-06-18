<template>
  <div id="home">
    <h1 class="title">Currency App</h1>
    <CurrenctForm />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import CurrenctForm from '../components/currenct-form.vue';

// Components

export default defineComponent({
  name: 'HomeView',

  components: {
    CurrenctForm
},
});
</script>
<style>
#home {
  max-height: 100vh;
}
.title {
    position: absolute;
    top:10%;
    left: 25%;
    right: 25%;
  text-align: center;
}
</style>
