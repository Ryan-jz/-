<template>
  <div class="ce-textpic ce-left ce-intext ce-nowrap layout--2">
    <div class="ce-gallery" data-ce-columns="1" data-ce-images="1">
      <div class="ce-column">
        <figure class="image">
          <img 
            loading="lazy" 
            alt="奶奶、母亲和孩子在厨房" 
            :src="packagingImg" 
            width="460" 
            height="699"
          >
        </figure>
      </div>
    </div>
    <div class="icon"></div>
    <div class="ce-bodytext">
      <h2>
        <span class="subtitle">包装</span>
        可持续性与我们每个人都息息相关。
      </h2>
      <p>承担责任意味着思考自身行为的后果。在当今世界，它还意味着致力于可持续发展。毕竟，我们都希望拥有一个值得居住的世界——无论现在还是将来。</p>
      <p>产品质量始终是我们的首要任务。但包装呢？作为德国最知名的食盐品牌，我们也希望在这方面履行好我们的责任。</p>
      <p>因此，我们非常自豪地宣布，巴特赖兴哈尔盐业公司超过 80% 的盐包装已采用可再生资源制成的纸板，并且这些纸板可以回收利用。我们的目标是到 2025 年实现所有包装均可回收或重复使用。</p>
    </div>
  </div>
</template>

<script setup>
import packagingImg from '@/assets/downloaded_images/02-verpackung.jpg'
</script>

<style scoped lang="scss">
    .icon{
      
        content: "";
        position: absolute;
        z-index: 1;
        top: -40px;
        right: auto;
        bottom: auto;
        left: 520px;
        display: block;
        width: 320px;
        height: 320px;
        background-image: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIzMTUuOTE3IiBoZWlnaHQ9IjMzMC44OTgiPjxwYXRoIGRhdGEtbmFtZT0iMDItdmVycGFja3VuZy1pY29uLXZlcnBhY2t1bmciIGQ9Ik0xNjEuNjE2IDMyMS44NWMtMS42MzIuMDU4IDMuNjkgMS43OTMtOTMuMDc3LTMzLjkzNWExNS45NSAxNS45NSAwIDAxLTEwLjQ1NC0xNC40NTVsLS42MDgtMTcuMzU2YTUuMzM2IDUuMzM2IDAgMTExMC42NjUtLjM3M2wuNjA4IDE3LjM1NWE1LjMxNiA1LjMxNiAwIDAwMy40ODQgNC44MThsODMuMTg0IDMwLjcwNy0xLjg5My01NC4wMmE1LjMzNiA1LjMzNiAwIDExMTAuNjY0LS4zNzRsMS45MDMgNTQuMzE3IDgwLjgzMi0zNi43NGE1LjMyMiA1LjMyMiAwIDAwMy4xMjQtNS4wNDZsLTEuNTI0LTQzLjQ4MS01Mi44OTUgMjMuODc0YTE1Ljk2OSAxNS45NjkgMCAwMS0xOC40MS0zLjc5OSA1LjMzNiA1LjMzNiAwIDAxNy44ODMtNy4xOTUgNS4zMjIgNS4zMjIgMCAwMDYuMTM2IDEuMjY4bDc0Ljg5Mi0zMy44MDJhNS4zMyA1LjMzIDAgMDAxLjcwMi04LjUwOGwtMTcuMDg2LTE4LjI2Ni04OS45MjMgNDEuMjE5LTIxLjA2MiAyNi41OTlhMTUuOTcgMTUuOTcgMCAwMS0xOC4wOTYgNS4wNzhMNDQuNTkgMjIxLjI1N2ExNi4wMiAxNi4wMiAwIDAxLTYuODgxLTI1LjEwN2MxOS41MjYtMjQuMDQ0IDE4LjI4NS0yMi41NDUgMTguNjYzLTIyLjkwNS45OS0uOTQtNC43ODggMS44MyA5Mi40NzUtNDIuMDM3YTE1LjkxMiAxNS45MTIgMCAwMTEyLjE2Ny0uNDFsOTIuODYgMzQuNTg2YTUuNDk4IDUuNDk4IDAgMDEyLjEwOCAxLjQzNWwxOS42NCAyMC45OTZhMTYuMDE5IDE2LjAxOSAwIDAxLTUuMTA3IDI1LjUyNGwtMTEuNDg0IDUuMTg1IDEuNjc2IDQ3Ljg1MmExNS45NTYgMTUuOTU2IDAgMDEtOS4zNzIgMTUuMTMyYy04NC40OSAzOC40LTg4LjMxOCA0MC4zMDgtODkuMzcgNDAuMzQ1LS4xMS4wMTEtLjE5NC0uMDA4LS4zNS0uMDAzem0tOC4zNzgtMTgwLjkxbC04OS44NDMgNDAuNTItMTcuNCAyMS40MjJhNS4zMzMgNS4zMzMgMCAwMDIuMjk0IDguMzdsNzcuMDcyIDI4LjQ3OGE1LjMyMyA1LjMyMyAwIDAwNi4wMzItMS42OTRsMTcuMzE4LTIxLjg3MS02Mi4zMjEtMjMuNDQyYTUuMzM2IDUuMzM2IDAgMTEzLjc1LTkuOTg3bDY3LjA4NSAyNS4yMzMgODAuODY5LTM3LjA3LTgwLjc5OS0zMC4wOTVhNS4zMSA1LjMxIDAgMDAtNC4wNTcuMTM2em05My41NjMtMTIuNzYxYzEzLjgwNi0yMS43NDcgNDIuNjUtMjIuNzIxIDU5LjE0MS0zLjc0OWE1LjMzNiA1LjMzNiAwIDExLTguMDU2IDdjLTExLjktMTMuNjg2LTMyLjE2OC0xMy4xNDUtNDIuMDggMi40NjlhNS4zMzUgNS4zMzUgMCAxMS05LjAwOC01Ljcyek01NC4zOCAxNDAuOTU2Yy0xMC45MzUtMTQuODEyLTMxLjE4My0xNC4wNzYtNDIuMTUuNDg0YTUuMzM3IDUuMzM3IDAgMDEtOC41MjYtNi40MmMxNS4xMzMtMjAuMDkxIDQzLjk3LTIxLjExNiA1OS4yNi0uNGE1LjMzNSA1LjMzNSAwIDExLTguNTg1IDYuMzM2em0zNy40NzctMTMuNjI3Yy0yLjQxLTcuODc2LTguNzEzLTIyLjYzLTIzLjg1LTMyLjk1YTUuMzM1IDUuMzM1IDAgMDE2LjAxMi04LjgxM2MxNy44MyAxMi4xNTcgMjUuMjIzIDI5LjQzIDI4LjA0NCAzOC42NDZhNS4zMzYgNS4zMzYgMCAxMS0xMC4yMDUgMy4xMnptMTE5LjIxMi0uNDA0YTUuMzM0IDUuMzM0IDAgMDEtMy45OTctNi40IDc1LjI0OSA3NS4yNDkgMCAwMTQuMDY0LTEyLjQ3NiA1LjMzNiA1LjMzNiAwIDAxOS44MiA0LjE3NyA2NC43MzMgNjQuNzMzIDAgMDAtMy40OTIgMTAuNzAxIDUuMzI1IDUuMzI1IDAgMDEtNi40IDMuOTk3em0tNzUuMjY1LTE5LjU1NWMtMS42My0xLjYyLTEuNTk0LTMuMzYtMS43NC01LjY0YTE3LjM0NiAxNy4zNDYgMCAwMC0zLjI0NS04LjE0Yy01LjQyLTcuMjU3LTE0Ljg1Ni03LjE4OC0xNS45MTMtNy4xNTdhNS4zMzYgNS4zMzYgMCAwMS0uMzQ2LTEwLjY2NSAzMS41ODggMzEuNTg4IDAgMDEyMC41OTIgNi45NTdjLjIxLTEuMTU0LjQ1NS0yLjMwMy43MjItMy40NDgtMTEuNTktMTEuODI0LTEzLjU4OC0zMC4zMDYtNC45MTYtNDguMTA1IDE2LjE3NS0zMy4xOTQgNDkuNTE1LTIxLjk1OSA2NC43LTI5LjYwNmExNS41NTkgMTUuNTU5IDAgMDEyMi4wNzQgMTAuNTQyYzUuMzY0IDIzLjk1OC0uODE3IDQ3LjYwOC0xNi4xMzQgNjEuNzIzYTY0Ljk5NCA2NC45OTQgMCAwMS0yMi4xNjYgMTMuOCA1LjMzNSA1LjMzNSAwIDExLTMuNTY2LTEwLjA1NiA1NC40NCA1NC40NCAwIDAwMTguNDk4LTExLjU5YzEyLjUyNy0xMS41NDIgMTcuNDg4LTMxLjI5MyAxMi45NTMtNTEuNTQ0YTQuODg4IDQuODg4IDAgMDAtNi44NjQtMy4zNDJjLTE4Ljc2MiA5LjQ1LTQ2LjYtMi41NTItNTkuOTA3IDI0Ljc0OC01LjU5OSAxMS40ODItNS43MTggMjIuNzg0LS43MDYgMzEuMjMyIDguNzY4LTIwLjcyNyAyNC44MDUtMzIuNjEzIDM0LjExMS0zOC4wNjVhNS4zMzUgNS4zMzUgMCAwMTUuMzk0IDkuMjA3Yy04Ljg4NiA1LjIwOC0yNC44MTggMTcuMTE0LTMxLjY2NCAzOC4zNS00Ljc2NCAxNC43NzgtMi40ODcgMjUuMjkyLTIuODUgMjcuNzg0YTUuMzIyIDUuMzIyIDAgMDEtNS4xMjEgNC41NjQgNS4yMzEgNS4yMzEgMCAwMS0zLjkwNS0xLjU0OXptOTEuMjgzLTIyLjc1MmE2OS42NiA2OS42NiAwIDAxMzIuOTQ5LTE3Ljg1OSA1LjMzNiA1LjMzNiAwIDExMi40MjMgMTAuMzkzIDU5LjAzMSA1OS4wMzEgMCAwMC0yNy45NDIgMTUuMTIgNS4zMzUgNS4zMzUgMCAxMS03LjQzLTcuNjU4em0tMTgzLjMxMS4yQTUuMzM1IDUuMzM1IDAgMDE0NS40NiA3NC4yOHEuNzk1LjEyNyAxLjY2OC4yOWE1LjMzNiA1LjMzNiAwIDAxLS43NzMgMTAuNTggMTIuMjcgMTIuMjcgMCAwMS0yLjU4LS4zMzN6IiBmaWxsPSIjNmE5Yjc1IiBvcGFjaXR5PSIuNiIvPjwvc3ZnPg==);
        background-position: 50%;
        background-size: contain;
        background-repeat: no-repeat;
    }
  .layout--2{
    margin-top: 200px !important;
    position: relative;
  }
.ce-textpic {
  padding: 80px 20px;
   display: flex;
  align-items: center;
  justify-content: center;

  margin: 0 auto;

  &.ce-left .ce-gallery {
    min-width: 450px;
    order: 0;
    .ce-column{
      width: 100%;
    }
  }

  .ce-bodytext {
    width: 820px;
  height: 100%;
      margin-top: -200px;
    padding: 90px 20px 90px 150px;
    background: #d7e9d4;
    background-clip: padding-box;
    border-image: url('@/assets/downloaded_images/brush-box-mint.png');
    border-image-slice: 25 54 18 53 fill;
    border-image-repeat: stretch;
    border-width: 25px 54px 18px 0;
    border-style: solid;
    border-radius: 15px;
    h2 {
      font-size: 42px;
      font-weight: 300;
      line-height: 1.3;
      margin: 0 0 25px;
      color: #2b2b2e;

      .subtitle {
        display: block;
        font-size: 90px;
        color: #ffc25c;
        text-transform: uppercase;
        letter-spacing: 2px;

        font-weight: 600;
      }
    }

    p {
      font-size: 18px;
      line-height: 1.8;
      color: #666;
      margin: 0 0 20px;

      &:last-child {
        margin-bottom: 0;
      }
    }
  }

  img {
    max-width: 100%;
    width: 100%;
    height: auto;
  }
}

@media (max-width: 1024px) {
  .ce-textpic {
    flex-direction: column;
    gap: 40px;

    &.ce-left .ce-gallery {
      min-width: auto;
    }

    .ce-bodytext {
      width: 100%;
      margin-top: 0;
      padding: 60px 20px;
    }
  }

  .layout--2 {
    margin-top: 100px !important;
  }

  .icon {
    width: 200px;
    height: 200px;
    left: 50%;
    transform: translateX(-50%);
    top: -100px;
  }
}

@media (max-width: 768px) {
  .ce-textpic {
    padding: 40px 15px;

    &.ce-left .ce-gallery {
      min-width: auto;
    }

    .ce-bodytext {
      width: 100%;
      margin-top: 0;
      padding: 40px 15px;
      border-width: 20px 30px 15px 0;

      h2 {
        font-size: 20px;

        .subtitle {
          font-size: 40px;
        }
      }

      p {
        font-size: 14px;
      }
    }
  }

  .layout--2 {
    margin-top: 0 !important;
  }

  .icon {
    position: static;
    width: 120px;
    height: 120px;
    margin: 0 auto 20px;
    transform: none;
  }
}
</style>


<style scoped lang="scss">
img {
  transition: transform 0.6s ease-out;
}
</style>
