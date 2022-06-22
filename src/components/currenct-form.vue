<template>
  <v-form v-model="valid" class="form-width">
    <v-container>
      <v-row>
        <v-col
          cols="12"

        >
          <v-text-field
            v-model="payloadAlert.email"
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
          v-model="payloadAlert.currency"
          label="Currency"
          dense
        ></v-select>
        </v-col>
          <v-col
          cols="6"

        >
            <v-select
          v-model="payloadAlert.threshold"
          :items="['<','>']"
          label="Threshold"
          dense
        ></v-select>
        </v-col>

        <v-col
          cols="12"

        >
          <v-text-field
            v-model="payloadAlert.money"
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
                @click="$emit('createdNotification',payloadAlert)"
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
<script>

  export default {
    data: () => ({
      valid: false,
      payloadAlert: {
        money:'',
        currency:'',
        threshold:'',
        email: '',
      },
      nameRules: [
        (        v) => !!v || 'Name is required',
        (        v) => v.length <= 10 || 'Name must be less than 10 characters',
      ],
      emailRules: [
        (        v) => !!v || 'E-mail is required',
        (        v) => /.+@.+/.test(v) || 'E-mail must be valid',
      ],
    }),
    methods: {
        reset() {
          this.payloadAlert.email = ''
          this.payloadAlert.money = ''
          this.payloadAlert.currency = ''
          this.payloadAlert.threshold = ''
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
