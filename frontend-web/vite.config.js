import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'
import UnoCSS from 'unocss/vite'
import { imagetools } from 'vite-imagetools'

// Vite 配置文件
export default defineConfig({
  base: './',
  plugins: [
    vue(),
    UnoCSS(),
    imagetools()
  ],
  
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src')
    }
  },
  
  server: {
   host: "0.0.0.0",
    port: 3001,
    open: true,
    cors: true,
    proxy: {
      '/api': {
        target: 'https://bad-reichenhaller.com.cn/api/v1',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      }
    }
  },
  
  build: {
    outDir: 'dist',
    assetsDir: 'assets',
    sourcemap: false
  }
})
