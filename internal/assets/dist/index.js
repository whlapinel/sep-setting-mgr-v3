
let activityTimeout;
let timeOutWarning;
let isSignedIn;

const warningTimer = 10 * 60000;
const signoutTimer = 15 * 60000;

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

function debounce(func, delay) {
    let timer;
    return function (...args) {
        clearTimeout(timer);
        timer = setTimeout(() => {
            func.apply(this, args);
        }, delay);
    };
}

function onActive() {
    console.log("active")
    // debounce this function

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
    document.addEventListener("mousemove", debounce(onActive, 1000))
    document.addEventListener("click", debounce(onActive, 1000))
    document.addEventListener("scroll", debounce(onActive, 1000))
    document.addEventListener("keydown", debounce(onActive, 1000))
    ResetActivityTimeout();
})

document.addEventListener("signout", (e) => {
    console.log("signout event triggered")
})

