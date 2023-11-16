// Function to open the login modal
function openLoginModal() {
  var loginModal = document.getElementById("login-form");
  console.log(loginModal); // This will print the DOM element to the console
  loginModal.style.display = "block";
}

// Get a reference to the login button
var loginButton = document.querySelector(".loginbtn");

// Add an event listener to the login button
loginButton.addEventListener("click", openLoginModal);

// Websockets
// const WebSocket = require("ws");
// const wss = new WebSocket.Server({ port: 8080 });

// let socket = new WebSocket("ws://localhost:8080/ws");
// console.log("Attempting Websocket Connection");

// socket.onopen = () => {
//   console.log("Successfully Connected");
//   socket.send("Hi From the Client");
// };

// socket.onmessage = (event) => {
//   console.log(`[message] Data received from server: ${event.data}`);
// };

// socket.onclose = (event) => {
//   console.log("Socket Closed Connection: ", event);
// };

// socket.onerror = (error) => {
//   console.log("Socket Error: ", error);
// };
