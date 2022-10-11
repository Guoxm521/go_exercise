<template>
  <div class="header_container">
    <div class="left">聊天室</div>
    <div class="right">
      <el-dropdown>
        <span class="el-dropdown-link">
          {{ user_name }}
          <el-icon class="el-icon--right">
            <arrow-down />
          </el-icon>
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item>
              <div @click="logout">退出登录</div>
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>
<script setup>
import { useUserStore } from "@/store/index";
import { computed } from "vue";
import { ArrowDown } from "@element-plus/icons-vue";
import { ElMessage, ElMessageBox } from "element-plus";
import "element-plus/es/components/message-box/style/css";
import { removeToken } from "@/utils/cookie";
import { useRouter } from "vue-router";
const userStore = useUserStore();
const router = useRouter();
const user_name = computed(() => {
  return userStore.user_name;
});
function logout() {
  let message = `退出登录当前账号,${user_name.value}?`;
  ElMessageBox.confirm(message, "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(() => {
      removeToken();
      router.push("/login");
    })
    .catch(() => {});
}
</script>
<style lang="less" scoped>
.header_container {
  position: fixed;
  top: 0;
  left: 0;
  height: 50px;
  width: 100%;
  background: rgba(172, 214, 208, 1);
  box-sizing: border-box;
  padding: 0 50px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  z-index: 100;
}
</style>