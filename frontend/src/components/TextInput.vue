<template>
  <div data-wails-no-drag>
    <van-field
      ref="input"
      v-on:keydown.enter="openSearch"
      :autofocus="true"
      v-model="message"
      rows="5"
      type="textarea"
      placeholder="请输入处理字符"
    />
    <van-field v-model="showTip" input-align="center" :disabled="true" />
    <van-row>
      <van-col class="history-btn-col" span="12">
        <button class="history-btn">上一版本</button
      ></van-col>
      <van-col class="history-btn-col" span="12">
        <button class="history-btn small">下一版本</button
      ></van-col>
    </van-row>
  </div>
</template>

<style lang="less">
.van {
  &-cell {
    background-color: transparent;
  }
  &-cell-group {
    background-color: transparent;
  }
  &-field__control {
    color: #ffffff !important;
  }
}
.history-btn {
  width: 80%;
  margin: 10px;
}
.history-btn-col {
  text-align: center;
}
</style>

<script>
export default {
  data() {
    return {
      message: "",
      history: [],
      showTip: `已进行了${history.length}次文本处理`,
    };
  },
  methods: {
    openSearch(e) {
      if (e.metaKey || e.altKey) {
        this.$emit("openSearch");
      }
    },
    focus() {
      this.$refs.input.focus();
    },
  },
};
</script>