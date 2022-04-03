const chat = document.getElementById(`chat`),
      usernameInput = document.getElementById(`username`),
      chatInputForm = document.getElementById(`chat-input-form`),
      chatInput = document.getElementById(`chat-input`);

function loginHandler() {
    fetch(`login?${new URLSearchParams({username: usernameInput.value,})}`) //connect to login handler and get a session id
    .then(response => response.text())
    .then(response => {
        if(isNaN(parseInt(response))) {
            alert(response);
        } else {
            const session = new URLSearchParams({username: usernameInput.value, sessionID: response}),
                  ws = new WebSocket(`wss://${location.hostname}:${location.port}/chat?${session}`);
            ws.onopen = () => alert(`connected to the backend`);
            ws.onclose = () => alert(`disconnected from the backend`);
            ws.onmessage = ({data}) => {
                const msg =  JSON.parse(data),
                    record = document.createElement(`li`);
                record.innerText = `${msg.sender}: ${msg.content}`;
                chat.appendChild(record);
            };
            chatInputForm.onsubmit = () => {
                ws.send(JSON.stringify({
                    content: chatInput.value
                }));
                return false;
            };
        }
    })
    .catch(alert);
    return false;
}