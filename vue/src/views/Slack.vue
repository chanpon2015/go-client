<template>
    <v-container fluid grid-list-md>
        <v-layout row wrap>
            <v-flex xs12>
                <v-textarea
                    hint="text"
                    v-model="message"
                ></v-textarea>
            </v-flex>
            <v-flex xs12>
                <v-btn
                    small
                    color="success"
                    @click="sendMessage"
                >送信</v-btn>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script lang="ts">
  import axios from 'axios';
  import { Component, Vue } from 'vue-property-decorator';
  @Component
  export default class Slack extends Vue {
    private message: string = '';
    private sendMessage() {
        const data: object = {
            token: localStorage.getItem('token') as string,
            channel: 'C5UDXLPT7',
            text: this.message,
        };
        axios.post(
            'https://slack.com/api/chat.postMessage',
            data,
        ).then(res => {
            alert('res');
        });
    }
  }
</script>