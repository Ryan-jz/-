<template>
  <div class="page-wrapper">
    <div class="fixed-bg">
         <section class="secondheader">
        <h1>极致享受，饱含<br>历史与爱意</h1>
        <p>是什么让"阿尔卑斯山之花"成为鉴赏家们钟爱的手工盐？它天然、产自当地、遵循传统工艺，并由工匠精心手工采摘。</p>
        <a href="#start" class="arrow-down">在这里了解制造过程</a>
      </section>
    </div>
    <div class="scroll-content">
      <PageHeader></PageHeader>

      <section class="headerbanner">
        <img src="@\assets\downloaded_images\headerbanner.jpg" alt="盐花">
      </section>

   

      <section class="infoslider" id="start">
        <div class="slider">
          <div class="slide" v-for="(step, i) in steps" :key="i" v-show="currentSlide === i">
            <img :src="step.img" :alt="step.title">
            <div class="content">
              <h3>步骤 {{ i + 1 }}：<br>{{ step.title }}</h3>
              <p>{{ step.text }}</p>
            </div>
          </div>
          <button @click="prevSlide" class="prev">‹</button>
          <button @click="nextSlide" class="next">›</button>
          <div class="dots">
            <span v-for="(_, i) in steps" :key="i" :class="{active: currentSlide === i}" @click="currentSlide = i"></span>
          </div>
        </div>
      </section>

      <section class="bigteaser">
        <img :src="saltFlowerImg" alt="盐花">
        <div class="preline">"阿尔卑斯山的海盐之花"</div>
        <h2>盐花</h2>
        <p>巴特赖兴哈尔的天然盐花质地和味道独特，是菜肴的完美点缀。我们称它们为"阿尔卑斯山的盐之花"，因为其历史悠久的制作工艺和精致的盐晶让人联想起地中海的盐田。</p>
        <!-- <div class="button-group">
          <button class="btn-primary">在线购买</button>
          <button class="btn-secondary">更多详情</button>
        </div> -->
      </section>

      <section class="products">
        <h2>新设计</h2>
        <div class="product-grid">
          <div class="product" v-for="p in products" :key="p.id" @click="activeProduct = p.id">
            <img :src="p.img" :alt="p.name">
            <div class="title">{{ p.name }}</div>
          </div>
        </div>
        <div class="product-detail" v-if="activeProduct">
          <div class="detail-content">
            <h3>{{ getProduct(activeProduct).name }}</h3>
            <p>{{ getProduct(activeProduct).desc }}</p>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import PageHeader from '../../components/PageHeader.vue'
import step1 from '@/assets/downloaded_images/herstellung-1.jpg'
import step2 from '@/assets/downloaded_images/herstellung-2.jpg'
import step3 from '@/assets/downloaded_images/herstellung-3.jpg'
import step4 from '@/assets/downloaded_images/herstellung-4.jpg'
// import productAlpine from '@/assets/downloaded_images/product-alpenblueten.png'
// import productHerbs from '@/assets/downloaded_images/product-bergkraeuter.png'
// import productSnack from '@/assets/downloaded_images/product-brotzeit.png'
// import productSteak from '@/assets/downloaded_images/product-steak-grill.png'
// import productChili from '@/assets/downloaded_images/product-chili.png'
import saltFlowerImg from '@/assets/downloaded_images/product-salzblueten.png'

const currentSlide = ref(0)
const activeProduct = ref(null)

const steps = ref([
  { title: '天然盐水', text: '距今已有2.5亿年历史，至今仍是人们津津乐道的话题：远古海洋。', img: step1 },
  { title: '盐晶体', text: '现在，手工制作在我们工厂开始了。', img: step2 },
  { title: '略读', text: '撇除盐粒的过程也遵循着几个世纪以来的传统。', img: step3 },
  { title: '自然风干', text: '天然高山盐被小心翼翼地铺成薄薄的一层。', img: step4 }
])

const products = ref([
  // { id: 'alpenblueten', name: '高山花卉', img: productAlpine, desc: '以万寿菊和接骨木花为主料' },
  // { id: 'bergkraeuter', name: '山地草药', img: productHerbs, desc: '精选有机香草和香料' },
  // { id: 'brotzeit', name: '小吃', img: productSnack, desc: '巴伐利亚特色香料' },
  // { id: 'steakgrill', name: '牛排烧烤', img: productSteak, desc: '略带辛辣的有机香料' },
  // { id: 'chili', name: '辣椒', img: productChili, desc: '富含有机辣椒片' }
])


const prevSlide = () => {
  currentSlide.value = currentSlide.value > 0 ? currentSlide.value - 1 : steps.value.length - 1
}

const nextSlide = () => {
  currentSlide.value = (currentSlide.value + 1) % steps.value.length
}

const getProduct = (id) => products.value.find(p => p.id === id)
</script>

<style scoped lang="scss">
.page-wrapper {
  position: relative;
  min-height: 100vh;
  color: #334046;
}

.fixed-bg {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: url('@/assets/downloaded_images/berge.jpg') center/fill no-repeat;
  z-index: -1;
}

.scroll-content {
  position: relative;
  z-index: 1;

}

.headerbanner img {
  width: 100%;
  display: block;
}

.secondheader {
 
  padding: 100px 20px;
  text-align: center;
  color: #fff;

  h1 {
    font-size: 48px;
    font-weight: 300;
    margin-bottom: 30px;
  }

  p {
    font-size: 18px;
    max-width: 800px;
    margin: 0 auto 40px;
  }

  .arrow-down {
    color: #fff;
    text-decoration: none;
    display: inline-block;
    padding: 15px 30px;
    border: 2px solid #fff;
    border-radius: 30px;
    transition: all 0.3s;

    &:hover {
      background: #fff;
      color: #124583;
    }
  }
}

.infoslider {
  padding: 80px 20px;
  background-image: url('@/assets/downloaded_images/bg-infoslider.png');
  color: #fff;
   margin-top: 80vh;
  .slider {
    max-width: 1200px;
    margin: 0 auto;
    position: relative;
        padding: 340px 0 180px 0;
        color: #124583;
    .slide {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 60px;
      align-items: center;

      img {
        width: 100%;
        border-radius: 8px;
      }

      h3 {
        font-size: 32px;
        margin-bottom: 20px;
      }

      p {
        font-size: 18px;
        line-height: 1.8;
      }
    }

    .prev, .next {
      position: absolute;
      top: 50%;
      transform: translateY(-50%);
      background: rgba(255,255,255,0.9);
      border: none;
      width: 50px;
      height: 50px;
      border-radius: 50%;
      font-size: 32px;
      cursor: pointer;
      color: #124583;
    }

    .prev { left: -70px; }
    .next { right: -70px; }

    .dots {
      text-align: center;
      margin-top: 40px;

      span {
        display: inline-block;
        width: 12px;
        height: 12px;
        border-radius: 50%;
        background: rgba(255,255,255,0.5);
        margin: 0 5px;
        cursor: pointer;

        &.active {
          background: #fff;
        }
      }
    }
  }
}

.bigteaser {
  padding: 100px 20px;
  text-align: center;
  background: url('@/assets/downloaded_images/bg-blue.jpg') center/fill;
  color: #fff;
  position: relative;
  height: 100vh;
  &::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(18, 69, 131, 0.7);
    z-index: 0;
  }

  > * {
    position: relative;
    z-index: 1;
  }

  img {
    max-width: 350px;
    margin: 0 auto 40px;
    filter: drop-shadow(0 10px 30px rgba(0,0,0,0.3));
  }

  .preline {
    font-size: 20px;
    margin-bottom: 15px;
    font-style: italic;
    opacity: 0.9;
  }

  h2 {
    font-size: 56px;
    font-weight: 300;
    margin-bottom: 30px;
    letter-spacing: 2px;
  }

  p {
    font-size: 18px;
    max-width: 900px;
    margin: 0 auto 40px;
    line-height: 1.8;
  }

  .button-group {
    display: flex;
    gap: 20px;
    justify-content: center;

    button {
      padding: 15px 40px;
      font-size: 16px;
      border: 2px solid #fff;
      border-radius: 30px;
      cursor: pointer;
      transition: all 0.3s;
      font-weight: 500;
    }

    .btn-primary {
      background: #fff;
      color: #124583;

      &:hover {
        background: transparent;
        color: #fff;
      }
    }

    .btn-secondary {
      background: transparent;
      color: #fff;

      &:hover {
        background: #fff;
        color: #124583;
      }
    }
  }
}

.products {
  padding: 80px 20px;
  max-width: 1400px;
  margin: 0 auto;

  h2 {
    text-align: center;
    font-size: 42px;
    margin-bottom: 60px;
    color: #124583;
  }

  .product-grid {
    display: grid;
    grid-template-columns: repeat(5, 1fr);
    gap: 30px;
    margin-bottom: 60px;
  }

  .product {
    text-align: center;
    cursor: pointer;
    transition: transform 0.3s;

    &:hover {
      transform: translateY(-10px);
    }

    img {
      width: 100%;
      margin-bottom: 15px;
    }

    .title {
      font-size: 18px;
      font-weight: 600;
    }
  }

  .product-detail {
    padding: 60px;
    background: #f8f9fa;
    border-radius: 8px;

    h3 {
      font-size: 32px;
      margin-bottom: 20px;
    }

    p {
      font-size: 18px;
      line-height: 1.8;
    }
  }
}

@media (max-width: 768px) {
  .secondheader h1 {
    font-size: 32px;
  }

  .infoslider .slider .slide {
    grid-template-columns: 1fr;
  }

  .products .product-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>
