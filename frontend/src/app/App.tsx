import type { AppProps } from "next/app";
import "./globals.css";
import { store } from "./store";
import { Provider } from "react-redux";
import { WithModalContext } from "@/hocs/WithModalContext";
import { AppModal } from "@/shared/contexts/Modal/Modal";
import { AppLayout } from "@/hocs/AppLayout";

export const App = ({ Component, pageProps }: AppProps) => {
  return (
    <Provider store={store}>
      <WithModalContext>
        <AppModal />
        <AppLayout>
          <Component {...pageProps} />
        </AppLayout>
      </WithModalContext>
    </Provider>
  );
};
