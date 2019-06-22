<template>
    <v-container fluid grid-list-md>
        <v-layout row wrap>
            <v-flex xs2>
                トークン
            </v-flex>
            <v-flex xs10>
                <v-text-field
                    v-model="token"
                ></v-text-field>
            </v-flex>
            <v-flex xs12>
                <v-btn
                    small
                    color="success"
                    @click="save"
                >保存</v-btn>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script lang="ts">
    import { Component, Vue } from 'vue-property-decorator';
  @Component
  export default class Settings extends Vue {
      private token: string = '';
      private mounted() {
          this.init();
      }
      private init() {
          this.token = localStorage.getItem('token') as string;
      }
      private save() {
          localStorage.setItem('token', this.token);
      }
  }
</script>