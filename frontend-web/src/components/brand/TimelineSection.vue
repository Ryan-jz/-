<template>
  <section class="timeline">
    <div class="container">
      <h2 class="section-title">品牌历史</h2>
      <div class="timeline-wrapper">
        <div class="timeline-line"></div>
        <div
          class="timeline-item"
          v-for="(item, index) in timeline"
          :key="index"
          :class="{ right: index % 2 === 1 }"
        >
          <div class="timeline-dot"></div>
          <div class="timeline-content">
            <div class="timeline-year">{{ item.year }}</div>
            <div class="timeline-image" v-if="item.image">
              <img :src="item.image" :alt="item.title" />
            </div>
            <h3>{{ item.title }}</h3>
            <p>{{ item.desc }}</p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import img01 from "@/assets/brand_histry/01.png";
import img02 from "@/assets/brand_histry/02.png";
import img03 from "@/assets/brand_histry/03.png";
import img04 from "@/assets/brand/mmexport1716774455983[1].jpg";
import img05 from "@/assets/brand_histry/05.jpg";
import img06 from "@/assets/brand_histry/06.jpeg";
import img07 from "@/assets/brand_histry/07.png";
import img08 from "@/assets/brand_histry/08.jpg";

const timeline = [
  {
    year: "公元前 2000 年",
    title: "发现盐水泉",
    desc: '凯尔特人成为最早的"专业"制盐者',
    image: img01,
  },
  {
    year: "1517 年",
    title: "皇家御令",
    desc: "德皇马克斯一世册封贝希特斯加登盐矿为皇家私产",
    image: img02,
  },
  {
    year: "1616-1619 年",
    title: "首条盐水管道",
    desc: "巴伐利亚公爵下令修建从赖兴哈尔到特劳恩施泰因的盐水管道",
    image: img03,
  },
  {
    year: "1890 年",
    title: "温泉小镇",
    desc: "赖兴哈尔正式更名为巴特赖兴哈尔",
    image: img04,
  },
  {
    year: "1950 年",
    title: "家庭包装",
    desc: "首次推出 500 克小包装食盐",
    image: img05,
  },
  {
    year: "1960 年",
    title: "碘盐问世",
    desc: "德国首次推出添加碘的食盐",
    image: img06,
  },
  {
    year: "2016 年",
    title: "品牌升级",
    desc: "MarkenSalz 更名为 AlpenSalz",
    image: img07,
  },
  {
    year: "2025 年",
    title: "高端系列",
    desc: "推出全新包装的阿尔卑斯之花盐系列",
    image: img08,
  },
];
</script>

<style scoped lang="scss">
.timeline {
  padding: 120px 0;
  background: linear-gradient(to bottom, #fff 0%, #f8f9fa 100%);
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
    content: "";
    position: absolute;
    bottom: -15px;
    left: 50%;
    transform: translateX(-50%);
    width: 80px;
    height: 3px;
    background: #d4af37;
  }
}

.timeline-wrapper {
  position: relative;
  padding: 40px 0;
}

.timeline-line {
  position: absolute;
  left: 50%;
  top: 0;
  bottom: 0;
  width: 2px;
  background: linear-gradient(to bottom, #d4af37, #f4e4b8, #d4af37);
  transform: translateX(-50%);
}

.timeline-item {
  position: relative;
  margin-bottom: 80px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 60px;
  align-items: center;

  &.right {
    .timeline-content {
      grid-column: 2;
      text-align: left;
    }

    .timeline-dot {
      left: auto;
      right: calc(50% - 10px);
    }
  }

  &:not(.right) {
    .timeline-content {
      grid-column: 1;
      text-align: right;
    }
  }
}

.timeline-dot {
  position: absolute;
  left: calc(50% - 10px);
  top: 50%;
  transform: translateY(-50%);
  width: 20px;
  height: 20px;
  background: #d4af37;
  border: 4px solid #fff;
  border-radius: 50%;
  box-shadow: 0 0 0 4px rgba(212, 175, 55, 0.2);
  z-index: 2;
  transition: all 0.3s ease;

  .timeline-item:hover & {
    transform: translateY(-50%) scale(1.3);
    box-shadow: 0 0 0 8px rgba(212, 175, 55, 0.3);
  }
}

.timeline-content {
  padding: 30px;
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  transition: all 0.4s ease;

  &:hover {
    transform: translateY(-5px);
    box-shadow: 0 8px 30px rgba(0, 0, 0, 0.12);
  }
}

.timeline-year {
  display: inline-block;
  padding: 8px 20px;
  background: linear-gradient(135deg, #d4af37 0%, #f4e4b8 100%);
  color: #1a1a2e;
  font-weight: 700;
  font-size: 0.9rem;
  border-radius: 20px;
  margin-bottom: 15px;
  letter-spacing: 1px;
}

.timeline-image {
  margin-bottom: 20px;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);

  img {
    width: 100%;
    height: auto;
    display: block;
    transition: transform 0.5s ease;
  }

  &:hover img {
    transform: scale(1.05);
  }
}

.timeline-content h3 {
  font-size: 1.5rem;
  font-weight: 600;
  color: #1a1a2e;
  margin-bottom: 10px;
}

.timeline-content p {
  font-size: 1rem;
  line-height: 1.7;
  color: #666;
}

@media (max-width: 768px) {
  .timeline {
    padding: 80px 0;
  }

  .section-title {
    font-size: 2rem;
  }

  .timeline-line {
    left: 20px;
  }

  .timeline-item {
    grid-template-columns: 1fr;
    gap: 0;
    padding-left: 60px;

    &.right .timeline-content,
    &:not(.right) .timeline-content {
      grid-column: 1;
      text-align: left;
    }

    &.right .timeline-dot,
    .timeline-dot {
      left: 10px;
      right: auto;
    }
  }
}
</style>
