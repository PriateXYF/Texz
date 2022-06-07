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
    <TextInput :config="config" ref="textInput" @openSearch="openSearch" />
  </div>
</template>

<script>
import { Toast } from "vant";
import ToolBar from "./components/ToolBar.vue";
import TextInput from "./components/TextInput.vue";
import Search from "./components/Search.vue";
import { GetModules, Handling, GetConfig } from "../wailsjs/go/main/App";
export default {
  name: "App",
  data() {
    return {
      modules: [],
      config: {},
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
        if (_this.config.autoCopy) _this.copyTextInput();
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
      }
    },
    undoTextInput() {
      if (this.$refs.textInput.isFocus) this.$refs.textInput.undo();
    },
  },
  async beforeCreate() {
    Toast.loading({
      message: "正在载入配置",
      duration: 0,
      forbidClick: true,
      loadingType: "spinner",
    });
    async function getModlues() {
      const modules = await GetModules();
      return modules;
    }
    async function getConfig() {
      const config = await GetConfig();
      try {
        JSON.parse(config);
      } catch (error) {
        Toast.clear();
        return Toast.fail("解析配置文件失败，请重启");
      }
      return JSON.parse(config);
    }
    this.modules = await getModlues();
    this.config = await getConfig();
    Toast.clear();
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