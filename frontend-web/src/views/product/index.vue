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
          <div class="title">最精妙的盐渍和调味</div>
          <div class="hero-content">
            <p class="subtitle">好东西需要时间：我们珍贵的阿尔卑斯山盐形成于大约2.5亿年前原始海洋蒸发之时。自那时起，这些盐就一直受到保护，静静地躺在巴伐利亚阿尔卑斯山基岩深处。这真是大自然的馈赠！</p>
            <p class="subtitle">我们从纯净的阿尔卑斯山盐水中提取珍贵的盐。这种盐水是由清澈的山泉水缓慢溶解岩石中的盐分而形成的。它是巴特赖兴哈尔各种盐产品的天然来源。这些产品包括添加或不添加维生素和微量元素的阿尔卑斯山盐、芳香调味盐以及各种特色盐。</p>
          </div>
        </div>
      </section>

      <!-- 阿尔卑斯盐产品板块 -->
      <section class="product-section-area">
        <div class="container">
          <ProductSection
            category="经典"
            title="阿尔卑斯盐"
            description="巴特赖兴哈勒阿尔卑斯盐产品世代以来都是厨房必备。无论过去还是现在，无论您喜欢烹饪什么菜肴，阿尔卑斯盐都是您烹饪和调味的得力助手。在如今注重营养的饮食文化中，我们采用纯正阿尔卑斯盐水制成的蒸发盐，风味恰到好处。此外，它们还添加了碘、氟、叶酸和硒等营养成分，提供人体必需的矿物质。"
            :products="alpineSaltProducts"
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
import { ref, onMounted } from 'vue'
import ProductSection from '@/components/ProductSection.vue'
import PageHeader from '@/components/PageHeader.vue'
import { getProductList } from '@/api/product'

const products = ref([])
const alpineSaltProducts = ref([])

const loadProducts = async () => {
  try {
    const res = await getProductList({ status: 1, page: 1, pageSize: 100 })
    if (res.data?.list) {
      products.value = res.data.list
      alpineSaltProducts.value = res.data.list.filter(p => p.categoryId === 2)
    }
  } catch (error) {
    console.error('加载产品失败:', error)
  }
}

onMounted(() => {
  loadProducts()
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

  padding-top:130px;
  color: #000;
  position: relative;
  
  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;

  }
  
  .title {
    font-size: 2.57143em;
    margin-top: 20px;
    margin-bottom: 42px;
    color: #000;
    text-align: center;
  }
  
  .hero-content {
    max-width: 1000px;
    margin: 0 auto;
    
    h1 {
      font-size: 56px;
      margin-bottom: 20px;
      font-weight: 700;
    }
    
    .subtitle {
      font-size: 14px;
      color: #000;
      text-align: left;
      line-height: 28px;
      margin-bottom: 26px;
    }
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
  
    padding: 40px 0;
    
    .container {
      padding: 0 15px;
    }
    
    .title {
      font-size: 1.8em !important;
      margin-top: 10px;
      margin-bottom: 20px;
      line-height: 1.3;
      text-align: center;
    }
    
    .hero-content {
      max-width: 100%;
      padding: 0 10px;
      
      h1 {
        font-size: 28px !important;
        line-height: 1.3;
      }
      
      .subtitle {
        font-size: 14px;
        line-height: 1.6;
        text-align: left;
        margin-bottom: 20px;
      }
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
    background-position: center;
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
    padding-top:130px;
    
    .container {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      padding: 0 10px;
    }
    
    .title {
      font-size: 1.5em !important;
    }
    
    .hero-content {
      .subtitle {
        font-size: 13px;
      }
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
