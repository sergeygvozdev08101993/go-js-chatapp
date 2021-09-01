let socket = new WebSocket("ws://localhost:3030/ws");

let connect = (callback) => {

    console.log("attempting connection...");

    socket.onopen = () => {
        console.log("successfully connected");
    };

    socket.onmessage = (msg) => {
      console.log(msg);
      callback(msg);
    };

    socket.onclose = (event) => {
        console.log("socket connection closed: ", event);
    };

    socket.onerror = (error) => {
        console.log("socket error: ", error);
    };
};

let sendMsg = (msg) => {
    console.log("sending message: ", msg);
    socket.send(msg);
}

export { connect, sendMsg };