<template>
  <div class="store-page">
    <PageHeader />
    <main class="main-content">
      <section class="hero">
        <div class="container">
          <h1 class="title">旗舰店</h1>
          <p class="subtitle">扫码进入官方旗舰店</p>
        </div>
      </section>

      <section class="content">
        <div class="container">
          <div class="store-grid">
            <article v-for="item in stores" :key="item.name" class="store-card">
              <div class="store-image">
                <picture>
                  <source :srcset="item.webpSrcset" type="image/webp" />
                  <img
                    :src="item.fallbackSrc"
                    :alt="item.name"
                    loading="lazy"
                    decoding="async"
                  />
                </picture>
              </div>
              <div class="store-meta">
                <div class="store-name">{{ item.name }}</div>
                <div class="store-hint">{{ item.hint }}</div>
              </div>
            </article>
          </div>
        </div>
      </section>
    </main>
  </div>
</template>

<script setup>
import PageHeader from '@/components/PageHeader.vue'

import tmallWebpSrcset from '@/assets/images/tianmao.png?w=160;320;480&format=webp&as=srcset'
import tmallFallbackSrc from '@/assets/images/tianmao.png?w=480'
import jdWebpSrcset from '@/assets/images/jingdong.jpg?w=160;320;480&format=webp&as=srcset'
import jdFallbackSrc from '@/assets/images/jingdong.jpg?w=480'
import douyinWebpSrcset from '@/assets/images/douyin.jpg?w=160;320;480&format=webp&as=srcset'
import douyinFallbackSrc from '@/assets/images/douyin.jpg?w=480'

const stores = [
  {
    name: '天猫旗舰店',
    hint: '打开天猫扫一扫或保存图片扫码',
    webpSrcset: tmallWebpSrcset,
    fallbackSrc: tmallFallbackSrc
  },
  {
    name: '京东旗舰店',
    hint: '打开京东扫一扫或保存图片扫码',
    webpSrcset: jdWebpSrcset,
    fallbackSrc: jdFallbackSrc
  },
  {
    name: '抖音旗舰店',
    hint: '打开抖音扫一扫或保存图片扫码',
    webpSrcset: douyinWebpSrcset,
    fallbackSrc: douyinFallbackSrc
  }
]
</script>

<style scoped lang="scss">
.store-page {
  min-height: 100%;
  background: #fff;
}

.main-content {
  width: 100%;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

.hero {
  padding: 56px 0 32px;
  background-image: url('@/assets/images/background.jpg');
  background-size: cover;
  background-position: center;
}

.title {
  font-size: 32px;
  color: #1a1a2e;
  text-align: center;
  font-weight: 700;
}

.subtitle {
  margin-top: 10px;
  text-align: center;
  color: rgba(26, 26, 46, 0.75);
  font-size: 14px;
}

.content {
  padding: 28px 0 80px;
}

.store-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 24px;
  justify-content: center;
}

.store-card {
  border-radius: 12px;
  overflow: hidden;
  background: #fff;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08);
}

.store-image {
  padding: 24px;
  display: flex;
  justify-content: center;
  background: #fff;

  img {
    width: min(320px, 100%);
    height: auto;
    display: block;
    border-radius: 10px;
  }
}

.store-meta {
  padding: 0 16px 18px;
  text-align: center;
}

.store-name {
  font-size: 16px;
  font-weight: 700;
  color: #1a1a2e;
}

.store-hint {
  margin-top: 8px;
  font-size: 13px;
  color: rgba(26, 26, 46, 0.7);
  line-height: 1.5;
}

@media (max-width: 768px) {
  .hero {
    padding: 36px 0 22px;
  }

  .title {
    font-size: 24px;
  }

  .store-grid {
    grid-template-columns: 1fr;
  }
}
</style>
