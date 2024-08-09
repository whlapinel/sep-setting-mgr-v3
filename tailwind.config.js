/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["**/*.templ", "**/*_templ.go"],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
  safelist: [
    'w-0',
    'w-60'
  ],
}

