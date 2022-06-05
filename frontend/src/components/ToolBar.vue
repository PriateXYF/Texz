<template>
  <van-row data-wails-drag type="flex" justify="space-between">
    <van-col span="6"></van-col>
    <van-col span="6">
      <van-icon name="setting" class="tool-bar-icon" />
      <van-icon name="description" class="tool-bar-icon" />
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