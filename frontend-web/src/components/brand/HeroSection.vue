<template>
  <section class="hero-carousel">
    <div class="slides-container" :style="{ transform: `translateY(-${currentSlide * 100}vh)` }">
      <div class="slide" v-for="(slide, index) in slides" :key="index">
        <div class="slide-bg" :style="{ backgroundImage: `url(${slide.bg})` }"></div>
        <div class="slide-content">
          <h1 v-if="slide.title" class="slide-title" v-html="slide.title"></h1>
          <p v-if="slide.subtitle" class="slide-subtitle">{{ slide.subtitle }}</p>
          <div v-if="slide.features" class="features-grid">
            <div class="feature" v-for="(f, i) in slide.features" :key="i">
              <h3>{{ f.title }}</h3>
              <p>{{ f.desc }}</p>
            </div>
          </div>
          <p v-if="slide.desc" class="slide-desc" v-html="slide.desc"></p>
        </div>
      </div>
    </div>
    <div class="nav-dots">
      <span v-for="(s, i) in slides" :key="i" :class="{ active: currentSlide === i }" @click="currentSlide = i"></span>
    </div>
    <div class="nav-arrows">
      <button class="arrow-up" @click="prev">↑</button>
      <button class="arrow-down" @click="next">↓</button>
    </div>
  </section>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import bg1 from '@/assets/brand/微信图片_20250620112836.jpg'
import bg2 from '@/assets/brand/2017-12-05 090614.jpg'
import bg3 from '@/assets/brand/20250612白金盐_袋装盐1-5_低钠盐6-8(1).jpg'
import bg4 from '@/assets/brand/微信图片_20250408203638.jpg'
import bg5 from '@/assets/brand/2018-08-14 162039.jpg'
import bg6 from '@/assets/brand/产品图.jpg'

const emit = defineEmits(['carousel-complete'])
const currentSlide = ref(0)
let timer = null

const slides = [
  {
    bg: bg1,
    title: '阳光·空气·水·盐<br><span class="gold">Alpen 白金盐</span><br>德国 No.1 岩盐品牌',
    subtitle: '食盐中的钻石，皇冠级纯净'
  },
  {
    bg: bg2,
    title: '皇室故事',
    desc: '1517 年，德皇马克斯一世挥笔御令，<br>将贝希特斯加登盐矿册封为"Königliche Saline"——皇家私产。<br>自此，白金岩盐只供宫廷餐桌、主教盛宴与凯旋军宴，<br>被欧洲史家称作"可食用的王冠"。'
  },
  {
    bg: bg3,
    title: '核心卖点',
    features: [
      { title: '👑 皇室专供矿脉', desc: '500 年不间断开采' },
      { title: '💎 99.9% 极致纯净', desc: '减盐 8% 不减味' },
      { title: '🌿 0 添加哲学', desc: '孕妇宝宝可安心食用' },
      { title: '☀️ 生命第四元素', desc: '每日不可替代' },
      { title: '🔍 皇冠溯源', desc: 'NFC 一触即达' },
      { title: '💰 百姓价享御用', desc: '平均 1 元 / 天' }
    ]
  },
  {
    bg: bg4,
    title: '感官大片',
    desc: '原始海洋蒸发，2.5 亿年阿尔卑斯岩层封存，<br>远离现代海洋塑料与微塑料污染，<br>像给地球按下"干净存档"。<br><br>倒在掌心，碎雪般晶莹；<br>轻抿舌尖，矿物甘甜味层层晕开；<br>这就是皇室御厨笔记里的"雪吻之咸"。'
  },
  {
    bg: bg5,
    title: '权威背书',
    features: [
      { title: 'WHO 2023 减盐模型', desc: '高纯低杂质盐，家庭钠摄入 ↓12%' },
      { title: '德国 BGR 报告', desc: '杂质密度仅为全球均值 1/5' },
      { title: '维也纳宫廷菜单 1812', desc: '御膳房唯一指定' }
    ]
  },
  {
    bg: bg6,
    title: '可持续与公益',
    features: [
      { title: '⚡ 皇室绿电', desc: '2025 起 100% 水电驱动，年减碳 2,100 吨' },
      { title: '♻️ 空瓶回收', desc: '6 个空瓶换 250 g 补充装' },
      { title: '❤️ 一克盐·一克爱', desc: '每售 1 罐，向山区捐 1 g 碘盐' }
    ]
  }
]

const next = () => {
  if (currentSlide.value < slides.length - 1) {
    currentSlide.value++
  } else {
    emit('carousel-complete')
  }
}

const prev = () => {
  if (currentSlide.value > 0) {
    currentSlide.value--
  }
}

const handleWheel = (e) => {
  e.preventDefault()
  if (e.deltaY > 0) {
    next()
  } else {
    prev()
  }
}

onMounted(() => {
  const carousel = document.querySelector('.hero-carousel')
  carousel.addEventListener('wheel', handleWheel, { passive: false })
})

onUnmounted(() => {
  const carousel = document.querySelector('.hero-carousel')
  if (carousel) {
    carousel.removeEventListener('wheel', handleWheel)
  }
  if (timer) clearInterval(timer)
})
</script>

<style scoped lang="scss">
.hero-carousel {
  position: relative;
  height: 100vh;
  overflow: hidden;
}

.slides-container {
  height: 100%;
  transition: transform 0.8s cubic-bezier(0.4, 0, 0.2, 1);
}

.slide {
  position: relative;
  height: 100vh;
  width: 100%;
}

.slide-bg {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-size: cover;
  background-position: center;
  filter: brightness(0.4);
}

.slide-content {
  position: relative;
  z-index: 2;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
  color: #fff;
  padding: 0 40px;
  max-width: 1400px;
  margin: 0 auto;
}

.slide-title {
  font-size: 4rem;
  font-weight: 700;
  margin-bottom: 30px;
  line-height: 1.3;

  :deep(.gold) {
    background: linear-gradient(135deg, #fff 0%, #d4af37 50%, #fff 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
  }
}

.slide-subtitle {
  font-size: 1.5rem;
  font-weight: 300;
  opacity: 0.9;
}

.slide-desc {
  font-size: 1.3rem;
  line-height: 1.9;
  font-weight: 300;
  max-width: 900px;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 30px;
  margin-top: 40px;
  width: 100%;
  max-width: 1200px;

  .feature {
    padding: 30px;
    background: rgba(255,255,255,0.1);
    backdrop-filter: blur(10px);
    border-radius: 12px;
    border: 1px solid rgba(255,255,255,0.2);

    h3 {
      font-size: 1.3rem;
      margin-bottom: 10px;
    }

    p {
      font-size: 1rem;
      opacity: 0.9;
    }
  }
}

.nav-dots {
  position: fixed;
  right: 40px;
  top: 50%;
  transform: translateY(-50%);
  z-index: 3;
  display: flex;
  flex-direction: column;
  gap: 12px;

  span {
    width: 12px;
    height: 12px;
    border-radius: 50%;
    background: rgba(255,255,255,0.4);
    cursor: pointer;
    transition: all 0.3s;

    &.active {
      background: #d4af37;
      height: 40px;
      border-radius: 6px;
    }
  }
}

.nav-arrows {
  position: fixed;
  right: 40px;
  bottom: 40px;
  z-index: 3;
  display: flex;
  flex-direction: column;
  gap: 15px;

  button {
    width: 50px;
    height: 50px;
    border-radius: 50%;
    background: rgba(255,255,255,0.2);
    backdrop-filter: blur(10px);
    border: 1px solid rgba(255,255,255,0.3);
    color: #fff;
    font-size: 2rem;
    cursor: pointer;
    transition: all 0.3s;

    &:hover {
      background: rgba(212,175,55,0.8);
      transform: scale(1.1);
    }

    &:disabled {
      opacity: 0.3;
      cursor: not-allowed;
    }
  }
}

@media (max-width: 768px) {
  .slide-title {
    font-size: 2rem;
  }

  .slide-subtitle {
    font-size: 1.1rem;
  }

  .slide-desc {
    font-size: 1rem;
  }

  .features-grid {
    grid-template-columns: 1fr;
    gap: 20px;
  }

  .nav-dots {
    right: 20px;
  }

  .nav-arrows {
    right: 20px;
    bottom: 20px;

    button {
      width: 40px;
      height: 40px;
      font-size: 1.5rem;
    }
  }
}
</style>
