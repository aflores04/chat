import {useEffect, useState} from "react";
import jwt_decode from "jwt-decode";

const socket = new WebSocket("ws://localhost:8010");

const App = () => {
    const [messages, setMessages] = useState([])

    let state = {
        message: ''
    }

    useEffect( () => {
        const fetchData = async () => {
            const response = await fetch(`http://localhost:3002/messages/list?sort_key=timestamp&sort_order=-1&amount=50`);
            await response.json().then((res) => {
                console.log(res)
                if (res.code == 400) {
                    alert(res.message)
                } else {
                    setMessages(res)
                }
            });
        }
        fetchData()
    }, [])

    const addMessage = (event) => {
        if (localStorage.key("token")) {
            let decoded = jwt_decode(localStorage.getItem("token"))

            socket.send(JSON.stringify({
                username: decoded.username,
                body: state.message
            }));
        } else {
            alert("login to publish messages")
        }
    }

    socket.onmessage = function (event) {
        let message = JSON.parse(event.data)
        setMessages([...messages, message.payload])
    };

    return (
    <>
        <div id="chat-box" style={{height: "auto"}}>
            {messages.map(message => {
                return <p>{message.username}: {message.body}</p>
            })}
        </div>
        <input id={"message"} onChange={(event) => {
            state.message = event.target.value
        }} />
        <button onClick={() => addMessage()}>Hola</button>
    </>
    );
}

export default App;
