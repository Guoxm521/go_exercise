<template>
  <div :class="['container', { reverse: reverse }]">
    <div class="left">
      <!-- <img src="@/assets/user/1.png" alt="" /> -->
      <img :src="avatar" alt="" />
    </div>
    <div :class="['right', reverse ? 'right_right' : 'right_left']">
      <h2>{{ content.user_name }}</h2>
      <p>{{ parseTime(content.time) }}</p>
      <div class="content">
        {{ content.content }}
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, computed, ref } from "vue";
import { parseTime, getUserAvatar } from "@/utils/filter";
const props = defineProps(["reverse", "content"]);
let content = props.content;
content.avatar = content.avatar || 4;
const avatar = require("./../assets/user/" + content.avatar + ".png");
</script>

<style lang="less" scoped>
.container {
  display: flex;
  justify-content: flex-start;
  margin-bottom: 20px;
  width: 80%;
}
.left img {
  display: block;
  width: 50px;
  height: 50px;
}
.right {
  h2 {
    font-size: 18px;
    font-weight: 400;
    line-height: 1.4;
  }
  p {
    line-height: 1.6;
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
  background-color: #d7ece7;
  padding: 5px;
  border-radius: 5px;
}
.reverse {
  flex-direction: row-reverse;
}
</style>>
