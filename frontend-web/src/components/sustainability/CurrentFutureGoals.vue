<template>
  <div class="two-columns">
    <div class="alternative-background layout-0">
      <h2 class="title">我们保证今天</h2>
      <div class="ce-textpic ce-left ce-intext layout--0">
        <div class="ce-gallery" data-ce-columns="1" data-ce-images="1">
          <div class="ce-column">
            <figure class="image">
              <img 
                loading="lazy" 
                alt="三包盐" 
                :src="todayPackagesImg" 
                width="200" 
                height="212"
              >
            </figure>
          </div>
        </div>
        <div class="ce-bodytext">
          <p>
            我们每年生产超过 <strong>6400 万个由可回收纸板</strong> 制成的盐包 。
          </p>
        </div>
      </div>
 
            <div class="ce-textpic ce-left ce-intext layout--0">
   
        <div class="ce-bodytext">
          <p>
             我们 超过 <strong>85%</strong>的包装材料均来自 可持续森林的<strong>可再生原材料。</strong>
          </p>
        </div>
             <div class="ce-gallery" data-ce-columns="1" data-ce-images="1">
          <div class="ce-column">
            <figure class="image">
           <div class="icon"></div>
            </figure>
          </div>
        </div>
      </div>
      <div class="ce-textpic ce-right ce-intext layout--0">
        <div class="ce-gallery" data-ce-columns="1" data-ce-images="1">
          <div class="ce-column">
            <figure class="image">
              <img 
                loading="lazy" 
                alt="盐罐" 
                :src="todayCanImg" 
                width="114" 
                height="229"
              >
            </figure>
          </div>
        </div>
        <div class="ce-bodytext">
          <p>我们的 <strong>可重复填充包装</strong>也 有助于保护环境。</p>
        </div>
      </div>
    </div>

    <div class="alternative-background2 layout-0">
      <h3  class="title"> 这就是我们希望在2025年达到的目标。</h3>
      <div class="ce-bodytext layout--0">
        <p>我们的包装<strong>可回收</strong>或<strong>可重复使用</strong>。</p>
      </div>
      <div class="ce-bodytext layout--0">
        <p>
          <br>
          所有用于初级包装的纸张均100 <strong>%</strong>
          <br>
          来自<strong>可持续森林</strong>。
        </p>
      </div>
      <div class="ce-textpic ce-right ce-intext ce-nowrap layout--0">
        <div class="ce-gallery" data-ce-columns="1" data-ce-images="1">
          <div class="ce-column">
            <figure class="image">
              <img 
                loading="lazy" 
                alt="用于包装的可持续材料" 
                :src="monomaterialImg" 
                width="249" 
                height="244"
              >
            </figure>
          </div>
        </div>
        <div class="ce-bodytext">
          <p>我们的塑料包装<strong>100%由单一材料</strong>制成。</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import todayPackagesImg from '@/assets/downloaded_images/05-nachhaltigkeit-heute-packungen.png'
import todayCanImg from '@/assets/downloaded_images/05-nachhaltigkeit-heute-dose.png'
import monomaterialImg from '@/assets/downloaded_images/05-nachhaltigkeit-2025-monomaterial.png'
</script>

<style scoped lang="scss">
  strong{
    font-family: Caveat, serif;
    font-weight: 700;
    color: #6a9b75 !important;
  }
  .title{
    text-align: center;
    color: #334046;
    font-size: 34px;
    font-weight: 500;
  }
.alternative-background {
    background: url('@/assets/downloaded_images/05-nachhaltigkeit-heute-paper-background.jpg') no-repeat center center;
    background-size: cover;
    .icon{
      background-image: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMDEuODgxIiBoZWlnaHQ9IjExNS40NDgiPjxwYXRoIGRhdGEtbmFtZT0iMDUtbmFjaGhhbHRpZ2tlaXQtaGV1dGUtaWNvbi1uYWNoZnVlbGxiYXIiIGQ9Ik0zOS43NiAxMTUuMDY3YTIuMDc3IDIuMDc3IDAgMDEtLjg3OS0xLjdWODguNjU1TDI3LjYzOSA3Ny40MTFhMi4wNzggMi4wNzggMCAwMTIuOTM5LTIuOTM5TDQyLjQzIDg2LjMyM2EyLjA3NCAyLjA3NCAwIDAxLjYwOCAxLjQ3djIyLjYzMmwxNS44LTUuNlY4Ny43ODlhMi4wNzggMi4wNzggMCAwMS42MS0xLjQ3bDM4LjI3MS0zOC4yNzN2LTguMjg4SDQuMTU3djcuMDcxSDc4LjhhMi4wNzkgMi4wNzkgMCAwMTAgNC4xNTdINy4xTDE5LjkyNSA2My44MWEyLjA4IDIuMDggMCAwMS0yLjk0IDIuOTQxTC42MSA1MC4zNzhhMi4wODMgMi4wODMgMCAwMS0uNjA4LTEuNDd2LTExLjIyYTIuMDc4IDIuMDc4IDAgMDEyLjA3OS0yLjA3OWgzMi44MjNhMi41ODIgMi41ODIgMCAwMTMuOTg0LTEuNDg0LjAzNi4wMzYgMCAwMS4wMS4wMDdsLjAzMS4wMjFhMi41NyAyLjU3IDAgMDExLjAwNyAxLjQ1N2g1OS44NjhhMi4wNzkgMi4wNzkgMCAwMTIuMDc5IDIuMDc5djExLjIyOGEyLjA3NiAyLjA3NiAwIDAxLS42MSAxLjQ3TDYzLjAwMiA4OC42NnYxNy42NDNhMi4wODIgMi4wODIgMCAwMS0xLjM4NSAxLjk2bC0xOS45NjEgNy4wNjlhMi4wODYgMi4wODYgMCAwMS0xLjg5NC0uMjYyem0tMTcuMzcyLTQyLjlhMi4wODEgMi4wODEgMCAxMTEuNDcxLjYwOCAyLjEgMi4xIDAgMDEtMS40NzEtLjYxMXptNjIuNjM3LTIxLjc4YTIuMDc1IDIuMDc1IDAgMTExLjQ3MS42MSAyLjEgMi4xIDAgMDEtMS40NzEtLjYwOXpNNjUuMjM3IDMxLjE2NGwtLjAwOS0uMDA3LS4wMzEtLjAyMWEyLjU4MSAyLjU4MSAwIDAxMi45Ny00LjIyMmwuMDEuMDA1LjAzMS4wMjJhMi41ODIgMi41ODIgMCAwMS0yLjk3MiA0LjIyM3ptLTM0LjExLTUuNDQzbC0uMDMxLS4wMjFhMi41ODIgMi41ODIgMCAwMTIuOTYzLTQuMjI5bC4wMjkuMDIyYTIuNTgxIDIuNTgxIDAgMTEtMi45NjEgNC4yMjh6TTE1LjIwOSAxNy4xNGwtLjAwOS0uMDA3LS4wMzEtLjAyMmEyLjU4MiAyLjU4MiAwIDAxMi45NzItNC4yMjIuMDE2LjAxNiAwIDAwLjAwOC4wMDdsLjAzMS4wMjJhMi41ODIgMi41ODIgMCAwMS0yLjk3MiA0LjIyM3ptMTguNzcyLTUuNjkxbC0uMDA4LS4wMDctLjAzMS0uMDIyYTIuNTgxIDIuNTgxIDAgMDEyLjk3LTQuMjIxbC4wMS4wMDUuMDMxLjAyMmEyLjU4MiAyLjU4MiAwIDAxLTIuOTcyIDQuMjIzem0xMS44My02LjczbC0uMDEtLjAwNy0uMDMxLS4wMjFBMi41ODIgMi41ODIgMCAwMTQ4Ljc0Mi40NjhhLjAzNi4wMzYgMCAwMS4wMS4wMDdsLjAzMS4wMjFhMi41ODIgMi41ODIgMCAwMS0yLjk3MSA0LjIyM3oiIGZpbGw9IiM2YTliNzUiLz48L3N2Zz4=);
       background-position: 50%;
       width: 100px;
       height: 115px;
    }
  }

  .alternative-background2 {
    background: url('@/assets/downloaded_images/05-nachhaltigkeit-2025-paper-background.jpg') no-repeat center center;
    background-size: cover;
  }

.two-columns {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 60px;
  max-width: 1400px;
  margin: 0 auto;
  padding: 80px 20px;

  > div {

    padding: 60px 40px;
    border-radius: 8px;
  }



  .ce-textpic {
    padding: 10px 0;
    display: flex;
    gap: 30px;
    align-items: center;

    &.ce-left {
      flex-direction: row;
    }

    &.ce-right {
      flex-direction: row-reverse;
    }

    .ce-gallery {
      flex: 0 0 auto;
    }

    .ce-bodytext {
      flex: 1;
      display: flex;
      p {
      font-family: Caveat, serif;
    font-weight: 700;
    color: #6a9b75;
    width: 70%;
      }
    }

    img {
      max-width: 100%;
      height: auto;
    }
  }

  .ce-bodytext.layout--0 {
    padding: 30px 0;

    p {
      font-size: 16px;
      line-height: 1.7;
      color: #666;
      margin: 0;
    }
  }

  strong {
    color: #2b2b2e;
    font-weight: 600;
  }
}

@media (max-width: 1024px) {
  .two-columns {
    grid-template-columns: 1fr;
    gap: 40px;
  }
}

@media (max-width: 768px) {
  .two-columns {
    padding: 60px 15px;

    > div {
      padding: 40px 20px;
    }

    h2, h3 {
      font-size: 24px;
    }
  }
}
</style>


<style scoped lang="scss">
img {
  transition: transform 0.6s ease-out;
}
</style>
