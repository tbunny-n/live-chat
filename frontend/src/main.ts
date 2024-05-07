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

