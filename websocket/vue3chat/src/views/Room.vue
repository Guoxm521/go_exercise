<template>
  <div class="home">
    <header>聊天室</header>
    <div class="main_box">
      <RoomItem
        :class="{ reverse: item % 2 }"
        v-for="item in 12"
        :key="item"
        :reverse="item % 2 ? true : false"
      ></RoomItem>
    </div>
  </div>
</template>

<script setup>
import RoomItem from "@/components/RoomItem.vue";
import { onMounted } from "vue";
onMounted(() => {
  setTimeout(() => {
    var height = document.body.scrollHeight;
    window.scroll({ top: height, left: 0, behavior: "smooth" });
  }, 2000);
});
</script>


<style lang="less" scoped>
.home {
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  background-image: radial-gradient(
    circle,
    #d5eae7,
    #d7ebe8,
    #d9ede9,
    #dbeeea,
    #ddefeb,
    #dbeeea,
    #daede9,
    #d8ece8,
    #d2e9e4,
    #cbe6e1,
    #c5e2dd,
    #bedfda
  );
}
.main_box {
  width: 900px;
  background-color: rgba(255, 255, 255, 0.5);
  padding: 50px 0;
  display: flex;
  flex-direction: column;
}
.reverse {
  align-self: flex-end;
}
header {
  position: fixed;
  top: 0;
  left: 0;
  height: 50px;
  width: 100%;
  background: #d5eae7;
  box-sizing: border-box;
  padding: 0 10px;
  display: flex;
  align-items: center;
}
</style>
