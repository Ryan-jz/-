<template>
  <section class="container layout--5">
    <div class="ce-bodytext layout--0">
      <h2>我们的证书</h2>
      <p>
        除了我们的指导方针和目标之外，我们还可以出示相关认证，以体现我们对生态效率的可持续追求。
      </p>
    </div>

    <div 
      v-for="(cert, index) in certificates" 
      :key="index"
      class="downloads"
    >
      <div class="preview">
        <img 
          :src="cert.previewSrc" 
          :width="cert.previewWidth" 
          :height="cert.previewHeight" 
          :alt="cert.previewAlt"
        >
      </div>
      <div class="description-and-files">
        <div class="description">
          <h3>{{ cert.title }}</h3>
          <p>{{ cert.description }}</p>
        </div>
        <ul class="files">
          <li 
            v-for="(file, fileIndex) in cert.files" 
            :key="fileIndex"
            class="pdf file"
          >
            <a 
              :href="file.href" 
              :title="file.title" 
              target="_blank"
            >
              {{ file.text }}
            </a>
          </li>
        </ul>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref } from 'vue'
import iso50001Img from '@/assets/downloaded_images/ISO_50001_Sammel_bis_2026_de_Seite_1.png'
import iso14001Img from '@/assets/downloaded_images/ISO_9001_140001_Kombi_SWS_D_aktuell.jpg'
import iso50001DePdf from '@/assets/downloaded_images/ISO_50001_de_bis_2026.pdf'
import iso50001EnPdf from '@/assets/downloaded_images/ISO_50001_en_bis_2026.pdf'
import iso14001DePdf from '@/assets/downloaded_images/ISO_9001_140001_Kombi_SWS_D_aktuell.pdf'
import iso14001EnPdf from '@/assets/downloaded_images/ISO_9001_140001_Kombi_SWS_E_aktuell.pdf'

const certificates = ref([
  {
    previewSrc: iso50001Img,
    previewWidth: 1052,
    previewHeight: 1492,
    previewAlt: '根据 ISO 50001 标准颁发的能源管理体系证书',
    title: 'ISO 50001',
    description: '（符合ISO 50001标准的能源管理体系认证）',
    files: [
      { 
        href: iso50001DePdf, 
        title: 'ISO 50001（德语）', 
        text: 'ISO 50001（德语）' 
      },
      { 
        href: iso50001EnPdf, 
        title: 'ISO 50001（英文）', 
        text: 'ISO 50001（英文）' 
      }
    ]
  },
  {
    previewSrc: iso14001Img,
    previewWidth: 595,
    previewHeight: 841,
    previewAlt: '西南德意志盐业股份公司获得了 TÜV 认证，该公司已实施并应用了质量和环境管理体系。',
    title: 'ISO 14001',
    description: '（符合ISO 14001标准的环境管理体系证书）',
    files: [
      { 
        href: iso14001DePdf, 
        title: 'ISO_9001_140001_Kombi_SWS_D_aktuell.pdf', 
        text: 'ISO_9001_140001_Kombi_SWS_D_aktuell.pdf' 
      },
      { 
        href: iso14001EnPdf, 
        title: 'ISO_9001_140001_Kombi_SWS_E_aktuell.pdf', 
        text: 'ISO_9001_140001_Kombi_SWS_E_aktuell.pdf' 
      }
    ]
  }
])
</script>

<style scoped lang="scss">
.container.layout--5 {
  max-width: 1400px;
  margin: 0 auto;
  padding: 100px 20px;

  .ce-bodytext.layout--0 {
    text-align: center;
    margin-bottom: 60px;

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

.downloads {
  display: flex;
  gap: 40px;
  margin-bottom: 50px;
  padding: 40px;
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  transition: box-shadow 0.3s;

  &:hover {
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.12);
  }

  .preview {
    flex: 0 0 220px;

    img {
      width: 100%;
      height: auto;
      border: 1px solid #e0e0e0;
      border-radius: 4px;
    }
  }

  .description-and-files {
    flex: 1;

    .description {
      margin-bottom: 25px;

      h3 {
        font-size: 24px;
        font-weight: 600;
        margin: 0 0 10px;
        color: #2b2b2e;
      }

      p {
        color: #888;
        font-size: 15px;
        margin: 0;
      }
    }

    .files {
      list-style: none;
      padding: 0;
      margin: 0;
      display: flex;
      gap: 15px;
      flex-wrap: wrap;

      li a {
        display: inline-flex;
        align-items: center;
        padding: 12px 24px;
        background: #c41e3a;
        color: #fff;
        text-decoration: none;
        border-radius: 6px;
        font-size: 15px;
        font-weight: 500;
        transition: all 0.3s;

        &:hover {
          background: #a01830;
          transform: translateY(-2px);
          box-shadow: 0 4px 12px rgba(196, 30, 58, 0.3);
        }

        &::before {
          content: '📄';
          margin-right: 10px;
          font-size: 18px;
        }
      }
    }
  }
}

@media (max-width: 768px) {
  .container.layout--5 {
    padding: 60px 15px;

    .ce-bodytext.layout--0 {
      margin-bottom: 40px;

      h2 {
        font-size: 28px;
      }

      p {
        font-size: 16px;
      }
    }
  }

  .downloads {
    flex-direction: column;
    padding: 25px;
    gap: 25px;

    .preview {
      flex: 0 0 auto;
    }

    .description-and-files .files {
      flex-direction: column;

      li a {
        justify-content: center;
      }
    }
  }
}
</style>
