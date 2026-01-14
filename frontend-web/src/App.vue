<!--
  前台根组件
-->
<template>
  <router-view />
  <Footer v-if="showFooter" />
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import Footer from '@/components/Footer.vue'

const route = useRoute()
const showFooter = computed(() => route.meta.layout !== false)
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

html, body, #app {
  height: 100%;
  font-family: 'Helvetica Neue', Helvetica, 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', Arial, sans-serif;
}
</style>
