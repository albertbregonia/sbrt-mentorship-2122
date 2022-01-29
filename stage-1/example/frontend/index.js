function loadLocalVideo() {
    navigator.mediaDevices.getUserMedia({ video: {
        width: 1280,
        height: 720,
    }, audio: true})
    .then(stream => {
        const localVideo = document.createElement(`video`);
        localVideo.srcObject = stream;
        localVideo.muted = localVideo.autoplay = true;
        document.getElementById(`videos`).appendChild(localVideo);
    });
}

const clock = document.getElementById(`clock`);
function clockHandler() {
    const defaultTimeFormat = {
        weekday: `long`, 
        year: `numeric`, 
        month: `long`, 
        day: `numeric`,
        hour: `numeric`,
        minute: `2-digit`,
        second: `2-digit`,
        timeZoneName: `short`
    };
    setInterval(() => {
        const now = new Date().toLocaleString(`en-US`, defaultTimeFormat),
        first = now.indexOf(`,`), last = now.lastIndexOf(`,`),
        day = now.substring(0, first),
        date = now.substring(first + 1, last),
        time = now.substring(last + 1);
        clock.innerHTML = time; //this is HTML so the white space doesn't matter
        clock.title = `${day}, ${date}`;
    }, 1000);
}

clockHandler();