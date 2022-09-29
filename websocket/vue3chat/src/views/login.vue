<template>
  <div class="login_container">
    <h1 class="title">等风来</h1>
    <div class="main">
      <img :src="select_avatar.url" class="user_top" alt="" />
      <div class="user_bottom">
        <img
          v-for="item in user_avatar_list"
          :key="item.index"
          :src="item.url"
          alt=""
          :class="{ selected: item.selected }"
          @click="changeAvatar(item)"
          @mouseenter="mouseenter(item)"
          @mouseleave="mouseleave"
        />
      </div>
      <div class="form_container">
        <el-form :model="form" label-position="top" label-width="120px">
          <el-form-item label="">
            <el-input v-model="form.name" placeholder="输入账号" />
          </el-form-item>
          <el-form-item label="">
            <el-input v-model="form.password" placeholder="输入密码" />
          </el-form-item>
          <el-form-item class="buttons">
            <el-button type="primary">登录/注册</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, reactive } from "vue";
const form = reactive({
  name: "",
  password: "",
});
const user_avatar_list = reactive([
  {
    index: 1,
    url: require("@/assets/user/1.png"),
    selected: false,
  },
  {
    index: 2,
    url: require("@/assets/user/2.png"),
    selected: false,
  },
  {
    index: 3,
    url: require("@/assets/user/3.png"),
    selected: false,
  },
  {
    index: 4,
    url: require("@/assets/user/4.png"),
    selected: false,
  },
  {
    index: 5,
    url: require("@/assets/user/5.png"),
    selected: false,
  },
  {
    index: 6,
    url: require("@/assets/user/6.png"),
    selected: false,
  },
  {
    index: 7,
    url: require("@/assets/user/7.png"),
    selected: false,
  },
  {
    index: 8,
    url: require("@/assets/user/8.png"),
    selected: false,
  },
  {
    index: 9,
    url: require("@/assets/user/9.png"),
    selected: false,
  },
  {
    index: 10,
    url: require("@/assets/user/10.png"),
    selected: false,
  },
  {
    index: 11,
    url: require("@/assets/user/11.png"),
    selected: false,
  },
  {
    index: 12,
    url: require("@/assets/user/12.png"),
    selected: false,
  },
]);
let select_avatar = reactive({
  index: 1,
  url: require("@/assets/user/1.png"),
});
let select_avatar_old = {};
onMounted(() => {
  user_avatar_list[0].selected = true;
  select_avatar.index = user_avatar_list[0].index;
  select_avatar.url = user_avatar_list[0].url;
  select_avatar_old = Object.assign(select_avatar_old, select_avatar);
});
let changeAvatar = (item) => {
  select_avatar.index = item.index;
  select_avatar.url = item.url;
  select_avatar_old = Object.assign(select_avatar_old, select_avatar);
  select_icon(item.index);
};
let mouseenter = (item) => {
  user_avatar_list.map((e) => {
    e.selected = false;
  });
  item.selected = true;
  select_avatar.index = item.index;
  select_avatar.url = item.url;
};
let mouseleave = () => {
  select_avatar = Object.assign(select_avatar, select_avatar_old);
  select_icon(select_avatar.index);
};
let select_icon = (index) => {
  user_avatar_list.map((e) => {
    e.selected = false;
    if (e.index == index) {
      e.selected = true;
    }
  });
};
</script>

<style scoped lang="less">
.login_container {
  position: relative;
  width: 100vw;
  height: 100vh;
  display: flex;
  justify-content: center;
  background: url("@/assets/images/cool-background.png");
  background-size: 100%;
  .main {
    width: 600px;
    padding: 128px 60px 48px;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
}
.title {
  position: fixed;
  font-size: 30px;
  font-weight: 500;
  top: 50px;
  left: 80px;
  color: #333;
}
:deep(.el-form-item) {
  width: 300px;
}

:deep(.buttons .el-form-item__content) {
  display: flex;
  align-items: center;
  justify-content: center;
}
.user_top {
  display: block;
  width: 158px;
  height: 158px;
  margin: 25px 0;
}
.user_bottom {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-wrap: wrap;
  margin-bottom: 30px;
  img {
    display: block;
    width: 30px;
    height: 30px;
    margin: 5px;
    cursor: pointer;
    resize: none;
    opacity: 0.5;
    &.selected {
      opacity: 1;
    }
  }
}
</style>