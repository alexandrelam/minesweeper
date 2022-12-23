import { createContext, useState } from "react";

interface WebSocketContextProps {
  websocket: WebSocket | null;
  setWebsocket: React.Dispatch<React.SetStateAction<WebSocket | null>>;
}

const WebSocketContext = createContext<WebSocketContextProps>({
  websocket: null,
  setWebsocket: () => {},
});

function WebSocketProvider(props: { children: React.ReactNode }) {
  const [websocket, setWebsocket] = useState<WebSocket | null>(null);

  return (
    <WebSocketContext.Provider value={{ websocket, setWebsocket }}>
      {props.children}
    </WebSocketContext.Provider>
  );
}

export { WebSocketContext, WebSocketProvider };
