const sendButton = document.querySelector("#send-message-button");

var chatbox: HTMLInputElement;

// On DOMContentLoaded
document.addEventListener("DOMContentLoaded", () => {
  chatbox = document.getElementById("chatbox") as HTMLInputElement;
  chatbox.focus();
});

/* We want the default action of the form, however due
to the websocket, we must manually clear the chatbox */
sendButton?.addEventListener("click", () => {
  setTimeout(() => {
    chatbox.value = "";
  }, 250);
  chatbox.focus();
});
