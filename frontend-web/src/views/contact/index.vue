<!--
  联系我们页面
-->
<template>
  <div class="contact-container">
   <PageHeader></PageHeader>

    <section class="content">
      <div class="container">
        <h1>联系我们</h1>
        <div class="contact-info">
          <div class="info-item">
            <h3>📧 邮箱</h3>
            <p>csm@alpensalt.com</p>
          </div>
          <div class="info-item">
            <h3>📱 电话</h3>
            <p>+86 400 820 7820</p>
            <p>18101843788</p>
          </div>
          <div class="info-item">
            <h3>📍 地址</h3>
            <p>Alpen 白金盐上海代表处
上海虹桥商务区
闵行盘阳路158弄5号105室</p>
          </div>
        </div>
      </div>
    </section>

    <footer class="footer">
      <div class="container">
        <!-- <p>&copy; 2024 GF Admin. All rights reserved.</p> -->
      </div>
    </footer>
  </div>
</template>

<script setup>
  import PageHeader from '@/components/PageHeader.vue'
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
