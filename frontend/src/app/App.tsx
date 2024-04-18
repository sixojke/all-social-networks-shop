import type { AppProps } from "next/app";
import "./globals.css";

export const App = ({ Component, pageProps }: AppProps) => {
  return <Component {...pageProps} />;
};
