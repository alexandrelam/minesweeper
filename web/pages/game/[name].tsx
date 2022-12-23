import { useRouter } from "next/router";
import { useState, useEffect, useContext } from "react";
import { ConnectedUsers } from "../../components/ConnectedUsers";
import { Mouse } from "../../components/Mouse";
import { mousemove } from "../../utils/mousemove";
import { handleMessages } from "../../utils/websocket";
import { User } from "../types/users";
import { WebSocketContext } from "../../provider/WebSocketProvider";

export default function Game() {
  const router = useRouter();
  const [connectedUsers, setConnectedUsers] = useState<User[]>([]);
  const { websocket, setWebsocket } = useContext(WebSocketContext);

  useEffect(() => {
    if (!router.query.name) {
      router.push("/");
    }

    const ws = new WebSocket(`ws://localhost:3001/ws/${router.query.name}`);

    setWebsocket(ws);

    ws.addEventListener("message", (event) => {
      handleMessages(event, setConnectedUsers);
    });

    return () => {
      ws.removeEventListener("message", (event) =>
        handleMessages(event, setConnectedUsers)
      );
      ws.close();
    };
  }, []);

  if (!websocket) {
    return <div>Loading...</div>;
  }

  return (
    <div onMouseMove={(e) => mousemove(e, websocket)} className="h-screen">
      <h1>Game</h1>
      <ConnectedUsers connectedUsers={connectedUsers} />
      {connectedUsers.map((user) => (
        <Mouse key={user.id} user={user} />
      ))}
    </div>
  );
}
