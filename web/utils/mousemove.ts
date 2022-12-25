import { throttle } from "./throttle";

export const mousemove = throttle((event: any, ws: WebSocket | undefined) => {
  if (!ws) return;

  const message = {
    action: "USER_MOUSE",
    mouseX: event.clientX,
    mouseY: event.clientY,
  };

  try {
    ws.send(JSON.stringify(message));
  } catch (error) {
    console.log(error);
  }
}, 100);
