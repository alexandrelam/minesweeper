import { useRouter } from "next/router";
import { useState, useEffect, useRef } from "react";
import { ConnectedUsers } from "../../components/ConnectedUsers";
import { mousemove } from "../../utils/mousemove";
import { handleMessages } from "../../utils/websocket";
import { User } from "../types/users";

export default function Game() {
  const router = useRouter();
  const [connectedUsers, setConnectedUsers] = useState<User[]>([]);
  const wsRef = useRef<WebSocket>();

  useEffect(() => {
    if (!router.query.name) {
      router.push("/");
    }

    wsRef.current = new WebSocket(
      `ws://localhost:3001/ws/${router.query.name}`
    );

    wsRef.current.addEventListener("message", (event) => {
      handleMessages(event, setConnectedUsers);
    });

    return () => {
      if (!wsRef.current) {
        return;
      }

      wsRef.current.removeEventListener("message", (event) =>
        handleMessages(event, setConnectedUsers)
      );
      wsRef.current.close();
    };
  }, []);

  return (
    <div onMouseMove={(e) => mousemove(e, wsRef.current)} className="h-screen">
      <h1>Game</h1>
      <ConnectedUsers connectedUsers={connectedUsers} />
    </div>
  );
}
