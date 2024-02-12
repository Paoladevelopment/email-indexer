import scrollbarPlugin from 'tailwind-scrollbar';
/** @type {import('tailwindcss').Config} */
export default {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  content: [],
  theme: {
    extend: {
      flexBasis: {
        '5/100': '5%',
      },
      width: {
        '95/100': '95%',
      },
      height: {
        '95/100': '95%',
        '9/10': '90%'
      },
      colors: {
        'black': "#000"
      }
    },
  },
  plugins: [
    scrollbarPlugin({ nocompatible: true })
  ],
}

