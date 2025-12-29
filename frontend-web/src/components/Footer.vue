<template>
  <footer class="footer">
    <div class="container">
      <div class="footer-content">
        <div class="footer-section">
          <h4>产品系列</h4>
          <ul>
            <li><router-link to="/product">产品列表</router-link></li>
          </ul>
        </div>
        
        <div class="footer-section">
          <h4>关于我们</h4>
          <ul>
            <li><router-link to="/about">品牌故事</router-link></li>
          </ul>
        </div>
        
        <div class="footer-section">
          <h4>客户服务</h4>
          <ul>
            <li><router-link to="/contact">联系我们</router-link></li>
          </ul>
        </div>
        
        <div class="footer-section">
          <h4>法律信息</h4>
          <ul>
            <li><a href="#">隐私政策</a></li>
            <li><a href="#">使用条款</a></li>
          </ul>
        </div>
        
        <div class="footer-section">
          <h4>关注我们</h4>
          <div class="footer-social">
            <a href="#" aria-label="Facebook">Facebook</a>
            <a href="#" aria-label="Instagram">Instagram</a>
            <a href="#" aria-label="TikTok">TikTok</a>
            <a href="#" aria-label="YouTube">YouTube</a>
          </div>
        </div>
      </div>
      
      <div class="footer-bottom">
        <p><a href="https://beian.miit.gov.cn/" target="_blank" rel="noopener noreferrer">沪ICP备2024052802号-2</a></p>
        <p class="footer-group">Part of International Salt Group</p>
      </div>
    </div>
  </footer>
</template>

<style lang="scss" scoped>
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

@media (max-width: 768px) {
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
}

@media (max-width: 480px) {
  .footer {
    padding: 30px 0 15px;
  }
}
</style>
