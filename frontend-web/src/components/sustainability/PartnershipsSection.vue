<template>
  <div>
    <a id="partnerschaften-und-projekte"></a>
    <div class="alternative-background layout-2">
      <div class="ce-bodytext layout--0">
        <h2>
          <span class="subtitle">我们还能做些什么？</span><br>
          合作伙伴关系和项目
        </h2>
        <p>
          环境责任包含诸多方面。除了产品和包装之外，我们也希望在其他领域践行可持续发展，并支持那些激励我们的项目。因为每一份对地球未来的承诺都至关重要！
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.alternative-background.layout-2 {
  background: #fff;
  padding: 80px 0;

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

      .subtitle {
        display: block;
        font-size: 16px;
        color: #c41e3a;
        text-transform: uppercase;
        letter-spacing: 2px;
        margin-bottom: 12px;
        font-weight: 600;
      }
    }

    p {
      font-size: 18px;
      line-height: 1.8;
      color: #666;
      margin: 0;
      max-width: 900px;
      margin: 0 auto;
    }
  }
}

@media (max-width: 768px) {
  .alternative-background.layout-2 {
    padding: 60px 0;

    .ce-bodytext {
      padding: 0 15px;

      h2 {
        font-size: 28px;

        .subtitle {
          font-size: 13px;
        }
      }

      p {
        font-size: 16px;
      }
    }
  }
}
</style>
