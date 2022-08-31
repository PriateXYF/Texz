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
import $ from 'jquery';
import { Toast } from "vant";
import ToolBar from "./components/ToolBar.vue";
import TextInput from "./components/TextInput.vue";
import Search from "./components/Search.vue";
import { GetModules, Handling, GetConfig, GetRemoteModlues } from "../wailsjs/go/main/App";
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
    dealResult(rawText, result, func){
      const funcName = func.name
        this.$refs.textInput.history.push({
          func: funcName,
          text: rawText,
        });
        this.$refs.textInput.history.push({
          func: funcName,
          text: result,
        });
        this.$refs.textInput.message = result;
        this.$refs.textInput.now = this.$refs.textInput.history.length;
        if (this.config.autoCopy) this.copyTextInput()
    },
    handling(func) {
      const _this = this;
      const rawText = this.$refs.textInput.message;
      if (!rawText) return Toast.fail("文本为空");
      const funcName = func.name;
      // 如果是客户端模块
      if (func.func) {
        const result = func.func(rawText)
        return this.dealResult(rawText, result, func)
      }
      Toast.loading({
        message: "正在处理文本...",
        duration: 0,
        forbidClick: true,
        loadingType: "spinner",
      });
      Handling(funcName, rawText).then((result) => {
        _this.dealResult(rawText, result, func)
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
    
    async function getCostomModules() {
      const remoteModlues = await GetRemoteModlues();
      window.CostomModules = {
          name : "Custom",
          note : "用户自定义模块",
          functions : []
      }
      for (const index in remoteModlues) {
        await $.getScript(remoteModlues[index])
      }
      console.log(window.CostomModules)
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
    await getCostomModules();
    this.modules.push(window.CostomModules)
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