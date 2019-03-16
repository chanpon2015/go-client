<template>
    <v-btn :color="color" v-on:click="buttonClick">
        <img src="../assets/gopher_head.png">
        {{ value }}
    </v-btn>
</template>

<script lang="ts">
    declare var astilectron: any;
    import { Component, Prop, Vue } from 'vue-property-decorator';
    @Component({
        props: {
            color: String,
            value: String,
            goProcName: String,
            goProcPaylod: String,
        }
    })
    export default class GoProcBtn extends Vue {
        public buttonClick() {
            let sendData = {
               name: this.$props.goProcName,
               payload: this.$props.goProcPaylod,
            };
            astilectron.sendMessage(sendData, (response: any) => {
                alert(response.payload);
            })
        }
    }
</script>