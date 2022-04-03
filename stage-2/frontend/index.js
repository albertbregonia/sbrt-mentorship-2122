const usernameInput = document.getElementById(`username`),
      chatInputForm = document.getElementById(`chat-input-form`),
      chatInput = document.getElementById(`chat-input`),
      chatLog = document.getElementById(`chat-log`);

function loginHandler() {
    fetch(`login?${new URLSearchParams({
        username: usernameInput.value //connect to the login handler in the backend
    })}`)
    .then(response => response.text())
    .then(msg => {
        if(isNaN(parseInt(msg))) { //if we got an error
            alert(msg);
        } else { //we got an actual sessionID
            const session = new URLSearchParams({username: usernameInput.value, sessionID: msg}),
                  ws = new WebSocket(`wss://${location.hostname}:${location.port}/chat?${session}`);
            ws.onopen = () => alert(`connected to the backend!`);
            ws.onclose = () => alert(`disconnected from the backend!`);
            ws.onmessage = ({data}) => {
                const msg = JSON.parse(data), //parse the msg
                      element = document.createElement(`li`);
                element.innerText = `${msg.sender}: ${msg.content}`; //set text format
                chatLog.appendChild(element); //display in html
            };
            chatInputForm.onsubmit = () => {
                ws.send(JSON.stringify({
                    content: chatInput.value
                }));
                return false; //prevents refresh
            };
        }
    });
    return false; //prevents refresh
}