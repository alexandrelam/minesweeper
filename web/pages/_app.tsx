import "../styles/globals.css";
import type { AppProps } from "next/app";
import { WebSocketProvider } from "../provider/WebSocketProvider";
import { HistoryProvider } from "../provider/HistoryProvider";

export default function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <HistoryProvider>
        <WebSocketProvider>
          <Component {...pageProps} />;
        </WebSocketProvider>
      </HistoryProvider>
    </>
  );
}

