<!--
  UnoCSS 测试页面
-->
<template>
  <div class="min-h-screen bg-gray-50 py-12 px-4">
    <div class="container-custom">
      <h1 class="text-4xl font-bold text-center mb-8 text-primary">
        UnoCSS 原子化 CSS 测试页面
      </h1>

      <!-- 按钮示例 -->
      <section class="mb-12 bg-white p-8 rounded-lg shadow">
        <h2 class="text-2xl font-semibold mb-4">按钮样式</h2>
        <div class="flex gap-4 flex-wrap">
          <button class="btn-primary">主要按钮</button>
          <button class="btn-secondary">次要按钮</button>
          <button class="px-6 py-3 bg-green-500 text-white rounded hover:bg-green-600 transition">
            成功按钮
          </button>
          <button class="px-6 py-3 bg-gray-500 text-white rounded hover:bg-gray-600 transition">
            默认按钮
          </button>
        </div>
      </section>

      <!-- 网格布局 -->
      <section class="mb-12 bg-white p-8 rounded-lg shadow">
        <h2 class="text-2xl font-semibold mb-4">网格布局</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          <div 
            v-for="i in 8" 
            :key="i"
            class="bg-gradient-to-br from-blue-400 to-purple-500 p-6 rounded-lg text-white text-center hover:scale-105 transition-transform cursor-pointer"
          >
            <div class="text-3xl mb-2">📦</div>
            <div class="font-semibold">卡片 {{ i }}</div>
          </div>
        </div>
      </section>

      <!-- 文字样式 -->
      <section class="mb-12 bg-white p-8 rounded-lg shadow">
        <h2 class="text-2xl font-semibold mb-4">文字样式</h2>
        <p class="text-sm text-gray-600 mb-2">小号文字 (text-sm)</p>
        <p class="text-base text-gray-700 mb-2">基础文字 (text-base)</p>
        <p class="text-lg text-gray-800 mb-2">大号文字 (text-lg)</p>
        <p class="text-xl font-bold text-primary mb-2">特大文字 (text-xl)</p>
        <p class="text-2xl font-bold bg-gradient-to-r from-primary to-pink-500 bg-clip-text text-transparent">
          渐变文字效果
        </p>
      </section>

      <!-- 间距和对齐 -->
      <section class="mb-12 bg-white p-8 rounded-lg shadow">
        <h2 class="text-2xl font-semibold mb-4">Flexbox 布局</h2>
        <div class="flex items-center justify-between bg-gray-100 p-4 rounded mb-4">
          <span>左侧内容</span>
          <span class="font-bold text-primary">中间内容</span>
          <span>右侧内容</span>
        </div>
        <div class="flex items-center justify-center gap-4 bg-gray-100 p-4 rounded">
          <div class="w-16 h-16 bg-red-500 rounded-full"></div>
          <div class="w-16 h-16 bg-green-500 rounded-full"></div>
          <div class="w-16 h-16 bg-blue-500 rounded-full"></div>
        </div>
      </section>

      <!-- 响应式设计 -->
      <section class="mb-12 bg-white p-8 rounded-lg shadow">
        <h2 class="text-2xl font-semibold mb-4">响应式设计</h2>
        <div class="bg-primary text-white p-4 rounded mb-4 text-center">
          <p class="text-sm md:text-base lg:text-xl">
            这段文字在不同屏幕尺寸下有不同大小
          </p>
        </div>
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
          <div v-for="i in 4" :key="i" class="bg-blue-100 p-4 rounded text-center">
            响应式卡片 {{ i }}
          </div>
        </div>
      </section>

      <!-- 悬停效果 -->
      <section class="bg-white p-8 rounded-lg shadow">
        <h2 class="text-2xl font-semibold mb-4">悬停和过渡效果</h2>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <div class="bg-red-500 text-white p-6 rounded text-center hover:bg-red-600 hover:scale-110 transition-all cursor-pointer">
            悬停放大
          </div>
          <div class="bg-green-500 text-white p-6 rounded text-center hover:rotate-6 transition-transform cursor-pointer">
            悬停旋转
          </div>
          <div class="bg-blue-500 text-white p-6 rounded text-center hover:shadow-2xl transition-shadow cursor-pointer">
            悬停阴影
          </div>
          <div class="bg-purple-500 text-white p-6 rounded text-center hover:rounded-full transition-all cursor-pointer">
            悬停圆形
          </div>
        </div>
      </section>

      <!-- 返回首页 -->
      <div class="mt-12 text-center">
        <router-link to="/" class="btn-primary text-lg">
          返回首页
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
// UnoCSS 测试页面
</script>
