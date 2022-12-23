import { useRouter } from "next/router";
import { useState, useEffect, useContext } from "react";
import { ConnectedUsers } from "../../components/ConnectedUsers";
import { Mouse } from "../../components/Mouse";
import { mousemove } from "../../utils/mousemove";
import { User } from "../../types/users";
import { WebSocketContext } from "../../provider/WebSocketProvider";
import { Tile as TileType } from "../../types/tile";
import { Message, MessageType } from "../../types/message";
import { Tile } from "../../components/Tile";

export default function Game() {
  const router = useRouter();
  const [connectedUsers, setConnectedUsers] = useState<User[]>([]);
  const [board, setBoard] = useState<TileType[][] | null>(null);
  const { websocket, setWebsocket } = useContext(WebSocketContext);

  function handleMessages(event: MessageEvent) {
    const message: Message = JSON.parse(event.data);

    switch (message.type) {
      case MessageType.CONNECTED_USERS:
        const data: User[] = message.data;
        setConnectedUsers(data);
      case MessageType.UPDATE_BOARD:
        const board: TileType[][] = message.data;
        if (board === null) return;

        setBoard(board);
    }
  }

  useEffect(() => {
    if (!router.query.name) {
      router.push("/");
    }

    const ws = new WebSocket(`ws://localhost:3001/ws/${router.query.name}`);

    setWebsocket(ws);

    ws.addEventListener("message", handleMessages);

    return () => {
      ws.removeEventListener("message", handleMessages);
      ws.close();
    };
  }, []);

  if (!websocket) {
    return <div>Loading...</div>;
  }

  return (
    <div onMouseMove={(e) => mousemove(e, websocket)} className="h-screen">
      <ConnectedUsers connectedUsers={connectedUsers} />
      {connectedUsers.map((user) => (
        <Mouse key={user.id} user={user} />
      ))}
      {board && board.length > 1 ? (
        <div className="flex flex-col gap-0.5">
          {board.map((row, i) => (
            <div key={i} className="flex gap-0.5">
              {row.map((tile, j) => (
                <Tile key={j} tile={tile} />
              ))}
            </div>
          ))}
        </div>
      ) : null}
    </div>
  );
}
