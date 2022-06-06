<template>
  <van-popup v-model="show" position="top">
    <van-search
      autofocus
      ref="search"
      @input="search"
      v-on:keydown.esc="closeSearch"
      v-model="value"
      :clearable="false"
      placeholder="查询处理方法"
    />
    <van-cell-group>
      <van-cell
        v-for="item in functions"
        :key="item.name"
        :title="item.name"
        :value="item.note"
        @click="selectFunction(item)"
      />
    </van-cell-group>
  </van-popup>
</template>
<style>
.van-cell {
  cursor: pointer;
}
.van-popup {
  background-color: transparent;
}
.van-search {
  background-color: transparent;
  border: solid 2px #ffffff !important;
  border-radius: 4px;
  margin: 20px;
  padding: 2px 2px !important;
}
.van-search__content {
  background-color: transparent;
}
.van-overlay {
  background-color: rgba(0, 0, 0, 0.9) !important;
}
</style>

<script>
export default {
  props: ["modules"],
  data() {
    return {
      show: false,
      value: "",
      functions: [],
    };
  },
  methods: {
    search(str) {
      const _this = this;
      this.functions = [];
      const key = str.toLowerCase();
      this.modules.forEach((mod) => {
        const funcs = mod.functions.filter((func) => {
          const funcName = func.name.toLowerCase();
          const funcNote = func.note.toLowerCase();
          return funcName.indexOf(key) >= 0 || funcNote.indexOf(key) >= 0;
        });
        _this.functions = [..._this.functions, ...funcs];
      });
    },
    openSearch() {
      const _this = this;
      this.show = true;
      setTimeout(() => {
        _this.$refs.search.querySelector("input").focus();
      }, 100);
    },
    closeSearch() {
      this.show = false;
      this.value = "";
      this.functions = [];
      this.$emit("focusTextInput");
    },
    // 选择一个处理方法
    selectFunction(item) {
      this.functions = [];
      this.$emit("handling", item);
      this.closeSearch();
    },
  },
};
</script>