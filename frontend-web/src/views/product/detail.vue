<!--
  产品详情页 - 参考 Bad Reichenhaller 产品页面布局
-->
<template>
  <div class="product-detail-page">
    <!-- 导航栏 -->
    <header class="header">
      <div class="container">
        <div class="logo">Brand Logo</div>
        <nav class="nav">
          <a href="#">Products</a>
          <a href="#">Recipes</a>
          <a href="#">About</a>
          <a href="#">Contact</a>
        </nav>
      </div>
    </header>

    <!-- 面包屑导航 -->
    <div class="breadcrumb">
      <div class="container">
        <span>Home</span>
        <span class="separator">/</span>
        <span>Products</span>
        <span class="separator">/</span>
        <span>Seasoned Salt</span>
        <span class="separator">/</span>
        <span class="current">BBQ Seasoned Salt for Poultry</span>
      </div>
    </div>

    <!-- 产品主要内容区 -->
    <div class="product-main">
      <div class="container">
        <!-- 左侧：产品图片 -->
        <div class="product-images">
          <div class="main-image">
            <img src="https://via.placeholder.com/500x600" alt="Product Image" />
          </div>
          <div class="thumbnail-list">
            <div class="thumbnail active">
              <img src="https://via.placeholder.com/100x100" alt="Thumbnail 1" />
            </div>
            <div class="thumbnail">
              <img src="https://via.placeholder.com/100x100" alt="Thumbnail 2" />
            </div>
            <div class="thumbnail">
              <img src="https://via.placeholder.com/100x100" alt="Thumbnail 3" />
            </div>
          </div>
        </div>

        <!-- 右侧：产品信息 -->
        <div class="product-info">
          <div class="product-category">Seasoned Salt</div>
          <h1 class="product-title">BBQ Seasoned Salt for Poultry</h1>
          
          <div class="product-description">
            <p>
              Perfect blend of herbs and spices specially crafted for poultry dishes. 
              This BBQ seasoned salt brings out the natural flavors of chicken, turkey, 
              and other poultry with a delicious smoky taste.
            </p>
          </div>

          <div class="product-features">
            <div class="feature-item">
              <span class="icon">✓</span>
              <span>Natural ingredients</span>
            </div>
            <div class="feature-item">
              <span class="icon">✓</span>
              <span>No artificial flavors</span>
            </div>
            <div class="feature-item">
              <span class="icon">✓</span>
              <span>Perfect for grilling</span>
            </div>
          </div>

          <div class="product-specs">
            <div class="spec-item">
              <span class="label">Content:</span>
              <span class="value">90g</span>
            </div>
            <div class="spec-item">
              <span class="label">Package:</span>
              <span class="value">Grinder</span>
            </div>
          </div>

          <div class="product-actions">
            <button class="btn btn-primary">Add to Cart</button>
            <button class="btn btn-secondary">Find in Store</button>
          </div>
        </div>
      </div>
    </div>

    <!-- 产品详细信息标签页 -->
    <div class="product-details-tabs">
      <div class="container">
        <div class="tabs-header">
          <button 
            class="tab-btn" 
            :class="{ active: activeTab === 'ingredients' }"
            @click="activeTab = 'ingredients'"
          >
            Ingredients
          </button>
          <button 
            class="tab-btn" 
            :class="{ active: activeTab === 'nutrition' }"
            @click="activeTab = 'nutrition'"
          >
            Nutritional Information
          </button>
          <button 
            class="tab-btn" 
            :class="{ active: activeTab === 'usage' }"
            @click="activeTab = 'usage'"
          >
            Usage & Tips
          </button>
        </div>

        <div class="tabs-content">
          <!-- 成分标签页 -->
          <div v-show="activeTab === 'ingredients'" class="tab-panel">
            <h3>Ingredients</h3>
            <p>
              Sea salt, paprika, garlic, onion, black pepper, herbs (thyme, rosemary, oregano), 
              natural smoke flavor, spices.
            </p>
          </div>

          <!-- 营养信息标签页 -->
          <div v-show="activeTab === 'nutrition'" class="tab-panel">
            <h3>Nutritional Information</h3>
            <table class="nutrition-table">
              <thead>
                <tr>
                  <th>Per 100g</th>
                  <th>Amount</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>Energy</td>
                  <td>50 kJ / 12 kcal</td>
                </tr>
                <tr>
                  <td>Fat</td>
                  <td>0.5 g</td>
                </tr>
                <tr>
                  <td>Carbohydrates</td>
                  <td>1.2 g</td>
                </tr>
                <tr>
                  <td>Protein</td>
                  <td>0.8 g</td>
                </tr>
                <tr>
                  <td>Salt</td>
                  <td>95 g</td>
                </tr>
              </tbody>
            </table>
          </div>

          <!-- 使用方法标签页 -->
          <div v-show="activeTab === 'usage'" class="tab-panel">
            <h3>Usage & Tips</h3>
            <ul class="usage-list">
              <li>Perfect for seasoning chicken, turkey, and other poultry before grilling or roasting</li>
              <li>Use as a dry rub: Apply generously to meat 30 minutes before cooking</li>
              <li>Great for marinades: Mix with olive oil and lemon juice</li>
              <li>Recommended amount: 1-2 teaspoons per 500g of meat</li>
              <li>Store in a cool, dry place</li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <!-- 相关产品推荐 -->
    <div class="related-products">
      <div class="container">
        <h2>You might also like</h2>
        <div class="products-grid">
          <div class="product-card" v-for="i in 4" :key="i">
            <div class="product-image">
              <img src="https://via.placeholder.com/250x300" alt="Related Product" />
            </div>
            <div class="product-card-info">
              <h4>BBQ Seasoned Salt for Beef</h4>
              <p class="product-size">90g</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 页脚 -->
    <footer class="footer">
      <div class="container">
        <div class="footer-content">
          <div class="footer-section">
            <h4>Products</h4>
            <ul>
              <li><a href="#">Seasoned Salt</a></li>
              <li><a href="#">Table Salt</a></li>
              <li><a href="#">Sea Salt</a></li>
            </ul>
          </div>
          <div class="footer-section">
            <h4>Company</h4>
            <ul>
              <li><a href="#">About Us</a></li>
              <li><a href="#">Contact</a></li>
              <li><a href="#">Careers</a></li>
            </ul>
          </div>
          <div class="footer-section">
            <h4>Support</h4>
            <ul>
              <li><a href="#">FAQ</a></li>
              <li><a href="#">Shipping</a></li>
              <li><a href="#">Returns</a></li>
            </ul>
          </div>
          <div class="footer-section">
            <h4>Follow Us</h4>
            <div class="social-links">
              <a href="#">Facebook</a>
              <a href="#">Instagram</a>
              <a href="#">Twitter</a>
            </div>
          </div>
        </div>
        <div class="footer-bottom">
          <p>&copy; 2024 Brand Name. All rights reserved.</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const activeTab = ref('ingredients')
</script>

<style lang="scss" scoped>
.product-detail-page {
  min-height: 100vh;
  background-color: #fff;
}

// 导航栏
.header {
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  padding: 20px 0;
  position: sticky;
  top: 0;
  z-index: 100;
  
  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  
  .logo {
    font-size: 24px;
    font-weight: bold;
    color: #c41e3a;
  }
  
  .nav {
    display: flex;
    gap: 30px;
    
    a {
      color: #333;
      text-decoration: none;
      font-size: 16px;
      transition: color 0.3s;
      
      &:hover {
        color: #c41e3a;
      }
    }
  }
}

// 面包屑导航
.breadcrumb {
  background-color: #f8f8f8;
  padding: 15px 0;
  font-size: 14px;
  
  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
  }
  
  .separator {
    margin: 0 10px;
    color: #999;
  }
  
  .current {
    color: #c41e3a;
    font-weight: 500;
  }
}

// 产品主要内容
.product-main {
  padding: 60px 0;
  
  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 60px;
  }
}

// 产品图片区域
.product-images {
  .main-image {
    width: 100%;
    margin-bottom: 20px;
    border-radius: 8px;
    overflow: hidden;
    background-color: #f8f8f8;
    
    img {
      width: 100%;
      height: auto;
      display: block;
    }
  }
  
  .thumbnail-list {
    display: flex;
    gap: 15px;
    
    .thumbnail {
      width: 80px;
      height: 80px;
      border-radius: 4px;
      overflow: hidden;
      cursor: pointer;
      border: 2px solid transparent;
      transition: border-color 0.3s;
      
      &.active,
      &:hover {
        border-color: #c41e3a;
      }
      
      img {
        width: 100%;
        height: 100%;
        object-fit: cover;
      }
    }
  }
}

// 产品信息区域
.product-info {
  .product-category {
    color: #999;
    font-size: 14px;
    text-transform: uppercase;
    letter-spacing: 1px;
    margin-bottom: 10px;
  }
  
  .product-title {
    font-size: 36px;
    color: #333;
    margin-bottom: 20px;
    line-height: 1.2;
  }
  
  .product-description {
    font-size: 16px;
    color: #666;
    line-height: 1.6;
    margin-bottom: 30px;
  }
  
  .product-features {
    margin-bottom: 30px;
    
    .feature-item {
      display: flex;
      align-items: center;
      gap: 10px;
      margin-bottom: 10px;
      
      .icon {
        color: #4caf50;
        font-weight: bold;
      }
    }
  }
  
  .product-specs {
    background-color: #f8f8f8;
    padding: 20px;
    border-radius: 8px;
    margin-bottom: 30px;
    
    .spec-item {
      display: flex;
      justify-content: space-between;
      margin-bottom: 10px;
      
      &:last-child {
        margin-bottom: 0;
      }
      
      .label {
        color: #666;
      }
      
      .value {
        font-weight: 600;
        color: #333;
      }
    }
  }
  
  .product-actions {
    display: flex;
    gap: 15px;
    
    .btn {
      flex: 1;
      padding: 15px 30px;
      font-size: 16px;
      border: none;
      border-radius: 4px;
      cursor: pointer;
      transition: all 0.3s;
      font-weight: 600;
      
      &.btn-primary {
        background-color: #c41e3a;
        color: #fff;
        
        &:hover {
          background-color: #a01828;
        }
      }
      
      &.btn-secondary {
        background-color: #fff;
        color: #c41e3a;
        border: 2px solid #c41e3a;
        
        &:hover {
          background-color: #c41e3a;
          color: #fff;
        }
      }
    }
  }
}

// 详细信息标签页
.product-details-tabs {
  background-color: #f8f8f8;
  padding: 60px 0;
  
  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
  }
  
  .tabs-header {
    display: flex;
    gap: 0;
    border-bottom: 2px solid #ddd;
    margin-bottom: 40px;
    
    .tab-btn {
      padding: 15px 30px;
      background: none;
      border: none;
      font-size: 16px;
      color: #666;
      cursor: pointer;
      transition: all 0.3s;
      border-bottom: 3px solid transparent;
      margin-bottom: -2px;
      
      &:hover {
        color: #333;
      }
      
      &.active {
        color: #c41e3a;
        border-bottom-color: #c41e3a;
        font-weight: 600;
      }
    }
  }
  
  .tabs-content {
    .tab-panel {
      h3 {
        font-size: 24px;
        margin-bottom: 20px;
        color: #333;
      }
      
      p {
        font-size: 16px;
        line-height: 1.6;
        color: #666;
      }
      
      .nutrition-table {
        width: 100%;
        max-width: 600px;
        border-collapse: collapse;
        
        th, td {
          padding: 12px;
          text-align: left;
          border-bottom: 1px solid #ddd;
        }
        
        th {
          background-color: #fff;
          font-weight: 600;
          color: #333;
        }
        
        td {
          color: #666;
        }
      }
      
      .usage-list {
        list-style: none;
        padding: 0;
        
        li {
          padding: 10px 0;
          padding-left: 25px;
          position: relative;
          color: #666;
          line-height: 1.6;
          
          &:before {
            content: "•";
            position: absolute;
            left: 0;
            color: #c41e3a;
            font-size: 20px;
          }
        }
      }
    }
  }
}

// 相关产品
.related-products {
  padding: 60px 0;
  
  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
  }
  
  h2 {
    font-size: 32px;
    margin-bottom: 40px;
    color: #333;
  }
  
  .products-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    gap: 30px;
  }
  
  .product-card {
    background-color: #fff;
    border-radius: 8px;
    overflow: hidden;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    transition: transform 0.3s;
    cursor: pointer;
    
    &:hover {
      transform: translateY(-5px);
      box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
    }
    
    .product-image {
      width: 100%;
      height: 300px;
      background-color: #f8f8f8;
      
      img {
        width: 100%;
        height: 100%;
        object-fit: cover;
      }
    }
    
    .product-card-info {
      padding: 20px;
      
      h4 {
        font-size: 18px;
        margin-bottom: 8px;
        color: #333;
      }
      
      .product-size {
        color: #999;
        font-size: 14px;
      }
    }
  }
}

// 页脚
.footer {
  background-color: #2c2c2c;
  color: #fff;
  padding: 60px 0 20px;
  
  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
  }
  
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
      
      li {
        margin-bottom: 10px;
        
        a {
          color: #ccc;
          text-decoration: none;
          transition: color 0.3s;
          
          &:hover {
            color: #fff;
          }
        }
      }
    }
    
    .social-links {
      display: flex;
      flex-direction: column;
      gap: 10px;
      
      a {
        color: #ccc;
        text-decoration: none;
        transition: color 0.3s;
        
        &:hover {
          color: #fff;
        }
      }
    }
  }
  
  .footer-bottom {
    border-top: 1px solid #444;
    padding-top: 20px;
    text-align: center;
    color: #999;
  }
}

// 响应式设计
@media (max-width: 768px) {
  .product-main .container {
    grid-template-columns: 1fr;
    gap: 30px;
  }
  
  .product-info .product-title {
    font-size: 28px;
  }
  
  .tabs-header {
    flex-wrap: wrap;
  }
  
  .products-grid {
    grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  }
}
</style>
