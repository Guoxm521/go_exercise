export default {
  account: [
    { required: true, message: "请输入账号", trigger: "blur" },
    { min: 3, max: 10, message: "长度在 3 到 10 个字符", trigger: "blur" },
  ],
  password: [
    { required: true, message: "请输入密码", trigger: "blur" },
    { min: 6, max: 10, message: "长度在 6 到 10 个字符", trigger: "blur" },
  ],
}
