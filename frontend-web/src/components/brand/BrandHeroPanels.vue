<template>
  <section class="brand-hero-panels">
    <div class="brand-hero-panels__container">
    <div
      v-for="(panel, index) in panels"
      :key="panel.id"
      :class="['brand-hero-panel', { 'brand-hero-panel--reverse': index % 2 === 1 }]"
    >
      <div class="brand-hero-panel__text">
        <h2
          class="brand-hero-panel__title"
          v-html="panel.title"
        ></h2>
        <p
          v-if="panel.subtitle"
          class="brand-hero-panel__subtitle"
        >
          {{ panel.subtitle }}
        </p>

        <!-- 纯文案型模块 -->
        <p
          v-if="panel.desc"
          class="brand-hero-panel__paragraph brand-hero-panel__paragraph--rich"
          v-html="panel.desc"
        ></p>

        <!-- 卖点 / 权威 / 公益 列表 -->
        <div
          v-if="panel.features"
          class="brand-hero-panel__features"
        >
          <div
            class="brand-hero-panel__feature-card"
            v-for="(item, i) in panel.features"
            :key="i"
          >
            <h3 class="brand-hero-panel__feature-title">
              {{ item.title }}
            </h3>
            <p class="brand-hero-panel__feature-desc">
              {{ item.desc }}
            </p>
          </div>
        </div>
      </div>

      <div class="brand-hero-panel__image">
          <img :src="panel.bg" :alt="panel.plainTitle" />
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import bg1 from '@/assets/brand/1_01 (1).jpg'
import bg2 from '@/assets/brand/1_02.jpg'
import bg3 from '@/assets/brand/1_02(01).jpg'
import bg4 from '@/assets/brand/1_01(4).jpg'
import bg5 from '@/assets/brand/1_05.jpg'
import bg6 from '@/assets/brand/1_06.jpg'

const panels = [
  {
    id: 'slogan',
    bg: bg1,
    title: '阳光·空气·水·盐<br><span class="gold">Alpen 白金盐</span><br>德国 No.1 岩盐品牌',
    plainTitle: '阳光·空气·水·盐 Alpen 白金盐 德国 No.1 岩盐品牌',
    subtitle: '食盐中的钻石，皇冠级纯净'
  },
  {
    id: 'royal',
    bg: bg2,
    title: '皇室故事',
    plainTitle: '皇室故事',
    desc: '1517 年，德皇马克斯一世挥笔御令，<br>将贝希特斯加登盐矿册封为“Königliche Saline”——皇家私产。<br>自此，白金岩盐只供宫廷餐桌、主教盛宴与凯旋军宴，<br>被欧洲史家称作“可食用的王冠”。'
  },
  {
    id: 'core',
    bg: bg3,
    title: '核心卖点',
    plainTitle: '核心卖点',
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
    id: 'senses',
    bg: bg4,
    title: '感官大片',
    plainTitle: '感官大片',
    desc: '原始海洋蒸发，2.5 亿年阿尔卑斯岩层封存，<br>远离现代海洋塑料与微塑料污染，<br>像给地球按下“干净存档”。<br><br>倒在掌心，碎雪般晶莹；<br>轻抿舌尖，矿物甘甜味层层晕开；<br>这就是皇室御厨笔记里的“雪吻之咸”。'
  },
  {
    id: 'authority',
    bg: bg5,
    title: '权威背书',
    plainTitle: '权威背书',
    features: [
      { title: 'WHO 2023 减盐模型', desc: '高纯低杂质盐，家庭钠摄入 ↓12%' },
      { title: '德国 BGR 报告', desc: '杂质密度仅为全球均值 1/5' },
      { title: '维也纳宫廷菜单 1812', desc: '御膳房唯一指定' }
    ]
  },
  {
    id: 'sustainability',
    bg: bg6,
    title: '可持续与公益',
    plainTitle: '可持续与公益',
    features: [
      { title: '⚡ 皇室绿电', desc: '2025 起 100% 水电驱动，年减碳 2,100 吨' },
      { title: '♻️ 空瓶回收', desc: '6 个空瓶换 250 g 补充装' },
      { title: '❤️ 一克盐·一克爱', desc: '每售 1 罐，向山区捐 1 g 碘盐' }
    ]
  }
]
</script>

<style scoped lang="scss">
.brand-hero-panels {

  margin: 0 auto;
  padding: 0 20px;

  overflow-y: auto;
  background-size: 100%;
  background-image: url('@/assets/images/background.jpg');
  background-repeat: no-repeat;

}
.brand-hero-panels__container{
    max-width: 1200px;
    margin: 0 auto;

}
.brand-hero-panel {
  display: flex;
  align-items: center;
  gap: 48px;
  min-height: 100vh;
  scroll-snap-align: start;

  &--reverse {
    flex-direction: row-reverse;
  }

  &__text {
    flex: 1 1 50%;
  }

  &__image {
    flex: 1 1 50%;
    text-align: center;

    img {
      max-width: 100%;
      height: auto;
      display: inline-block;
    }
  }

  &__title {
    font-size: 32px;
    font-weight: 700;
    margin-bottom: 12px;
    color: #2b2b2e;

    :deep(.gold) {
      background: linear-gradient(135deg, #fff 0%, #d4af37 50%, #fff 100%);
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
    }
  }

  &__subtitle {
    font-size: 20px;
    font-weight: 500;
    margin-bottom: 20px;
    color: #666;
  }

  &__paragraph {
    font-size: 16px;
    line-height: 1.8;
    color: #444;
    margin-bottom: 10px;

    &--rich {
      :deep(br) {
        content: '';
      }
    }
  }

  &__features {
    display: grid;
    grid-template-columns: repeat(3, minmax(0, 1fr));
    gap: 20px;
    margin-top: 24px;
  }

  &__feature-card {
    padding: 18px 16px;
    background: #fff;
    border-radius: 12px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
    border: 1px solid rgba(0, 0, 0, 0.02);
  }

  &__feature-title {
    font-size: 15px;
    font-weight: 600;
    margin-bottom: 8px;
    color: #2b2b2e;
  }

  &__feature-desc {
    font-size: 14px;
    line-height: 1.7;
    color: #555;
  }
}

@media (max-width: 1024px) {
  .brand-hero-panel__features {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (max-width: 768px) {
  .brand-hero-panel {
    flex-direction: column;

    &--reverse {
      flex-direction: column;
    }
  }

  .brand-hero-panel__features {
    grid-template-columns: 1fr;
  }
}
</style>

