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
    <p class="show-tip" v-show="now === 0">首次处理</p>
    <p class="show-tip" v-show="now != 0">
      {{ history.length > 0 ? history[now - 1].func : "" }}
      的{{ now % 2 == 0 ? "输出" : "输入" }}
      {{ now + "/" + history.length }}
    </p>
    <van-row>
      <van-col class="history-btn-col" span="12">
        <button v-show="now > 1" class="history-btn" @click="backVersion">
          上一版本
        </button>
        <button disabled v-show="now <= 1" class="history-btn">
          暂无数据
        </button></van-col
      >
      <van-col class="history-btn-col" span="12">
        <button
          v-show="now < history.length"
          class="history-btn"
          @click="nextVersion"
        >
          下一版本
        </button>
        <button disabled v-show="now >= history.length" class="history-btn">
          暂无数据
        </button>
      </van-col>
    </van-row>
  </div>
</template>

<style lang="less">
.show-tip {
  text-align: center;
}
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
      now: 0,
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
    backVersion() {
      this.message = this.history[--this.now - 1].text;
    },
    nextVersion() {
      this.message = this.history[++this.now - 1].text;
    },
  },
};
</script>