<template>
    <div>
        <img src="../assets/shika.png" :height="cpu" :width="cpu">
    </div>
</template>

<script lang="ts">
    declare var astilectron: any;
    import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
    @Component
    export default class Mascot extends Vue {
        public cpu: Number = 100;
        public wcpu: Number = 100;
        @Watch('wcpu')
        public changeUsingCpu(newValue: number, oldValue: number): void {
            console.log(newValue);
            console.log(oldValue);
            let diff: number = newValue - oldValue
            if (0 < diff) {
                for (let i = oldValue; i < newValue; i++) {
                    this.animation(i);
                }
            } else {
                for (let i = newValue; i < oldValue; i++) {
                    this.animation(i);
                }
            }
        }
        public mounted() {
            document.addEventListener('astilectron-ready', this.resize);
        }
        public resize() {
            astilectron.onMessage((data: any) => {
                this.wcpu = data.payload;
            });
        }
        public animation(wcpu: number) {
            this.cpu = wcpu;
        }
    } 
</script>