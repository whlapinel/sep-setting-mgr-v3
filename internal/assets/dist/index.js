
let activityTimeout;
let timeOutWarning;
let isSignedIn;

const warningTimer = 60000; // 1 minute
const signoutTimer = 120000; // 2 minutes

function setIsSignedIn(signedIn) {
    isSignedIn = signedIn;
}

async function signout() {
    console.log("signOut")
    // dispatches custom event to trigger signout
    // the event listener for this event is in the user-status component
    // this component will send request to server to signout
    document.querySelector("div#user-status")
        .dispatchEvent(new CustomEvent("signout", { bubbles: true, cancelable: true }));
    clearTimeout(timeOutWarning);
    clearTimeout(activityTimeout);
    document.removeEventListener("mousemove", onActive)
    document.removeEventListener("click", onActive)
    document.removeEventListener("scroll", onActive)
    document.removeEventListener("keydown", onActive)
    setIsSignedIn(false)
}

async function onSignin() {
    console.log("signIn")
    ResetActivityTimeout();
    setIsSignedIn(true);
}

function onActive() {
    console.log("active")
    ResetActivityTimeout();
}

function ResetActivityTimeout() {
    clearTimeout(activityTimeout);
    clearTimeout(timeOutWarning);
    timeOutWarning = setTimeout(() => {
        console.log("timeOutWarning")
    }, warningTimer);
    activityTimeout = setTimeout(() => {
        signout();
    }, signoutTimer);
}

document.addEventListener("signin", (e) => {
    console.log("signin event triggered")
    document.addEventListener("mousemove", onActive)
    document.addEventListener("click", onActive)
    document.addEventListener("scroll", onActive)
    document.addEventListener("keydown", onActive)
    ResetActivityTimeout();
})

document.addEventListener("signout", (e) => {
    console.log("signout event triggered")
})

