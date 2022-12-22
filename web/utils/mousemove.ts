import { throttle } from "./throttle";

export const mousemove = throttle((event: any, ws: WebSocket | undefined) => {
  if (!ws) return;

  const message = {
    type: "USER_MOUSE",
    data: {
      mouseX: event.clientX,
      mouseY: event.clientY,
    },
  };
  console.log(message);
}, 150);
