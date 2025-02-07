const {heroui} = require("@heroui/theme");

/** @type {import('tailwindcss').Config} */
export default {
    content: [
        './node_modules/@heroui/theme/dist/**/*.{js,ts,jsx,tsx}',
        "./index.html",
        "./src/**/*.{js,ts,jsx,tsx}",
    ],
    theme: {
        extend: {
            colors: {
                primary: {
                    DEFAULT: '#13293d',
                    dark: '#00a8e8'
                }
            }
        },
    },
    darkMode: 'class',
    plugins: [
        heroui()
    ],
}

