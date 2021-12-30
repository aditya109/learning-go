let socket = new WebSocket("ws://localhost:8000/ws");
console.log("Attempting WebSocket Connection");

socket.onopen = () => {
    console.log("Successfully Connected");
    socket.send("Hi From The Client !");
}
socket.onclose = (event) => {
    console.log("Socket Closed Connection: ", event);
}
socket.onmessage = (msg) => {
    console.log(msg);
}
socket.onerror = (error) => {
    console.log("Socket Error:", error);
}