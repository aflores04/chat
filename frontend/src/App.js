import {useState} from "react";

const socket = new WebSocket("ws://localhost:8000");

const App = () => {
  const [messages, setMessages] = useState([])

  const addMessage = (message) => {
      socket.send(JSON.stringify(message));
  }

socket.onmessage = function (event) {
    setMessages([...messages, event.data])
};

  return (
    <>
        <div id="chat-box">
            <ul>
                {messages.map(message => {
                    return <li key={message.username}>{message.body}</li>
                })}
            </ul>
        </div>
        <button onClick={() => addMessage({
            room_id: 1,
            username: "johndoe",
            body: "I'm a message"
        })}>Hola</button>
    </>
  );
}

export default App;
