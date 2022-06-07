<template>
  <div>
    <Search
      class="select-disabled"
      :modules="modules"
      ref="search"
      @handling="handling"
      @focusTextInput="focusTextInput"
    />
    <ToolBar
      class="select-disabled"
      @emptyTextInput="emptyTextInput"
      @openSearch="openSearch"
      @copyTextInput="copyTextInput"
      @undoTextInput="undoTextInput"
    />
    <TextInput ref="textInput" @openSearch="openSearch" />
  </div>
</template>

<script>
import { Toast } from "vant";
import ToolBar from "./components/ToolBar.vue";
import TextInput from "./components/TextInput.vue";
import Search from "./components/Search.vue";
import { GetModules, Handling } from "../wailsjs/go/main/App";
export default {
  name: "App",
  data() {
    return {
      modules: [],
    };
  },
  components: {
    ToolBar,
    TextInput,
    Search,
  },
  methods: {
    openSearch() {
      this.$refs.textInput.isFocus = false;
      this.$refs.search.openSearch();
    },
    focusTextInput() {
      this.$refs.textInput.isFocus = true;
      this.$refs.textInput.focus();
    },
    handling(func) {
      const _this = this;
      const rawText = this.$refs.textInput.message;
      if (!rawText) return Toast.fail("文本为空");
      const funcName = func.name;
      Toast.loading({
        message: "正在处理文本...",
        duration: 0,
        forbidClick: true,
        loadingType: "spinner",
      });
      Handling(funcName, rawText).then((result) => {
        _this.$refs.textInput.history.push({
          func: funcName,
          text: rawText,
        });
        _this.$refs.textInput.history.push({
          func: funcName,
          text: result,
        });
        _this.$refs.textInput.message = result;
        _this.$refs.textInput.now = _this.$refs.textInput.history.length;
        Toast.clear();
      });
    },
    // 清空文本框
    emptyTextInput() {
      if (this.$refs.textInput.isFocus) this.$refs.textInput.empty();
    },
    copyTextInput() {
      if (this.$refs.textInput.isFocus) {
        this.$refs.textInput.copy();
        Toast.success("复制成功");
      }
    },
    undoTextInput(){
      if (this.$refs.textInput.isFocus) this.$refs.textInput.undo();
    }
  },
  beforeCreate() {
    Toast.loading({
      message: "正在加载模块",
      duration: 0,
      forbidClick: true,
      loadingType: "spinner",
    });
    GetModules().then((result) => {
      this.modules = result;
      Toast.clear();
    });
  },
  mounted() {
    var _this = this;
    document.addEventListener("keydown", function (e) {
      if (
        e.metaKey &&
        (e.code.toLowerCase() === "enter" || e.key.toLowerCase() === "f")
      ) {
        _this.openSearch();
      }
    });
  },
};
</script>