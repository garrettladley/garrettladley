/** @type {import('tailwindcss').Config} */
export default {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go",],
  theme: {
    extend: {
      colors: {
        midnight: '#040b26',
      },
      fontFamily: {
        mono: ['Courier New', 'Courier', 'monospace'],
      },
    },
  },
  plugins: [],
};
