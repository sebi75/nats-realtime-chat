/* eslint-disable @typescript-eslint/no-unsafe-assignment */
import { useCallback, useEffect, useRef, useState } from "react";

export const useChatWebsocket = (channelId: string) => {
  const [data, setData] = useState<string>("");
  const [error, setError] = useState<string>("");
  const websocketConnection = useRef<WebSocket | null>(null);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      setError("Unauthorized");
      return;
    }

    const ws = new WebSocket(
      `ws://localhost:8080/connect?channelId=${channelId}&token=${token}`
    );
    websocketConnection.current = ws;

    ws.onmessage = (event) => {
      const data = event.data;
      setData(data as string);
    };

    return () => {
      ws.close();
    };
  }, [channelId]);

  const send = useCallback(
    (message: string, options: Record<string, unknown>) => {
      if (!websocketConnection.current) {
        return;
      }
      const data = {
        type: options.type,
        data: {
          message: message,
        },
      };

      websocketConnection.current.send(JSON.stringify(data));
    },
    []
  );

  return { data, error, send };
};
