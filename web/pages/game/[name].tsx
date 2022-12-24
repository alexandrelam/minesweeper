import { useRouter } from "next/router";
import { useState, useEffect, useContext } from "react";
import { ConnectedUsers } from "../../components/ConnectedUsers";
import { Mouse } from "../../components/Mouse";
import { mousemove } from "../../utils/mousemove";
import { User } from "../../types/users";
import { WebSocketContext } from "../../provider/WebSocketProvider";
import { Tile as TileType } from "../../types/tile";
import { EventMessage, Message, MessageType } from "../../types/message";
import { MemoizedTile } from "../../components/Tile";

export default function Game() {
  const router = useRouter();
  const [connectedUsers, setConnectedUsers] = useState<User[]>([]);
  const [board, setBoard] = useState<TileType[][] | null>(null);
  const { websocket, setWebsocket } = useContext(WebSocketContext);

  function handleMessages(event: MessageEvent) {
    let message: Message = JSON.parse(event.data);

    switch (message.type) {
      case MessageType.CONNECTED_USERS:
        const data: User[] = message.data;
        setConnectedUsers(data);
        break;
      case MessageType.UPDATE_BOARD:
        setBoard((message as Message<EventMessage>).data.board);
        break;
      case MessageType.GAME_LOST:
        setBoard((message as Message<EventMessage>).data.board);
        alert("Dommage, vous avez perdu !");
        break;
      case MessageType.GAME_WON:
        setBoard((message as Message<EventMessage>).data.board);
        alert("Bravo, vous avez gagné !");
        break;
      case MessageType.HISTORY:
        console.log(message.data);
        break;
    }
  }

  useEffect(() => {
    if (!router.query.name) {
      router.push("/");
    }

    const ws = new WebSocket(`ws://10.0.0.237:3001/ws/${router.query.name}`);

    setWebsocket(ws);

    ws.addEventListener("message", handleMessages);

    return () => {
      ws.removeEventListener("message", handleMessages);
      ws.close();
    };
  }, []);

  function createNewGame() {
    if (!websocket) return;
    setBoard(null);
    websocket.send(JSON.stringify({ action: MessageType.CREATE_GAME }));
  }

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
        <div className="flex flex-col gap-0.5 ml-10 mt-10">
          {board.map((row, i) => (
            <div key={i} className="flex gap-0.5">
              {row.map((tile, j) => (
                <MemoizedTile key={`${i}${j}`} tile={tile} row={i} column={j} />
              ))}
            </div>
          ))}
        </div>
      ) : null}
      <button
        onClick={createNewGame}
        className="ml-10 mt-10 px-5 py-2 bg-indigo-200
                   text-indigo-500 rounded font-bold
                   hover:bg-indigo-300 hover:text-indigo-600 hover:shadow-md"
      >
        Nouvelle partie
      </button>
    </div>
  );
}
