<template>
  <div id="home">
    <h1 class="title">Currency App</h1>
    <Transition>
      <currency-notify v-if="isRun" :typeOfAlert="this.typeOfAlert" :alertValue="alertValue"/>
    </Transition>
    <CurrenctForm @createdNotification="createdNotification" />
  </div>
</template>

<script>
import CurrenctForm from '../components/currenct-form.vue';
import CurrencyNotify from "@/components/currency-notify.vue";
import apiService from "@/service/apiService";



export default {
  data: () => ({
    isRun: false,
    typeOfAlert:"success",
    alertValue:"Alert pomyślnie ustawiono"
  }),
  name: 'HomeView',
  components: {
    CurrencyNotify,
    CurrenctForm
},
  methods: {


    createdNotification(payloadAlert) {
      payloadAlert.money = parseFloat(payloadAlert.money)
      console.log(payloadAlert)
      apiService.postAlert(payloadAlert)
      .then(()=>{
        this.typeOfAlert = "success"
        this.alertValue = "Alert pomyślnie ustawiono"
      })
      .catch(()=> {
        this.typeOfAlert = "error"
        this.alertValue = "Błąd w przesyłaniu danych"
      })
      this.isRun = true
      setTimeout( () =>this.isRun = false,3000)
    }
  }
};
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
.v-enter-active,
.v-leave-active {
  transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>
