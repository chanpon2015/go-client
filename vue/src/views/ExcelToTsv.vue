<template>
    <div>
        <DragOnDropFile
            :files="files"
            text="ここにファイルをドロップしてください"
        ></DragOnDropFile>
    </div>
</template>

<script lang="ts">
    import DragOnDropFile from '@/components/DragOnDropFile.vue';
    import { Component, Vue } from 'vue-property-decorator';
    @Component({
        components: {
            DragOnDropFile,
        },
    })
    export default class Settings extends Vue {
        files: File[] = [];
    }
</script>