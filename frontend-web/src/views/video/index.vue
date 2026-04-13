<template>
  <div class="video-page">
    <PageHeader />
    <main class="main-content">
      <section class="hero">
        <div class="container">
          <h1 class="title">视频</h1>
          <p class="subtitle">品牌活动与产品内容</p>
        </div>
      </section>

      <section class="content">
        <div class="container">
          <div class="video-grid">
            <article v-for="item in videoList" :key="item.src" class="video-card">
              <div class="video-wrapper">
                <video controls playsinline preload="metadata">
                  <source :src="item.src" type="video/mp4" />
                </video>
              </div>
              <div class="video-meta">
                <div class="video-name">{{ item.title }}</div>
              </div>
            </article>
          </div>
        </div>
      </section>

      <section class="china-market">
        <div class="container china-market-inner">
          <div class="china-market-content">
            <span class="section-label">CHINA</span>
            <h2 class="section-title">2017 年<br />白金盐进入中国市场</h2>
            <p class="text-body">进口商：中盐集团、广盐集团</p>
            <ul class="china-market-points">
              <li>广泛上架中高端商超、知名连锁超市、百货公司销售</li>
              <li>尤其覆盖北上广深地区</li>
            </ul>
          </div>
          <div class="china-market-gallery">
            <picture class="china-market-photo">
              <source :srcset="china01WebpSrcset" type="image/webp" />
              <img :src="china01FallbackSrc" alt="2017 白金盐进入中国市场" loading="lazy" decoding="async" />
            </picture>
            <picture class="china-market-photo">
              <source :srcset="china02WebpSrcset" type="image/webp" />
              <img :src="china02FallbackSrc" alt="渠道上架展示" loading="lazy" decoding="async" />
            </picture>
            <picture class="china-market-photo">
              <source :srcset="china03WebpSrcset" type="image/webp" />
              <img :src="china03FallbackSrc" alt="商超渠道展示" loading="lazy" decoding="async" />
            </picture>
            <picture class="china-market-photo">
              <source :srcset="china04WebpSrcset" type="image/webp" />
              <img :src="china04FallbackSrc" alt="北上广深地区销售" loading="lazy" decoding="async" />
            </picture>
          </div>
        </div>
      </section>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import PageHeader from '@/components/PageHeader.vue'

const videoModules = import.meta.glob('@/assets/video/*.{mp4,webm,ogg}', {
  eager: true,
  query: '?url',
  import: 'default'
})

const videoList = computed(() => {
  const parseDate = (t) => {
    const toTs = (y, m = 1, d = 1) => new Date(Number(y), Number(m) - 1, Number(d)).getTime()
    let m
    m = t.match(/(\d{4})年(\d{1,2})月(\d{1,2})日?/)
    if (m) return toTs(m[1], m[2], m[3])
    m = t.match(/(\d{4})年(\d{1,2})月/)
    if (m) return toTs(m[1], m[2], 1)
    m = t.match(/(\d{4})[-\/\.](\d{1,2})[-\/\.](\d{1,2})/)
    if (m) return toTs(m[1], m[2], m[3])
    m = t.match(/(\d{4})年/)
    if (m) return toTs(m[1], 1, 1)
    return 0
  }

  const items = Object.entries(videoModules).map(([path, url]) => {
    const fileName = path.split('/').pop() || path
    const title = fileName.replace(/\.(mp4|webm|ogg)$/i, '')
    const date = parseDate(title)
    return { title, src: url, date }
  })

  return items.sort((a, b) => b.date - a.date || b.title.localeCompare(a.title, 'zh-Hans-CN'))
})

import china01WebpSrcset from '@/assets/brand/_20260323134841_1644_1 (1).jpg?w=320;480;640&format=webp&quality=78&as=srcset'
import china01FallbackSrc from '@/assets/brand/_20260323134841_1644_1 (1).jpg?w=640&format=jpg&quality=82'
import china02WebpSrcset from '@/assets/brand/_20260323134841_1644_1 (2).jpg?w=320;480;640&format=webp&quality=78&as=srcset'
import china02FallbackSrc from '@/assets/brand/_20260323134841_1644_1 (2).jpg?w=640&format=jpg&quality=82'
import china03WebpSrcset from '@/assets/brand/_20260323134841_1644_1 (3).jpg?w=320;480;640&format=webp&quality=78&as=srcset'
import china03FallbackSrc from '@/assets/brand/_20260323134841_1644_1 (3).jpg?w=640&format=jpg&quality=82'
import china04WebpSrcset from '@/assets/brand/_20260323134841_1644_1 (4).jpg?w=320;480;640&format=webp&quality=78&as=srcset'
import china04FallbackSrc from '@/assets/brand/_20260323134841_1644_1 (4).jpg?w=640&format=jpg&quality=82'
</script>

<style scoped lang="scss">
.video-page {
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

.video-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 24px;
}

.video-card {
  border-radius: 12px;
  overflow: hidden;
  background: #fff;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08);
}

.video-wrapper {
  background: #000;

  video {
    width: 100%;
    height: auto;
    display: block;
  }
}

.video-meta {
  padding: 14px 16px 16px;
}

.video-name {
  font-size: 14px;
  font-weight: 600;
  color: #1a1a2e;
  line-height: 1.4;
}

.china-market {
  padding: 20px 0 60px;
}

.china-market-inner {
  display: grid;
  grid-template-columns: 1.05fr 0.95fr;
  gap: 44px;
  align-items: start;
  border: 1px solid rgba(212, 175, 55, 0.18);
  border-radius: 18px;
  padding: 24px;
  background: linear-gradient(180deg, rgba(212,175,55,0.08), #fff 60%);
}

.section-label {
  display: block;
  font-size: 12px;
  letter-spacing: 4px;
  text-transform: uppercase;
  color: #d4af37;
  margin-bottom: 10px;
  font-weight: 700;
}

.section-title {
  font-size: 32px;
  line-height: 1.2;
  color: #0f172a;
  margin-bottom: 12px;
}

.text-body {
  font-size: 16px;
  color: #334155;
  margin-bottom: 10px;
}

.china-market-points {
  margin: 0;
  padding: 0;
  list-style: none;
  display: grid;
  gap: 10px;

  li {
    position: relative;
    padding-left: 18px;
    font-size: 15px;
    color: #334155;
    line-height: 1.7;
  }

  li::before {
    content: "";
    position: absolute;
    left: 0;
    top: 10px;
    width: 8px;
    height: 8px;
    border-radius: 999px;
    background: #d4af37;
    box-shadow: 0 0 0 4px rgba(212, 175, 55, 0.16);
  }
}

.china-market-gallery {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;
}

.china-market-photo {
  display: block;
  border-radius: 14px;
  overflow: hidden;
  background: rgba(0, 0, 0, 0.03);
  box-shadow: 0 10px 26px rgba(0, 0, 0, 0.1);
  aspect-ratio: 4/3;
}

.china-market-photo img {
  width: 100%;
  height: 100%;
  display: block;
  object-fit: cover;
  transform: scale(1.001);
  transition: transform 0.6s ease;
}

.china-market-photo:hover img {
  transform: scale(1.05);
}

@media (max-width: 1024px) {
  .video-grid {
    grid-template-columns: repeat(2, 1fr);
  }

  .china-market-inner {
    grid-template-columns: 1fr;
    gap: 22px;
  }
}

@media (max-width: 768px) {
  .hero {
    padding: 36px 0 22px;
  }

  .title {
    font-size: 24px;
  }

  .video-grid {
    grid-template-columns: 1fr;
  }

  .china-market-gallery {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>
