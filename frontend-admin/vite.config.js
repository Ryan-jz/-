import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// Vite 配置文件
// 配置项目构建和开发服务器选项
export default defineConfig({
  // Vue 插件配置
  plugins: [vue()],
  
  // 路径别名配置
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src')
    }
  },
  
  // 开发服务器配置
  server: {
    port: 3000,           // 开发服务器端口
    open: true,           // 自动打开浏览器
    cors: true,           // 允许跨域
    
    // 代理配置 - 解决开发环境跨域问题
    proxy: {
      '/api': {
        target: 'http://8.133.175.112:8000',   // 后端 API 地址
        changeOrigin: true,                // 改变请求源
        // 不需要 rewrite，保持 /api 前缀
      }
    }
  },
  
  // 构建配置
  build: {
    outDir: 'dist',       // 输出目录
    assetsDir: 'assets',  // 静态资源目录
    sourcemap: false,     // 不生成 sourcemap
    
    // 代码分割配置
    rollupOptions: {
      output: {
        manualChunks: {
          'vue-vendor': ['vue', 'vue-router', 'pinia'],
          'element-plus': ['element-plus']
        }
      }
    }
  }
})
