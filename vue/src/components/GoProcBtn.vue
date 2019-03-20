<template>
    <v-btn :color="color" v-on:click="buttonClick" class="go_proc_btn">
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
            goProcPayload: String,
        }
    })
    export default class GoProcBtn extends Vue {
        public buttonClick() {
            let sendData = {
               name: this.$props.goProcName,
               payload: this.$props.goProcPayload,
            };
            astilectron.sendMessage(sendData, (response: any) => {
                console.log(response);
            })
        }
    }
</script>

<style scoped>
.go_proc_btn {
    background: url(../assets/gopher_head.png) center no-repeat;
}
</style>