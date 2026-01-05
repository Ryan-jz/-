<template>
  <div>
    <a id="verpackungsziele"></a>
    <div class="ce-bodytext layout--1">
      <h2>我们的包装目标</h2>
      <p>到 2025 年，我们所有的包装都将是可回收或可重复使用的。</p>
    </div>

    <section class="container layout--2">
      <div class="ce-bodytext layout--0">
        <h3>纸</h3>
        <p>
          我们用于初级包装的纸张100%来自可持续管理的森林。我们的外包装纸箱和二级展示架（用于保护产品在运往零售商途中以及在店内进行展示）仍然采用100%再生纸板制成。
        </p>
      </div>
      <div class="ce-bodytext layout--0">
        <h3 >玻璃</h3>
        <p>我们所有的玻璃包装均可重复灌装和使用。我们将继续使用玻璃包装来包装高端产品。</p>
      </div>
      <div class="ce-bodytext layout--0">
        <h3>塑料</h3>
        <p>
          我们现在只使用单一材料包装，即由单一类型塑料制成的包装。这种材料可以回收利用，作为再生材料用于新的塑料包装。我们避免使用由不可回收材料制成的包装。我们尽可能地探索使用再生材料。
        </p>
      </div>
    </section>
  </div>
</template>

<style scoped lang="scss">
  .layout--0{
       background-clip: padding-box;
    border-image: url('@/assets/downloaded_images/brush-box-mint.png');
    border-image-slice: 25 54 18 53 fill;
    border-image-repeat: stretch;
    border-width: 25px 54px 18px 0;
    border-style: solid;
    border-radius: 15px;
  }
.ce-bodytext.layout--1 {
  max-width: 1400px;
  margin: 0 auto;
  padding: 80px 20px;
  text-align: center;

  h2 {
    font-size: 42px;
    font-weight: 300;
    line-height: 1.3;
    margin: 0 0 25px;
    color: #ffc25c;
  }

  p {
    font-size: 18px;
    line-height: 1.8;
    color: #666;
    margin: 0;
  }
}

.container.layout--2 {
  max-width: 1400px;
  margin: 0 auto;

  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 60px;

  .ce-bodytext {
    h3 {
      font-size: 28px;
      font-weight: 500;
      line-height: 1.4;
      margin: 0 0 20px;
      color: #2b2b2e;
      text-align: center;
    }

    p {
      font-size: 18px;
      line-height: 1.8;
      color: #666;
      margin: 0;
    }
  }
}

@media (max-width: 1024px) {
  .container.layout--2 {
    grid-template-columns: 1fr;
    gap: 40px;
  }
}

@media (max-width: 768px) {
  .ce-bodytext.layout--1 {
    padding: 60px 15px;

    h2 {
      font-size: 28px;
    }

    p {
      font-size: 16px;
    }
  }

  .container.layout--2 {
    padding: 60px 15px;

    .ce-bodytext h3 {
      font-size: 22px;
    }

    .ce-bodytext p {
      font-size: 16px;
    }
  }
}
</style>
