<template>
  <div class="stage-slideshow">
    <div class="main-image">
      <img 
        :src="slides[currentSlide].imgSrc" 
        :alt="slides[currentSlide].alt" 
        :title="slides[currentSlide].title"
      >
      <nav class="controls">
        <a 
          href="javascript:void(0)" 
          class="previous control" 
          :class="{ disabled: currentSlide === 0 }"
          @click="prevSlide"
        >
          后退
        </a>
        <a 
          href="javascript:void(0)" 
          class="next control" 
          :class="{ disabled: currentSlide === slides.length - 1 }"
          @click="nextSlide"
        >
          前
        </a>
      </nav>
    </div>
    <div class="thumbnails">
      <div 
        v-for="(slide, index) in slides" 
        :key="index"
        :class="['thumbnail', { active: index === currentSlide }]"
        @click="currentSlide = index"
      >
        <img :src="slide.imgSrc" :alt="slide.alt">
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
  slides: {
    type: Array,
    required: true
  }
})

const currentSlide = ref(0)

const prevSlide = () => {
  if (currentSlide.value > 0) {
    currentSlide.value--
  }
}

const nextSlide = () => {
  if (currentSlide.value < props.slides.length - 1) {
    currentSlide.value++
  }
}
</script>

<style scoped lang="scss">
.stage-slideshow {
  max-width: 1400px;
  margin: 60px auto;
  padding: 0 20px;

  .main-image {
    position: relative;
    width: 100%;
    height:600px;
    border-radius: 8px;
    overflow: hidden;
    margin-bottom: 20px;

    img {
      width: 100%;
      height: 100%;
    
      display: block;
    }

    .controls {
      position: absolute;
      top: 50%;
      left: 0;
      right: 0;
      transform: translateY(-50%);
      display: flex;
      justify-content: space-between;
      padding: 0 20px;
      pointer-events: none;

      .control {
        pointer-events: all;
        width: 50px;
        height: 50px;
        background: rgba(255, 255, 255, 0.9);
        border-radius: 50%;
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
        transition: all 0.3s;
        box-shadow: 0 2px 10px rgba(0, 0, 0, 0.15);
        font-size: 0;

        &:hover {
          background: #fff;
          transform: scale(1.1);
        }

        &.disabled {
          opacity: 0.3;
          cursor: not-allowed;
        }

        &::after {
          content: '';
          width: 12px;
          height: 12px;
          border-top: 3px solid #2b2b2e;
          border-right: 3px solid #2b2b2e;
        }

        &.previous::after {
          transform: rotate(-135deg);
          margin-left: 3px;
        }

        &.next::after {
          transform: rotate(45deg);
          margin-right: 3px;
        }
      }
    }
  }

  .thumbnails {
    display: flex;
    gap: 15px;
    justify-content: center;

    .thumbnail {
      width: 120px;
      height: 80px;
      border-radius: 4px;
      overflow: hidden;
      cursor: pointer;
      border: 3px solid transparent;
      transition: all 0.3s;

      &:hover {
        border-color: #6a9b75;
      }

      &.active {
        border-color: #6a9b75;
      }

      img {
        width: 100%;
        height: 100%;
        object-fit: cover;
        display: block;
      }
    }
  }
}

@media (max-width: 768px) {
  .stage-slideshow {
    padding: 0 15px;

    .main-image {
      height: 300px;

      .controls .control {
        width: 40px;
        height: 40px;

        &::after {
          width: 10px;
          height: 10px;
          border-width: 2px;
        }
      }
    }

    .thumbnails {
      gap: 10px;
      overflow-x: auto;
      justify-content: flex-start;

      .thumbnail {
        width: 80px;
        height: 60px;
        flex-shrink: 0;
      }
    }
  }
}
</style>
