<template>
  <section class="scene">
    <div class="container">
      <h2 class="section-title">场景化应用</h2>
      <div class="scenes-grid">
        <div class="scene-card" v-for="(scene, index) in scenes" :key="index">
          <div class="scene-time">{{ scene.time }}</div>
          <div class="scene-content">
            <h3>{{ scene.title }}</h3>
            <p>{{ scene.desc }}</p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
const scenes = [
  {
    time: '早 7:00',
    title: '国王的"晨间仪式"',
    desc: '一杯 35℃ 淡盐温水，唤醒肠道蠕动'
  },
  {
    time: '午 12:30',
    title: '茜茜公主的"午后轻食"',
    desc: '0.3 g 撒在牛油果上，咸度恰好平衡单宁'
  },
  {
    time: '晚 19:00',
    title: '王储的"第一口辅盐"',
    desc: '宝宝辅食最后一道撒 2 粒，妈妈不再用舌尖"二次调味"'
  },
  {
    time: '深夜 23:00',
    title: '皇室骑师的"夜训配方"',
    desc: '健身后的自制电解质水，500 ml 水 + 0.5 g Alpen + 5 ml 柠檬汁，补钠不水肿'
  }
]
</script>

<style scoped lang="scss">
.scene {
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

.scenes-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 40px;
}

.scene-card {
  padding: 40px;
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0,0,0,0.08);
  transition: all 0.4s ease;
  border-left: 4px solid #d4af37;

  &:hover {
    transform: translateY(-5px);
    box-shadow: 0 12px 40px rgba(0,0,0,0.15);
  }
}

.scene-time {
  font-size: 2rem;
  font-weight: 700;
  color: #d4af37;
  margin-bottom: 20px;
}

.scene-content {
  h3 {
    font-size: 1.5rem;
    font-weight: 600;
    color: #1a1a2e;
    margin-bottom: 15px;
  }

  p {
    font-size: 1.1rem;
    line-height: 1.8;
    color: #666;
  }
}

@media (max-width: 768px) {
  .scene {
    padding: 80px 0;
  }

  .section-title {
    font-size: 2rem;
  }

  .scenes-grid {
    grid-template-columns: 1fr;
  }
}
</style>
