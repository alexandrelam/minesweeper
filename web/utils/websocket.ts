import { Message, MessageType } from "../pages/types/message";

export function handleMessages(
  event: MessageEvent<any>,
  setConnectedUsers: (users: any) => void
) {
  const message: Message = JSON.parse(event.data);

  switch (message.type) {
    case MessageType.CONNECTED_USERS:
      setConnectedUsers(message.data);
    case MessageType.USER_MOUSE:
      setConnectedUsers(message.data);
  }
}
