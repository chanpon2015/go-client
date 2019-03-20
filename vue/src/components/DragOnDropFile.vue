<template>
    <div>
        <label
            @dragover.prevent="dragver=true"
            @dragleave.prevent="dragover=false"
            @drop.prevent="onDrop"
            :class="{[$style['drag-area']]: true, [$style.dragover]: dragover}"
        >
        <span
            v-if="files.length==0">
        {{ text }}
        </span>
        <span v-if="files.length > 0">{{ files[0].name}}</span>    
            <input
                type="file"
                style="display: none;"
                @change="onDrop">
        </label>
    </div>
</template>

<script lang="ts">
    import { Component, Prop, Vue } from 'vue-property-decorator';
    @Component
    export default class DragOnDropFile extends Vue {
        /** data */
        dragover: boolean = false;
        /** prop */
        @Prop() files!: File[];
        @Prop() text!: String;
        onDrop(event: any) {
            let fileList = event.target.files ?
                event.target.files :
                event.dataTransfer.files;
            for (let i = 0; i < fileList.length; i++) {
                this.files.push(fileList[i]);
            }
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
    background-color: azure;
}
.drag-area.dragover {
    border: 1px dashed #2ca9e1;
}
.drag-area.hover {
    border: 1px dashed #999;
}
</style>