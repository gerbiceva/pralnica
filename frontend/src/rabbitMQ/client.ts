import { Client } from "@stomp/stompjs";
import { useState, useEffect } from "react";

const client = new Client({
  brokerURL: import.meta.env.VITE_BACKEND_RABBITMQ,
  logRawCommunication: true,
  connectHeaders: {
    login: import.meta.env.VITE_RABBITMQ_USER,
    passcode: import.meta.env.VITE_RABBITMQ_PASS,
  },
});

interface LiveDataProps {
  updatedAt: Date;
  weather: string;
}

export const useLiveData = () => {
  const [error, serError] = useState<Error | null>(null);
  const [isConnected, setIsConnected] = useState(false);
  const [data, setData] = useState<LiveDataProps>();

  // On connect, subscribe to the lectureId and languageCode topics
  client.onConnect = () => {
    console.log("connected");
    setIsConnected(true);

    client.subscribe(`/topic/weather`, (m) => {
      const msg: LiveDataProps = JSON.parse(m.body);
      setData(msg);
      // console.log({ msg }, m);
    });

    client.publish({
      destination: "/topic/weather",
      body: JSON.stringify({ weather: "sadasdasd" }),
    });
  };

  client.onDisconnect = () => {
    console.log("disconnected");
    setIsConnected(false);
  };

  client.onStompError = (err) => {
    console.log("error");
    serError(new Error(err.body));
  };

  client.onChangeState = (state) => {
    console.log({ state });
    // if (state == 2) {

    // }
  };

  client.onWebSocketError = (err) => {
    console.log({ err });
  };

  client.onUnhandledFrame = (frame) => {
    console.log({ frame });
  };

  client.activate();

  useEffect(() => {
    return () => {
      client.deactivate();
    };
  }, []);

  // client.publish({
  //   destination: "/topic/weather",
  //   body: "sadasdasd",
  // });

  return {
    isConnected,
    data,
    error,
  };
};
