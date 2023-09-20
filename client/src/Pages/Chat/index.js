import "./style.css";
import useWebSocket from 'react-use-websocket';

const WS_URL = 'ws://127.0.0.1:8080/ws';

const Chat = () => {
    useWebSocket(WS_URL, {
        onOpen: () => {
          console.log('WebSocket connection established.');
        }
    });

    return (
        <>
            <div className='content'>
                <div className='content-label'>
                    <h1> Simple Chat Go-React Websocket </h1>
                </div>
            </div>
        </>
    )
}

export default Chat