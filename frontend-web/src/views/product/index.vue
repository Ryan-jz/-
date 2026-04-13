<!--
  前台首页 - 参考 Bad Reichenhaller 经典垂直布局
-->
<template>
  <div class="home-container">
    <!-- 高对比度模式切换按钮 -->


    <PageHeader />

    <!-- 主内容区 -->
    <main class="main-content">
 
      <!-- Hero Banner -->
      <section class="hero-banner">
        <div class="container">
          <div class="hero-panel">
            <div class="hero-kicker">来自阿尔卑斯山的珍宝！</div>
            <h1 class="hero-title">
              Bad Reichenhaller
              <span class="hero-title-sub">阿尔卑斯山白金盐</span>
            </h1>

            <div class="hero-subline">
              <span class="hero-en">Treasures from the ALPS</span>
              <span class="hero-dot">·</span>
              <span class="hero-cn">大自然的恩赐</span>
            </div>

            <div class="hero-badges">
              <span class="hero-badge hero-badge-strong">Alpen 白金盐·德国 No.1</span>
              <span class="hero-badge">食盐中的钻石</span>
              <span class="hero-badge hero-badge-outline">纯净度 99.9% ｜ 越纯净越健康</span>
            </div>

            <div class="hero-grid">
              <div class="hero-copy">
                <p>源自阿尔卑斯山的白金（盐），完美生态，无价之宝！</p>
                <p>原始海洋蒸发 2.5 亿年前产生，食盐沉睡在阿尔卑斯山深处，千百年来，宝贵的盐随山泉水缓释流出，形成天然纯净珍贵的高山盐水。</p>
                <p>德国西南盐业股份公司是德国最大的制盐企业，德国市场领导品牌，500 年精湛技艺传承。</p>
                <p>贝希特斯加登盐矿，建于 1517 年，也是德国最重要的旅游景点之一。</p>
              </div>

              <div class="hero-facts">
                <div class="facts-lead">每个德国家庭都知道的优质产品！高品质健康生活从这里开始！</div>
                <ul class="facts">
                  <li>
                    <span class="fact-num">60%+</span>
                    <span class="fact-label">德国市场份额</span>
                  </li>
                  <li>
                    <span class="fact-num">90%+</span>
                    <span class="fact-label">品牌知名度</span>
                  </li>
                  <li>
                    <span class="fact-text">原始海洋蒸发储存在阿尔卑斯山的珍宝</span>
                  </li>
                </ul>

                <div class="essentials">
                  <div class="essentials-head">
                    <span class="essentials-title">阳光·空气·水·盐</span>
                    <span class="essentials-sub">每日健康必须</span>
                  </div>
                  <div class="essentials-body">食盐（NaCl）是人类第四个维持生命的要素，每日 5g，不可缺少，不可替代！</div>
                </div>
              </div>
            </div>

            <div class="hero-footer">
              <div class="hero-slogan">少吃盐·吃好盐！世界卫生组织倡导减盐丨您的“盐”值关系到全家健康！</div>
              <div class="hero-cta">把健康带回家！Alpen 白金盐·赋予您有滋有味的健康生活！</div>
            </div>
          </div>
        </div>
      </section>

      <!-- 阿尔卑斯盐产品板块 -->
      <section class="product-section-area">
        <div class="container">
          <ProductSection
            v-for="category in categories"
            :key="category.id"
            :category="category.name"
            :title="category.name"
            :description="category.description"
            :products="category.products"
            :columns="5"
            max-width="100%"
            spacing="60px"
          />
        </div>
      </section>

    
    </main>

 
  </div>
</template>

<script setup>
import { computed, ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import ProductSection from '@/components/ProductSection.vue'
import PageHeader from '@/components/PageHeader.vue'
import { getCategoryWithProducts } from '@/api/product'

const route = useRoute()
const categories = ref([])

const categoryId = computed(() => {
  const categoryParam = Array.isArray(route.query.category)
    ? route.query.category[0]
    : route.query.category
  const id = typeof categoryParam === 'string' && categoryParam !== '' ? Number(categoryParam) : NaN
  return Number.isFinite(id) ? id : undefined
})

let loadSeq = 0

const loadCategories = async (id) => {
  try {
    const seq = ++loadSeq

    const params = { status: 1 }
    if (Number.isFinite(id)) {
      params.category = id
    }

    const res = await getCategoryWithProducts(params)
    if (seq !== loadSeq) return
    if (res.data?.list) {
      categories.value = res.data.list.filter(item => item.id===categoryId.value)
    }
  } catch (error) {
    console.error('加载分类失败:', error)
  }
}

onMounted(() => {
  loadCategories(categoryId.value)
})
</script>


<style lang="scss" scoped>
.main-content {
  width: 100%;
}
// 容器
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

.home-container {

     background-image:url('@/assets/images/background.jpg');
         background-repeat: no-repeat;
    
     background-size:100% ;


}

// 高对比度切换按钮
.high-contrast-toggle {
  position: fixed;
  top: 10px;
  right: 10px;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background-color: rgba(0, 0, 0, 0.1);
  border: none;
  cursor: pointer;
  z-index: 1000;
  transition: all 0.3s;
  display: flex;
  align-items: center;
  justify-content: center;
  
  .icon {
    font-size: 20px;
  }
  
  &:hover {
    background-color: rgba(0, 0, 0, 0.2);
  }
  
  &.active {
    background-color: #000;
    color: #fff;
  }
}

// 页头
.header {

  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

// 顶部栏
.top-bar {

  padding: 2px 0;
  height:112px;
  display:flex;
  align-items:center;
  .container {
    display: flex;
    justify-content: center;
    align-items: center;
    height:100%;
  }
  
  .logo {
    height:100%;
    img {
      height: 100%;
      display: block;
    }
  }
  
  .top-bar-right {
    display: flex;
    align-items: center;
    gap: 30px;
  }
  
  .social-links {
    display: flex;
    gap: 15px;
    
    a {
      width: 32px;
      height: 32px;
      display: flex;
      align-items: center;
      justify-content: center;
      border-radius: 50%;
      background-color: #f0f0f0;
      color: #333;
      text-decoration: none;
      transition: all 0.3s;
      
      &:hover {
        background-color: #c41e3a;
        color: #fff;
      }
    }
  }
  
  .language-switcher {
    display: flex;
    align-items: center;
    gap: 8px;
    
    button {
      background: none;
      border: none;
      color: #666;
      cursor: pointer;
      font-size: 14px;
      transition: color 0.3s;
      
      &:hover,
      &.active {
        color: #c41e3a;
        font-weight: 600;
      }
    }
    
    .separator {
      color: #ccc;
    }
  }
}

// 主导航
.primary-navigation {
  position: relative;
  transition: all 0.3s;
  height:58px;
  border-top:2px solid #fff;
  border-bottom:3px solid #fff;
  display:flex;
  align-items:center;
  justify-content:center;
background: linear-gradient( 90deg, #92121B 0%, #D5061C 25%, #D5061C 75%,#92121B 100%);
  &.affixed {
    position: sticky;
    top: 0;
    z-index: 999;
    background-color: #fff;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  }
  
  .container {
    position: relative;
  }
  
  .menu-toggle {
    display: none;
    flex-direction: column;
    gap: 5px;
    background: none;
    border: none;
    cursor: pointer;
    padding: 10px;
    z-index: 10;
    
    span {
      width: 25px;
      height: 3px;
      background-color: #fff;
      transition: all 0.3s;
      display: block;
    }
    
    &:hover span {
      background-color: rgba(255, 255, 255, 0.8);
    }
  }
  
  .nav-menu {
    display: flex;
    justify-content: center;
    list-style: none;
    padding: 0;
    margin: 0;
    
    .nav-item {
      position: static;
      
      > a,
      .nav-link {
        display: block;
        padding: 20px 25px;
        color: #fff;
        text-decoration: none;
        font-size: 16px;
        font-weight: 500;
        transition: all 0.3s;
        cursor: pointer;
        
        &:hover {
          opacity: 0.8;
        }
        
        .arrow {
          font-size: 10px;
          margin-left: 5px;
          transition: transform 0.3s;
        }
      }
      
      &.has-dropdown {
        &:hover {
          .nav-link .arrow {
            transform: rotate(180deg);
          }
          
          .dropdown-menu-wrapper {
            opacity: 1;
            visibility: visible;
            transform: translateY(0);
          }
        }
        
        .dropdown-menu-wrapper {
          position: fixed;
          top: 170px;
          left: 0;
          right: 0;
          width: 100vw;
          opacity: 0;
          visibility: hidden;
          transform: translateY(-10px);
          transition: all 0.3s ease;
          z-index: 1000;
          
          .dropdown-menu {
            background-color: #fff;
            box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
            padding: 40px 0;
            
            .dropdown-container {
              max-width: 1400px;
              margin: 0 auto;
              padding: 0 40px;
            }
            
            .dropdown-content {
              display: grid;
              grid-template-columns: repeat(4, 1fr);
              gap: 40px;
            }
            
            .dropdown-column {
              .dropdown-title {
                font-size: 16px;
                font-weight: 600;
                color: #2b2b2e;
                margin: 0 0 20px 0;
                padding-bottom: 10px;
                border-bottom: 2px solid #c41e3a;
              }
              
              ul {
                list-style: none;
                padding: 0;
                margin: 0;
                
                li {
                  margin-bottom: 12px;
                  
                  a {
                    display: block;
                    padding: 8px 0;
                    color: #666;
                    text-decoration: none;
                    font-size: 14px;
                    transition: all 0.3s;
                    position: relative;
                    padding-left: 15px;
                    
                    &::before {
                      content: '';
                      position: absolute;
                      left: 0;
                      top: 50%;
                      transform: translateY(-50%);
                      width: 6px;
                      height: 6px;
                      background: #c41e3a;
                      border-radius: 50%;
                      opacity: 0;
                      transition: opacity 0.3s;
                    }
                    
                    &:hover {
                      color: #c41e3a;
                      padding-left: 20px;
                      
                      &::before {
                        opacity: 1;
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}

// Hero Banner
.hero-banner {
  padding-top: 130px;
  padding-bottom: 24px;
  color: #0e0e10;
  position: relative;

  .hero-panel {
    max-width: 1040px;
    margin: 0 auto;
    padding: 44px 48px;
    border-radius: 26px;
    background: rgba(255, 255, 255, 0.78);
    border: 1px solid rgba(0, 0, 0, 0.08);
    box-shadow:
      0 26px 55px rgba(18, 22, 28, 0.14),
      0 2px 0 rgba(255, 255, 255, 0.55) inset;
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    overflow: hidden;
    position: relative;

    &::before {
      content: '';
      position: absolute;
      inset: -40% -30% auto -30%;
      height: 280px;
      background:
        radial-gradient(closest-side, rgba(196, 30, 58, 0.16), rgba(196, 30, 58, 0) 72%),
        radial-gradient(closest-side, rgba(10, 90, 140, 0.12), rgba(10, 90, 140, 0) 70%);
      pointer-events: none;
      filter: blur(2px);
      opacity: 0.95;
    }
  }

  .hero-kicker {
    position: relative;
    font-size: 13px;
    letter-spacing: 0.18em;
    text-transform: uppercase;
    color: rgba(14, 14, 16, 0.62);
    margin-bottom: 14px;
  }

  .hero-title {
    position: relative;
    margin: 0;
    font-size: clamp(34px, 4.4vw, 56px);
    line-height: 1.06;
    font-weight: 760;
    letter-spacing: -0.02em;
    text-align: center;
  }

  .hero-title-sub {
    display: block;
    margin-top: 10px;
    font-size: clamp(18px, 2.35vw, 28px);
    font-weight: 650;
    letter-spacing: 0.08em;
    color: rgba(14, 14, 16, 0.86);
  }

  .hero-subline {
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 10px;
    margin-top: 14px;
    color: rgba(14, 14, 16, 0.62);
    font-size: 14px;
    letter-spacing: 0.06em;
  }

  .hero-en {
    font-weight: 650;
  }

  .hero-dot {
    opacity: 0.5;
  }

  .hero-badges {
    position: relative;
    display: flex;
    justify-content: center;
    flex-wrap: wrap;
    gap: 10px;
    margin-top: 18px;
  }

  .hero-badge {
    display: inline-flex;
    align-items: center;
    padding: 9px 12px;
    border-radius: 999px;
    font-size: 13px;
    line-height: 1;
    color: rgba(14, 14, 16, 0.76);
    background: rgba(255, 255, 255, 0.7);
    border: 1px solid rgba(0, 0, 0, 0.08);
  }

  .hero-badge-strong {
    background: rgba(196, 30, 58, 0.1);
    border-color: rgba(196, 30, 58, 0.22);
    color: #7b0d1a;
    font-weight: 700;
  }

  .hero-badge-outline {
    background: rgba(255, 255, 255, 0.45);
    color: rgba(14, 14, 16, 0.74);
  }

  .hero-grid {
    position: relative;
    margin-top: 26px;
    display: grid;
    grid-template-columns: 1.35fr 0.95fr;
    gap: 26px;
    align-items: start;
  }

  .hero-copy {
    p {
      margin: 0 0 14px;
      font-size: 15px;
      line-height: 1.85;
      color: rgba(14, 14, 16, 0.82);
    }
  }

  .hero-facts {
    padding: 18px 18px 16px;
    border-radius: 18px;
    background: rgba(255, 255, 255, 0.6);
    border: 1px solid rgba(0, 0, 0, 0.08);
  }

  .facts-lead {
    font-size: 14px;
    line-height: 1.6;
    font-weight: 650;
    color: rgba(14, 14, 16, 0.86);
    margin-bottom: 12px;
  }

  .facts {
    list-style: none;
    padding: 0;
    margin: 0;
    display: grid;
    gap: 10px;

    li {
      display: flex;
      align-items: baseline;
      gap: 10px;
      padding: 10px 12px;
      border-radius: 14px;
      background: rgba(255, 255, 255, 0.62);
      border: 1px solid rgba(0, 0, 0, 0.06);
    }
  }

  .fact-num {
    flex: none;
    font-size: 24px;
    font-weight: 780;
    letter-spacing: -0.02em;
    color: #c41e3a;
  }

  .fact-label,
  .fact-text {
    font-size: 13px;
    line-height: 1.55;
    color: rgba(14, 14, 16, 0.76);
  }

  .essentials {
    margin-top: 14px;
    padding: 14px 14px 12px;
    border-radius: 16px;
    background: linear-gradient(180deg, rgba(10, 90, 140, 0.09), rgba(10, 90, 140, 0.04));
    border: 1px solid rgba(10, 90, 140, 0.14);
  }

  .essentials-head {
    display: flex;
    flex-wrap: wrap;
    gap: 10px;
    align-items: baseline;
    margin-bottom: 8px;
  }

  .essentials-title {
    font-size: 14px;
    font-weight: 780;
    color: rgba(14, 14, 16, 0.86);
    letter-spacing: 0.08em;
  }

  .essentials-sub {
    font-size: 13px;
    color: rgba(14, 14, 16, 0.6);
  }

  .essentials-body {
    font-size: 13px;
    line-height: 1.65;
    color: rgba(14, 14, 16, 0.78);
  }

  .hero-footer {
    position: relative;
    margin-top: 22px;
    padding-top: 18px;
    border-top: 1px solid rgba(0, 0, 0, 0.08);
    display: grid;
    gap: 10px;
  }

  .hero-slogan {
    font-size: 14px;
    line-height: 1.7;
    color: rgba(14, 14, 16, 0.82);
  }

  .hero-cta {
    font-size: 15px;
    line-height: 1.65;
    font-weight: 780;
    color: rgba(14, 14, 16, 0.92);
  }
}

// 产品板块区域
.product-section-area {
  padding: 60px 0;

  
  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
  }
}

// 产品板块
.product-section {
  padding: 80px 0;
  
  &.alt-bg {
    background-color: #f8f8f8;
  }
  
  .anchor {
    display: block;
    position: relative;
    top: -100px;
    visibility: hidden;
  }
  
  .section-header {
    margin-bottom: 40px;
    
    h3 {
      font-size: 42px;
      color: #333;
      line-height: 1.3;
      
      .subtitle {
        display: block;
        font-size: 16px;
        color: #999;
        text-transform: uppercase;
        letter-spacing: 2px;
        margin-bottom: 10px;
        font-weight: 400;
      }
    }
  }
  
  .ce-bodytext {
    max-width: 800px;
    margin-bottom: 50px;
    
    p {
      font-size: 18px;
      line-height: 1.8;
      color: #666;
    }
  }
  
  .ce-gallery {
    display: grid;
    gap: 30px;
    
    &[data-ce-columns="3"] {
      grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    }
    
    &[data-ce-columns="4"] {
      grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    }
    
    .gallery-item {
      figure {
        margin: 0;
        background-color: #fff;
        border-radius: 8px;
        overflow: hidden;
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
        transition: all 0.3s;
        
        &:hover {
          transform: translateY(-8px);
          box-shadow: 0 8px 20px rgba(0, 0, 0, 0.15);
        }
        
        a {
          display: block;
          
          img {
            width: 100%;
            height: auto;
            display: block;
          }
        }
        
        figcaption {
          padding: 20px;
          text-align: center;
          font-size: 16px;
          color: #333;
          font-weight: 500;
        }
      }
    }
  }
}

// 可持续发展板块
.sustainability-section {
  padding: 80px 0;
  background: linear-gradient(135deg, #4a7c59 0%, #2d5a3d 100%);
  color: #fff;
  
  .section-header {
    text-align: center;
    margin-bottom: 60px;
    
    h2 {
      font-size: 42px;
      margin: 0;
    }
  }
  
  .sustainability-content {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
    gap: 40px;
  }
  
  .sustainability-item {
    text-align: center;
    padding: 30px;
    background-color: rgba(255, 255, 255, 0.1);
    border-radius: 8px;
    transition: all 0.3s;
    
    &:hover {
      background-color: rgba(255, 255, 255, 0.15);
      transform: translateY(-5px);
    }
    
    .icon {
      font-size: 48px;
      margin-bottom: 20px;
    }
    
    h4 {
      font-size: 22px;
      margin-bottom: 15px;
    }
    
    p {
      font-size: 16px;
      opacity: 0.9;
      line-height: 1.6;
    }
  }
}

// 页脚
.footer {
  background-color: #2c2c2c;
  color: #fff;
  padding: 60px 0 20px;
  
  .footer-content {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: 40px;
    margin-bottom: 40px;
  }
  
  .footer-section {
    h4 {
      font-size: 18px;
      margin-bottom: 20px;
      color: #fff;
    }
    
    ul {
      list-style: none;
      padding: 0;
      margin: 0;
      
      li {
        margin-bottom: 12px;
        
        a {
          color: #ccc;
          text-decoration: none;
          transition: color 0.3s;
          font-size: 14px;
          
          &:hover {
            color: #fff;
          }
        }
      }
    }
    
    .footer-social {
      display: flex;
      flex-direction: column;
      gap: 12px;
      
      a {
        color: #ccc;
        text-decoration: none;
        transition: color 0.3s;
        font-size: 14px;
        
        &:hover {
          color: #fff;
        }
      }
    }
  }
  
  .footer-bottom {
    border-top: 1px solid #444;
    padding-top: 30px;
    text-align: center;
    
    p {
      margin: 5px 0;
      color: #999;
      font-size: 14px;
      
      a {
        color: #999;
        text-decoration: none;
        transition: color 0.3s;
        
        &:hover {
          color: #fff;
        }
      }
    }
    
    .footer-group {
      font-size: 13px;
      color: #777;
    }
  }
}

// 响应式设计
@media (max-width: 768px) {
  // 容器在移动端减少内边距
  .container {
    padding: 0 15px;
  }

  // 顶部栏移动端优化
  .top-bar {
    height: auto;
    padding: 10px 0;
    
    .container {
      flex-direction: column;
      gap: 10px;
      align-items: center;
    }
    
    .logo {
      height: 60px;
      
      img {
        height: 100%;
      }
    }
    
    .top-bar-right {
      gap: 15px;
      
      .social-links {
        gap: 10px;
        
        a {
          width: 28px;
          height: 28px;
          font-size: 14px;
        }
      }
      
      .language-switcher {
        font-size: 12px;
      }
    }
  }

  // 主导航移动端优化
  .primary-navigation {
    height: auto;
    padding: 10px 0;
    
    .container {
      width: 100%;
    }
    
    .menu-toggle {
      display: flex;
      position: absolute;
      left: 15px;
      top: 50%;
      transform: translateY(-50%);
      z-index: 10;
      
      span {
        background-color: #fff;
      }
    }
    
    .nav-menu {
      position: absolute;
      top: 100%;
      left: 0;
      right: 0;
      background: linear-gradient(90deg, #92121B 0%, #D5061C 100%);
      flex-direction: column;
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
      max-height: 0;
      overflow: hidden;
      transition: max-height 0.3s;
      
      &.open {
        max-height: 500px;
      }
      
      .nav-item {
        border-bottom: 1px solid rgba(255, 255, 255, 0.1);
        
        > a {
          padding: 15px 20px;
          color: #fff;
          font-size: 14px;
        }
        
        .dropdown-menu {
          position: static;
          box-shadow: none;
          background-color: rgba(0, 0, 0, 0.2);
          
          li a {
            color: #fff;
            font-size: 13px;
            padding: 10px 30px;
          }
        }
      }
    }
  }

  // Hero Banner 移动端优化
  .hero-banner {
  
    padding-top: 110px;
    padding-bottom: 16px;

    .hero-panel {
      padding: 26px 18px;
      border-radius: 20px;
    }

    .hero-kicker {
      font-size: 12px;
      margin-bottom: 10px;
    }

    .hero-title {
      font-size: 32px;
    }

    .hero-title-sub {
      font-size: 18px;
      letter-spacing: 0.06em;
    }

    .hero-subline {
      font-size: 13px;
      gap: 8px;
    }

    .hero-badges {
      justify-content: flex-start;
    }

    .hero-grid {
      grid-template-columns: 1fr;
      gap: 16px;
    }

    .hero-copy p {
      font-size: 14px;
      margin-bottom: 12px;
    }

    .hero-facts {
      padding: 14px;
    }

    .facts li {
      padding: 9px 10px;
    }

    .fact-num {
      font-size: 22px;
    }

    .hero-cta {
      font-size: 14px;
    }
  }

  // 主内容区移动端优化
  .main-content {
    width: 100%;
  }
  
  // 产品板块区域移动端优化
  .product-section-area {
    padding: 40px 0;
    
    .container {
      padding: 0 15px;
    }
  }
  
  // 产品板块移动端优化
  .product-section {
    padding: 40px 0;
    
    .section-header {
      margin-bottom: 25px;
      
      h3 {
        font-size: 28px;
        
        .subtitle {
          font-size: 12px;
          margin-bottom: 8px;
        }
      }
    }
    
    .ce-bodytext {
      margin-bottom: 30px;
      
      p {
        font-size: 14px;
        line-height: 1.6;
      }
    }
    
    .ce-gallery {
      gap: 20px;
      
      &[data-ce-columns="3"],
      &[data-ce-columns="4"] {
        grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
      }
      
      .gallery-item {
        figure {
          figcaption {
            padding: 12px;
            font-size: 13px;
          }
        }
      }
    }
  }

  // 可持续发展板块移动端优化
  .sustainability-section {
    padding: 40px 0;
    
    .section-header {
      margin-bottom: 30px;
      
      h2 {
        font-size: 28px;
      }
    }
    
    .sustainability-content {
      grid-template-columns: 1fr;
      gap: 20px;
    }
    
    .sustainability-item {
      padding: 20px;
      
      .icon {
        font-size: 36px;
      }
      
      h4 {
        font-size: 18px;
      }
      
      p {
        font-size: 14px;
      }
    }
  }

  // 页脚移动端优化
  .footer {
    padding: 40px 0 20px;
    
    .footer-content {
      grid-template-columns: 1fr;
      gap: 25px;
    }
    
    .footer-section {
      h4 {
        font-size: 16px;
        margin-bottom: 15px;
      }
      
      ul li {
        margin-bottom: 8px;
        
        a {
          font-size: 13px;
        }
      }
      
      .footer-social {
        flex-direction: row;
        flex-wrap: wrap;
        gap: 15px;
      }
    }
    
    .footer-bottom {
      padding-top: 20px;
      
      p {
        font-size: 12px;
      }
    }
  }

  // 背景图片移动端优化
  .home-container {
    background-size: fill;
    // background-position: center;
  }
}

// 超小屏幕优化 (480px 以下)
@media (max-width: 480px) {
  .container {
    padding: 0 10px;
  }

  .top-bar {
    .logo {
      height: 50px;
    }
  }

  .hero-banner {
    padding-top: 104px;
    padding-bottom: 14px;

    .hero-panel {
      padding: 22px 14px;
      border-radius: 18px;
    }

    .hero-title {
      font-size: 28px;
      text-align: left;
    }

    .hero-title-sub {
      font-size: 16px;
      letter-spacing: 0.05em;
    }

    .hero-subline {
      justify-content: flex-start;
    }

    .hero-badges {
      justify-content: flex-start;
    }

    .hero-badge {
      font-size: 12px;
      padding: 8px 10px;
    }

    .fact-num {
      font-size: 20px;
    }
  }

  .product-section-area {
    padding: 30px 0;
    
    .container {
      padding: 0 10px;
    }
  }

  .product-section {
    padding: 30px 0;
    
    .section-header h3 {
      font-size: 24px;
    }
    
    .ce-gallery {
      gap: 15px;
      
      &[data-ce-columns="3"],
      &[data-ce-columns="4"] {
        grid-template-columns: 1fr;
      }
    }
  }

  .sustainability-section {
    padding: 30px 0;
    
    .section-header h2 {
      font-size: 24px;
    }
  }

  .footer {
    padding: 30px 0 15px;
  }
}

// 高对比度模式样式
:global(body.high-contrast) {
  .home-container {
    background-color: #000;
    color: #fff;
  }
  
  .header,
  .primary-navigation {
    background-color: #000 !important;
    border-bottom: 2px solid #fff;
  }
  
  .nav-menu a {
    color: #fff !important;
  }
  
  .product-section {
    background-color: #000 !important;
    color: #fff;
    
    h3, p {
      color: #fff !important;
    }
  }
}
</style>
