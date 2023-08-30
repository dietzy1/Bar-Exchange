/** @format */

import ENV from "@/constants/env";
import useWebSocket from "react-use-websocket";

import { useState } from "react";

const useWebsocket = () => {
  const [socketUrl, setSocketUrl] = useState(ENV.WEBSOCKET_URL);

  if (ENV.WEBSOCKET_URL === "") {
    setSocketUrl("ws://localhost:10000/ws");
  }

  const websocketConfig = {
    url: socketUrl,
    options: {
      shouldReconnect: () => true,
      reconnectInterval: 5000,
      reconnectAttempts: 30,
    },
    connect: true,
  };

  const { lastJsonMessage, readyState } = useWebSocket(
    websocketConfig.url,
    websocketConfig.options,
    websocketConfig.connect
  );

  const isWebSocketReady = readyState === 1;

  return { lastJsonMessage, isWebSocketReady };
};
export default useWebsocket;
