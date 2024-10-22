import "./App.css";
import React, { useEffect, useRef } from "react";

function App() {
  const socketUrl = "ws://localhost:8080/";

  const connection = useRef(null);
  useEffect(() => {
    const socket = new WebSocket(socketUrl);

    socket.addEventListener("open", (event) => {
      socket.send("Connection established");
    });

    // Listen for messages
    socket.addEventListener("message", (event) => {
      console.log("Message from server ", event.data);
    });

    connection.current = socket;

    return () => connection.current.close();
  }, []);

  const handleKeyDown = (event) => {
    connection.current.send(event.key);
  };

  const handleKeyUp = (event) => {
    connection.current.send("u" + event.key);
  };

  return (
    <div>
      <input
        type="text"
        onPaste={(e) => e.preventDefault()}
        onKeyDown={handleKeyDown}
        onKeyUp={handleKeyUp}
      />
    </div>
  );
}

export default App;
