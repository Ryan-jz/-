<template>
  <div id="app">
    <!-- 头部导航 -->
    <header class="site-header">
      <div class="header-container">
 
        <label for="navigation-expanded" class="nav-toggle">
          <span>{{ navExpanded ? '收起菜单' : '展开菜单' }}</span>
        </label>

        <!-- Logo -->
        <a class="logo" href="/" title="访问巴特赖兴哈勒网站">
          <img alt="巴特赖兴哈勒的标志" src="data:image/png;base64,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">
        </a>

        <!-- 主导航 -->
        <nav class="main-nav">
          <ul class="nav-list">
            <li class="nav-item" v-for="(group, index) in navGroups" :key="index">
              <input 
                :id="`node-${group.id}`" 
                type="radio" 
                :name="`nav-group`" 
                v-model="activeNavGroup" 
                :value="group.id"
              >
              <label :for="`node-${group.id}`">{{ group.label }}</label>
              <ul class="submenu" v-if="activeNavGroup === group.id">
                <li v-for="(item, itemIndex) in group.items" :key="itemIndex">
                  <a :href="item.href" :title="item.title">{{ item.text }}</a>
                </li>
              </ul>
            </li>
          </ul>

          <!-- 社交媒体链接 -->
          <ul class="social-media">
            <li v-for="(social, index) in socialMedia" :key="index">
              <a :href="social.url" :title="social.title" target="_blank" rel="noopener noreferrer">
                {{ social.title }}
              </a>
            </li>
          </ul>
        </nav>
      </div>
    </header>

    <!-- 主内容区域 -->
    <main class="main-content">
      <section class="hero-section" ref="heroSection">
        <div class="hero-container" :style="{ transform: `translateY(-${currentSection * 100}vh)` }">
          <div class="hero-slide first-slide">
            <div class="hero-content">
              <div class="hero-text">
                <div class="logo-badge">
                  <img src="@/assets/images/logo.png"/>
                </div>
                <h1 class="hero-title">
                  <span class="title-line">不含碘</span>
                  <span class="subtitle-line">你是不是漏掉了什么？</span>
                </h1>
              </div>
              <div class="hero-image">
                <img src="@/assets/downloaded_images/01-packung-fallback-mobile@2x.png" alt="阿尔卑斯碘盐">
              </div>
            </div>
          </div>
          
          <div class="hero-slide second-slide">
            <div class="info-content">
              <h2 class="info-title">碘很重要</h2>
              <div class="info-text">
                <p>理想情况下，我们的土壤应该富含碘，以便动植物能够吸收。<strong>然而，由于德国土壤中碘含量极低甚至不含碘，我们的食物中也缺乏碘。</strong>尽管过去20年来，通过教育和使用碘盐，碘摄入量有所提高，但专家表示，目前的摄入量仍然远未达到理想水平。每日碘摄入量不足约占德国营养学会（DGE）推荐量的三分之二。<strong>德国是一个碘缺乏地区！</strong>对于那些因自身情况或饮食习惯需要依赖碘补充摄入的人来说，提高对碘缺乏的认识尤为重要。如果您感于以下列出的群体之一，请密切关注您的碘摄入水平。</p>
              </div>
              <div class="learn-more">
                <button @click="showMoreInfo = !showMoreInfo">
                  <span>{{ showMoreInfo ? '收起' : '了解更多' }}</span>
                </button>
                <div class="references" v-if="showMoreInfo">
                  <ul>
                    <li><a href="https://jodmangel.de/ausreichende-jodversorgung/#situation" target="_blank">jodmangel.de</a></li>
                    <li><a href="https://www.bmleh.de" target="_blank">bmleh.de</a></li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- 为什么需要碘 -->
      <section class="content-section" id="warum-jod">
        <div class="section-container">
          <div class="text-with-image">
            <img src="@/assets\downloaded_images/01-warum-brauchen-wir-jod-mobile.jpg" alt="我们为什么需要碘">
            <div class="text-content">
              <h3>我们为什么需要碘？</h3>
              <p>
                <strong>碘主要作用于我们的甲状腺。</strong>
                这个形似蝴蝶的小器官位于颈部前方，喉咙下方。甲状腺对新陈代谢至关重要。
                碘是甲状腺激素的组成部分，因此在能量代谢、心律、血压以及生长发育和大脑发育中都发挥着关键作用。
              </p>
            </div>
          </div>
        </div>
      </section>

      <!-- 碘缺乏症状 -->
      <section class="content-section" id="jodmangel">
        <div class="section-container">
          <div class="text-with-image alternative">
            <!-- <img src="@/assets\downloaded_images/iodine-deficiency.jpg" alt="碘缺乏症"> -->
            <div class="text-content">
              <h3>碘缺乏症会怎样？</h3>
              <p>体内碘含量不足会扰乱新陈代谢，导致器官损伤和疾病。持续疲劳、注意力难以集中和缺乏精力可能是碘缺乏的症状。</p>
              
              <div class="symptoms-grid">
                <div class="symptom-group">
                  <h4>成人碘缺乏的可能症状：</h4>
                  <ul>
                    <li v-for="(symptom, index) in adultSymptoms" :key="`adult-${index}`">
                      {{ symptom }}
                    </li>
                  </ul>
                </div>
                
                <div class="symptom-group">
                  <h4>儿童碘缺乏的可能症状：</h4>
                  <ul>
                    <li v-for="(symptom, index) in childSymptoms" :key="`child-${index}`">
                      {{ symptom }}
                    </li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- 碘摄入量计算器 -->
      <section class="calculator-section" id="jod-und-nahrung">
        <div class="section-container">
          <h2><span>碘</span><span>营养</span></h2>
          
          <div class="calculator">
            <div class="calculator-instructions">
              <h3>我应该摄入多少碘？</h3>
              <p>如今，食物种类之丰富以及人们对美味健康食品的热情空前高涨。然而，碘缺乏症仍然可能发生。</p>
            </div>
            
            <form class="calculator-form">
              <div class="form-group">
                <fieldset>
                  <legend>人群</legend>
                  <div class="radio-group">
                    <label v-for="(group, index) in iodineGroups" :key="index">
                      <input 
                        type="radio" 
                        v-model="selectedGroup" 
                        :value="group.id"
                      >
                      {{ group.label }}
                    </label>
                  </div>
                </fieldset>
              </div>
              
              <div class="form-group" v-if="selectedGroup">
                <fieldset>
                  <legend>年龄</legend>
                  <select v-model="selectedAge" class="age-select">
                    <option v-for="(age, index) in getAgesForGroup" :key="index" :value="age.value">
                      {{ age.label }}
                    </option>
                  </select>
                </fieldset>
              </div>
              
              <div class="result" v-if="recommendedIodine">
                <h4>您的碘摄入量建议：</h4>
                <div class="iodine-amount">
                  <span class="amount">{{ recommendedIodine.amount }}</span>
                  <span class="unit">{{ recommendedIodine.unit }}</span>
                </div>
                <p class="note">*此为估算值</p>
              </div>
            </form>
          </div>
        </div>
      </section>

      <!-- 食物碘含量 -->
      <section class="food-section" id="lebensmittel">
        <div class="section-container">
          <h3>食物中含有多少碘？</h3>
          <div class="food-grid">
            <div class="food-item" v-for="(food, index) in foods" :key="index">
              <img :src="food.image" :alt="food.name">
              <div class="food-info">
                <h4>{{ food.name }}</h4>
                <p class="iodine-content">{{ food.iodine }} 微克/100克</p>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- 特殊人群 -->
      <section class="groups-section" id="personen">
        <div class="section-container">
          <h2><span>有</span><span>碘需求增加</span></h2>
          
          <div class="groups-tabs">
            <div class="tab-header">
              <button 
                v-for="(group, index) in specialGroups" 
                :key="index"
                :class="{ active: activeTab === group.id }"
                @click="activeTab = group.id"
              >
                {{ group.label }}
              </button>
            </div>
            
            <div class="tab-content">
              <div v-for="(group, index) in specialGroups" :key="index">
                <div class="group-info" v-if="activeTab === group.id">
                  <div class="group-image">
                    <img :src="group.image" :alt="group.label">
                  </div>
                  <div class="group-text">
                    <h3>{{ group.title }}</h3>
                    <p>{{ group.description }}</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- FAQ -->
      <section class="faq-section" id="faq">
        <div class="section-container">
          <h2><span>常问问题</span></h2>
          
          <div class="faq-grid">
            <div 
              class="faq-item" 
              v-for="(faq, index) in faqs" 
              :key="index"
              @click="toggleFaq(index)"
            >
              <div class="faq-question">
                <h3>{{ faq.question }}</h3>
                <span class="faq-toggle">{{ activeFaqIndex === index ? '−' : '+' }}</span>
              </div>
              <div class="faq-answer" v-if="activeFaqIndex === index">
                <p v-html="faq.answer"></p>
              </div>
            </div>
          </div>
        </div>
      </section>
    </main>

    <!-- 页脚 -->
    <footer class="site-footer">
      <div class="footer-container">
        <nav class="footer-nav">
          <ul>
            <li><a href="/impressum">印记</a></li>
            <li><a href="/datenschutz">数据保护</a></li>
          </ul>
        </nav>
        
        <div class="footer-claim">
          <span>阿尔卑斯山的宝藏</span>
        </div>
        
        <div class="footer-social">
          <a 
            v-for="(social, index) in socialMedia" 
            :key="index"
            :href="social.url" 
            :title="social.title"
            target="_blank"
            rel="noopener noreferrer"
          >
            {{ social.title }}
          </a>
        </div>
      </div>
    </footer>

    <!-- 返回顶部按钮 -->
    <a 
      v-show="showBackToTop" 
      href="#top" 
      class="back-to-top" 
      @click.prevent="scrollToTop"
    >
      向上
    </a>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'

// 响应式数据
const navExpanded = ref(false)
const activeNavGroup = ref(null)
const showMoreInfo = ref(false)
const selectedGroup = ref('adults')
const selectedAge = ref('25-51')
const activeTab = ref('pregnant')
const activeFaqIndex = ref(null)
const showBackToTop = ref(false)
const currentSection = ref(0)
const heroSection = ref(null)
let isScrolling = false

// 导航数据
const navGroups = [
  {
    id: 'warum-jod',
    label: '为什么要加碘？',
    items: [
      { href: '#warum-jod', title: '为什么要用碘？', text: '为什么要用碘？' },
      { href: '#jodmangel', title: '碘缺乏症', text: '碘缺乏症' },
      { href: '#schilddruese', title: '甲状腺功能', text: '甲状腺功能' }
    ]
  },
  {
    id: 'jod-und-nahrung',
    label: '碘与食物',
    items: [
      { href: '#jod-und-nahrung', title: '碘与食物', text: '碘与食物' },
      { href: '#wieviel', title: '碘含量是多少？', text: '碘含量是多少？' },
      { href: '#lebensmittel', title: '食物中的碘', text: '食物中的碘' },
      { href: '#jodsalz', title: '碘盐', text: '碘盐' }
    ]
  },
  {
    id: 'personen',
    label: '人',
    items: [
      { href: '#personen', title: '人', text: '人' },
      { href: '#schwangere-und-stillende', title: '孕妇和哺乳期妇女', text: '孕妇和哺乳期妇女' },
      { href: '#kinder', title: '孩子们', text: '孩子们' },
      { href: '#sportler', title: '运动员', text: '运动员' },
      { href: '#vegetarier-und-veganer', title: '素食者和纯素食者', text: '素食者和纯素食者' },
      { href: '#senioren', title: '老年人', text: '老年人' }
    ]
  },
  {
    id: 'faq',
    label: '常问问题',
    items: [
      { href: '#faq', title: '常问问题', text: '常问问题' }
    ]
  }
]

// 社交媒体
const socialMedia = [
//   { url: 'https://www.facebook.com/Bad-Reichenhaller-677439229118646/', title: 'Facebook上的Bad Reichenhaller' },
//   { url: 'https://www.instagram.com/badreichenhaller', title: 'Instagram上的Bad Reichenhaller' },
//   { url: 'https://www.tiktok.com/@bad_reichenhaller', title: 'TikTok上的Bad Reichenhaller' },
//   { url: 'https://www.pinterest.com/reichenhaller/', title: 'Pinterest上的Bad Reichenhaller' },
//   { url: 'https://www.youtube.com/channel/UCSTi2WaLc6MgKWKSOHrPXiA', title: 'YouTube上的Bad Reichenhaller' }
]

// 症状数据
const adultSymptoms = [
  '注意力不集中',
  '缺乏动力',
  '持续疲劳',
  '抑郁情绪',
  '易怒',
  '干燥、脱皮的皮肤',
  '消化不良',
  '免疫防御能力弱',
  '水分保持'
]

const childSymptoms = [
  '精神运动发育障碍',
  '听力障碍',
  '智力表现下降',
  '学习困难',
  '情绪波动',
  '消化不良'
]

// 碘摄入量计算器数据
const iodineGroups = [
  { id: 'infants', label: '婴儿', ages: [
    { label: '0至4个月以下', value: '0-4', amount: 80 },
    { label: '4至12个月以下', value: '4-12', amount: 80 }
  ]},
  { id: 'children', label: '儿童', ages: [
    { label: '1岁至4岁以下', value: '1-4', amount: 90 },
    { label: '4岁至7岁以下', value: '4-7', amount: 90 },
    { label: '7岁至10岁以下', value: '7-10', amount: 120 },
    { label: '10岁至13岁以下', value: '10-13', amount: 120 },
    { label: '13岁至15岁以下', value: '13-15', amount: 150 }
  ]},
  { id: 'adults', label: '青少年及成人', ages: [
    { label: '15岁至19岁以下', value: '15-19', amount: 150 },
    { label: '19岁至25岁以下', value: '19-25', amount: 150 },
    { label: '25岁至51岁以下', value: '25-51', amount: 150 },
    { label: '51岁至65岁以下', value: '51-65', amount: 150 },
    { label: '65岁及以上', value: '65+', amount: 150 }
  ]},
  { id: 'pregnant', label: '孕妇', ages: [
    { label: '所有孕妇', value: 'all', amount: 220 }
  ]},
  { id: 'breastfeeding', label: '哺乳期妇女', ages: [
    { label: '所有哺乳期妇女', value: 'all', amount: 230 }
  ]}
]

// 食物数据
const foods = [
//   { name: '鳕鱼', image: '@/assets\downloaded_images/codfish.jpg', iodine: '258' },
//   { name: '西兰花', image: '@/assets\downloaded_images/broccoli.jpg', iodine: '17' },
//   { name: '鸡蛋', image: '@/assets\downloaded_images/egg.jpg', iodine: '9' },
//   { name: '三文鱼', image: '@/assets\downloaded_images/salmon.jpg', iodine: '6' },
//   { name: '桃子', image: '@/assets\downloaded_images/peach.jpg', iodine: '3' }
]

// 特殊人群数据
const specialGroups = [
  { 
    id: 'pregnant', 
    label: '孕妇和哺乳期妇女',
    title: '孕妇和哺乳期妇女',
    // image: '@/assets\downloaded_images/pregnant-women.jpg',
    description: '孕妇碘需求量的增加源于甲状腺激素分泌的增加以及胎儿对碘供应的迫切需求...'
  },
  { 
    id: 'children', 
    label: '儿童',
    title: '儿童',
    // image: '@/assets\downloaded_images/children.jpg',
    description: '孩子们精力充沛！为了让他们尽情奔跑玩耍，并帮助他们在学校集中注意力...'
  },
  // ... 其他人群
]

// FAQ数据
const faqs = [
  {
    question: '碘有什么作用？',
    answer: '甲状腺利用微量元素碘和蛋白质作为构建原料来合成甲状腺激素。这些激素调节着许多重要的生理功能...'
  },
  {
    question: '如何判断是否缺碘？',
    answer: '以下症状可能是甲状腺功能亢进或甲状腺功能减退的早期征兆：注意力难以集中、缺乏动力...'
  },
  // ... 其他FAQ
]

// 计算属性
const getAgesForGroup = computed(() => {
  const group = iodineGroups.find(g => g.id === selectedGroup.value)
  return group ? group.ages : []
})

const recommendedIodine = computed(() => {
  if (!selectedGroup.value || !selectedAge.value) return null
  
  const group = iodineGroups.find(g => g.id === selectedGroup.value)
  if (!group) return null
  
  const ageInfo = group.ages.find(a => a.value === selectedAge.value)
  if (!ageInfo) return null
  
  return {
    amount: ageInfo.amount,
    unit: '微克/天'
  }
})

// 方法
const toggleFaq = (index) => {
  activeFaqIndex.value = activeFaqIndex.value === index ? null : index
}

const scrollToTop = () => {
  window.scrollTo({
    top: 0,
    behavior: 'smooth'
  })
}

const handleScroll = () => {
  showBackToTop.value = window.scrollY > 300
}

const handleWheel = (e) => {
  if (!heroSection.value || isScrolling) return
  
  const heroRect = heroSection.value.getBoundingClientRect()
  const heroTop = heroRect.top
  
  if (heroTop <= 10 && heroTop >= -10) {
    if (e.deltaY > 0 && currentSection.value === 0) {
      e.preventDefault()
      isScrolling = true
      currentSection.value = 1
      setTimeout(() => { isScrolling = false }, 600)
    } else if (e.deltaY < 0 && currentSection.value === 1) {
      e.preventDefault()
      isScrolling = true
      currentSection.value = 0
      setTimeout(() => { isScrolling = false }, 600)
    } else if (e.deltaY < 0 && currentSection.value === 0) {
      return
    }
  }
}

// 生命周期钩子
onMounted(() => {
  window.addEventListener('scroll', handleScroll)
  window.addEventListener('wheel', handleWheel, { passive: false })
  
  onUnmounted(() => {
    window.removeEventListener('scroll', handleScroll)
    window.removeEventListener('wheel', handleWheel)
  })
})
</script>

<style scoped>
/* 基础样式 */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  line-height: 1.6;
  color: #333;
}

#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

/* 头部样式 */
.site-header {
  background: #367ec1;

  position: sticky;
  top: 0;
  z-index: 1000;
}

.header-container {
  max-width: 100%;
  margin: 0;
  padding: 0.5rem 2rem;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.logo img {
  height: 40px;
  width: auto;
}

.nav-toggle {
  display: none;
  cursor: pointer;
  color: white;
}

.nav-toggle span {
  padding: 0.5rem 1rem;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 4px;
  color: white;
}

/* 主导航 */
.main-nav {
  display: flex;
  align-items: center;
  gap: 2rem;
}

.nav-list {
  display: flex;
  list-style: none;
  gap: 1.5rem;
}

.nav-item {
  position: relative;
}

.nav-item label {
  cursor: pointer;
  padding: 0.5rem 1rem;
  display: block;
  color: white;
  font-weight: 500;
  transition: opacity 0.3s;
}

.nav-item label:hover {
  opacity: 0.8;
}

.nav-item input[type="radio"] {
  display: none;
}

.submenu {
  position: absolute;
  top: 100%;
  left: 0;
  background: white;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
  border-radius: 4px;
  min-width: 200px;
  padding: 0.5rem 0;
  list-style: none;
  display: none;
}

.nav-item input:checked + label + .submenu {
  display: block;
}

.submenu li a {
  display: block;
  padding: 0.5rem 1rem;
  color: #333;
  text-decoration: none;
  transition: background 0.3s;
}

.submenu li a:hover {
  background: #f5f5f5;
}

.social-media {
  display: flex;
  list-style: none;
  gap: 1rem;
}

.social-media li a {
  color: white;
  text-decoration: none;
  font-size: 1.2rem;
  transition: opacity 0.3s;
}

.social-media li a:hover {
  opacity: 0.8;
}

/* 主内容区域 */
.main-content {
  flex: 1;
}

/* Hero区域 */
.hero-section {
  background: #367ec1;
  color: white;
  height: 100vh;
  position: relative;
  overflow: hidden;
}

.hero-container {
  height: 100vh;
  width: 100%;
  transition: transform 0.6s cubic-bezier(0.4, 0, 0.2, 1);
}

.hero-slide {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 3rem 2rem;
}

.first-slide .hero-content {
  max-width: 1200px;
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 4rem;
}

.hero-text {
  flex: 1;
}

.logo-badge {
  color: #d32f2f;
  padding: 0.5rem 1.5rem;
  border-radius: 20px;
  display: inline-block;
  margin-bottom: 2rem;
  font-weight: bold;
  font-size: 0.9rem;
}

.logo-badge img {
  height: 30px;
  width: auto;
}

.hero-title {
  font-size: 3.5rem;
  line-height: 1.2;
  margin: 0;
}

.title-line {
  display: block;
  font-weight: 700;
}

.subtitle-line {
  display: block;
  font-size: 2rem;
  font-weight: 400;
  margin-top: 0.5rem;
  border-bottom: 3px solid white;
  padding-bottom: 0.5rem;
  display: inline-block;
}

.hero-image {
  flex: 0 0 auto;
}

.hero-image img {
  max-width: 300px;
  height: auto;
  filter: drop-shadow(0 10px 30px rgba(0,0,0,0.3));
}

.second-slide {
  background: linear-gradient(135deg, #367ec1 0%, #2c6aa3 100%);
}

.info-content {
  max-width: 900px;
  text-align: center;
}

.info-title {
  font-size: 3rem;
  margin-bottom: 2rem;
  background: white;
  color: #367ec1;
  padding: 1rem 2rem;
  display: inline-block;
  border-radius: 8px;
}

.info-text {
  font-size: 1.1rem;
  line-height: 1.8;
  margin-bottom: 2rem;
}

.info-text strong {
  color: #ffd700;
}

.learn-more {
  margin-top: 2rem;
}

.learn-more button {
  background: white;
  color: #367ec1;
  border: none;
  padding: 1rem 2rem;
  font-size: 1rem;
  border-radius: 25px;
  cursor: pointer;
  font-weight: bold;
  transition: transform 0.3s;
}

.learn-more button:hover {
  transform: scale(1.05);
}

.references {
  margin-top: 1rem;
  background: rgba(255,255,255,0.1);
  padding: 1rem;
  border-radius: 8px;
}

.references ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.references li {
  margin: 0.5rem 0;
}

.references a {
  color: white;
  text-decoration: underline;
}

/* 内容区域通用样式 */
.content-section, 
.calculator-section,
.food-section,
.groups-section,
.faq-section {
  padding: 4rem 2rem;
}

.section-container {
  max-width: 1200px;
  margin: 0 auto;
}

.text-with-image {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 3rem;
  align-items: center;
}

.text-with-image.alternative {
  grid-template-columns: 1fr 1fr;
  direction: rtl;
}

.text-with-image.alternative .text-content {
  direction: ltr;
}

.text-with-image img {
  width: 100%;
  height: auto;
  border-radius: 8px;
}

.text-content h3 {
  font-size: 2rem;
  margin-bottom: 1.5rem;
  color: #2c5282;
}

.text-content p {
  margin-bottom: 1.5rem;
  font-size: 1.1rem;
  line-height: 1.8;
}

/* 症状网格 */
.symptoms-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
  margin-top: 2rem;
}

.symptom-group h4 {
  margin-bottom: 1rem;
  color: #2d3748;
}

.symptom-group ul {
  list-style-position: inside;
}

.symptom-group li {
  margin-bottom: 0.5rem;
}

/* 计算器样式 */
.calculator {
  background: #f8fafc;
  border-radius: 12px;
  padding: 2rem;
  margin-top: 2rem;
}

.calculator-instructions {
  margin-bottom: 2rem;
}

.calculator-form {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.form-group fieldset {
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  padding: 1rem;
}

.form-group legend {
  padding: 0 0.5rem;
  font-weight: 600;
  color: #4a5568;
}

.radio-group {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
}

.radio-group label {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  cursor: pointer;
  padding: 0.5rem 1rem;
  background: white;
  border: 2px solid #e2e8f0;
  border-radius: 4px;
  transition: all 0.3s;
}

.radio-group label:hover {
  border-color: #4299e1;
}

.radio-group input[type="radio"]:checked + span {
  background: #4299e1;
  border-color: #4299e1;
  color: white;
}

.age-select {
  width: 100%;
  padding: 0.75rem;
  border: 2px solid #e2e8f0;
  border-radius: 4px;
  font-size: 1rem;
}

.result {
  background: white;
  border-radius: 8px;
  padding: 1.5rem;
  text-align: center;
  margin-top: 1rem;
}

.iodine-amount {
  font-size: 2.5rem;
  font-weight: bold;
  color: #2d3748;
  margin: 1rem 0;
}

.iodine-amount .unit {
  font-size: 1.5rem;
  color: #718096;
}

.note {
  font-size: 0.9rem;
  color: #a0aec0;
  margin-top: 0.5rem;
}

/* 食物网格 */
.food-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 2rem;
  margin-top: 2rem;
}

.food-item {
  background: white;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 4px 6px rgba(0,0,0,0.1);
  transition: transform 0.3s;
}

.food-item:hover {
  transform: translateY(-4px);
}

.food-item img {
  width: 100%;
  height: 150px;
  object-fit: cover;
}

.food-info {
  padding: 1rem;
  text-align: center;
}

.food-info h4 {
  margin-bottom: 0.5rem;
  color: #2d3748;
}

.iodine-content {
  font-weight: bold;
  color: #2c5282;
}

/* 特殊人群标签页 */
.groups-tabs {
  margin-top: 2rem;
}

.tab-header {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin-bottom: 2rem;
}

.tab-header button {
  padding: 0.75rem 1.5rem;
  background: #e2e8f0;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.3s;
  font-size: 1rem;
}

.tab-header button.active {
  background: #4299e1;
  color: white;
}

.tab-header button:hover:not(.active) {
  background: #cbd5e0;
}

.tab-content .group-info {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 3rem;
  align-items: center;
}

.group-image img {
  width: 100%;
  height: auto;
  border-radius: 8px;
}

.group-text h3 {
  font-size: 2rem;
  margin-bottom: 1rem;
  color: #2c5282;
}

.group-text p {
  font-size: 1.1rem;
  line-height: 1.8;
}

/* FAQ样式 */
.faq-grid {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  margin-top: 2rem;
}

.faq-item {
  background: white;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  cursor: pointer;
}

.faq-question {
  padding: 1.5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #f8fafc;
  transition: background 0.3s;
}

.faq-question:hover {
  background: #edf2f7;
}

.faq-question h3 {
  margin: 0;
  font-size: 1.25rem;
  color: #2d3748;
}

.faq-toggle {
  font-size: 1.5rem;
  font-weight: bold;
  color: #4299e1;
}

.faq-answer {
  padding: 1.5rem;
  border-top: 1px solid #e2e8f0;
  animation: fadeIn 0.3s ease;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

/* 页脚样式 */
.site-footer {
  background: #2d3748;
  color: white;
  padding: 3rem 2rem;
}

.footer-container {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2rem;
}

.footer-nav ul {
  display: flex;
  list-style: none;
  gap: 2rem;
}

.footer-nav a {
  color: #cbd5e0;
  text-decoration: none;
  transition: color 0.3s;
}

.footer-nav a:hover {
  color: white;
}

.footer-claim {
  font-size: 1.5rem;
  font-weight: bold;
  text-align: center;
}

.footer-social {
  display: flex;
  gap: 1rem;
  flex-wrap: wrap;
  justify-content: center;
}

.footer-social a {
  color: #cbd5e0;
  text-decoration: none;
  padding: 0.5rem 1rem;
  border: 1px solid #4a5568;
  border-radius: 4px;
  transition: all 0.3s;
}

.footer-social a:hover {
  color: white;
  border-color: white;
}

/* 返回顶部按钮 */
.back-to-top {
  position: fixed;
  bottom: 2rem;
  right: 2rem;
  background: #4299e1;
  color: white;
  width: 3rem;
  height: 3rem;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  text-decoration: none;
  box-shadow: 0 2px 8px rgba(0,0,0,0.2);
  transition: all 0.3s;
  z-index: 1000;
}

.back-to-top:hover {
  background: #3182ce;
  transform: translateY(-2px);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .nav-toggle {
    display: block;
  }
  
  .main-nav {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100vh;
    background: white;
    flex-direction: column;
    padding: 2rem;
    transform: translateX(-100%);
    transition: transform 0.3s ease;
  }
  
  #navigation-expanded:checked ~ .main-nav {
    transform: translateX(0);
  }
  
  .nav-list {
    flex-direction: column;
    width: 100%;
  }
  
  .text-with-image,
  .text-with-image.alternative,
  .tab-content .group-info {
    grid-template-columns: 1fr;
  }
  
  .symptoms-grid {
    grid-template-columns: 1fr;
  }
  
  .tab-header {
    flex-direction: column;
  }
  
  .tab-header button {
    width: 100%;
    text-align: left;
  }
  
  .slide-content h1 {
    font-size: 2rem;
  }
  
  .slide-content h1 .subline {
    font-size: 1.2rem;
  }
}

@media (max-width: 480px) {
  .food-grid {
    grid-template-columns: 1fr;
  }
  
  .content-section,
  .calculator-section,
  .food-section,
  .groups-section,
  .faq-section {
    padding: 2rem 1rem;
  }
  
  .calculator {
    padding: 1rem;
  }
}
</style>