import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'
import UnoCSS from 'unocss/vite'

// Vite 配置文件
export default defineConfig({
  base: './',
  plugins: [
    vue(),
    UnoCSS()
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
        target: 'http://8.133.175.112:8000',
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
