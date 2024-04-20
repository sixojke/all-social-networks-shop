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
      "main-light-blue": "#E6F0FF",
      "main-blue": "#99C2FF",
      "main-blue-gray": "#44517E",
      "main-error-red": "#E72734",
      "main-error-dark-red": "#9C304A",
      "main-access-green": "#9FCD93",
      "main-access-dark-green": "1D8C02",
      "main-white": "#FFFFFF",
      "main-black": "#001433",
      "main-light-gray": "#999DA6",
      "main-dark-blue": "#0065FD",
    },
  },
  plugins: [],
};
export default config;
