<template>
  <div :class="['container', { reverse: reverse }]">
    <div class="left">
      <img src="@/assets/user/1.png" alt="" />
    </div>
    <div :class="['right', reverse ? 'right_right' : 'right_left']">
      <h2>名称</h2>
      <p>2022-06-03 13:24:36</p>
      <div class="content">
        你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from "vue";
const props = defineProps(["reverse"]);
console.log(props.reverse);
</script>

<style lang="less" scoped>
.container {
  display: flex;
  justify-content: flex-start;
  margin-bottom: 10px;
  width: 80%;
}
.left img {
  display: block;
  width: 60px;
  height: 60px;
}
.right {
  h2 {
    font-size: 20px;
    font-weight: 400;
    line-height: 1.4;
  }
  p {
    line-height: 1.4;
  }
}
.right_left {
  text-align: left;
}
.right_right {
  text-align: right;
}
.content {
  line-height: 1.5;
}
.reverse {
  flex-direction: row-reverse;
}
</style>>
