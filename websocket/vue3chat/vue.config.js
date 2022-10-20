const { defineConfig } = require("@vue/cli-service")
const AutoImport = require("unplugin-auto-import/webpack")
const Components = require("unplugin-vue-components/webpack")
const { ElementPlusResolver } = require("unplugin-vue-components/resolvers")
const production = process.env.NODE_ENV === "production"
module.exports = defineConfig({
  publicPath: process.env.NODE_ENV === "production" ? "/socket/" : "./",
  assetsDir: "static",
  outputDir: "dist",
  indexPath: "index.html",
  lintOnSave: false,
  parallel: false,
  transpileDependencies: true,
  configureWebpack: {
    plugins: [
      AutoImport({
        resolvers: [ElementPlusResolver()],
      }),
      Components({
        resolvers: [ElementPlusResolver()],
      }),
    ],
  },
})
