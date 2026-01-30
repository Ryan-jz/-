<template>
  <section class="qa">
    <div class="container">
      <h2 class="section-title">常见问题</h2>
      <div class="qa-list">
        <div class="qa-item" v-for="(item, index) in qaList" :key="index">
          <div class="question">{{ item.q }}</div>
          <div class="answer">{{ item.a }}</div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
const qaList = [
  {
    q: '它是海盐吗？',
    a: '岩盐。2.5 亿年前海水蒸发后被岩层封存，远离现代海洋塑料与核污染话题区；当年德皇因"远离海岸阴谋"而指定贝希特斯加登为唯一皇家盐源。'
  },
  {
    q: '纯度高到底好在哪？',
    a: 'NaCl 99.9%，用量减少 8-12% 即可达到同样味觉峰值，帮助家庭无痛减盐；皇室御厨记录"用量减半，风味更醇"。'
  },
  {
    q: '与普通井盐、湖盐差价值不值？',
    a: '进口运输+关税+矿石开采成本确实高，但按每日 5 g 计算，家庭年均花费≈1.5 杯拿铁，换来的是"减盐不减味"的长期健康收益+皇室同款身份符号。'
  },
  {
    q: '如何快速验证正品？',
    a: '① 瓶身手机扫一扫官网溯源页；② 海关溯源码；③ 发货时可见海关报关清关。'
  }
]
</script>

<style scoped lang="scss">
.qa {
  padding: 120px 0;
  background: #fff;
}

.container {
  max-width: 1000px;
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

.qa-list {
  display: grid;
  gap: 30px;
}

.qa-item {
  padding: 40px;
  background: linear-gradient(135deg, #f8f9fa 0%, #fff 100%);
  border-radius: 12px;
  border-left: 4px solid #d4af37;
  transition: all 0.3s ease;

  &:hover {
    transform: translateX(10px);
    box-shadow: 0 8px 30px rgba(0,0,0,0.1);
  }
}

.question {
  font-size: 1.5rem;
  font-weight: 600;
  color: #1a1a2e;
  margin-bottom: 15px;

  &::before {
    content: 'Q: ';
    color: #d4af37;
    font-weight: 700;
  }
}

.answer {
  font-size: 1.1rem;
  line-height: 1.8;
  color: #555;
  padding-left: 30px;

  &::before {
    content: 'A: ';
    color: #d4af37;
    font-weight: 700;
    margin-left: -30px;
    margin-right: 10px;
  }
}

@media (max-width: 768px) {
  .qa {
    padding: 80px 0;
  }

  .section-title {
    font-size: 2rem;
  }

  .qa-item {
    padding: 30px 20px;
  }

  .question {
    font-size: 1.3rem;
  }

  .answer {
    font-size: 1rem;
    padding-left: 20px;
  }
}
</style>
