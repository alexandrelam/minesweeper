import { useRouter } from "next/router";
import { useState, useEffect } from "react";
import { ConnectedUsers } from "../../components/ConnectedUsers";
import { Message, MessageType } from "../types/message";
import { User } from "../types/users";

export default function Game() {
  const router = useRouter();
  const [connectedUsers, setConnectedUsers] = useState<User[]>([]);

  useEffect(() => {
    if (!router.query.name) {
      router.push("/");
    }

    const ws = new WebSocket(`ws://localhost:3001/ws/${router.query.name}`);
    ws.addEventListener("open", () => {
      console.log("Connected");
    });

    ws.addEventListener("message", (event) => {
      const message: Message = JSON.parse(event.data);

      switch (message.type) {
        case MessageType.CONNECTED_USERS:
          setConnectedUsers(message.data);
      }
    });

    return () => {
      ws.close();
    };
  }, []);

  return (
    <div>
      <h1>Game</h1>
      <ConnectedUsers connectedUsers={connectedUsers} />
    </div>
  );
}
