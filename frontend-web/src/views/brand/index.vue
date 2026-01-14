<template>
  <PageHeader />

  <iframe 
    src="/brand.html" 
    frameborder="0"
    class="brand-iframe"
  ></iframe>
</template>
<script setup>
  import PageHeader from "@/components/PageHeader.vue"
</script>
<style scoped>
.brand-iframe {
  width: 100%;
  height: 100vh;
  border: none;
  display: block;
}
</style>
