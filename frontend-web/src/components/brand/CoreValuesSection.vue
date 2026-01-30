<template>
  <section class="core-values">
    <div class="container">
      <h2 class="section-title">核心卖点</h2>
      <div class="hero-image">
        <img src="@/assets/brand/20250612白金盐_袋装盐1-5_低钠盐6-8(1).jpg" alt="产品展示" />
      </div>
      <div class="values-grid">
        <div class="value-card" v-for="(value, index) in values" :key="index">
          <div class="value-icon">{{ value.icon }}</div>
          <h3>{{ value.title }}</h3>
          <p>{{ value.desc }}</p>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
const values = [
  {
    icon: '👑',
    title: '皇室专供矿脉',
    desc: '500 年不间断开采，德国仅存"活盐矿"'
  },
  {
    icon: '💎',
    title: '99.9% 极致纯净',
    desc: '入口 0.6 秒崩解，减盐 8% 不减味'
  },
  {
    icon: '🌿',
    title: '0 添加哲学',
    desc: '无麸质、无重金属、无核污染，孕妇宝宝可安心食用'
  },
  {
    icon: '☀️',
    title: '生命第四元素',
    desc: '阳光、空气、水、盐，每日不可替代'
  },
  {
    icon: '🔍',
    title: '皇冠溯源',
    desc: '瓶身 NFC 一触即达皇室档案，真伪立判'
  },
  {
    icon: '💰',
    title: '百姓价享御用',
    desc: '平均 1 元 / 天，把国王的"味"搬进自家厨房'
  }
]
</script>

<style scoped lang="scss">
.core-values {
  padding: 120px 0;
  background: #fff;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

.section-title {
  font-size: 3rem;
  font-weight: 700;
  color: #1a1a2e;
  text-align: center;
  margin-bottom: 60px;
  position: relative;

  &::after {
    content: '';
    position: absolute;
    bottom: -15px;
    left: 50%;
    transform: translateX(-50%);
    width: 80px;
    height: 3px;
    background: #d4af37;
  }
}

.values-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 40px;
}

.value-card {
  padding: 50px 30px;
  background: linear-gradient(135deg, #f8f9fa 0%, #fff 100%);
  border-radius: 16px;
  text-align: center;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  border: 2px solid transparent;
  position: relative;
  overflow: hidden;

  &::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: linear-gradient(135deg, #d4af37 0%, #f4e4b8 100%);
    opacity: 0;
    transition: opacity 0.4s ease;
  }

  &:hover {
    transform: translateY(-10px);
    border-color: #d4af37;
    box-shadow: 0 20px 40px rgba(212, 175, 55, 0.2);

    &::before {
      opacity: 0.05;
    }

    .value-icon {
      transform: scale(1.2) rotate(5deg);
    }
  }
}

.value-icon {
  font-size: 4rem;
  margin-bottom: 20px;
  transition: transform 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  z-index: 1;
}

.value-card h3 {
  font-size: 1.4rem;
  font-weight: 600;
  color: #1a1a2e;
  margin-bottom: 15px;
  position: relative;
  z-index: 1;
}

.value-card p {
  font-size: 1rem;
  line-height: 1.7;
  color: #666;
  position: relative;
  z-index: 1;
}

@media (max-width: 1024px) {
  .values-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 768px) {
  .core-values {
    padding: 80px 0;
  }

  .section-title {
    font-size: 2rem;
    margin-bottom: 50px;
  }

  .values-grid {
    grid-template-columns: 1fr;
    gap: 30px;
  }

  .value-card {
    padding: 40px 25px;
  }
}
</style>
