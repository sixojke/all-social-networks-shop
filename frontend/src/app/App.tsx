import type { AppProps } from "next/app";
import "./globals.css";
import { store } from "@/app/store";
import { Provider } from "react-redux";
import { WithModalContext } from "@/hocs/WithModalContext";
import { AppModal } from "@/shared/contexts/Modal/Modal";
import { AppLayout } from "@/hocs/AppLayout";
import { AppInitializer } from "./AppInitializer";

export const App = ({ Component, pageProps }: AppProps) => {
  return (
    <Provider store={store}>
      <AppInitializer>
        <WithModalContext>
          <AppModal />
          <AppLayout>
            <Component {...pageProps} />
          </AppLayout>
        </WithModalContext>
      </AppInitializer>
    </Provider>
  );
};
