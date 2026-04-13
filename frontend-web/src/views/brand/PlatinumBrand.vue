<template>
  <div class="platinum-brand-page">
    <nav
      class="platinum-nav"
      :class="{ scrolled: isScrolled, 'mobile-menu-open': isMobileMenuOpen }"
    >
      <div class="nav-container">
        <router-link to="/" class="nav-logo">
          <img
            src="@/assets/images/logo.png"
            alt="ALPENSALZ"
            class="logo-img"
          />
        </router-link>

        <!-- Mobile Menu Toggle -->
        <button
          class="mobile-menu-toggle"
          @click="toggleMobileMenu"
          aria-label="Menu"
        >
          <span class="hamburger-line"></span>
          <span class="hamburger-line"></span>
          <span class="hamburger-line"></span>
        </button>

        <div class="nav-links">
          <a href="#" @click.prevent="scrollToSection('heritage')">皇室传承</a>
          <a href="#" @click.prevent="scrollToSection('purity')">极境纯净</a>
          <a href="#" @click.prevent="scrollToSection('products')">白金系列</a>
          <a href="#" @click.prevent="scrollToSection('sustainability')"
            >可持续</a
          >
        </div>
      </div>
    </nav>

    <!-- Hero Section -->
    <header class="hero-section" id="hero">
      <div
        class="hero-bg"
        :style="{
          backgroundImage: `url(${heroBgLoaded ? heroBg : heroBgPlaceholder})`,
        }"
      ></div>
      <div class="hero-content">
        <h1 class="hero-title" data-aos="fade-up">Alpen 白金盐</h1>
        <p class="hero-subtitle" data-aos="fade-up" data-aos-delay="200">
          阿尔卑斯山的“白金”：一天 5g 的健康奢侈品
        </p>
        <div class="hero-separator"></div>
        <p class="hero-desc" data-aos="fade-up" data-aos-delay="400">
          Alpen 白金盐 · 为健康代“盐”
        </p>
      </div>
      <div class="scroll-indicator">
        <span>SCROLL</span>
        <div class="line"></div>
      </div>
    </header>

    <!-- Heritage Section -->
    <section id="heritage" class="section heritage-section">
      <div class="container">
        <div class="grid-2">
          <div class="content-box" data-aos="fade-right">
            <span class="section-label">TREASURES FROM THE ALPS</span>
            <h2 class="section-title">
              原始海洋蒸发<br />储存在阿尔卑斯山的珍宝
            </h2>
            <p class="text-body">
              沉睡在阿尔卑斯山深处！千百年来，宝贵的盐随山泉水缓释流出，形成天然纯净珍贵的高山盐水。
            </p>
            <p class="text-body">
              阿尔卑斯山白金盐 AlpenSalz 作为德国 Bad Reichenhaller
              的标志性产品，以其独特的蓝色包装和白色菱形图案而闻名，代表了阿尔卑斯山的财富。
            </p>
            <!-- <p class="text-body">
              贝希特斯加登盐矿在 1517
              年由贝希特斯加登国王马克斯一世创立，供皇室享用。
              贝希特斯加登盐矿，建于 1517 年，500
              多年采矿历史，也是德国最重要的旅游景点之一。
            </p> -->
            <p class="text-body">
              德国西南盐业股份公司是德国最大的制盐企业，德国市场领导品牌。
            </p>
            <div class="quote-box">
              "食盐（NaCl）人类第四个维持生命的要素，每日健康必须。"
              <cite>— 品牌代言人茜茜公主 (德国公主奥地利皇后)</cite>
            </div>
          </div>
          <div class="image-box" data-aos="fade-left">
            <img
              :src="heritageImg"
              alt="Ancient Mine"
              class="parallax-img"
              loading="lazy"
              decoding="async"
            />
          </div>
        </div>

        <!-- Timeline -->
        <div class="timeline-wrapper" data-aos="fade-up">
          <div class="timeline-track"></div>
          <div
            class="timeline-item"
            v-for="(item, index) in timeline"
            :key="index"
            :class="{ reverse: index % 2 !== 0 }"
          >
            <div class="timeline-dot"></div>
            <div class="timeline-date">{{ item.year }}</div>
            <div class="timeline-content">
              <div class="timeline-image-container" v-if="item.image">
                <img
                  :src="item.image"
                  :alt="item.title"
                  loading="lazy"
                  decoding="async"
                />
              </div>
              <h3>{{ item.title }}</h3>
              <p>{{ item.desc }}</p>
            </div>
          </div>
        </div>

        <div class="china-market" data-aos="fade-up">
          <div class="china-market-content">
            <span class="section-label">CHINA</span>
            <h2 class="section-title">2017 年<br />白金盐进入中国市场</h2>
            <p class="text-body">进口商：中盐集团、广盐集团</p>
            <ul class="china-market-points">
              <li>广泛上架中高端商超、知名连锁超市、百货公司销售</li>
              <li>尤其覆盖北上广深地区</li>
            </ul>
          </div>
          <div class="china-market-gallery">
            <picture class="china-market-photo">
              <source :srcset="china01WebpSrcset" type="image/webp" />
              <img
                :src="china01FallbackSrc"
                alt="2017 白金盐进入中国市场"
                loading="lazy"
                decoding="async"
              />
            </picture>
            <picture class="china-market-photo">
              <source :srcset="china02WebpSrcset" type="image/webp" />
              <img
                :src="china02FallbackSrc"
                alt="渠道上架展示"
                loading="lazy"
                decoding="async"
              />
            </picture>
            <picture class="china-market-photo">
              <source :srcset="china03WebpSrcset" type="image/webp" />
              <img
                :src="china03FallbackSrc"
                alt="商超渠道展示"
                loading="lazy"
                decoding="async"
              />
            </picture>
            <picture class="china-market-photo">
              <source :srcset="china04WebpSrcset" type="image/webp" />
              <img
                :src="china04FallbackSrc"
                alt="北上广深地区销售"
                loading="lazy"
                decoding="async"
              />
            </picture>
          </div>
        </div>
      </div>
    </section>
    <section id="timeline" class="animate-section">
      <TimelineSection />
    </section>
    <!-- Purity Section (Dark Mode) -->
    <section id="purity" class="section purity-section">
      <div class="purity-bg-overlay"></div>
      <div class="container">
        <div class="centered-header" data-aos="fade-up">
          <span class="section-label gold">PURITY & QUALITY</span>
          <h2 class="section-title white">食盐中的钻石<br />纯净度 99.9%</h2>
          <p class="section-subtitle white">
            越纯净越健康 · 把健康带回家！<br />Alpen 白金盐 ·
            赋予您有滋有味的健康生活！
          </p>
        </div>

        <!-- Core Features -->
        <div class="features-grid">
          <div class="feature-card" data-aos="fade-up" data-aos-delay="100">
            <div class="feature-icon">💎</div>
            <h3>99.9% 纯度</h3>
            <p>
              原始海洋蒸发储存在阿尔卑斯山的珍宝，隔绝了现代污染，完全不受核辐射影响。
            </p>
          </div>
          <div class="feature-card" data-aos="fade-up" data-aos-delay="200">
            <div class="feature-icon">🛡️</div>
            <h3>少吃盐·吃好盐</h3>
            <p>世界卫生组织倡导减盐！您的“盐”值关系到全家健康！</p>
          </div>
          <div class="feature-card" data-aos="fade-up" data-aos-delay="300">
            <div class="feature-icon">🌿</div>
            <h3>每日 5g 必需</h3>
            <p>
              食盐（NaCl）人类第四个维持生命的要素，每日健康必须。不可缺少，不可替代！
            </p>
          </div>
        </div>

        <!-- Authority Endorsement -->
        <div class="authority-grid" data-aos="fade-up" data-aos-delay="400">
          <div class="authority-item">
            <div class="auth-icon">🏛️</div>
            <h4>WHO 减盐模型</h4>
            <p>2023指南：高纯度低杂质盐，保持咸度前提下可减少 8-12% 钠摄入。</p>
          </div>
          <div class="authority-item">
            <div class="auth-icon">📊</div>
            <h4>德国市场领导品牌</h4>
            <p>
              德国西南盐业股份公司是德国最大的制盐企业，Alpen白金盐·德国No.1。
            </p>
          </div>
          <div class="authority-item">
            <div class="auth-icon">🔬</div>
            <h4>TU Dresden 实验室</h4>
            <p>双盲测试：苦度值 ↓23%，咸度圆润度 ↑19%，后味金属感 ↓31%。</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Visual Interlude -->
    <section
      class="visual-interlude"
      :style="{ backgroundImage: `url(${interludeBg})` }"
    >
      <div class="interlude-content">
        <h3>"可以品尝的钻石"</h3>
        <p>一闪一闪亮晶晶<br />一盐一世界 · 一味一安康</p>
      </div>
    </section>

    <!-- Scenes Section -->
    <section id="scenes" class="section scenes-section">
      <div class="container">
        <div class="centered-header" data-aos="fade-up">
          <span class="section-label">DAILY ESSENTIAL</span>
          <h2 class="section-title">阳光·空气·水·盐<br />每日健康必需</h2>
        </div>
        <div class="scenes-grid">
          <div
            class="scene-card"
            v-for="(scene, index) in scenes"
            :key="index"
            data-aos="fade-up"
            :data-aos-delay="index * 100"
          >
            <div class="scene-time">{{ scene.time }}</div>
            <div class="scene-info">
              <h3>{{ scene.title }}</h3>
              <p>{{ scene.desc }}</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Products Section -->
    <section id="products" class="section products-section">
      <div class="container">
        <div class="grid-2 reverse-mobile">
          <div class="image-box product-showcase" data-aos="zoom-in">
            <img
              :src="productImg"
              alt="Platinum Salt Product"
              loading="lazy"
              decoding="async"
            />
          </div>
          <div class="content-box" data-aos="fade-left">
            <span class="section-label">COLLECTION</span>
            <h2 class="section-title">白金系列<br />现代御膳</h2>
            <p class="text-body">
              500年采矿史、德国工匠精神，全部浓缩成这一罐。
              今天，你舀起的那一勺，正是昨日御厨呈给国王的同一座山、同一层矿脉。
            </p>
            <ul class="product-features">
              <li><span>富含矿物质</span> — 铁、钙、镁、钾等多种微量元素</li>
              <li><span>天然抑菌</span> — 恒温24℃、湿度45%的矿洞环境</li>
              <li><span>细腻口感</span> — 独特的晶体结构，咸而不苦</li>
            </ul>
            <!-- <button class="cta-button">探索全系列</button> -->
          </div>
        </div>
      </div>
    </section>

    <!-- QA Section -->
    <section id="qa" class="section qa-section">
      <div class="container">
        <div class="grid-2">
          <div class="qa-header" data-aos="fade-right">
            <span class="section-label">FAQ</span>
            <h2 class="section-title">常见问题<br />解析</h2>
          </div>
          <div class="qa-list" data-aos="fade-left">
            <div class="qa-item" v-for="(item, index) in qaList" :key="index">
              <div class="qa-question">
                <span class="q-mark">Q.</span>
                {{ item.q }}
              </div>
              <div class="qa-answer">{{ item.a }}</div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Sustainability Section -->
    <section
      id="sustainability"
      class="section sustain-section"
      :style="{ backgroundImage: `url(${sustainBg})` }"
    >
      <div class="container">
        <div class="sustain-content" data-aos="fade-up">
          <span class="section-label">SUSTAINABILITY</span>
          <h2 class="section-title">承诺未来</h2>
          <div class="sustain-grid">
            <div class="sustain-item">
              <div class="big-number">100%</div>
              <h4>皇室绿电</h4>
              <p>2025 起 100% 巴伐利亚水电驱动<br />年减碳 2,100 吨</p>
            </div>
            <div class="sustain-item">
              <div class="big-number">♻️</div>
              <h4>空瓶回收</h4>
              <p>6 个空瓶换 250 g 补充装<br />再制"皇冠盐勺"</p>
            </div>
            <div class="sustain-item">
              <div class="big-number">1g=1g</div>
              <h4>一克盐·一克爱</h4>
              <p>每售 1 罐，向山区捐 1 g 碘盐<br />累计捐赠 7.3 吨</p>
            </div>
          </div>

          <div class="brand-manifesto" data-aos="fade-up" data-aos-delay="200">
            <h3>"岁岁年年 | 安安康康"</h3>
            <p>
              ALPEN 白金盐始终在您身边<br />伴您三餐四季 ·
              守护您和家人的幸福滋味！
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- Bottom History Section -->
    <section class="section bottom-history-section">
      <div class="container grid-2">
        <div class="bottom-carousel-wrapper" data-aos="fade-right">
          <button class="nav prev" @click="prevBottom">‹</button>
          <div class="bottom-carousel-window">
            <img
              :src="currentBottomSrc"
              alt="History slide"
              loading="lazy"
              decoding="async"
            />
          </div>
          <button class="nav next" @click="nextBottom">›</button>
        </div>
        <div class="history-text" data-aos="fade-left">
          <p class="text-body">
            贝希特斯加登盐矿在 1517
            年由贝希特斯加登国王马克斯一世创立，供皇室享用。
            贝希特斯加登盐矿，建于 1517 年，500
            多年采矿历史，也是德国最重要的旅游 景点之一。
          </p>
        </div>
      </div>
    </section>

    <footer class="brand-footer">
      <div class="container">
        <div class="footer-logo">ALPENSALZ</div>
        <!-- <p class="copyright">© 2024 Bad Reichenhaller. All Rights Reserved.</p> -->
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, computed } from "vue";
import heroBg from "@/assets/brand/banner.png?w=1600&format=webp&quality=75";
import heroBgPlaceholder from "@/assets/brand/banner.png?w=48&format=webp&quality=30&blur=20&inline";
import heritageImg from "@/assets/brand/01_03_promotionslider_brh-keyvisual_desktop_blauundgelb_01.jpg?w=1200&format=webp&quality=80";
import productImg from "@/assets/brand/20250626袋装盐6-10(4).png?w=1200&format=webp&quality=80";
import interludeBg from "@/assets/images/background.jpg";
import sustainBg from "@/assets/brand/Salzbergwerk-Berchtesgaden-Highlights.194559e5.e2f073ed.jpg";
import TimelineSection from "@/components/brand/TimelineSection.vue";
// Timeline images
import time1517 from "@/assets/brand/Pressebild-Bergler-Sprengung.9f9680d3.jpg?w=1200&format=webp&quality=80";
import time1706 from "@/assets/brand/Salzbergwerk-Berchtesgaden-Highlights.194559e5.e2f073ed.jpg?w=1200&format=webp&quality=80";
import time1960 from "@/assets/brand/Spiegelsee-Familie.b01c4e2b.jpg?w=1200&format=webp&quality=80";
import time2025 from "@/assets/brand/20250612白金盐_袋装盐1-5_低钠盐6-8.png";
import img1 from "@/assets/brand/Spiegelsee-Familie.b01c4e2b.jpg?w=1200&format=webp&quality=80";
import img2 from "@/assets/brand/2017-12-05 090626.jpg?w=1200&format=webp&quality=80";
import img3 from "@/assets/brand/微信图片_20250408203638.jpg?w=1200&format=webp&quality=80";
import img4 from "@/assets/brand/Content-Veranstaltung-Salzkathedrale.feba5f13.jpg?w=1200&format=webp&quality=80";
import china01WebpSrcset from "@/assets/brand/_20260323134841_1644_1 (1).jpg?w=320;480;640&format=webp&quality=78&as=srcset";
import china01FallbackSrc from "@/assets/brand/_20260323134841_1644_1 (1).jpg?w=640&format=jpg&quality=82";
import china02WebpSrcset from "@/assets/brand/_20260323134841_1644_1 (2).jpg?w=320;480;640&format=webp&quality=78&as=srcset";
import china02FallbackSrc from "@/assets/brand/_20260323134841_1644_1 (2).jpg?w=640&format=jpg&quality=82";
import china03WebpSrcset from "@/assets/brand/_20260323134841_1644_1 (3).jpg?w=320;480;640&format=webp&quality=78&as=srcset";
import china03FallbackSrc from "@/assets/brand/_20260323134841_1644_1 (3).jpg?w=640&format=jpg&quality=82";
import china04WebpSrcset from "@/assets/brand/_20260323134841_1644_1 (4).jpg?w=320;480;640&format=webp&quality=78&as=srcset";
import china04FallbackSrc from "@/assets/brand/_20260323134841_1644_1 (4).jpg?w=640&format=jpg&quality=82";

const isScrolled = ref(false);
const isMobileMenuOpen = ref(false);
const heroBgLoaded = ref(false);

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value;
};

const handleScroll = () => {
  isScrolled.value = window.scrollY > 50;
};

const scrollToSection = (sectionId) => {
  isMobileMenuOpen.value = false;
  const element = document.getElementById(sectionId);
  if (element) {
    const navHeight = 80; // 导航栏高度
    const elementPosition = element.getBoundingClientRect().top;
    const offsetPosition = elementPosition + window.pageYOffset - navHeight;

    window.scrollTo({
      top: offsetPosition,
      behavior: "smooth",
    });
  }
};

onMounted(() => {
  window.addEventListener("scroll", handleScroll);

  const img = new Image();
  img.src = heroBg;
  img.onload = () => {
    heroBgLoaded.value = true;
  };

  // Simple intersection observer for fade-in animations if AOS is not available
  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          entry.target.classList.add("visible");
        }
      });
    },
    { threshold: 0.1 },
  );

  document.querySelectorAll("[data-aos]").forEach((el) => {
    el.classList.add("aos-init");
    observer.observe(el);
  });
});

onUnmounted(() => {
  window.removeEventListener("scroll", handleScroll);
});

const timeline = [
  {
    year: "2.5 亿年前",
    title: "原始海洋",
    desc: '蒸发、挤压、隆起，海水被时间浓缩成 1800 米厚的盐层，与阿尔卑斯山一起长高。没有塑料微粒，没有工业排放和核污染——这是地球留给人类的"干净存档"。',
    image: img1,
  },

  {
    year: "1706 年",
    title: '茜茜公主的"白金盐"礼盒',
    desc: '巴伐利亚国王路德维希二世为表姐茜茜公主定制"白金盐"婚嫁礼盒，被维也纳宫廷女官记入《御膳笔记》："入口即化，喉间回甘，比匈牙利埃斯泰尔哈吉家族的珍珠项链更珍贵。"',
    image: img3,
  },
  {
    year: "今日",
    title: "每秒 0.5 米的皇家电梯",
    desc: '矿工仍沿用手工预选，剔除杂色矿脉；电梯 72 秒直达井下 650 米，恒温 24℃、湿度 45%，天然抑菌——这是德国仅存的可参观"活盐矿"。出井前，每一粒盐都要经过"皇盐仪式"。',
    image: img4,
  },
];

const heritageImages = [heritageImg, time1517, time1706, time1960];

// carousel state
const currentHeritage = ref(0);
const currentHeritageSrc = computed(
  () => heritageImages[currentHeritage.value],
);
const prevHeritage = () => {
  currentHeritage.value =
    (currentHeritage.value + heritageImages.length - 1) % heritageImages.length;
};
const nextHeritage = () => {
  currentHeritage.value = (currentHeritage.value + 1) % heritageImages.length;
};

const bottomImages = [time1517, time1706, time1960];

// bottom carousel state
const currentBottom = ref(0);
const currentBottomSrc = computed(() => bottomImages[currentBottom.value]);
const prevBottom = () => {
  currentBottom.value =
    (currentBottom.value + bottomImages.length - 1) % bottomImages.length;
};
const nextBottom = () => {
  currentBottom.value = (currentBottom.value + 1) % bottomImages.length;
};

const scenes = [
  {
    time: "07:00",
    title: "晨间仪式",
    desc: "晨起第一杯矿物质水。在温水中加入少量天然纯净的Alpen白金盐饮用，杀菌消炎·增强肠胃功能·促进血液循环！",
  },
  {
    time: "12:30",
    title: "午后轻食",
    desc: "0.3 g 撒在牛油果上，咸度恰好平衡单宁",
  },
  {
    time: "19:00",
    title: "第一口辅盐",
    desc: '宝宝辅食最后一道撒 2 粒，妈妈不再用舌尖"二次调味"',
  },
  { time: "23:00", title: "夜训配方", desc: "健身后自制电解质水，补钠不水肿" },
];

const qaList = [
  {
    q: "它是海盐吗？",
    a: "岩盐。原始海洋蒸发储存在阿尔卑斯山的珍宝，原始自然生态，原理一切污染与核辐射。",
  },
  {
    q: "纯度高到底好在哪？",
    a: "德国精湛技术与工艺、纯净度99.9%越纯越入味·越纯越健康! 健康之道在于·无污染·天然无添加，在于纯净度!食盐不纯，视同吃工业盐!",
  },
  {
    q: "如何快速验证正品？",
    a: "① 瓶身手机扫一扫官网溯源页；② 海关溯源码；③ 发货时可见海关报关清关。",
  },
];
</script>

<style scoped lang="scss">
@import url("https://fonts.googleapis.com/css2?family=Cinzel:wght@400;600;700&family=Lato:wght@300;400;700&family=Playfair+Display:ital,wght@0,400;0,600;0,700;1,400&display=swap");

/* Design System Variables */
.platinum-brand-page {
  --gold-primary: #d4af37;
  --gold-light: #e5c568;
  --dark-bg: #0f172a;
  --dark-card: #1e293b;
  --text-main: #334155;
  --text-light: #94a3b8;
  --white: #ffffff;
  --font-serif: "Playfair Display", serif;
  --font-display: "Cinzel", serif;
  --font-sans: "Lato", sans-serif;

  font-family: var(--font-sans);
  color: var(--text-main);
  background-color: var(--white);
  overflow-x: hidden;
  width: 100%;
}

/* Reset & Utilities */
h1,
h2,
h3,
h4,
h5,
h6 {
  margin: 0;
}
a {
  text-decoration: none;
  color: inherit;
  transition: 0.3s;
}
img {
  max-width: 100%;
  display: block;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 24px;
}

.section {
  padding: 120px 0;
  position: relative;
}

.section-label {
  display: block;
  font-family: var(--font-sans);
  font-size: 12px;
  letter-spacing: 4px;
  text-transform: uppercase;
  color: var(--gold-primary);
  margin-bottom: 16px;
  font-weight: 700;

  &.gold {
    color: var(--gold-primary);
  }
}

.section-title {
  font-family: var(--font-serif);
  font-size: 48px;
  line-height: 1.2;
  color: var(--dark-bg);
  margin-bottom: 32px;

  &.white {
    color: var(--white);
  }
}

.section-subtitle {
  font-family: var(--font-sans);
  font-size: 18px;
  line-height: 1.6;
  color: var(--text-light);
  margin-top: -16px;
  margin-bottom: 32px;
  font-weight: 300;

  &.white {
    color: rgba(255, 255, 255, 0.8);
  }
}

.text-body {
  font-size: 18px;
  line-height: 1.8;
  color: var(--text-main);
  margin-bottom: 24px;
  font-weight: 300;
}

/* Navigation */
.platinum-nav {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 1000;
  padding: 24px 0;
  transition: all 0.4s ease;

  &.scrolled {
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(10px);
    padding: 16px 0;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.05);

    .nav-logo {
      color: var(--dark-bg);
    }
    .nav-links a {
      color: var(--dark-bg);
    }
  }
}

.nav-container {
  max-width: 1400px;
  margin: 0 auto;
  padding: 0 40px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.nav-logo {
  display: block;
  height: 40px;

  .logo-img {
    height: 100%;
    width: auto;
    // filter: brightness(0) invert(1);
    transition: filter 0.3s;
  }
}

.platinum-nav.scrolled {
  .logo-img {
    filter: none;
  }
}

.mobile-menu-toggle {
  display: none;
  background: none;
  border: none;
  cursor: pointer;
  flex-direction: column;
  gap: 6px;
  z-index: 1001;
  padding: 8px;

  .hamburger-line {
    width: 24px;
    height: 2px;
    background-color: var(--white);
    transition: all 0.3s ease;
  }
}

.platinum-nav.scrolled .mobile-menu-toggle .hamburger-line {
  background-color: var(--dark-bg);
}

.nav-links {
  display: flex;
  gap: 40px;

  a {
    color: var(--white);
    font-size: 14px;
    letter-spacing: 1px;
    text-transform: uppercase;
    position: relative;

    &::after {
      content: "";
      position: absolute;
      bottom: -4px;
      left: 0;
      width: 0;
      height: 1px;
      background: var(--gold-primary);
      transition: width 0.3s;
    }

    &:hover::after {
      width: 100%;
    }
  }
}

/* Hero Section */
.hero-section {
  height: 100vh;
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  color: var(--white);
  overflow: hidden;
}

.hero-bg {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-position: center;
  background-color: #000;
  background-size: cover;
  background-repeat: no-repeat;
  filter: brightness(0.6);
  z-index: 0;
  /* Simple zoom effect */
  animation: zoomBg 20s infinite alternate;
}

@keyframes zoomBg {
  from {
    transform: scale(1);
  }
  to {
    transform: scale(1.1);
  }
}

.hero-content {
  position: relative;
  z-index: 1;
  max-width: 800px;
  padding: 0 20px;
}

.hero-title {
  font-family: var(--font-display);
  font-size: 80px;
  font-weight: 400;
  letter-spacing: 0.1em;
  margin-bottom: 16px;
  background: linear-gradient(to right, #fff, #e0e0e0);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.hero-subtitle {
  font-family: var(--font-sans);
  font-size: 14px;
  letter-spacing: 0.3em;
  color: var(--gold-primary);
  margin-bottom: 32px;
}

.hero-separator {
  width: 1px;
  height: 60px;
  background: var(--gold-primary);
  margin: 0 auto 32px;
}

.hero-desc {
  font-family: var(--font-serif);
  font-size: 24px;
  line-height: 1.6;
  font-style: italic;
}

.scroll-indicator {
  position: absolute;
  bottom: 40px;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12px;
  z-index: 1;

  span {
    font-size: 10px;
    letter-spacing: 2px;
    color: var(--white);
    opacity: 0.7;
  }

  .line {
    width: 1px;
    height: 40px;
    background: var(--white);
    opacity: 0.5;
  }
}

/* Heritage Section */
.heritage-section {
  background: var(--white);
}

.grid-2 {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 80px;
  align-items: center;

  &.reverse-mobile {
    @media (max-width: 768px) {
      display: flex;
      flex-direction: column-reverse;
    }
  }
}

.quote-box {
  margin-top: 40px;
  padding-left: 24px;
  border-left: 2px solid var(--gold-primary);
  font-family: var(--font-serif);
  font-size: 20px;
  font-style: italic;
  color: var(--dark-bg);

  cite {
    display: block;
    margin-top: 16px;
    font-size: 14px;
    font-family: var(--font-sans);
    font-style: normal;
    color: var(--text-light);
    letter-spacing: 1px;
    text-transform: uppercase;
  }
}

/* carousel for heritage images */

/* bottom section full-width carousel */
.bottom-history-section {
  background: var(--white);
  padding: 80px 0;
  display: flex;
  align-items: center;
}

/* keep full viewport for standalone, but when inside grid restrict to normal flow */
.bottom-history-section .bottom-carousel-wrapper {
  position: relative;
  width: 100%;
  left: auto;
  right: auto;
  margin: 0 auto;
  display: flex;
  align-items: stretch;
  justify-content: center;
}

.bottom-carousel-window {
  width: 80%;
  max-width: 1200px;
  overflow: hidden;
  /* fill available height */
  height: 100%;
}

.bottom-carousel-window img {
  width: 100%;
  display: block;
  border-radius: 4px;
}

.bottom-carousel-wrapper .nav {
  background: none;
  border: none;
  font-size: 64px;
  color: var(--gold-primary);
  cursor: pointer;
  padding: 0 16px;
  transition: color 0.3s;
}

.bottom-carousel-wrapper .nav:hover {
  color: var(--gold-light);
}

.debug-path {
  margin-top: 8px;
  font-size: 12px;
  color: var(--text-light);
  word-break: break-all;
}

/* Timeline */
.timeline-wrapper {
  margin-top: 120px;
  position: relative;
  padding: 40px 0;
}

.timeline-track {
  position: absolute;
  left: 50%;
  top: 0;
  bottom: 0;
  width: 1px;
  background: linear-gradient(
    to bottom,
    transparent,
    var(--gold-primary),
    transparent
  );
  transform: translateX(-50%);
}

.timeline-item {
  display: flex;
  justify-content: flex-end;
  padding-right: 50%;
  position: relative;
  margin-bottom: 60px;

  &.reverse {
    justify-content: flex-start;
    padding-right: 0;
    padding-left: 50%;

    .timeline-dot {
      left: -6px;
      right: auto;
    }
    .timeline-date {
      left: -100px;
      right: auto;
      text-align: right;
    }
    .timeline-content {
      text-align: left;
      align-items: flex-start;
    }
  }
}

.timeline-dot {
  width: 12px;
  height: 12px;
  background: var(--gold-primary);
  border-radius: 50%;
  position: absolute;
  right: -6px;
  top: 0;
  box-shadow: 0 0 0 4px rgba(212, 175, 55, 0.2);
}

.timeline-date {
  position: absolute;
  right: -100px;
  top: -8px;
  font-family: var(--font-display);
  font-size: 24px;
  color: var(--gold-primary);
  width: 80px;
}

.timeline-content {
  padding: 0 40px;
  text-align: right;
  display: flex;
  flex-direction: column;
  align-items: flex-end;

  h3 {
    font-family: var(--font-serif);
    font-size: 20px;
    margin-bottom: 8px;
    color: var(--dark-bg);
  }

  p {
    color: var(--text-main);
    font-size: 14px;
    line-height: 1.6;
    max-width: 300px;
  }
}

.timeline-image-container {
  margin-bottom: 16px;
  overflow: hidden;
  border-radius: 4px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition: transform 0.5s ease;
  }

  &:hover img {
    transform: scale(1.05);
  }
}

.china-market {
  margin-top: 80px;
  display: grid;
  grid-template-columns: 1.05fr 0.95fr;
  gap: 44px;
  align-items: start;
  padding: 32px;
  border-radius: 18px;
  border: 1px solid rgba(212, 175, 55, 0.18);
  background: linear-gradient(
    180deg,
    rgba(212, 175, 55, 0.08) 0%,
    rgba(255, 255, 255, 1) 55%,
    rgba(255, 255, 255, 1) 100%
  );
}

.china-market-content .section-title {
  font-size: 40px;
  margin-bottom: 18px;
}

.china-market-content .text-body {
  margin-bottom: 14px;
}

.china-market-points {
  margin: 0;
  padding: 0;
  list-style: none;
  display: grid;
  gap: 10px;

  li {
    position: relative;
    padding-left: 18px;
    font-size: 15px;
    color: var(--text-main);
    line-height: 1.7;

    &::before {
      content: "";
      position: absolute;
      left: 0;
      top: 10px;
      width: 8px;
      height: 8px;
      border-radius: 999px;
      background: var(--gold-primary);
      box-shadow: 0 0 0 4px rgba(212, 175, 55, 0.16);
    }
  }
}

.china-market-gallery {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;
}

.china-market-photo {
  display: block;
  border-radius: 14px;
  overflow: hidden;
  background: rgba(0, 0, 0, 0.03);
  box-shadow: 0 10px 26px rgba(0, 0, 0, 0.1);
  aspect-ratio: 4 / 3;

  img {
    width: 100%;
    height: 100%;
    display: block;
    object-fit: cover;
    transform: scale(1.001);
    transition: transform 0.6s ease;
  }

  &:hover img {
    transform: scale(1.05);
  }
}

/* Purity Section */
.purity-section {
  background: var(--dark-bg);
  color: var(--white);
  overflow: hidden;
}

.purity-bg-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: radial-gradient(
    circle at 30% 50%,
    rgba(212, 175, 55, 0.05),
    transparent 50%
  );
  pointer-events: none;
}

.centered-header {
  text-align: center;
  margin-bottom: 80px;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 40px;
  margin-bottom: 80px;
}

.feature-card {
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.05);
  padding: 40px;
  text-align: center;
  transition: transform 0.4s;

  &:hover {
    transform: translateY(-10px);
    background: rgba(255, 255, 255, 0.05);
    border-color: var(--gold-primary);
  }

  .feature-icon {
    font-size: 40px;
    margin-bottom: 24px;
  }

  h3 {
    font-family: var(--font-serif);
    font-size: 24px;
    margin-bottom: 16px;
    color: var(--gold-primary);
  }

  p {
    color: var(--text-light);
    line-height: 1.6;
  }
}

.authority-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 30px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  padding-top: 60px;
}

.authority-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;

  .auth-icon {
    font-size: 32px;
    margin-bottom: 16px;
    opacity: 0.8;
  }

  h4 {
    font-family: var(--font-serif);
    font-size: 18px;
    color: var(--white);
    margin-bottom: 12px;
  }

  p {
    font-size: 13px;
    color: var(--text-light);
    line-height: 1.6;
    max-width: 280px;
  }
}

/* Visual Interlude */
.visual-interlude {
  height: 60vh;
  background-position: center;
  background-size: cover;
  background-attachment: fixed;
  background-repeat: no-repeat;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;

  &::before {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.1);
  }
}

.interlude-content {
  position: relative;
  z-index: 1;
  text-align: center;

  h3 {
    font-family: var(--font-display);
    font-size: 4vw;
    color: var(--white);
    text-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
    letter-spacing: 0.1em;
    margin-bottom: 24px;
  }

  p {
    font-family: var(--font-serif);
    font-size: 24px;
    color: var(--white);
    line-height: 1.6;
    font-style: italic;
  }
}

/* Scenes Section */
.scenes-section {
  background: var(--white);
}

.scenes-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
}

.scene-card {
  background: #fafafa;
  padding: 32px 24px;
  border-top: 2px solid var(--gold-primary);
  transition: all 0.3s;

  &:hover {
    transform: translateY(-8px);
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.05);
    background: var(--white);
  }
}

.scene-time {
  font-family: var(--font-display);
  color: var(--gold-primary);
  font-size: 24px;
  margin-bottom: 16px;
  font-weight: 700;
}

.scene-info {
  h3 {
    font-family: var(--font-serif);
    font-size: 18px;
    margin-bottom: 8px;
    color: var(--dark-bg);
  }

  p {
    font-size: 13px;
    color: var(--text-main);
    line-height: 1.6;
  }
}

/* Products Section */
.products-section {
  background: linear-gradient(to bottom, #fff, #f8f9fa);
}

.product-showcase img {
  transform: scale(1.1);
  transition: transform 0.8s ease;
}

.product-features {
  list-style: none;
  padding: 0;
  margin: 32px 0;

  li {
    margin-bottom: 16px;
    padding-left: 24px;
    position: relative;
    font-size: 16px;
    color: var(--text-main);

    &::before {
      content: "•";
      color: var(--gold-primary);
      position: absolute;
      left: 0;
      font-size: 24px;
      line-height: 1;
      top: -2px;
    }

    span {
      font-weight: 700;
      color: var(--dark-bg);
    }
  }
}

.cta-button {
  background: var(--dark-bg);
  color: var(--white);
  border: none;
  padding: 16px 40px;
  font-family: var(--font-sans);
  font-size: 14px;
  letter-spacing: 2px;
  text-transform: uppercase;
  cursor: pointer;
  transition: all 0.3s;

  &:hover {
    background: var(--gold-primary);
    transform: translateY(-2px);
    box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
  }
}

/* QA Section */
.qa-section {
  background: var(--white);
}

.qa-header {
  padding-right: 40px;
}

.qa-list {
  display: flex;
  flex-direction: column;
  gap: 32px;
}

.qa-item {
  border-bottom: 1px solid rgba(0, 0, 0, 0.05);
  padding-bottom: 32px;

  &:last-child {
    border-bottom: none;
  }
}

.qa-question {
  font-family: var(--font-serif);
  font-size: 20px;
  color: var(--dark-bg);
  margin-bottom: 12px;
  display: flex;
  align-items: center;
  gap: 12px;

  .q-mark {
    color: var(--gold-primary);
    font-family: var(--font-display);
    font-weight: 700;
  }
}

.qa-answer {
  font-size: 15px;
  color: var(--text-main);
  line-height: 1.8;
  padding-left: 36px;
}

/* Sustainability */
.sustain-section {
  text-align: center;
  background: #f8f9fa;
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  /* overlay for readability */
  position: relative;
}

.sustain-section::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(248, 249, 250, 0.85);
  pointer-events: none;
}

.sustain-grid {
  display: flex;
  justify-content: center;
  gap: 60px;
  margin-top: 60px;
  margin-bottom: 80px;
}

.sustain-item {
  flex: 1;
  max-width: 300px;

  .big-number {
    font-family: var(--font-display);
    font-size: 56px;
    color: var(--gold-primary);
    line-height: 1;
    margin-bottom: 16px;
  }

  h4 {
    font-family: var(--font-serif);
    font-size: 20px;
    margin-bottom: 8px;
    color: var(--dark-bg);
  }

  p {
    font-size: 14px;
    color: var(--text-main);
    line-height: 1.6;
  }
}

.brand-manifesto {
  max-width: 800px;
  margin: 0 auto;
  padding: 60px 0 0;
  border-top: 1px solid rgba(0, 0, 0, 0.05);

  h3 {
    font-family: var(--font-serif);
    font-size: 28px;
    font-style: italic;
    color: var(--dark-bg);
    margin-bottom: 16px;
  }

  p {
    font-size: 14px;
    letter-spacing: 1px;
    color: var(--text-light);
    text-transform: uppercase;
  }
}

/* Footer */
.brand-footer {
  background: var(--dark-bg);
  color: var(--white);
  padding: 80px 0;
  text-align: center;
}

.footer-logo {
  font-family: var(--font-display);
  font-size: 32px;
  margin-bottom: 24px;
  letter-spacing: 4px;
}

.copyright {
  color: var(--text-light);
  font-size: 12px;
}

/* AOS substitute styles */
[data-aos] {
  opacity: 0;
  transform: translateY(30px);
  transition:
    opacity 1s ease,
    transform 1s ease;

  &.visible {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Responsive */
@media (max-width: 768px) {
  /* Navigation */
  .mobile-menu-toggle {
    display: flex;
  }

  .nav-container {
    padding: 0 24px;
  }

  .nav-links {
    position: fixed;
    top: 0;
    right: -100%;
    width: 100%;
    height: 100vh;
    background: var(--dark-bg);
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 32px;
    transition: right 0.4s ease;
    z-index: 999;
    padding: 40px;
  }

  .platinum-nav.mobile-menu-open .nav-links {
    right: 0;
  }

  .nav-links a {
    color: var(--white) !important;
    font-size: 24px;
  }

  /* Toggle Animation */
  .platinum-nav.mobile-menu-open .mobile-menu-toggle {
    .hamburger-line:nth-child(1) {
      transform: rotate(45deg) translate(5px, 6px);
    }
    .hamburger-line:nth-child(2) {
      opacity: 0;
    }
    .hamburger-line:nth-child(3) {
      transform: rotate(-45deg) translate(5px, -6px);
    }

    .hamburger-line {
      background-color: var(--white) !important;
    }
  }

  /* Hero */
  .hero-title {
    font-size: 12vw;
    line-height: 1.1;
  }
  .hero-subtitle {
    font-size: 12px;
    letter-spacing: 2px;
  }
  .hero-content {
    padding: 0 24px;
  }

  /* Grids */
  .grid-2 {
    grid-template-columns: 1fr;
    gap: 40px;
  }
  .features-grid {
    grid-template-columns: 1fr;
    gap: 24px;
  }
  .sustain-grid {
    flex-direction: column;
    gap: 40px;
  }
  .scenes-grid {
    grid-template-columns: 1fr;
    gap: 24px;
  }
  .authority-grid {
    grid-template-columns: 1fr;
    gap: 40px;
  }

  .section {
    padding: 60px 0;
  }

  .image-box .year-badge {
    right: 0;
    bottom: -10px;
    padding: 10px 20px;
    font-size: 14px;
  }

  /* Timeline */
  .timeline-track {
    left: 20px;
    transform: none;
  }
  .timeline-item,
  .timeline-item.reverse {
    flex-direction: column;
    padding: 0 0 0 50px;
    justify-content: flex-start;
    align-items: flex-start;
    margin-bottom: 40px;

    .timeline-dot {
      left: 14px;
      right: auto;
    }
    .timeline-date {
      position: relative;
      left: 0;
      top: 0;
      right: auto;
      text-align: left;
      margin-bottom: 8px;
    }
    .timeline-content {
      padding: 0;
      text-align: left;
      align-items: flex-start;

      p {
        max-width: 100%;
      }
    }
  }

  .china-market {
    margin-top: 50px;
    grid-template-columns: 1fr;
    gap: 22px;
    padding: 20px;
  }

  .china-market-content .section-title {
    font-size: 32px;
  }

  .china-market-gallery {
    grid-template-columns: repeat(2, 1fr);
    gap: 10px;
  }

  /* QA */
  .qa-header {
    padding-right: 0;
    margin-bottom: 40px;
  }
  .qa-list {
    gap: 24px;
  }

  /* General Text */
  .section-title {
    font-size: 32px;
    margin-bottom: 24px;
  }
  .text-body {
    font-size: 16px;
  }
  .hero-desc {
    font-size: 18px;
  }
}
</style>
