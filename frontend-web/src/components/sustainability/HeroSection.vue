<template>
  <div class="ce-textpic ce-center ce-above bg">
    <div class="ce-gallery">
      <img :src="headerImg" alt="母亲怀着孩子">
    </div>
    <div class="ce-bodytext">
      <div class="title">
        <span class="subtitle">我们承担</span>
        责任
      </div>
      <div class="icon">
        
      </div>
      <p class="content">可持续发展一直是公司理念的核心组成部分。作为一家创新型盐生产商，我们深感自身肩负着环境保护和健康保护的重任。这体现在我们的产品和项目中，也体现在我们负责任地利用资源上。想了解更多？我们已在此总结了最重要的几点。</p>
    </div>
  </div>
</template>

<script setup>
import headerImg from '@/assets/downloaded_images/01-nachhaltigkeit-header-form.png'
</script>

<style scoped lang="scss">
    .icon{
        content: "";
        position: absolute;
        z-index: 1;
       left:20%;
       transform:translateY(50px);
        display: block;
        width: 270px;
        height: 270px;
        background-image: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIyNzMiIGhlaWdodD0iMjczIj48cGF0aCBkYXRhLW5hbWU9IjAxLW5hY2hoYWx0aWdrZWl0LWljb24tcGxhbmV0IiBkPSJNMjY4Ljg5NCAzMS44MTdhMTA5LjggMTA5LjggMCAwMC05LjUyOC0yMS42NTggNC4wODUgNC4wODUgMCAwMC00LjMxOS0xLjgyMSAxMDkuOTY2IDEwOS45NjYgMCAwMC0yMi4xNzkgOC4yNjhjLTE1LjczMiA3Ljg0OC0yNi43MjIgMTguMzI2LTMxLjc4NCAzMC4yOTQtNC4yMTQgOS45NjgtNC44MTkgMjIuMTUzLTEuODI5IDM1LjYxNWExMTMuNDc0IDExMy40NzQgMCAwMC0yMS4zMzktMTguNTU3IDE2NS4xNzIgMTY1LjE3MiAwIDAwLTE4LjU1LTI4LjQwOGMtMy4yNDctMy45ODctOS40OTIgMS4wMTYtNi4yMTIgNS4wNDRBMTYwLjExNCAxNjAuMTE0IDAgMDExNjcuOTI2IDYyLjNsLTM2LjQ2LTcuMUw5Ni44MjggNy45ODhhNTYuNTEgNTYuNTEgMCAwMTEwLjEgMS4wM2MxMC42NDggMi4wODMgMjEuMTY1IDcuNSAzMS4yNjYgMTYuMSAzLjk1NyAzLjM3MSA5LjEtMi43NjQgNS4xOTEtNi4wODktMTEuMTYxLTkuNDk4LTIyLjkwOS0xNS41MS0zNC45MTktMTcuODYxLTE4LjM0Mi0zLjU4Ni0zOS43MyAxLjEyMi02MS44NjIgMTMuNjE4YTE1OS43ODkgMTU5Ljc4OSAwIDAwLTI3Ljk4OCAxOS45NzYgNC4yMyA0LjIzIDAgMDAtLjQ3NSA1LjQ5IDE2My4xNiAxNjMuMTYgMCAwMDE3Ljk1MiAyOC4xNTYgMTI1LjEgMTI1LjEgMCAwMDE4LjkyIDE5LjQzMnEtMi44IDIuMzQyLTUuNDM1IDQuOTEyYTkyLjQzMiA5Mi40MzIgMCAwMDI5LjkzNSAxNTEuNzUzYzQuNzY2IDEuOTI4IDcuNzc5LTUuNDc5IDMuMDEtNy40MTJhODQuNTIzIDg0LjUyMyAwIDAxLTE3LjMxMy05LjUxNiA0NS4wMyA0NS4wMyAwIDAwLTYuOC0zMS41NDJjLTEuMjY5LTEuOTc1IDIuNTI1LTkuNzI5IDQuNzg5LTE0LjM2MSAyLjMzOC00Ljc3NCA0Ljc1MS05LjcwOSA1Ljg3OC0xNC4zNTcgMS4yLTQuOTM3LS4xNS05LjA4MS0zLjg5NC0xMS45ODktMi43NDgtMi4xMzUtNi4zNDUtMy4yMzktOS41MTUtNC4yMTQtMi4wNTktLjYzMy00LjM5NC0xLjM1LTUuMjI4LTIuMDQ0LTEuNzE4LTEuNDM4LTQuNTIzLTYuNzUyLTYuNzc3LTExLjAxOS0yLjUtNC43MzQtNC44ODMtOS4yMzktNy4zMTQtMTIuMDE0QTg1LjA1IDg1LjA1IDAgMDE2MS42MiA5Mi43NzZjMTAuMDU0IDYuODQzIDIxLjUyOSAxMS41NjUgMzMuNzI2IDEyLjUyNC0xNC4xNDUgMTAuOTUxLTE4LjcyIDI0LjkzNi0xMi4xOCAzOC4zNjkgNy42MzggMTUuNjg2IDIzLjM1OCAxMy4xOCAzNy4yMjcgMTAuOTcyIDIuNC0uMzgzIDQuODgzLS43NzggNy4yNjEtMS4wNzUuODkyLS4xMTIgMS4xNjEuMTMyIDEuMzQxLjI5IDMuMDM3IDIuNyAyLjk0NyAxNS41OTEgMi44ODIgMjUuMDA2LS4wODggMTMuMDIyLS4xNzIgMjUuMzIyIDQuNjgzIDMxLjYzN2ExMi4yMjIgMTIuMjIyIDAgMDAxMC4wNTEgNC44NjUgMjAuMTkgMjAuMTkgMCAwMDcuOS0xLjc4NmMxMi4wNzItNS4xNjUgMjQuNjItMjAuNzEyIDIyLjcwNi0zMy4wMTItMS42NzMtMTAuNzU0LTIuNzgtMTcuODU5LS4xNy0yMy41ODYgMi40MTMtNS4zIDguMjQ2LTEwLjIwOCAxOS43MzQtMTYuNDY5YTg0LjQ3OSA4NC40NzkgMCAwMS05MS44NDUgMTAyLjIzM2MtNS4xMDgtLjU1My01Ljk5IDcuMzg4LS44NzUgNy45NTFhOTMuMDgzIDkzLjA4MyAwIDAwNzQuOTE1LTI1LjggOTIuMzUyIDkyLjM1MiAwIDAwLTE0LjIzNi0xNDMuNTFjNS4yODgtMy43NjcgOS4yMjctNy4wMjUgMTEuNDExLTguOTE2YTEwNi40NjMgMTA2LjQ2MyAwIDAxMjQuNDkyIDI0LjUyNyAxMDYuMTUxIDEwNi4xNTEgMCAwMS0xMDMuNjc3IDE2Ni41NiAxMDYuMTY0IDEwNi4xNjQgMCAwMS04Ny41MTUtMTIyLjAzYy44My01LjAzNC03LjA2Mi02LjM1NS03LjktMS4zYTExNC4xNTYgMTE0LjE1NiAwIDAwOTQuMTEgMTMxLjIyMyAxMTUuOTE0IDExNS45MTQgMCAwMDE4Ljg2IDEuNTUxIDExNC4yMTQgMTE0LjIxNCAwIDAwMTA0LjU0NS0xNTkuNzM4QTExNS4zMjUgMTE1LjMyNSAwIDAwMjM3LjUyMSAxMDZjMTUuNzI5LTcuODQ4IDI2LjcyLTE4LjMyNiAzMS43ODEtMzAuMjk0czQuOTI1LTI3LjE0OS0uNDA4LTQzLjg4OXpNMzYuNTg4IDE0MS43ODRjMi45NzggNS42NDIgNS43OTUgMTAuOTc0IDguNzIyIDEzLjQyIDIuMDczIDEuNzMxIDQuOTU4IDIuNjE4IDguMDExIDMuNTU2IDcuODIzIDIuNCA4LjczMiAzLjU4MiA3Ljk4MSA2LjY3My0uOTIxIDMuOC0zLjE0MiA4LjM0LTUuMjg4IDEyLjcyNy00LjEzNCA4LjQ1MS04LjAzOCAxNi40MjktNC4zMzcgMjIuMTk1YTM4LjAxNCAzOC4wMTQgMCAwMTYuMDU3IDIxLjIyIDg1LjE5MyA4NS4xOTMgMCAwMS0yNC42NjgtODYuMjQzYzEuMTc3IDIuMDE2IDIuNDE1IDQuMzU5IDMuNTIyIDYuNDUyem0xOTMuNjQ5LTc4LjU5M2wtMTEuMzcxLTI3LjU0M2E3NC45MDkgNzQuOTA5IDAgMDExNy40MjYtMTEuODA1IDEwNi44NzUgMTA2Ljg3NSAwIDAxMTIuNjQ3LTUuMzJ6bTI2LjA3OC00MS41NjhhMTA3LjIyNSAxMDcuMjI1IDAgMDE1LjAwNiAxMi43ODcgNzYuNDA3IDc2LjQwNyAwIDAxMy42MjUgMTkuMjhMMjM4LjI5IDY0LjY3OXpNMjA4LjQ2IDUwLjAxNWEzOC4wODEgMzguMDgxIDAgMDE0LjQ2OS03Ljc2M0wyMjUuODc4IDczLjZsLTExLjYyIDI3Ljc2M2ExMDYuMTM4IDEwNi4xMzggMCAwMS01LjE4Ni0xMy4xN2MtMy4yNzMtMTAuMzM3LTYuMDM1LTI1LjM1Ny0uNjEyLTM4LjE3OHptMjUuNjQzIDQ4Ljc0OGExMDYuMzg2IDEwNi4zODYgMCAwMS0xMi4yNDggNS4xOGwxMi4wNTgtMjguOCAzMC42OTMtMTIuNjZhMzguNTYgMzguNTYgMCAwMS0yLjY3NiAxMC4xMDhjLTUuNDIyIDEyLjgyMi0xOC4xMjUgMjEuMzA5LTI3LjgyNyAyNi4xNzJ6TTg5Ljc3OCA5Ni40ODdhNTQuMjQgNTQuMjQgMCAwMS00LjE4Ni0xbDQ0LjMtMzIuNDQ1IDM2LjQ3NSA3LjFjLTEzLjI4OSAxMC40NjMtNDUuNjI2IDMyLjQwNi03Ni41ODkgMjYuMzQ1ek00OS4yOTMgNzEuNDY0YTE1NS43IDE1NS43IDAgMDEtMjAuNC0yOC4xbDQ2LjQxMyA5LjA0N3ptOC4yLTUzLjM1OWwxOS40MTkgMjYuNDY5LTQ2LjcxOC05LjFhMTU3LjggMTU3LjggMCAwMTI3LjI5Ny0xNy4zNjh6bTMxIDI4LjcyNUw2NC45NDcgMTQuNzRhOTAuMDg1IDkwLjA4NSAwIDAxMjIuMzcxLTYuMTg1bDMyLjU3MSA0NC4zODl6bS0yLjQ5MSA3LjY2M2wzMy4yIDYuNDY1LTQyLjYwNSAzMS4yQTgyLjgxIDgyLjgxIDAgMDE1NSA3Ny4yem0xMDguNjA3IDc4LjFjLTE0LjQzIDcuNjgxLTIxLjUgMTMuNzE1LTI0Ljg0NyAyMS4wNzUtMy42IDcuOS0yLjI4MSAxNi4zODYtLjQ1NSAyOC4xMjYgMS4wNzkgNi45MjMtNy4xNTYgMTkuODE1LTE3Ljk0NSAyNC40MzItMy45OTIgMS43MDgtNi44MzUgMS41MDYtOC40NTUtLjYtMy4xOC00LjEzNi0zLjA5My0xNi42NTMtMy4wMjUtMjYuNzA4LjEtMTMuOTEzLjE3OC0yNS45MjktNS41NjYtMzEuMDM2YTkuNDQgOS40NCAwIDAwLTcuNjUtMi4yNTZjLTIuNTE2LjMxMy01LjA2OS43Mi03LjUzNiAxLjExMy0xNC43MTMgMi4zNDMtMjMuOTcyIDMuMjgxLTI4Ljc2OC02LjU3Mi03Ljc3MS0xNS45NTggNy45NTktMjcuNzc4IDE1LjA4My0zMi4xMTlhMjUuMzUxIDI1LjM1MSAwIDAwNC4wNDctMy4wODRjMTcuMjE4LTIuMDU2IDMzLjQ4Ny05LjQzOCA0Ny45ODUtMTguNjg0YTg0LjkwNiA4NC45MDYgMCAwMTM3LjEzIDQ2LjMxMXptMCAwIiBmaWxsPSIjNmE5Yjc1IiBvcGFjaXR5PSIuNiIvPjwvc3ZnPg==);
        background-position: 50%;
        background-size: contain;
        background-repeat: no-repeat;
    }
.bg{
    background-image: url('@/assets/downloaded_images/tile-paper-background.jpg');
    height: 110vh;
}
.ce-textpic.ce-center.ce-above {
  text-align: center;
  position: relative;
  padding: 0;
    
  .ce-gallery {
    width: 100%;
    margin-bottom: 0;
    position: relative;

    img {
      width: 100%;
      height: auto;
      display: block;
    }
  }

  .ce-bodytext {
    position: absolute;
    bottom: -10px;
    left: 0;
    right: 0;
    background: url('data:image/svg+xml;charset=utf-8;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIyNTIwIiBoZWlnaHQ9IjE3My41ODQiPjxwYXRoIGRhdGEtbmFtZT0iMDEtbmFjaGhhbHRpZ2tlaXQtc2Nod3VuZy1wYXBpZXIiIGQ9Ik0yNTIwIDE3My41ODNIMTE1OC4wNTdjMzcuODg2LS40IDc3LjU1NC0xLjQxOCAxMTcuOS0zLjAzMSAzOC41ODQtMS41NDMgNzguNzUxLTMuNjY3IDExOS4zODQtNi4zMTIgNzUuNjYzLTQuOTI2IDE1NS44LTExLjg0IDIzOC4xOTMtMjAuNTQ5IDEzOC43ODktMTQuNjcgMjg2LjQ4My0zNC42NDIgNDM4Ljk4MS01OS4zNjFhOTMxNC42OTMgOTMxNC42OTMgMCAwMDE4MC43MjgtMzEuMTM4YzQ4LTguNzYgOTUuNjE0LTE3Ljg0MiAxNDEuNTEtMjYuOTkzIDc3LjcxNS0xNS41IDEyNC43NzgtMjYuMDk0IDEyNS4yNDUtMjYuMnYxNzMuNTg0em0tMTQzNi4zODYgMEgwVi4wMDFjLjMuMSAzMC4xMTUgMTAuNjI0IDgzLjI2MiAyNi4yIDMxLjIzNyA5LjE1MyA2NC4yOTQgMTguMjM1IDk4LjI1NiAyNi45OTIgNDIuNDM4IDEwLjk0NCA4Ni40NTYgMjEuNDIxIDEzMC44MyAzMS4xMzkgMTEyLjg4NiAyNC43MiAyMjguNTc3IDQ0LjY5MSAzNDMuODU4IDU5LjM2MWEzODgzLjE3IDM4ODMuMTcgMCAwMDEwMC4yNTkgMTEuNDM2IDM3OTIuMjAxIDM3OTIuMjAxIDAgMDAxMDUuMDA2IDkuMTEzYzM2LjQgMi42NDYgNzMuMTYgNC43NyAxMDkuMjY5IDYuMzEzYTM1NTMuOTc0IDM1NTMuOTc0IDAgMDAxMTIuODczIDMuMDI5eiIgZmlsbD0iI2ZmZiIvPjwvc3ZnPg==') no-repeat bottom center;
    background-size: 100% auto;
    height: 100%;
    padding-top: 36%;

    margin: 0 auto;

    .title {
      line-height: 1;
        text-transform: uppercase;
        color: #6a9b75;
      margin-bottom: 30px;
      font-size: 88px;
      font-weight: 300;
      line-height: 1.2;
        max-width: 800px;
        text-align: start;
        font-weight: 500;
        margin-top: 20px;
        margin: 0 auto;
      .subtitle {
        display: block;
        font-size: 90px;
        color: #fff;
        text-transform: uppercase;
        letter-spacing: 3px;
        margin-bottom: 15px;
        font-weight: 600;
      }
    }

    p {
      color: #000;
      max-width:1200px;
      margin: 0 auto;
      font-size: 22px;
      line-height: 1.8;
      font-weight: 500;
      padding-left: 400px;
      line-height: 1.8em;
     color: #000;
     text-align: start;
    }
  }
}

@media (max-width: 768px) {
  .bg {
    height: auto;
  }

  .icon {
    position: static;
    width: 120px;
    height: 120px;
    margin: 0 auto 20px;
    transform: none;
  }

  .ce-textpic.ce-center.ce-above {
    .ce-bodytext {
      position: relative;
      padding: 40px 15px;
      padding-top: 10%;

      .title {
        font-size: 32px;
        text-align: center;
        margin: 0 auto 20px;

        .subtitle {
          font-size: 36px;
        }
      }

      p {
        font-size: 14px;
        padding-left: 0;
        text-align: center;
        max-width: 100%;
      }
    }
  }
}
</style>
