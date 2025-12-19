import { defineConfig, presetUno, presetAttributify, presetIcons } from 'unocss'

export default defineConfig({
  presets: [
    presetUno(), // 默认预设，包含 Tailwind CSS / Windi CSS 的常用工具类
    presetAttributify(), // 属性化模式，可以将工具类写成属性
    presetIcons({
      scale: 1.2,
      warn: true
    })
  ],
  shortcuts: {
    // 自定义快捷方式
    'btn': 'px-4 py-2 rounded inline-block cursor-pointer transition-all duration-300',
    'btn-primary': 'btn bg-red-600 text-white hover:bg-red-700',
    'btn-secondary': 'btn bg-white text-red-600 border-2 border-red-600 hover:bg-red-600 hover:text-white',
    'container-custom': 'max-w-1200px mx-auto px-5'
  },
  theme: {
    colors: {
      primary: '#c41e3a',
      'primary-dark': '#a01828'
    }
  }
})
