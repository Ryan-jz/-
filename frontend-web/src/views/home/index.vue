<!--
  前台首页
-->
<template>
  <div class="home-container">
    <!-- 导航栏 -->
    <header class="header">
      <div class="container">
        <div class="logo">GF Admin</div>
        <nav class="nav">
          <router-link to="/">首页</router-link>
          <router-link to="/about">关于我们</router-link>
          <router-link to="/contact">联系我们</router-link>
        </nav>
      </div>
    </header>

    <!-- 主横幅 -->
    <section class="hero">
      <div class="container">
        <h1>欢迎使用 GF Admin 管理系统</h1>
        <p>基于 GoFrame + Vue3 的现代化后台管理解决方案</p>
        <div class="hero-buttons">
          <button class="btn btn-primary">立即开始</button>
          <button class="btn btn-secondary">了解更多</button>
        </div>
      </div>
    </section>

    <!-- 特性介绍 -->
    <section class="features">
      <div class="container">
        <h2>核心特性</h2>
        <div class="feature-grid">
          <div class="feature-item">
            <div class="feature-icon">🚀</div>
            <h3>高性能</h3>
            <p>基于 GoFrame 框架，提供高性能的后端服务</p>
          </div>
          <div class="feature-item">
            <div class="feature-icon">🎨</div>
            <h3>现代化 UI</h3>
            <p>使用 Vue3 + Element Plus 构建美观的用户界面</p>
          </div>
          <div class="feature-item">
            <div class="feature-icon">🔒</div>
            <h3>安全可靠</h3>
            <p>JWT 认证，完善的权限管理体系</p>
          </div>
          <div class="feature-item">
            <div class="feature-icon">📱</div>
            <h3>响应式设计</h3>
            <p>完美适配各种设备和屏幕尺寸</p>
          </div>
        </div>
      </div>
    </section>

    <!-- 页脚 -->
    <footer class="footer">
      <div class="container">
        <p>&copy; 2024 GF Admin. All rights reserved.</p>
      </div>
    </footer>
  </div>
</template>

<script setup>
// 首页组件
</script>

<style lang="scss" scoped>
.home-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.header {
  background-color: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  padding: 20px 0;
  
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
    color: #409eff;
  }
  
  .nav {
    display: flex;
    gap: 30px;
    
    a {
      color: #333;
      text-decoration: none;
      font-size: 16px;
      transition: color 0.3s;
      
      &:hover,
      &.router-link-active {
        color: #409eff;
      }
    }
  }
}

.hero {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #fff;
  padding: 100px 0;
  text-align: center;
  
  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
  }
  
  h1 {
    font-size: 48px;
    margin-bottom: 20px;
  }
  
  p {
    font-size: 20px;
    margin-bottom: 40px;
    opacity: 0.9;
  }
  
  .hero-buttons {
    display: flex;
    gap: 20px;
    justify-content: center;
  }
  
  .btn {
    padding: 12px 30px;
    font-size: 16px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    transition: all 0.3s;
    
    &.btn-primary {
      background-color: #fff;
      color: #667eea;
      
      &:hover {
        transform: translateY(-2px);
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
      }
    }
    
    &.btn-secondary {
      background-color: transparent;
      color: #fff;
      border: 2px solid #fff;
      
      &:hover {
        background-color: #fff;
        color: #667eea;
      }
    }
  }
}

.features {
  padding: 80px 0;
  background-color: #f5f7fa;
  
  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
  }
  
  h2 {
    text-align: center;
    font-size: 36px;
    margin-bottom: 60px;
    color: #333;
  }
  
  .feature-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 40px;
  }
  
  .feature-item {
    text-align: center;
    padding: 30px;
    background-color: #fff;
    border-radius: 10px;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
    transition: transform 0.3s;
    
    &:hover {
      transform: translateY(-5px);
    }
    
    .feature-icon {
      font-size: 48px;
      margin-bottom: 20px;
    }
    
    h3 {
      font-size: 24px;
      margin-bottom: 15px;
      color: #333;
    }
    
    p {
      color: #666;
      line-height: 1.6;
    }
  }
}

.footer {
  background-color: #333;
  color: #fff;
  padding: 30px 0;
  text-align: center;
  margin-top: auto;
  
  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
  }
}
</style>
