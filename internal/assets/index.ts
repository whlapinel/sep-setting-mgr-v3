


function warningTime(sessionExpiration: number, cushion: number): number {
    return sessionExpiration - cushion - new Date().getTime();
}

function signoutTime(sessionExpiration: number): number {
    return sessionExpiration - new Date().getTime();
}

let signoutTimerID = 0;
let warningTimerID = 0;
let logInterval = 0;

document.addEventListener("signin", (e: any) => {
    clearTimeout(signoutTimerID)
    clearTimeout(warningTimerID)
    clearInterval(logInterval)
    console.log("signin event triggered")

    console.log(e)
    console.log(e.detail)
    console.log(e.detail.expiration);

    const sessionExpiration = e.detail.expiration;
    const cushion = 5000
    console.log("setting warningTimer for: ", warningTime(sessionExpiration, cushion))
    warningTimerID = setTimeout((cushion: number) => {
        let now = new Date().getTime()
        alert(`You will be signed out in ${(sessionExpiration - now) / 1000} seconds.`)
    }, warningTime(sessionExpiration, cushion))

    console.log("setting signoutTimer for: ", signoutTime(sessionExpiration))
    signoutTimerID = setTimeout(() => {
        signout();
    }, signoutTime(sessionExpiration))


    logInterval = setInterval(() => {
        console.log("expiration: ", new Date(sessionExpiration))
        let timeRemaining = (sessionExpiration - new Date().getTime()) / 1000
        console.log("time remaining: ", timeRemaining)
    }, 2000)
});

document.addEventListener("userSignout", (e) => {
    console.log("userSignout event triggered")
    clearTimeout(signoutTimerID)
    clearTimeout(warningTimerID)
    clearInterval(logInterval)
});

document.addEventListener("autoSignout", (e) => {
    console.log("autoSignout event triggered")
});

// auto-signout function
async function signout() {
    clearTimeout(signoutTimerID)
    clearTimeout(warningTimerID)
    clearInterval(logInterval)
    // dispatches custom event to trigger signout
    // the event listener for this event is in the user-signout component
    // this component will send request to server to signout
    const signoutDiv = document.querySelector("div#user-signout");
    signoutDiv?.dispatchEvent(new CustomEvent("autoSignout", { bubbles: true, cancelable: true }));
}

const userMenuButton = document.querySelector("button#user-menu-button");
userMenuButton?.addEventListener("click", () => {
    const userMenu = document.querySelector("div#user-menu");
    userMenu?.classList.toggle("hidden");
});





