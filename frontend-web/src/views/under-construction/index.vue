<!--
  网站建设中页面 - 单页版本
-->
<template>
  <div class="construction-page" @mousemove="handleMouseMove">



    <!-- 页头 -->
    <header class="header">
      <div class="top-bar">
        <div class="container">
          <router-link to="/" class="logo">
            <img src="@/assets/images/logo.png" alt="Brand Logo" />
          </router-link>
        </div>
      </div>
      
      <nav class="primary-navigation">
        <div class="container">
          <ul class="nav-menu">
            <li class="nav-item">
              <router-link to="/" @mouseenter="isHovering = true" @mouseleave="isHovering = false">
                敬请期待
              </router-link>
            </li>
          </ul>
        </div>
      </nav>
    </header>

    <!-- 主内容 -->
    <main class="main-content">
      <div class="container">
        <div class="content-wrapper">
          <!-- 建设中徽章 -->
          <!-- <div class="construction-badge" @mouseenter="isHovering = true" @mouseleave="isHovering = false">
            <svg viewBox="0 0 120 120" xmlns="http://www.w3.org/2000/svg">
              <defs>
                <linearGradient id="badgeGradient" x1="0%" y1="0%" x2="100%" y2="100%">
                  <stop offset="0%" style="stop-color:#c41e3a" />
                  <stop offset="100%" style="stop-color:#92121B" />
                </linearGradient>
              </defs>
              <circle cx="60" cy="60" r="50" fill="none" stroke="url(#badgeGradient)" stroke-width="3" class="badge-circle" />
              <text x="60" y="70" text-anchor="middle" fill="url(#badgeGradient)" font-size="40" font-weight="bold" class="badge-text">建设中</text>
            </svg>
          </div> -->
          
          <!-- 标题区域 -->
          <h1 class="main-title">页面精心打造中</h1>
          <p class="subtitle">就像我们精心制作每一罐阿尔卑斯盐一样</p>
          <p class="description">我们正在为您准备一个完美的体验</p>
          
          <!-- 进度展示 -->
          <div class="progress-section">
            <div class="progress-label">
              <span>开发进度</span>
              <span class="progress-percent">{{ overallProgress }}%</span>
            </div>
            <div class="progress-bar-wrapper">
              <div class="progress-bar-fill" :style="{ width: overallProgress + '%' }">
                <div class="progress-shine"></div>
              </div>
            </div>
          </div>
          
          <!-- 店铺二维码  -->
          <div class="qrcode-section">
            <h3 class="qrcode-title">扫码进入店铺</h3>
            <div class="qrcode-grid">
              <div class="qrcode-item" @mouseenter="isHovering = true" @mouseleave="isHovering = false">
                <div class="qrcode-image">
                  <img src="@/assets/images/douyin.jpg" alt="抖音店铺" />
                </div>
                <div class="qrcode-label">抖音店铺</div>
              </div>
              <div class="qrcode-item" @mouseenter="isHovering = true" @mouseleave="isHovering = false">
                <div class="qrcode-image">
                  <img src="@/assets/images/tianmao.png" alt="天猫店铺" />
                </div>
                <div class="qrcode-label">天猫店铺</div>
              </div>
              <div class="qrcode-item" @mouseenter="isHovering = true" @mouseleave="isHovering = false">
                <div class="qrcode-image">
                  <img src="@/assets/images/jingdong.jpg" alt="京东店铺" />
                </div>
                <div class="qrcode-label">京东店铺</div>
              </div>
            </div>
          </div>
          
          <!-- 倒计时 -->
          <!-- <div class="countdown-section">
            <h3 class="countdown-title">预计上线时间</h3>
            <div class="countdown-grid">
              <div class="countdown-item" @mouseenter="isHovering = true" @mouseleave="isHovering = false">
                <div class="countdown-value">{{ countdown.days }}</div>
                <div class="countdown-label">天</div>
              </div>
              <div class="countdown-separator">:</div>
              <div class="countdown-item" @mouseenter="isHovering = true" @mouseleave="isHovering = false">
                <div class="countdown-value">{{ countdown.hours }}</div>
                <div class="countdown-label">时</div>
              </div>
              <div class="countdown-separator">:</div>
              <div class="countdown-item" @mouseenter="isHovering = true" @mouseleave="isHovering = false">
                <div class="countdown-value">{{ countdown.minutes }}</div>
                <div class="countdown-label">分</div>
              </div>
              <div class="countdown-separator">:</div>
              <div class="countdown-item" @mouseenter="isHovering = true" @mouseleave="isHovering = false">
                <div class="countdown-value">{{ countdown.seconds }}</div>
                <div class="countdown-label">秒</div>
              </div>
            </div>
          </div> -->
          
          <!-- 操作按钮 -->
          <!-- <div class="action-buttons">
            <router-link 
              to="/" 
              class="btn btn-primary"
              @mouseenter="isHovering = true" 
              @mouseleave="isHovering = false"
            >
              返回首页
            </router-link>
            <router-link 
              to="/contact" 
              class="btn btn-secondary"
              @mouseenter="isHovering = true" 
              @mouseleave="isHovering = false"
            >
              联系我们
            </router-link>
          </div> -->
        </div>
      </div>
    </main>

    <!-- 页脚 -->
    <footer class="footer">
      <div class="container">
        <p>&copy; 2024 Brand Name. All rights reserved.</p>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

// 光标
const cursorX = ref(0)
const cursorY = ref(0)
const isHovering = ref(false)
const hoveredFeature = ref(null)

// 进度
const overallProgress = ref(0)

// 特性数据
const features = ref([
  { icon: '🎨', title: '精美设计', desc: '延续品牌一贯的优雅风格' },
  { icon: '⚡', title: '快速加载', desc: '优化性能，提升用户体验' },
  { icon: '📱', title: '响应式布局', desc: '完美适配各种设备屏幕' },
  { icon: '🔒', title: '安全可靠', desc: '保护您的数据和隐私' }
])

// 倒计时
const countdown = ref({
  days: 15,
  hours: 8,
  minutes: 30,
  seconds: 45
})

// 鼠标移动
const handleMouseMove = (e) => {
  cursorX.value = e.clientX
  cursorY.value = e.clientY
}

// 特性悬停
const handleFeatureHover = (index, hover) => {
  hoveredFeature.value = hover ? index : null
  isHovering.value = hover
}

// 进度条动画
let progressInterval = null

// 倒计时更新
let countdownInterval = null
const updateCountdown = () => {
  countdown.value.seconds--
  if (countdown.value.seconds < 0) {
    countdown.value.seconds = 59
    countdown.value.minutes--
  }
  if (countdown.value.minutes < 0) {
    countdown.value.minutes = 59
    countdown.value.hours--
  }
  if (countdown.value.hours < 0) {
    countdown.value.hours = 23
    countdown.value.days--
  }
}

onMounted(() => {

  
  progressInterval = setInterval(() => {
    if (overallProgress.value < 75) {
      overallProgress.value += Math.random() * 1.5
    }
  }, 100)
  
  countdownInterval = setInterval(updateCountdown, 1000)
})

onUnmounted(() => {

  if (progressInterval) clearInterval(progressInterval)
  if (countdownInterval) clearInterval(countdownInterval)
})
</script>

<style lang="scss" scoped>


.construction-page {
  min-height: 100vh;
  position: relative;
  background-image: url('@/assets/images/background.jpg');
  background-size: cover;
  background-position: center;
  background-attachment: fixed;
  
  &::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;

    z-index: 0;
  }
}

// 自定义光标
.custom-cursor {
  position: fixed;
  pointer-events: none;
  z-index: 10000;
  transition: transform 0.15s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  
  .cursor-dot {
    position: absolute;
    width: 8px;
    height: 8px;
    background: #c41e3a;
    border-radius: 50%;
    transform: translate(-50%, -50%);
    transition: all 0.15s ease;
  }
  
  .cursor-ring {
    position: absolute;
    width: 40px;
    height: 40px;
    border: 2px solid #c41e3a;
    border-radius: 50%;
    transform: translate(-50%, -50%);
    transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  }
  
  &.cursor-hover {
    .cursor-dot {
      transform: translate(-50%, -50%) scale(2);
    }
    .cursor-ring {
      transform: translate(-50%, -50%) scale(1.5);
    }
  }
}

// 垂直条纹背景
.stripe-background {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  z-index: 0;
  opacity: 0.3;
  
  .stripe {
    flex: 1;
    background: linear-gradient(180deg, 
      rgba(196, 30, 58, 0.02) 0%, 
      rgba(213, 6, 28, 0.05) 50%, 
      rgba(146, 18, 27, 0.02) 100%
    );
    animation: stripeMove 4s ease-in-out infinite;
    
    &:nth-child(even) {
      animation-direction: reverse;
    }
  }
}

@keyframes stripeMove {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-30px);
  }
}

// 页头
.header {
  position: relative;
  // background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  z-index: 100;
}

.top-bar {
  padding: 15px 0;
  height: 112px;
  display: flex;
  align-items: center;
  
  .container {
    max-width: 1400px;
    margin: 0 auto;
    padding: 0 40px;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
  }
  
  .logo {
    height: 100%;
    img {
      height: 100%;
      display: block;
    }
  }
}

.primary-navigation {
  height: 58px;
  border-top: 2px solid #fff;
  border-bottom: 3px solid #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(90deg, #92121B 0%, #D5061C 25%, #D5061C 75%, #92121B 100%);
  
  .container {
    max-width: 1400px;
    margin: 0 auto;
    padding: 0 40px;
  }
  
  .nav-menu {
    display: flex;
    justify-content: center;
    list-style: none;
    padding: 0;
    margin: 0;
    
    .nav-item a {
      display: block;
      padding: 20px 25px;
      color: #fff;
      text-decoration: none;
      font-size: 16px;
      font-weight: 500;
      transition: opacity 0.3s;
      
      &:hover {
        opacity: 0.8;
      }
    }
  }
}

// 主内容
.main-content {
  position: relative;
  z-index: 1;
  padding: 80px 0;
  min-height: calc(100vh - 170px - 60px);
  
  .container {
    max-width: 1000px;
    margin: 0 auto;
    padding: 0 40px;
  }
  
  .content-wrapper {
    text-align: center;
  }
}

// 建设中徽章
.construction-badge {
  width: 120px;
  height: 120px;
  margin: 0 auto 40px;
  
  svg {
    width: 100%;
    height: 100%;
    filter: drop-shadow(0 4px 12px rgba(196, 30, 58, 0.2));
  }
  
  .badge-circle {
    stroke-dasharray: 314;
    stroke-dashoffset: 314;
    animation: drawBadge 2s ease-out forwards;
  }
  
  .badge-text {
    opacity: 0;
    animation: fadeInText 0.5s ease-out 1.5s forwards;
  }
}

@keyframes drawBadge {
  to { stroke-dashoffset: 0; }
}

@keyframes fadeInText {
  to { opacity: 1; }
}

// 标题区域
.main-title {
  font-size: 48px;
  font-weight: bold;
  color: #2b2b2e;
  margin: 0 0 20px 0;
}

.subtitle {
  font-size: 20px;
  color: #666;
  margin: 0 0 10px 0;
}

.description {
  font-size: 16px;
  color: #999;
  margin: 0 0 50px 0;
}

// 进度条
.progress-section {
  max-width: 600px;
  margin: 0 auto 60px;
  
  .progress-label {
    display: flex;
    justify-content: space-between;
    margin-bottom: 15px;
    font-size: 16px;
    font-weight: 600;
    color: #2b2b2e;
    
    .progress-percent {
      color: #c41e3a;
    }
  }
  
  .progress-bar-wrapper {
    height: 12px;
    background: #e0e0e0;
    border-radius: 10px;
    overflow: hidden;
    
    .progress-bar-fill {
      height: 100%;
      background: linear-gradient(90deg, #92121B 0%, #D5061C 50%, #c41e3a 100%);
      border-radius: 10px;
      transition: width 1.5s cubic-bezier(0.645, 0.045, 0.355, 1);
      position: relative;
      overflow: hidden;
      
      .progress-shine {
        position: absolute;
        top: 0;
        left: -100%;
        width: 100%;
        height: 100%;
        background: linear-gradient(90deg, transparent, rgba(255,255,255,0.3), transparent);
        animation: shine 2s infinite;
      }
    }
  }
}

@keyframes shine {
  to { left: 100%; }
}

// 店铺二维码
.qrcode-section {
  margin-bottom: 60px;
  
  .qrcode-title {
    font-size: 24px;
    font-weight: 600;
    color: #2b2b2e;
    margin: 0 0 30px 0;
  }
  
  .qrcode-grid {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 40px;
    flex-wrap: wrap;
  }
  
  .qrcode-item {
    text-align: center;
    transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
    
    &:hover {
      transform: translateY(-10px);
      
      .qrcode-image {
        box-shadow: 0 12px 32px rgba(196, 30, 58, 0.2);
      }
    }
    
    .qrcode-image {
      width: 260px;
      height: 260px;
      background: #fff;
      padding: 20px;
      border-radius: 16px;
      box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
      margin-bottom: 15px;
      transition: all 0.3s ease;
      
      img {
        width: 100%;
        height: 100%;
        object-fit: contain;
        display: block;
      }
    }
    
    .qrcode-label {
      font-size: 16px;
      font-weight: 600;
      color: #2b2b2e;
      
      &::before {
        content: '';
        display: block;
        width: 40px;
        height: 3px;
        background: linear-gradient(90deg, #92121B 0%, #D5061C 100%);
        margin: 0 auto 10px;
        border-radius: 2px;
      }
    }
  }
}

// 特性卡片
.features-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
  margin-bottom: 60px;
}

.feature-card {
  background: #fff;
  padding: 30px 20px;
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  text-align: center;
  transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  
  &.card-hover {
    transform: translateY(-10px);
    box-shadow: 0 12px 32px rgba(196, 30, 58, 0.15);
  }
  
  .feature-icon {
    font-size: 48px;
    margin-bottom: 15px;
  }
  
  .feature-title {
    font-size: 18px;
    font-weight: 600;
    color: #2b2b2e;
    margin: 0 0 10px 0;
  }
  
  .feature-desc {
    font-size: 14px;
    color: #666;
    line-height: 1.6;
    margin: 0;
  }
}

// 倒计时
.countdown-section {
  margin-bottom: 50px;
  
  .countdown-title {
    font-size: 20px;
    font-weight: 600;
    color: #2b2b2e;
    margin: 0 0 30px 0;
  }
  
  .countdown-grid {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 20px;
  }
  
  .countdown-item {
    background: #fff;
    padding: 24px 32px;
    border-radius: 12px;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
    text-align: center;
    transition: all 0.3s ease;
    
    &:hover {
      transform: scale(1.1);
      box-shadow: 0 8px 24px rgba(196, 30, 58, 0.15);
    }
    
    .countdown-value {
      font-size: 36px;
      font-weight: bold;
      color: #c41e3a;
      line-height: 1;
      margin-bottom: 8px;
    }
    
    .countdown-label {
      font-size: 13px;
      color: #666;
    }
  }
  
  .countdown-separator {
    font-size: 36px;
    color: rgba(43, 43, 46, 0.3);
    font-weight: bold;
  }
}

// 操作按钮
.action-buttons {
  display: flex;
  gap: 20px;
  justify-content: center;
  
  .btn {
    padding: 14px 40px;
    font-size: 16px;
    font-weight: 600;
    border-radius: 8px;
    text-decoration: none;
    transition: all 0.3s ease;
    
    &.btn-primary {
      background: linear-gradient(90deg, #92121B 0%, #D5061C 100%);
      color: #fff;
      
      &:hover {
        transform: translateY(-2px);
        box-shadow: 0 8px 20px rgba(196, 30, 58, 0.3);
      }
    }
    
    &.btn-secondary {
      background: #fff;
      color: #c41e3a;
      border: 2px solid #c41e3a;
      
      &:hover {
        background: #c41e3a;
        color: #fff;
      }
    }
  }
}

// 页脚
.footer {
  position: relative;
  background-color: #2c2c2c;
  color: #fff;
  padding: 20px 0;
  text-align: center;
  z-index: 100;
  
  .container {
    max-width: 1400px;
    margin: 0 auto;
    padding: 0 40px;
  }
  
  p {
    margin: 0;
    font-size: 14px;
    color: #999;
  }
}

// 响应式
@media (max-width: 768px) {
  .top-bar {
    height: auto;
    padding: 10px 0;
    
    .container {
      padding: 0 20px;
    }
    
    .logo {
      height: 60px;
    }
  }
  
  .main-content {
    padding: 50px 0;
    
    .container {
      padding: 0 20px;
    }
  }
  
  .construction-badge {
    width: 100px;
    height: 100px;
  }
  
  .main-title {
    font-size: 32px;
  }
  
  .subtitle {
    font-size: 18px;
  }
  
  .features-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 16px;
  }
  
  .qrcode-section {
    .qrcode-grid {
      gap: 30px;
    }
    
    .qrcode-item {
      .qrcode-image {
        width: 220px;
        height: 220px;
        padding: 18px;
      }
      
      .qrcode-label {
        font-size: 14px;
      }
    }
  }
  
  .countdown-grid {
    flex-wrap: wrap;
    gap: 12px;
    
    .countdown-item {
      padding: 20px 24px;
      
      .countdown-value {
        font-size: 28px;
      }
    }
    
    .countdown-separator {
      font-size: 28px;
    }
  }
  
  .action-buttons {
    flex-direction: column;
    
    .btn {
      width: 100%;
    }
  }
}

@media (max-width: 480px) {
  .features-grid {
    grid-template-columns: 1fr;
  }
  
  .qrcode-section {
    .qrcode-grid {
      gap: 20px;
    }
    
    .qrcode-item {
      .qrcode-image {
        width: 200px;
        height: 200px;
        padding: 15px;
      }
      
      .qrcode-label {
        font-size: 13px;
      }
    }
  }
}
</style>
