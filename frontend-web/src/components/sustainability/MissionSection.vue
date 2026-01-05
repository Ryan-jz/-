<template>
  <div class="alternative-background layout-1">
    <div class="ce-textpic ce-right ce-intext ce-nowrap layout--0">
      <div class="ce-gallery" data-ce-columns="1" data-ce-images="1">
        <div class="ce-column img-1">
          <figure class="image">
            <img 
              loading="lazy" 
              alt="巴特赖兴哈勒盐包装盒被当作服装，用来浇花" 
              :src="missionImg" 
              width="444" 
              height="625"
            >
          </figure>
        </div>
      </div>
      <div class="ce-bodytext">
        <h2>我们的使命</h2>
        <p>
          为了保护自然环境，造福子孙后代，我们秉持负责任的态度，采用可持续的包装解决方案。在此过程中，我们既支持和促进了循环经济，又确保了产品的保护和质量。
        </p>
        <div class="icon"></div>
      </div>
    </div>

    <div class="ce-textpic ce-left ce-intext ce-nowrap layout--0">
      <div class="ce-gallery" data-ce-columns="1" data-ce-images="1">
        <div class="ce-column">
          <figure class="image">
            <img 
              loading="lazy" 
              alt="两只手，盐从中流过" 
              :src="guidelinesImg" 
              width="552" 
              height="492"
            >
          </figure>
        </div>
      </div>
      <div class="ce-bodytext">
        <h2>我们的指导原则</h2>
        <p>我们尽可能使用纸张，必要时使用塑料，当我们的品牌理念要求时才使用玻璃。</p>
        <div class="icon2"></div>
      </div>
    </div>
  </div>
</template>

<script setup>
import missionImg from '@/assets/downloaded_images/03-mission.jpg'
import guidelinesImg from '@/assets/downloaded_images/03-leitsaetze.jpg'
</script>

<style scoped lang="scss">
.icon {
        content: "";
        position: absolute;
        z-index: 1;
        bottom: auto;
        left: auto;
        display: block;
        width: 270px;
        height: 250px;
        background-image: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIyNzUuODQzIiBoZWlnaHQ9IjI1MC41MDUiPjxwYXRoIGRhdGEtbmFtZT0iMDMtbWlzc2lvbi1pY29uLW1pc3Npb24iIGQ9Ik0yNzQuNTYyIDIzNy4zNzVsLTU3Ljk4LTkzLjkxNGE4LjY2IDguNjYgMCAwMC0xNC43MjUgMGwtNy41MjcgMTIuMTkzYTQuMDQzIDQuMDQzIDAgMDA2Ljg5NCA0LjIyN2w3LjUyOC0xMi4xOTRhLjU1Mi41NTIgMCAwMS45MzggMGw1Ny45OCA5My45MTRhLjU0OC41NDggMCAwMS0uNDY5LjgzNUgyMzUuMDlMMTQ1LjI4NCA5Ni45NzRhOC42MDggOC42MDggMCAwMC0zLjIxLTMuMDQ5di0yNy4zN2EzNy41MjkgMzcuNTI5IDAgMDA5Ljg1OCAxLjNjOS4yMyAwIDE3Ljg3My0zLjIgMjYuMzI5LTYuMzI2IDExLjcxNy00LjMzNSAyMi43ODQtOC40MjkgMzQuNzU0LTMuNjQ1YTUgNSAwIDAwNi44NjktNC42MzNWNy42ODlhNC45ODUgNC45ODUgMCAwMC0zLjA0NC00LjZjLTE1LjE5NS02LjQ1OC0yOC41MTUtMS41MzEtNDEuNCAzLjIzNC0xMi4wNjcgNC40NjQtMjMuNDY0IDguNjgtMzUuODc1IDMuMTUzYTQuMDQ2IDQuMDQ2IDAgMDAtNS43IDMuNjgydi4xNWE0LjEyMSA0LjEyMSAwIDAwLS4xMDkuOTMydjc5LjY4OGE4LjYxNiA4LjYxNiAwIDAwLTMuMjA2IDMuMDQ3bC00Mi42MjggNjkuMDYtMTMuOTM2LTIyLjU3NGE4LjY2IDguNjYgMCAwMC0xNC43MjUgMGwtNTcuOTggOTMuOTE0QTguNjE0IDguNjE0IDAgMDA4LjY0MyAyNTAuNUgyNjcuMmE4LjYxNCA4LjYxNCAwIDAwNy4zNjItMTMuMTI5ek0xNDEuOTY5IDE4Ljg4OGMxMy4xIDMuNTgxIDI0Ljg2OC0uNzcyIDM2LjMtNSAxMS4zMDYtNC4xODMgMjIuMDE3LTguMTQ0IDMzLjUyOS00LjF2MzkuMTg1Yy0xMy4xNC0zLjU5MS0yNC45MTkuNzY3LTM2LjM0NiA0Ljk5My0xMS4yODcgNC4xNzUtMjEuOTg4IDguMTM0LTMzLjQ3OSA0LjEwNXptLTkxLjcxMyAyMjMuNTVsMTMuODg0LTIyLjQ4OWE0LjA0MyA0LjA0MyAwIDAwLTYuODkzLTQuMjI3bC0xNi40OTQgMjYuNzE2SDguNjQzYS41NDguNTQ4IDAgMDEtLjQ2OS0uODM1bDU3Ljk4LTkzLjkxNGEuNTUxLjU1MSAwIDAxLjkzNyAwbDE2LjA3OSAyNi4wNDMtMTYuMDU0IDI2YTQuMDQzIDQuMDQzIDAgMDA2Ljg5MyA0LjIyN0wxMzcuNDUzIDEwMS4yYS41NTEuNTUxIDAgMDEuOTM3IDBsODcuMiAxNDEuMjM3eiIgZmlsbD0iI2Q3ZTlkNCIvPjwvc3ZnPg==);
        background-position: 50%;
        background-size: contain;
        background-repeat: no-repeat;
  }
.icon2 {
        content: "";
        position: absolute;
        z-index: 1;
        bottom: 0;
        top: auto;
        left: auto;
        display: block;
        width: 270px;
        height: 250px;
        background-image: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIyMDMuOTI4IiBoZWlnaHQ9IjI2NiI+PHBhdGggZGF0YS1uYW1lPSIwMy1sZWl0c2FldHplLWljb24tbGVpdHNhZXR6ZSIgZD0iTTkzLjA5NiAyNjZhNC40MzUgNC40MzUgMCAwMS00LjQzNS00LjQzM3YtODQuMjM0SDEzLjI5NmE0LjQzNCA0LjQzNCAwIDAxLTQuNDM1LTQuNDMzdi02Mi4wNjdhNC40MzQgNC40MzQgMCAwMTQuNDM1LTQuNDMzaDc1LjM2NVY4OC42NjNIMzEuMDI4YTQuNDMyIDQuNDMyIDAgMDEtMy4zNjQtMS41NTZsLTI2LjYtMzEuMDM1YTQuNDMzIDQuNDMzIDAgMDEwLTUuNzYybDI2LjYtMzEuMDMzYTQuNDMxIDQuNDMxIDAgMDEzLjM2NC0xLjU0OGg1Ny42MzN2LTQuNDMzYTEzLjMgMTMuMyAwIDAxMjYuNiAwdjQuNDMzaDc1LjM2N2E0LjQzNCA0LjQzNCAwIDAxNC40MzMgNC40MzNWODQuMjNhNC40MzIgNC40MzIgMCAwMS00LjQzMyA0LjQzM2gtNzUuMzY3djE3LjczNGg1Ny42MzNhNC40MzkgNC40MzkgMCAwMTMuMzY3IDEuNTQ2bDI2LjYgMzEuMDM1YTQuNDM2IDQuNDM2IDAgMDEwIDUuNzcxbC0yNi42IDMxLjAzM2E0LjQzNCA0LjQzNCAwIDAxLTMuMzY3IDEuNTQ4aC01Ny42MzN2ODQuMjMzYTQuNDM0IDQuNDM0IDAgMDEtNC40MzMgNC40MzN6bTQuNDMzLTguODY4aDguODY2di03OS44aC04Ljg2NnptLTc5LjgtODguNjY2aDE1My4xMjdsMjIuOC0yNi42LTIyLjgtMjYuNkgxNy43Mjl6bTc5LjgtNjIuMDY2aDguODY2Vjg4LjY2M2gtOC44NjZ6TTEwLjI2NyA1My4ybDIyLjggMjYuNmgxNTMuMTI4VjI2LjZIMzMuMDY4em05MS42OTUtNDQuMzM3YTQuNDMyIDQuNDMyIDAgMDAtNC40MzMgNC40MzN2NC40MzNoOC44NjZ2LTQuNDMzYTQuMjQ4IDQuMjQ4IDAgMDAtMS4yNzYtMy4wNzYgNC4zMSA0LjMxIDAgMDAtMy4xMzktMS4zNTd6bS01My4yIDE0Ni4zYTQuNDMzIDQuNDMzIDAgMDEwLTguODY2aDEwNi40YTQuNDMzIDQuNDMzIDAgMDEwIDguODY2em0wLTE3LjczMmE0LjQzMyA0LjQzMyAwIDAxMC04Ljg2NmgxMDYuNGE0LjQzMyA0LjQzMyAwIDAxMCA4Ljg2NnptMC03MC45MzRhNC40MzMgNC40MzMgMCAwMTAtOC44NjZoMTA2LjRhNC40MzMgNC40MzMgMCAwMTAgOC44NjZ6bTAtMTcuNzM0YTQuNDMzIDQuNDMzIDAgMDEwLTguODY2aDEwNi40YTQuNDMzIDQuNDMzIDAgMDEwIDguODY2eiIgZmlsbD0iI2Q3ZTlkNCIvPjwvc3ZnPg==);
        background-position: 50%;
        background-size: contain;
        background-repeat: no-repeat;
  }
  .image{
    border-radius: 20px 20px 20px 20px;
    overflow: hidden;
    img{
      border-radius: 20px 20px 20px 20px;
    overflow: hidden;
    }
  }
  .img-1{
   position: relative;
   height: 450px;
   img{
    position: absolute;
    top: -200px;
   }
  }
.alternative-background {
  border-image-source: url('@/assets/downloaded_images/brush-fullwidth-box-green.png');
  background-repeat: no-repeat;
      background-clip: padding-box;
    border-image-slice: 196 0 144 fill;
    border-image-repeat: repeat;
    border-width: 100px 0 80px;
    border-style: solid;

  background-position: center;
  background-size: 100% auto;
  padding: 80px 0;

  // &.layout-1 {
  //   background: linear-gradient(180deg, #fff 0%, #f8f8f8 100%);
  // }
}

.ce-textpic {
  padding: 0 20px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 60px;

  max-width: 1400px;
  margin: 0 auto 60px;

  &:last-child {
    margin-bottom: 0;
  }

  &.ce-right .ce-gallery {
    position: relative;
    order: 2;
  }

  .ce-bodytext {
    position: relative;
    h2 {
      font-size: 42px;
      font-weight: 300;
      line-height: 1.3;
      margin: 0 0 25px;
      color: #ffcf25;
    }

    p {
      font-size: 18px;
      line-height: 1.8;
      color: #fff;
      margin: 0;
    }
  }

  img {
    max-width: 100%;
    height: auto;
  }
}

@media (max-width: 1024px) {
  .ce-textpic {
    grid-template-columns: 1fr;
    gap: 40px;

    .ce-gallery {
      order: 0 !important;
    }
  }
}

@media (max-width: 768px) {
  .alternative-background {
    padding: 60px 0;
  }

  .ce-textpic {
    padding: 0 15px;
    margin-bottom: 40px;

    .ce-bodytext h2 {
      font-size: 28px;
    }

    .ce-bodytext p {
      font-size: 16px;
    }
  }
}
</style>


<style scoped lang="scss">
img {
  transition: transform 0.6s ease-out;
}
</style>
