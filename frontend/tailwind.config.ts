import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    colors: {
      "main-light-green": "#E6FFFA",
      "main-blue": "#99C2FF",
      "main-green-gray": "#999DA6",
      "main-error-red": "#E72734",
      "main-error-dark-red": "#9C304A",
      "main-access-green": "#9FCD93",
      "main-access-dark-green": "1D8C02",
      "main-white": "#FFFFFF",
      "main-black": "#013229",
      "main-light-gray": "#999DA6",
      "main-dark-green": "#018476",
    },
  },
  plugins: [],
};
export default config;
