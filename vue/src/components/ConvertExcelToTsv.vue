<template>
    <div>
        <div>
            <label
                @dragover.prevent="dragver=true"
                @dragleave.prevent="dragover=false"
                @drop.prevent="onDrop"
                :class="{[$style['drag-area']]: true, [$style.dragover]: dragover}"
            >
            <span v-if="files.length==0">ここにファイルをドロップしてください</span>
            <span v-if="files.length > 0">{{ files[0].name}}</span>    
                <input
                    type="file"
                    style="display: none;"
                    @change="onDrop">
            </label>
        </div>
        <GoProcBtn
            color="red"
            goProcName="excel_to_tsv"
            :goProcPayload="payload"
        ></GoProcBtn>
    </div>
</template>

<script lang="ts">
    import GoProcBtn from '@/components/GoProcBtn.vue';
    import { Component, Vue } from 'vue-property-decorator';
    @Component({
        inheritAttrs: false,
        components: {
            GoProcBtn,
        },
    })
    export default class ConvertExcelToTsv extends Vue {
        dragover: boolean = false;
        files: any[] = [];
        filePath: string = '';
        payload: any;
        onDrop(event: any) {
            let fileList = event.target.files ?
                event.target.files :
                event.dataTransfer.files;
            for (let i = 0; i < fileList.length; i++) {
                this.files.push(fileList[i]);
            }
            if (0 < this.files.length) {
                this.filePath = this.files[0].path;
                this.payload = {
                    xlsx_path: this.filePath,
                    export_dir: "/test/",
                }
            }
        }
        test(response: any) {
            alert(response);
        }
    }
</script>

<style module>
.drag-area {
    width: 300px;
    padding: 30px;
    text-align: center;
    border: 1px dashed #000;
    border-radius: 2px;
    display: inline-block;
    cursor: pointer;
}
.drag-area.dragover {
    border: 1px dashed #2ca9e1;
}
.drag-area.hover {
    border: 1px dashed #999;
}
</style>