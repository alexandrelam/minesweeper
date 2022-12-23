import { useContext, useState } from "react";
import { Message, MessageType } from "../types/message";
import { User } from "../types/users";
import { WebSocketContext } from "../provider/WebSocketProvider";

type Props = {
  user: User;
};

export function Mouse({ user }: Props) {
  const [mouseX, setMouseX] = useState(0);
  const [mouseY, setMouseY] = useState(0);
  const { websocket } = useContext(WebSocketContext);

  const style = {
    left: mouseX,
    top: mouseY,
  };

  if (!websocket) return null;

  websocket.addEventListener("message", (event) => {
    const message: Message<User> = JSON.parse(event.data);

    if (message.type === MessageType.USER_MOUSE) {
      if (message.data.id !== user.id) return;

      setMouseX(message.data.mouseX);
      setMouseY(message.data.mouseY);
    }
  });

  return (
    <span
      style={style}
      className="absolute px-3 py-0.5 rounded 
                 bg-indigo-500/20 pointer-events-none
                 text-indigo-500 text-xs font-semibold"
    >
      {user.name}
    </span>
  );
}
