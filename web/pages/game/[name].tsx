import { useRouter } from "next/router";
import { useEffect } from "react";

export default function Game() {
  const router = useRouter();

  useEffect(() => {
    if (!router.query.name) {
      router.push("/");
    }

    const ws = new WebSocket(`ws://localhost:3001/ws/${router.query.name}`);
    ws.addEventListener("open", () => {
      console.log("Connected");
    });

    ws.addEventListener("message", (event) => {
      // parse JSON message received from server
      const data = JSON.parse(event.data);
      console.log(data);
    });

    return () => {
      ws.close();
    };
  }, []);

  return (
    <div>
      <h1>Game</h1>
    </div>
  );
}
