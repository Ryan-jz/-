<template>
  <div class=" alternative-background">
    <a id="recycling-tipps"></a>
    <div class=" layout-3">
      <div class="ce-bodytext layout--0">
        <h2>回收小贴士</h2>
        <p>谁将成为知识问答冠军？阿尔皮和乔迪将挑战回收利用相关问题。</p>
      </div>
    </div>

    <div class="columns">
      <div 
        class="video" 
        v-for="(video, index) in videos" 
        :key="index"
      >
        <div class="video-embed">
          <video 
            :ref="el => videoRefs[index] = el"
            playsinline
            @click="togglePlay(index)"
          >
            <source :src="video.src" type="video/webm">
          </video>
          <div 
            class="play-button" 
            v-show="!playingStates[index]"
            @click="togglePlay(index)"
          >
            ▶
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import video1 from '@/assets/downloaded_images/BRH_nachhaltigkeit_recycling_quiz_1.webm'
import video2 from '@/assets/downloaded_images/BRH_nachhaltigkeit_recycling_quiz_3.webm'
import video3 from '@/assets/downloaded_images/BRH_nachhaltigkeit_recycling_quiz_2.webm'

const videos = ref([
  { src: video1 },
  { src: video2 },
  { src: video3 }
])

const videoRefs = ref([])
const playingStates = ref([false, false, false])

const togglePlay = (index) => {
  const video = videoRefs.value[index]
  if (video.paused) {
    video.play()
    playingStates.value[index] = true
  } else {
    video.pause()
    playingStates.value[index] = false
  }
}
</script>

<style scoped lang="scss">
  .alternative-background{
    position: relative;
      padding: 80px 0;
  border-image-slice: 196 0 144 fill;
    border-image-repeat: repeat;
    border-width: 100px 0 80px;
    border-style: solid;
    border-image-source: url('@/assets/downloaded_images/brush-fullwidth-box-blue.png');
    margin-bottom: 400px;
  }

.layout-3 {


  .ce-bodytext {
    max-width: 1400px;
    margin: 0 auto;
    padding: 0 20px;
    text-align: center;

    h2 {
      font-size: 42px;
      font-weight: 300;
      line-height: 1.3;
      margin: 0 0 25px;
      color: #2b2b2e;
    }

    p {
      font-size: 18px;
      line-height: 1.8;
      color: #666;
      margin: 0;
    }
  }
}

.columns {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 40px;
  max-width: 1400px;
  margin: 60px auto;
  padding: 0 20px;
  position: absolute;
 top: 350px;
  left: 50%;
  transform: translate(-50%, -50%);
}

.video {
  width: 100%;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
 
  .video-embed {
    position: relative;
    cursor: pointer;

    video {
      width: 100%;
      height: 400px;
      display: block;
    }

    .play-button {
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      width: 80px;
      height: 80px;
      background: rgba(255, 255, 255, 0.9);
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 32px;
      color: #6a9b75;
      transition: all 0.3s;
      padding-left: 8px;

      &:hover {
        background: #fff;
        transform: translate(-50%, -50%) scale(1.1);
      }
    }
  }
}

@media (max-width: 1024px) {
  .columns {
    grid-template-columns: 1fr;
    gap: 30px;
  }
}

@media (max-width: 768px) {
  .alternative-background.layout-3 {
    padding: 60px 0;

    .ce-bodytext {
      padding: 0 15px;

      h2 {
        font-size: 28px;
      }

      p {
        font-size: 16px;
      }
    }
  }

  .columns {
    padding: 0 15px;
  }
}
</style>
