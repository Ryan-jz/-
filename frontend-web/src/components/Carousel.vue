<template>
  <div class="carousel" @mouseenter="pauseAutoplay" @mouseleave="resumeAutoplay">
    <div class="carousel-container">
      <!-- 轮播项 -->
      <transition-group name="slide" tag="div" class="carousel-wrapper">
        <div
          v-for="(item, index) in items"
          v-show="index === currentIndex"
          :key="item.id || index"
          class="carousel-item"
        >
          <!-- 图片 -->
          <img 
            v-if="item.mediaType === 1 || !item.mediaType"
            :src="item.mediaUrl || item.image" 
            :alt="item.title || `Slide ${index + 1}`"
            class="carousel-image"
          />
          <!-- 视频 -->
          <video
            v-else-if="item.mediaType === 2"
            :src="item.mediaUrl"
            class="carousel-video"
            autoplay
            muted
            loop
            playsinline
          ></video>
          
          <div v-if="showContent && (item.title || item.description)" class="carousel-content">
            <h2 v-if="item.title" class="carousel-title">{{ item.title }}</h2>
            <p v-if="item.description" class="carousel-description">{{ item.description }}</p>
            <button v-if="item.buttonText" class="carousel-button" @click="handleButtonClick(item)">
              {{ item.buttonText }}
            </button>
          </div>
        </div>
      </transition-group>

      <!-- 左右箭头 -->
      <button 
        v-if="showArrows" 
        class="carousel-arrow carousel-arrow-left" 
        @click="prev"
        aria-label="Previous slide"
      >
        ‹
      </button>
      <button 
        v-if="showArrows" 
        class="carousel-arrow carousel-arrow-right" 
        @click="next"
        aria-label="Next slide"
      >
        ›
      </button>

      <!-- 指示器 -->
      <div v-if="showIndicators" class="carousel-indicators">
        <button
          v-for="(item, index) in items"
          :key="index"
          :class="['indicator', { active: index === currentIndex }]"
          @click="goTo(index)"
          :aria-label="`Go to slide ${index + 1}`"
        ></button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watch } from 'vue'

const props = defineProps({
  items: {
    type: Array,
    required: true,
    default: () => []
  },
  autoplay: {
    type: Boolean,
    default: true
  },
  interval: {
    type: Number,
    default: 5000
  },
  showArrows: {
    type: Boolean,
    default: true
  },
  showIndicators: {
    type: Boolean,
    default: true
  },
  showContent: {
    type: Boolean,
    default: true
  },
  height: {
    type: String,
    default: '100%'
  },
  aspectRatio: {
    type: String,
    default: '16/9'
  }
})

const emit = defineEmits(['change', 'button-click'])

const currentIndex = ref(0)
let autoplayTimer = null
const isAutoplayPaused = ref(false)

const next = () => {
  currentIndex.value = (currentIndex.value + 1) % props.items.length
  emit('change', currentIndex.value)
}

const prev = () => {
  currentIndex.value = (currentIndex.value - 1 + props.items.length) % props.items.length
  emit('change', currentIndex.value)
}

const goTo = (index) => {
  currentIndex.value = index
  emit('change', currentIndex.value)
}

const handleButtonClick = (item) => {
  emit('button-click', item)
}

const startAutoplay = () => {
  if (props.autoplay && !isAutoplayPaused.value) {
    autoplayTimer = setInterval(() => {
      next()
    }, props.interval)
  }
}

const stopAutoplay = () => {
  if (autoplayTimer) {
    clearInterval(autoplayTimer)
    autoplayTimer = null
  }
}

const pauseAutoplay = () => {
  isAutoplayPaused.value = true
  stopAutoplay()
}

const resumeAutoplay = () => {
  isAutoplayPaused.value = false
  startAutoplay()
}

watch(() => props.autoplay, (newVal) => {
  if (newVal) {
    startAutoplay()
  } else {
    stopAutoplay()
  }
})

onMounted(() => {
  startAutoplay()
})

onUnmounted(() => {
  stopAutoplay()
})
</script>

<style scoped>
.carousel {
  width: 100%;
  height: 100%;
  position: relative;
  overflow: hidden;
}

.carousel-container {
  position: relative;
  width: 100%;
  height: v-bind(height);
  aspect-ratio: v-bind(aspectRatio);
}

.carousel-wrapper {
  position: relative;
  width: 100%;
  height: 100%;
}

.carousel-item {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

.carousel-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.carousel-video {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.carousel-content {
  position: absolute;
  bottom: 8%;
  left: 50%;
  transform: translateX(-50%);
  text-align: center;
  color: white;
  max-width: min(800px, 90%);
  padding: 0 20px;
  z-index: 2;
}

.carousel-title {
  font-size: clamp(24px, 5vw, 48px);
  font-weight: bold;
  margin: 0 0 clamp(10px, 2vw, 20px) 0;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
}

.carousel-description {
  font-size: clamp(14px, 2vw, 20px);
  margin: 0 0 clamp(15px, 3vw, 30px) 0;
  line-height: 1.6;
  text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.5);
}

.carousel-button {
  background-color: #e31e24;
  color: white;
  border: none;
  padding: clamp(10px, 1.5vw, 15px) clamp(25px, 4vw, 40px);
  font-size: clamp(14px, 1.5vw, 16px);
  font-weight: 600;
  cursor: pointer;
  border-radius: 4px;
  transition: all 0.3s ease;
}

.carousel-button:hover {
  background-color: #c01a1f;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(227, 30, 36, 0.4);
}

.carousel-arrow {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  background-color: rgba(255, 255, 255, 0.8);
  color: #333;
  border: none;
  width: clamp(40px, 5vw, 50px);
  height: clamp(40px, 5vw, 50px);
  font-size: clamp(24px, 3vw, 32px);
  cursor: pointer;
  z-index: 3;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
}

.carousel-arrow:hover {
  background-color: rgba(255, 255, 255, 1);
  transform: translateY(-50%) scale(1.1);
}

.carousel-arrow-left {
  left: clamp(10px, 2vw, 20px);
}

.carousel-arrow-right {
  right: clamp(10px, 2vw, 20px);
}

.carousel-indicators {
  position: absolute;
  bottom: clamp(15px, 2vw, 20px);
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  gap: clamp(8px, 1vw, 10px);
  z-index: 3;
}

.indicator {
  width: clamp(10px, 1.2vw, 12px);
  height: clamp(10px, 1.2vw, 12px);
  border-radius: 50%;
  background-color: rgba(255, 255, 255, 0.5);
  border: none;
  cursor: pointer;
  transition: all 0.3s ease;
  padding: 0;
}

.indicator:hover {
  background-color: rgba(255, 255, 255, 0.8);
}

.indicator.active {
  background-color: #e31e24;
  width: clamp(24px, 3vw, 30px);
  border-radius: clamp(5px, 0.6vw, 6px);
}

/* 过渡动画 */
.slide-enter-active,
.slide-leave-active {
  transition: all 0.6s ease;
}

.slide-enter-from {
  opacity: 0;
  transform: translateX(100%);
}

.slide-leave-to {
  opacity: 0;
  transform: translateX(-100%);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .carousel-content {
    bottom: 5%;
  }
}

@media (max-width: 480px) {
  .carousel-content {
    bottom: 3%;
    padding: 0 15px;
  }
  
  .carousel-arrow {
    opacity: 0.7;
  }
}
</style>
