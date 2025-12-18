<!--
  联系我们页面
-->
<template>
  <div class="contact-container">
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

    <section class="content">
      <div class="container">
        <h1>联系我们</h1>
        <div class="contact-info">
          <div class="info-item">
            <h3>📧 邮箱</h3>
            <p>admin@example.com</p>
          </div>
          <div class="info-item">
            <h3>📱 电话</h3>
            <p>+86 138-0000-0000</p>
          </div>
          <div class="info-item">
            <h3>📍 地址</h3>
            <p>中国 北京市 朝阳区</p>
          </div>
        </div>
      </div>
    </section>

    <footer class="footer">
      <div class="container">
        <p>&copy; 2024 GF Admin. All rights reserved.</p>
      </div>
    </footer>
  </div>
</template>

<script setup>
</script>

<style lang="scss" scoped>
.contact-container {
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

.content {
  flex: 1;
  padding: 60px 0;
  
  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
  }
  
  h1 {
    font-size: 36px;
    margin-bottom: 40px;
    color: #333;
    text-align: center;
  }
  
  .contact-info {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 40px;
    margin-top: 60px;
    
    .info-item {
      text-align: center;
      padding: 40px;
      background-color: #f5f7fa;
      border-radius: 10px;
      
      h3 {
        font-size: 24px;
        margin-bottom: 15px;
        color: #333;
      }
      
      p {
        font-size: 16px;
        color: #666;
      }
    }
  }
}

.footer {
  background-color: #333;
  color: #fff;
  padding: 30px 0;
  text-align: center;
  
  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
  }
}
</style>
