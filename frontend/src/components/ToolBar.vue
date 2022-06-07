<template>
  <van-row data-wails-drag type="flex" justify="space-between">
    <van-col span="12"></van-col>
    <van-col span="12">
      <van-icon name="setting" class="tool-bar-icon" />
      <van-icon name="replay" class="tool-bar-icon" />
      <van-icon name="delete" class="tool-bar-icon" @click="emptyTextInput" />
      <van-icon name="description" class="tool-bar-icon" @click="copyTextInput"/>
      <van-icon name="search" class="tool-bar-icon" @click="openSearch"/>
    </van-col>
  </van-row>
</template>

<style>
.tool-bar-icon {
  cursor: pointer;
  font-size: 1.4rem;
  float: right;
  margin-top: 6px;
  margin-right: 10px;
}
</style>

<script>
export default {
  methods: {
    emptyTextInput() {
      this.$emit("emptyTextInput");
    },
    openSearch(){
      this.$emit("openSearch");
    },
    copyTextInput(){
      this.$emit("copyTextInput");
    }
  },
  mounted() {
    var _this = this;
    document.addEventListener("keydown", function (e) {
      if (e.metaKey && e.key.toLowerCase() === "d") {
        _this.emptyTextInput();
      }
      if (e.altKey && e.code.toLowerCase() === "keyc") {
        e.preventDefault()
        _this.copyTextInput();
      }
    });
  },
};
</script>