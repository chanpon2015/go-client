<template>
    <div>
        <label for="upload">
            <input type="file" id="upload" webkitdirectory>
        </label>
    </div>
</template>

<script lang="ts">
    import { Component, Vue } from 'vue-property-decorator';
    @Component
    export default class Settings extends Vue {
    }
</script>

<style>
label {
    background: #000;
    color: #fff;
    cursor: pointer;
    padding: 15px 30px;
    margin: 0 10px 0 0;
    display: inline-block;
}
</style>