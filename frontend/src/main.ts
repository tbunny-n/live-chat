const sendButton = document.querySelector("#send-message-button");

sendButton?.addEventListener("click", (event: Event) => {
  event.preventDefault();
  const chatbox = document.getElementById("chatbox") as HTMLInputElement;
  chatbox.value = "";
});
