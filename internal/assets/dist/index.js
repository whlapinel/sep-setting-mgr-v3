
let activityTimeout;
let timeOutWarning;
let isSignedIn;

function setIsSignedIn(signedIn) {
    isSignedIn = signedIn;
}

async function signout() {
    console.log("signOut")
    // dispatch custom event to trigger signout
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
    }, 240000); // 4 minutes
    activityTimeout = setTimeout(() => {
        document.querySelector("div#user-status").dispatchEvent(new CustomEvent("signout", { bubbles: true, cancelable: true }));
        signout();
    }, 300000); // 5 minutes
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
    clearTimeout(activityTimeout);
})

document.querySelector("body").addEventListener("htmx:afterRequest", (e) => {
    const form = document.querySelector("#form");
    console.log("afterRequest event triggered");
    console.log(e.target)

    if (e.target.matches("#form")) {
        var xhr = e.detail.xhr;
        console.log(xhr.status);

        if (form.reportValidity()) {
            console.log("clearing forms, removing required, and closing dialog")
            document.querySelector("form").reset()
            document.querySelector("dialog").close();
        }
    }
})


document.querySelector("body").addEventListener("click", (e) => {
    if (e.target.id === "open-dialog") {
        document.querySelector("dialog").showModal();
    }
})

document.querySelector("body").addEventListener("click", (e) => {
    const closeButton = document.querySelector("dialog button#cancel");
    if (e.target === closeButton) {
        e.preventDefault();
        document.querySelector("dialog").close();
    }
})
