<template>
  <section class="royal-story">
    <div class="container">
      <div class="story-header">
        <h2 class="section-title">皇室传承</h2>
        <p class="section-subtitle">从御膳房到餐桌：一粒"皇盐"的 2.5 亿年加冕之路</p>
      </div>

      <div class="story-content">
        <div class="story-item" v-for="(chapter, index) in chapters" :key="index">
          <div class="story-year">{{ chapter.year }}</div>
          <div class="story-body">
            <img v-if="chapter.image" :src="chapter.image" :alt="chapter.title" />
            <h3>{{ chapter.title }}</h3>
            <p>{{ chapter.content }}</p>
          </div>
        </div>
      </div>

      <div class="royal-quote">
        <div class="quote-mark">"</div>
        <p>1517 年，德皇马克斯一世挥笔御令，将贝希特斯加登盐矿册封为"Königliche Saline"——皇家私产。自此，白金岩盐只供宫廷餐桌、主教盛宴与凯旋军宴，被欧洲史家称作"可食用的王冠"。</p>
        <div class="quote-mark">"</div>
      </div>
    </div>
  </section>
</template>

<script setup>
import img1 from '@/assets/brand/2017-12-05 090614.jpg'
import img2 from '@/assets/brand/2017-12-05 090626.jpg'
import img3 from '@/assets/brand/微信图片_20250408203638.jpg'
import img4 from '@/assets/brand/2018-08-14 162039.jpg'

const chapters = [
  {
    year: '2.5 亿年前',
    title: '原始海洋',
    content: '蒸发、挤压、隆起，海水被时间浓缩成 1800 米厚的盐层，与阿尔卑斯山一起长高。没有塑料微粒，没有工业排放和核污染——这是地球留给人类的"干净存档"。',
    image: img1
  },
  {
    year: '1517 年',
    title: '德皇马克斯一世签发"白色黄金"御令',
    content: '皇家纹章盖在羊皮纸上，贝希特斯加登盐矿成为"Königliche Saline"——专属皇家御膳房。盐砖被马拉雪橇送往慕尼黑、维也纳的宫廷，被称作"可以品尝的钻石"。当时，1 斤矿盐 = 3 斤黑胡椒的市价，只有皇室与主教餐桌才有资格享用。',
    image: img2
  },
  {
    year: '1706 年',
    title: '茜茜公主的"白金盐"礼盒',
    content: '巴伐利亚国王路德维希二世为表姐茜茜公主定制"白金盐"婚嫁礼盒，被维也纳宫廷女官记入《御膳笔记》："入口即化，喉间回甘，比匈牙利埃斯泰尔哈吉家族的珍珠项链更珍贵。"',
    image: img3
  },
  {
    year: '今日',
    title: '每秒 0.5 米的皇家电梯',
    content: '矿工仍沿用手工预选，剔除杂色矿脉；电梯 72 秒直达井下 650 米，恒温 24℃、湿度 45%，天然抑菌——这是德国仅存的可参观"活盐矿"。出井前，每一粒盐都要经过"皇盐仪式"。',
    image: img4
  },
  {
    year: '你家厨房',
    title: '1 元钱／天的"现代御膳"',
    content: '500 年采矿史、99.9% 纯度、德国工匠精神，全部浓缩成 1 天 1 元的成本——把"皇室特供"变成"寻常百姓的健康底味"。今天，你舀起的那一勺，正是昨日御厨呈给国王的同一座山、同一层矿脉。',
    image: img1
  }
]
</script>

<style scoped lang="scss">
.royal-story {
  padding: 120px 0;
  background: linear-gradient(to bottom, #f8f9fa 0%, #fff 100%);
  position: relative;

  &::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    height: 1px;
    background: linear-gradient(to right, transparent, #d4af37, transparent);
  }
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

.story-header {
  text-align: center;
  margin-bottom: 80px;
}

.section-title {
  font-size: 3rem;
  font-weight: 700;
  color: #1a1a2e;
  margin-bottom: 20px;
  position: relative;
  display: inline-block;

  &::after {
    content: '';
    position: absolute;
    bottom: -10px;
    left: 50%;
    transform: translateX(-50%);
    width: 60px;
    height: 3px;
    background: #d4af37;
  }
}

.section-subtitle {
  font-size: 1.2rem;
  color: #666;
  font-weight: 300;
  margin-top: 30px;
}

.story-content {
  display: grid;
  gap: 60px;
  margin-bottom: 80px;
}

.story-item {
  display: grid;
  grid-template-columns: 200px 1fr;
  gap: 40px;
  align-items: start;
  padding: 40px;
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0,0,0,0.08);
  transition: all 0.4s ease;

  &:hover {
    transform: translateY(-5px);
    box-shadow: 0 8px 30px rgba(0,0,0,0.12);
  }
}

.story-year {
  font-size: 1.5rem;
  font-weight: 700;
  color: #d4af37;
  padding: 20px;
  background: linear-gradient(135deg, #f8f9fa 0%, #fff 100%);
  border-radius: 8px;
  text-align: center;
  border: 2px solid #d4af37;
}

.story-body {
  img {
    width: 100%;
    height: auto;
    border-radius: 8px;
    margin-bottom: 20px;
  }

  h3 {
    font-size: 1.5rem;
    font-weight: 600;
    color: #1a1a2e;
    margin-bottom: 15px;
  }

  p {
    font-size: 1.05rem;
    line-height: 1.8;
    color: #555;
  }
}

.royal-quote {
  position: relative;
  padding: 60px 80px;
  background: linear-gradient(135deg, #1a1a2e 0%, #16213e 100%);
  border-radius: 12px;
  color: #fff;
  text-align: center;
  margin-top: 60px;

  .quote-mark {
    font-size: 6rem;
    font-family: Georgia, serif;
    color: #d4af37;
    opacity: 0.3;
    position: absolute;
    line-height: 1;

    &:first-child {
      top: 20px;
      left: 40px;
    }

    &:last-child {
      bottom: 20px;
      right: 40px;
      transform: rotate(180deg);
    }
  }

  p {
    font-size: 1.3rem;
    line-height: 1.9;
    font-weight: 300;
    position: relative;
    z-index: 1;
  }
}

@media (max-width: 768px) {
  .royal-story {
    padding: 80px 0;
  }

  .section-title {
    font-size: 2rem;
  }

  .story-item {
    grid-template-columns: 1fr;
    gap: 20px;
    padding: 30px 20px;
  }

  .story-year {
    font-size: 1.2rem;
    padding: 15px;
  }

  .royal-quote {
    padding: 40px 30px;

    p {
      font-size: 1.1rem;
    }

    .quote-mark {
      font-size: 4rem;

      &:first-child {
        top: 10px;
        left: 20px;
      }

      &:last-child {
        bottom: 10px;
        right: 20px;
      }
    }
  }
}
</style>
