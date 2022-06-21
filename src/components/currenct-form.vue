<template>
  <v-form v-model="valid" class="form-width">
    <v-container>
      <v-row>
        <v-col
          cols="12"

        >
          <v-text-field
            v-model="alert.email"
            :rules="emailRules"
            label="E-mail"
            filled
            required
          ></v-text-field>
        </v-col>

        <v-col
          cols="6"

        >
            <v-select
            filled
          :items="['EUR','USD','CHF']"
          v-model="alert.currency"
          label="Currency"
          dense
        ></v-select>
        </v-col>
          <v-col
          cols="6"

        >
            <v-select
          v-model="alert.threshold"
          :items="['<','>']"
          label="Threshold"
          dense
        ></v-select>
        </v-col>

        <v-col
          cols="12"

        >
          <v-text-field
            v-model="alert.money"
            label="Cash limit"
            filled
            required
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row>
          <v-col>
            <v-btn
                class="mr-4 buttons-width"
                color="success"
                @click="$emit('createdNotification',alert)"
                >
                submit
             </v-btn>
          </v-col>
          <v-col>
              <v-btn color="info" @click="reset" class="buttons-width">
      clear
    </v-btn>
          </v-col>
      </v-row>
    </v-container>
  </v-form>
</template>
<script lang="ts">
import apiService from "@/service/apiService"
type alert = {
  money: string,
  currency: string,
  threshold: string,
  email: string,
}
  export default {
    data: () => ({
      valid: false,
      alert: {
        money:'',
        currency:'',
        threshold:'',
        email: '',
      } as alert,
      nameRules: [
        (        v: string) => !!v || 'Name is required',
        (        v: string) => v.length <= 10 || 'Name must be less than 10 characters',
      ],
      emailRules: [
        (        v: string) => !!v || 'E-mail is required',
        (        v: string) => /.+@.+/.test(v) || 'E-mail must be valid',
      ],
    }),
    methods: {
      alert: {
        money:'',
        currency:'',
        threshold:'',
        email: '',
      } as alert,

        reset(): void {
            this.alert.email = ''
            this.alert.money = ''
            this.alert.currency = ''
            this.alert.threshold = ''
        },
    }
  }
</script>
<style lang="scss">
    .buttons-width {
        width: 100%;
    }
    .form-width {
        width: 50%;

        position: absolute;
        bottom: 25%;
        top: 25%;
        left: 25%;
        right: 25%;
    }

</style>
