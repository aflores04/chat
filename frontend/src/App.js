import {useState} from "react";

const socket = new WebSocket("ws://localhost:8010");

const App = () => {
  const [messages, setMessages] = useState([])

  const addMessage = (message) => {
      socket.send(JSON.stringify(message));
  }

socket.onmessage = function (event) {
      let message = JSON.parse(event.data)
    setMessages([...messages, message.payload])
};

  return (
    <>
        <div id="chat-box">
            {messages.map(message => {
                return <p>{message.username}: {message.body}</p>
            })}
        </div>
        <input id={"message"} />
        <button onClick={() => addMessage({
            room_id: "1",
            username: "johndoe",
            body: "I'm a message"
        })}>Hola</button>
    </>
  );
}

export default App;
