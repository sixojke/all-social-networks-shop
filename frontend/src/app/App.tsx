import type { AppProps } from "next/app";
import "./globals.css";
import { store } from "./store";
import { Provider } from "react-redux";

export const App = ({ Component, pageProps }: AppProps) => {
  return (
    <Provider store={store}>
      <Component {...pageProps} />
    </Provider>
  );
};
