/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
        fontFamily:{
            poppins:["Poppins", "sans-serif"],
            grotesque: ["Darker grotesque", "sans-serif"],
        },
        colors: {
          'custompink': 'rgba(244, 239, 255, 1)',
        },
    },
  },
  plugins: [],
}

