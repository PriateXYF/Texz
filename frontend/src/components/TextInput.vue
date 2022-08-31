<template>
  <div data-wails-no-drag @click="focus">
    <van-field
      class="text-input"
      ref="input"
      :autofocus="true"
      @blur="focus()"
      @dragenter="dragenter"
      @dragleave="dragleave"
      @drop="drop"
      :maxlength="2 ** 19"
      show-word-limit
      v-model="message"
      rows="6"
      type="textarea"
      :placeholder="placeholder"
    />
    <p class="show-tip select-disabled" v-show="now === 0">
      输入字符串后按 CMD/Ctrl + Enter 选择处理方法
    </p>
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
  &-field__word-num {
    color: #ffffff !important;
  }
  &-field__word-limit {
    color: rgba(255, 255, 255, 0.7) !important;
  }
}
.history-btn {
  width: 80%;
  margin: 10px;
}
.history-btn-col {
  text-align: center;
}

.text-input textarea::-webkit-input-placeholder {
  text-align: center;
  line-height: 140px;
}
.text-input textarea:-moz-placeholder {
  text-align: center;
  line-height: 140px;
}
.text-input textarea::-moz-placeholder {
  text-align: center;
  line-height: 140px;
}
.text-input textarea:-ms-input-placeholder {
  text-align: center;
  line-height: 140px;
}
</style>

<script>
import { Toast } from "vant";
import { Copy } from "../../wailsjs/go/main/App";
export default {
  props: ["config"],
  data() {
    return {
      isFocus: true,
      message: "",
      history: [],
      now: 0,
      placeholder: "输入文字或拖入文本文件",
    };
  },
  methods: {
    focus() {
      if (this.isFocus) this.$refs.input.focus();
    },
    empty() {
      this.message = "";
    },
    copy() {
      Copy(this.message).then((err) => {
        if (!err) Toast.success("复制成功");
        else Toast.fail("复制失败");
      });
    },
    undo() {
      if (this.now > 0) this.message = this.history[this.now - 1].text;
    },
    backVersion() {
      if (this.now > 1) this.message = this.history[--this.now - 1].text;
    },
    nextVersion() {
      if (this.now < this.history.length)
        this.message = this.history[++this.now - 1].text;
    },
    dragenter() {
      this.empty();
      this.placeholder = "松开鼠标导入文件";
    },
    dragleave() {
      this.undo();
      this.placeholder = "输入文字或拖入文本文件";
    },
    drop(e) {
      e.preventDefault();
      const _this = this;
      this.placeholder = "输入文字或拖入文本文件";
      if (e.dataTransfer.files.length != 1)
        return Toast.fail("只能拖入一个文件");
      const file = e.dataTransfer.files[0];
      console.log(file.type)
      const whiteFileList = ["text", "application/json", "application/x-yaml", "application/x-javascript", "application/x-sh"];
      const fileExt = file.name.split(".")[file.name.split(".").length - 1].toLowerCase();
      if (!whiteFileList.some((white) => file.type.indexOf(white) == 0))
        if (this.config.whiteFileList){
          const notInWhiteFileList = !this.config.whiteFileList.some((configWhite) => fileExt === configWhite.toLowerCase())
          if(notInWhiteFileList) return Toast.fail("只能拖入文本文件")
        }else{
          return Toast.fail("只能拖入文本文件")
        }
      var reader = new FileReader();
      reader.readAsText(file);
      Toast.loading({
        message: "正在读取文件",
        duration: 0,
        forbidClick: true,
      });
      reader.onload = function () {
        Toast.clear();
        if (reader.result.length > 2 ** 19) return Toast.fail("字数过多");
        _this.message = reader.result;
      };
      reader.onerror = function () {
        Toast.clear();
        Toast.fail("解析文件失败");
      };
    },
  },
  mounted() {
    var _this = this;
    document.addEventListener("keydown", function (e) {
      if (e.metaKey && e.key.toLowerCase() === "1") {
        _this.backVersion();
      }
      if (e.metaKey && e.key.toLowerCase() === "2") {
        _this.nextVersion();
      }
    });
  },
};
</script>