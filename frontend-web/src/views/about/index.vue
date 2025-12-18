<!--
  关于我们页面
-->
<template>
  <div class="about-container">
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
        <h1>关于我们</h1>
        <div class="about-content">
          <h2>项目介绍</h2>
          <p>GF Admin 是一个基于 GoFrame v2 和 Vue 3 开发的现代化后台管理系统。</p>
          
          <h2>技术栈</h2>
          <ul>
            <li>后端：GoFrame v2.6.4 + MySQL 8.0 + Redis</li>
            <li>前端：Vue 3 + Vite + Element Plus + Pinia</li>
            <li>认证：JWT Token</li>
          </ul>
          
          <h2>功能特性</h2>
          <ul>
            <li>用户管理：完整的用户增删改查功能</li>
            <li>角色管理：灵活的角色权限配置</li>
            <li>菜单管理：动态菜单配置</li>
            <li>部门管理：组织架构管理</li>
          </ul>
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
.about-container {
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
  }
  
  .about-content {
    h2 {
      font-size: 24px;
      margin: 30px 0 15px;
      color: #333;
    }
    
    p, li {
      font-size: 16px;
      line-height: 1.8;
      color: #666;
    }
    
    ul {
      padding-left: 20px;
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
