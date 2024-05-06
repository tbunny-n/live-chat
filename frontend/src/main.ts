/* We want the default action of the form, however due
to the websocket, we must manually clear the chatbox */
const sendButton = document.querySelector("#send-message-button");
sendButton?.addEventListener("click", () => {
  setTimeout(() => {
    chatbox.value = "";
  }, 250);
  chatbox.focus();
});

var chatbox: HTMLInputElement;
// On DOMContentLoaded
document.addEventListener("DOMContentLoaded", () => {
  chatbox = document.getElementById("chatbox") as HTMLInputElement;
  chatbox.focus();
});

// Set the client's username [temp]
let username = "snuzzers";

const messageContainer = document.getElementById("chat-messages") as HTMLElement;
// Handle recieving websocket messages
document.body.addEventListener("htmx:wsAfterMessage", (event: any) => {
  console.log("WebSocket message processed:", event.detail.message);
  // Add new message to chat display
  const messageInfo = JSON.parse(event.detail.message);
  const message = messageInfo.chatbox;

  const messageElement = document.createElement("div");
  messageElement.textContent = `<${username}> ${message}`;

  messageContainer.insertAdjacentHTML("beforeend", `<div hx-swap-oob="beforeend">${messageElement.outerHTML}</div>`);
});
