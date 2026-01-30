<template>
  <section class="quality">
    <div class="container">
      <div class="quality-content">
        <div class="quality-text">
          <h2>对质量的承诺</h2>
          <p>我们的巴特赖兴哈勒阿尔卑斯盐（Bad Reichenhaller AlpenSalz）名副其实。作为德国最知名的盐品牌，它代表着来自阿尔卑斯山脉深处的高品质盐，由纯净的高山盐水精心提炼而成。</p>
          <p>我们始终将产品质量放在首位。定期的内部和外部审核确保我们始终符合高标准的质量要求。所有巴特赖兴哈勒产品均适合纯素食者食用，且不含增味剂、防腐剂、色素和纳米颗粒。</p>
          <div class="quality-features">
            <div class="feature-item" v-for="(feature, index) in features" :key="index">
              <div class="feature-icon">✓</div>
              <span>{{ feature }}</span>
            </div>
          </div>
        </div>
        <div class="quality-image">
          <img src="@/assets/brand/微信图片_20250408203638.jpg" alt="Quality" />
        </div>
      </div>

      <div class="authority-section">
        <h3>权威背书</h3>
        <div class="authority-grid">
          <div class="authority-card">
            <div class="authority-icon">🏛️</div>
            <h4>WHO 减盐模型</h4>
            <p>2023 版《WHO Sodium Reduction Guideline》指出：若家庭把"普通精盐"更换为"高纯度低杂质盐"，在保持咸度前提下，可减少 8-12% 钠摄入，相当于每天少吃 0.6 g 盐；曾为奥地利皇室医院临床验证。</p>
          </div>
          <div class="authority-card">
            <div class="authority-icon">📊</div>
            <h4>德国 BGR 矿产报告（2022）</h4>
            <p>贝希特斯加登矿岩盐层厚度 1800 m，杂质密度 0.04%，为全球岩矿平均值的 1/5；当年德皇因此把"皇家银器"熔铸成矿道铁轨，只为加速运盐。</p>
          </div>
          <div class="authority-card">
            <div class="authority-icon">🔬</div>
            <h4>感官评测数据</h4>
            <p>德国 TU Dresden 食品感官实验室双盲测试：苦度值 ↓23%、咸度圆润度 ↑19%、后味金属感 ↓31%。皇室御厨团评价："重现 19 世纪宫廷菜单的纯净底味"。</p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
const features = [
  '99.9% 纯度保证',
  '无麸质、无添加剂',
  '无重金属、无核污染',
  '适合纯素食者',
  '孕妇宝宝可安心食用',
  '通过 ISO 9001 & 14001 认证'
]
</script>

<style scoped lang="scss">
.quality {
  padding: 120px 0;
  background: #fff;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

.quality-content {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 80px;
  align-items: center;
  margin-bottom: 100px;
}

.quality-text {
  h2 {
    font-size: 2.5rem;
    font-weight: 700;
    color: #1a1a2e;
    margin-bottom: 30px;
    position: relative;
    padding-bottom: 20px;

    &::after {
      content: '';
      position: absolute;
      bottom: 0;
      left: 0;
      width: 60px;
      height: 3px;
      background: #d4af37;
    }
  }

  p {
    font-size: 1.1rem;
    line-height: 1.8;
    color: #555;
    margin-bottom: 20px;
  }
}

.quality-features {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 15px;
  margin-top: 40px;
}

.feature-item {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 1rem;
  color: #333;

  .feature-icon {
    width: 24px;
    height: 24px;
    background: #d4af37;
    color: #fff;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-weight: 700;
    flex-shrink: 0;
  }
}

.quality-image {
  position: relative;
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 20px 60px rgba(0,0,0,0.15);

  img {
    width: 100%;
    height: 100%;
    object-fit: contain;
    display: block;
    transition: transform 0.6s ease;
  }

  &:hover img {
    transform: scale(1.05);
  }
}

.authority-section {
  h3 {
    font-size: 2.5rem;
    font-weight: 700;
    color: #1a1a2e;
    text-align: center;
    margin-bottom: 60px;
    position: relative;

    &::after {
      content: '';
      position: absolute;
      bottom: -15px;
      left: 50%;
      transform: translateX(-50%);
      width: 80px;
      height: 3px;
      background: #d4af37;
    }
  }
}

.authority-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 40px;
}

.authority-card {
  padding: 40px 30px;
  background: linear-gradient(135deg, #f8f9fa 0%, #fff 100%);
  border-radius: 12px;
  text-align: center;
  border: 2px solid #e9ecef;
  transition: all 0.4s ease;

  &:hover {
    transform: translateY(-5px);
    border-color: #d4af37;
    box-shadow: 0 10px 30px rgba(212, 175, 55, 0.15);
  }

  .authority-icon {
    font-size: 3rem;
    margin-bottom: 20px;
  }

  h4 {
    font-size: 1.3rem;
    font-weight: 600;
    color: #1a1a2e;
    margin-bottom: 15px;
  }

  p {
    font-size: 1rem;
    line-height: 1.7;
    color: #666;
  }
}

@media (max-width: 1024px) {
  .authority-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 768px) {
  .quality {
    padding: 80px 0;
  }

  .quality-content {
    grid-template-columns: 1fr;
    gap: 50px;
    margin-bottom: 80px;
  }

  .quality-text h2 {
    font-size: 2rem;
  }

  .quality-features {
    grid-template-columns: 1fr;
  }

  .authority-section h3 {
    font-size: 2rem;
  }
}
</style>
