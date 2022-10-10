<template>
  <div class="home">
    <header>聊天室</header>
    <div class="container">
      <div class="item" v-for="item in 6" :key="item">
        <div class="item_container">
          <h2>{{ item }}</h2>
          <div class="user_box">
            <el-icon color="#32CCBC"><User /></el-icon>
            <span>23</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { getSocketInfo } from "@/api/index";
getSocketInfo({}).then((res) => {
  console.log(res);
});
</script>


<style lang="less" scoped>
.home {
  width: 100vw;
  height: 100vh;
  display: flex;
  justify-content: center;
  background: url("@/assets/images/main-background.png");
  background-size: 100%;
  position: relative;
}
.container {
  width: 800px;
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
  height: fit-content;
  padding-top: 150px;
  .item {
    width: 250px;
    height: 250px;
    padding: 10px;
    margin-top: 25px;
    box-sizing: border-box;
    border-radius: 8px;
    background-color: #fff;
  }
  .item:hover {
    border: 1px solid salmon;
  }
  .item_container {
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    position: relative;
    h2 {
      font-size: 46px;
      color: #333;
    }
  }
}
.user_box {
  position: absolute;
  bottom: 0px;
  right: 0px;
  display: flex;
  align-items: center;
  color: #333;
  span {
    margin-left: 10px;
    color: #32ccbc;
  }
}

header {
  position: fixed;
  top: 0;
  left: 0;
  height: 50px;
  width: 100%;
  background: rgba(172, 214, 208, 0.5);
  box-sizing: border-box;
  padding: 0 10px;
  display: flex;
  align-items: center;
}
</style>
