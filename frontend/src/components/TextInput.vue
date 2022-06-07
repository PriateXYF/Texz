<template>
  <div data-wails-no-drag @click="focus">
    <van-field
      ref="input"
      :autofocus="true"
      @blur="focus()"
      v-model="message"
      rows="5"
      type="textarea"
      placeholder="请输入处理字符"
    />
    <p class="show-tip select-disabled" v-show="now === 0">输入字符串后按 CMD/Ctrl + Enter 选择处理方法</p>
    <p class="show-tip select-disabled" v-show="now != 0">
      {{ history.length > 0 ? history[now - 1].func : "" }}
      的{{ now % 2 == 0 ? "输出" : "输入" }}
      {{ now + "/" + history.length }}
    </p>
    <van-row class="select-disabled">
      <van-col class="history-btn-col" span="12">
        <button :disabled="now <= 1" class="history-btn" @click="backVersion">
          上一版本
        </button>
      </van-col>
      <van-col class="history-btn-col" span="12">
        <button
          :disabled="now >= history.length"
          class="history-btn"
          @click="nextVersion"
        >
          下一版本
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
      isFocus: true,
      message: "",
      history: [],
      now: 0,
    };
  },
  methods: {
    focus() {
      if(this.isFocus) this.$refs.input.focus();
    },
    empty(){
      this.message = ""
    },
    copy(){
      navigator.clipboard.writeText(this.message);
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