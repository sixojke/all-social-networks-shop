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
      "main-dark-gray": "#B3B3B3",
      "main-light-gray": "#EDEDED",
      "main-black": "#262626",
      "main-white": "#FBFCFF",
      "main-error-red": "#E72734",
      "main-error-dark-red": "#9C304A",
      "main-access-green": "#9FCD93",
      white: "#FFFFFF",
      "main-access-dark-green": "1D8C02",
    },
  },
  plugins: [],
};
export default config;
