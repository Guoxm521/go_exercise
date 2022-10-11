<template>
  <div>
    <TopHeader></TopHeader>
    <div class="main_container">
      <router-view></router-view>
    </div>
  </div>
</template>

<script setup>
import { useUserStore } from "@/store/index";
import TopHeader from "./TopHeader.vue";
</script>

<style lang="less" scoped>
.main_container {
  margin-top: 50px;
  width: 100%;
  overflow: hidden;
}
</style>