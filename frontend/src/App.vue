<template>
  <div>
    <Search
      :modules="modules"
      ref="search"
      @handling="handling"
      @focusTextInput="focusTextInput"
    />
    <ToolBar />
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
      this.$refs.search.openSearch();
    },
    focusTextInput() {
      this.$refs.textInput.focus();
    },
    handling(func) {
      const _this = this
      const rawText = this.$refs.textInput.message;
      const funcName = func.name;
      Toast.loading({
        message: "正在处理文本...",
        duration: 0,
        forbidClick: true,
        loadingType: "spinner",
      });
      Handling(funcName, rawText).then((result) => {
        _this.$refs.textInput.history.push({
          func : funcName,
          text : rawText
        })
        _this.$refs.textInput.history.push({
          func : funcName,
          text : result
        })
        _this.$refs.textInput.message = result
        _this.$refs.textInput.now = _this.$refs.textInput.history.length
        Toast.clear();
      });
    },
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
};
</script>