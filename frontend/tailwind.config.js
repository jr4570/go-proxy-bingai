/** @type {import('tailwindcss').Config} */
module.exports = {
  // 禁用預加載，修復tailwind樣式 與 naive-ui button等樣式等衝突問題
  corePlugins: {
    preflight: false,
  },
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {},
  },
  plugins: [],
};
