<template>
  <section class="sustainability">
    <div class="container">
      <h2 class="section-title">可持续与公益</h2>
      <div class="sustainability-grid">
        <div class="sustainability-card" v-for="(item, index) in items" :key="index">
          <div class="card-icon">{{ item.icon }}</div>
          <h3>{{ item.title }}</h3>
          <p>{{ item.desc }}</p>
        </div>
      </div>
      <div class="final-quote">
        <h3>把世界的辽阔，浓缩成一粒加冕的盐</h3>
        <div class="brand-name">Alpen 白金盐</div>
        <p class="tagline">让 2.5 亿年的纯净，继续为国王、王后，也为你我，守口如瓶地咸下去</p>
      </div>
    </div>
  </section>
</template>

<script setup>
const items = [
  {
    icon: '⚡',
    title: '皇室绿电',
    desc: '2025 起 100% 巴伐利亚水电驱动，年减碳 2,100 吨'
  },
  {
    icon: '♻️',
    title: '空瓶回收',
    desc: '6 个空瓶换 250 g 补充装，再制"皇冠盐勺"'
  },
  {
    icon: '❤️',
    title: '一克盐·一克爱',
    desc: '每售 1 罐，向云南山区捐 1 g 碘盐，累计 7.3 吨'
  }
]
</script>

<style scoped lang="scss">
.sustainability {
  padding: 120px 0;
  background: linear-gradient(135deg, #f8f9fa 0%, #fff 100%);
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
  margin-bottom: 80px;
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

.sustainability-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 40px;
  margin-bottom: 100px;
}

.sustainability-card {
  padding: 50px 30px;
  background: #fff;
  border-radius: 16px;
  text-align: center;
  box-shadow: 0 4px 20px rgba(0,0,0,0.08);
  transition: all 0.4s ease;

  &:hover {
    transform: translateY(-10px);
    box-shadow: 0 12px 40px rgba(0,0,0,0.15);
  }

  .card-icon {
    font-size: 4rem;
    margin-bottom: 25px;
  }

  h3 {
    font-size: 1.5rem;
    font-weight: 600;
    color: #1a1a2e;
    margin-bottom: 15px;
  }

  p {
    font-size: 1.05rem;
    line-height: 1.7;
    color: #666;
  }
}

.final-quote {
  text-align: center;
  padding: 80px 40px;
  background: linear-gradient(135deg, #1a1a2e 0%, #16213e 100%);
  border-radius: 20px;
  color: #fff;
  position: relative;
  overflow: hidden;

  &::before {
    content: '';
    position: absolute;
    top: -50%;
    left: -50%;
    width: 200%;
    height: 200%;
    background: radial-gradient(circle, rgba(212,175,55,0.1) 0%, transparent 70%);
    animation: rotate 20s linear infinite;
  }

  h3 {
    font-size: 2.5rem;
    font-weight: 700;
    margin-bottom: 20px;
    position: relative;
    z-index: 1;
  }

  p {
    font-size: 1.3rem;
    font-weight: 300;
    margin-bottom: 30px;
    position: relative;
    z-index: 1;
  }

  .brand-name {
    font-size: 3rem;
    font-weight: 700;
    background: linear-gradient(135deg, #fff 0%, #d4af37 50%, #fff 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    margin: 30px 0;
    position: relative;
    z-index: 1;
  }

  .tagline {
    font-size: 1.5rem;
    font-weight: 400;
    margin-top: 20px;
  }
}

@keyframes rotate {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

@media (max-width: 1024px) {
  .sustainability-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 768px) {
  .sustainability {
    padding: 80px 0;
  }

  .section-title {
    font-size: 2rem;
  }

  .final-quote {
    padding: 60px 30px;

    h3 {
      font-size: 1.8rem;
    }

    p {
      font-size: 1.1rem;
    }

    .brand-name {
      font-size: 2rem;
    }

    .tagline {
      font-size: 1.2rem;
    }
  }
}
</style>
