<template>
  <div class="alternative-background layout-4">
    <h2>
      <span class="subtitle">与……合作</span><br>
      贝希特斯加登国家公园
    </h2>
    <div class="ce-textpic ce-right ce-intext ce-nowrap layout--3">
      <div class="ce-gallery" data-ce-columns="1" data-ce-images="1">
        <div class="ce-column">
          <figure class="image">
            <img 
              loading="lazy" 
              alt="贝希特斯加登国家公园标志" 
              :src="nationalParkLogoImg" 
              width="318" 
              height="66"
            >
          </figure>
        </div>
      </div>
      <div class="ce-bodytext">
        <p>
          保护我们珍贵的自然环境有很多方法。因为我们格外珍视自己的家园，所以我们非常乐意支持贝希特斯加登国家公园。它是德国唯一一个与阿尔卑斯山脉接壤的国家公园。除了自然保护之外，它的主要任务之一是保护濒危动物，包括金雕。一项保护这些雄伟鸟类的项目于25年前启动。该项目包括研究它们的繁殖行为、监测它们的后代以及确保鸟类的安全。
        </p>
        <p>
          我们阿尔卑斯山脉系列产品的部分销售收入将用于支持这个项目。例如，部分收益用于购买观景镜，以便更好地观测树线以上的"天空之王"。如需了解更多信息或预订导览游，请访问
          <a href="http://www.nationalpark-berchtesgaden.de" target="_blank" rel="noreferrer">
            www.nationalpark-berchtesgaden.de
          </a>
        </p>
      </div>
    </div>

    <ImageSlideshow :slides="slides" />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import ImageSlideshow from './ImageSlideshow.vue'
import nationalParkLogoImg from '@/assets/downloaded_images/08-kooperation-nationalpark-berchtesgaden-logo-nationalpark-berchtesgarden.svg'
import img1 from '@/assets/downloaded_images/08-kooperation-nationalpark-berchtesgaden-img-1.jpg'
import img2 from '@/assets/downloaded_images/08-kooperation-nationalpark-berchtesgaden-img-2.jpg'
import img3 from '@/assets/downloaded_images/08-kooperation-nationalpark-berchtesgaden-img-3.jpg'
import img4 from '@/assets/downloaded_images/08-kooperation-nationalpark-berchtesgaden-img-4.jpg'
import img5 from '@/assets/downloaded_images/08-kooperation-nationalpark-berchtesgaden-img-5.jpg'

const slides = ref([
  { imgSrc: img1, alt: '鹰头', title: '鹰头' },
  { imgSrc: img2, alt: '摄影师', title: '摄影师' },
  { imgSrc: img3, alt: '山脉前的雄鹰', title: '山脉前的雄鹰' },
  { imgSrc: img4, alt: '山地景观', title: '山地景观' },
  { imgSrc: img5, alt: '鹰', title: '鹰' }
])
</script>

<style scoped lang="scss">
.alternative-background.layout-4 {
  background: #fafafa;
  padding: 80px 0;

  h2 {
    font-size: 42px;
    font-weight: 300;
    line-height: 1.3;
    margin: 0 0 60px;
    color: #2b2b2e;
    text-align: center;

    .subtitle {
      display: block;
      font-size: 16px;
      color: #c41e3a;
      text-transform: uppercase;
      letter-spacing: 2px;
      margin-bottom: 12px;
      font-weight: 600;
    }
  }
}

.ce-textpic {
  padding: 0 20px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 60px;
  align-items: center;
  max-width: 1400px;
  margin: 0 auto 60px;

  &.ce-right .ce-gallery {
    order: 2;
  }

  .ce-bodytext {
    p {
      font-size: 18px;
      line-height: 1.8;
      color: #666;
      margin: 0 0 20px;

      &:last-child {
        margin-bottom: 0;
      }
    }

    a {
      color: #c41e3a;
      text-decoration: none;
      transition: opacity 0.3s;

      &:hover {
        opacity: 0.8;
      }
    }
  }

  img {
    max-width: 100%;
    height: auto;
  }
}

@media (max-width: 1024px) {
  .ce-textpic {
    grid-template-columns: 1fr;
    gap: 40px;

    .ce-gallery {
      order: 0 !important;
    }
  }
}

@media (max-width: 768px) {
  .alternative-background.layout-4 {
    padding: 60px 0;

    h2 {
      font-size: 28px;
      margin-bottom: 40px;

      .subtitle {
        font-size: 13px;
      }
    }
  }

  .ce-textpic {
    padding: 0 15px;

    .ce-bodytext p {
      font-size: 16px;
    }
  }
}
</style>


<style scoped lang="scss">
img {
  transition: transform 0.6s ease-out;
}
</style>
