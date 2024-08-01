"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
function warningTime(sessionExpiration, cushion) {
    return sessionExpiration - cushion - new Date().getTime();
}
function signoutTime(sessionExpiration) {
    return sessionExpiration - new Date().getTime();
}
let signoutTimerID = 0;
let warningTimerID = 0;
let logInterval = 0;
document.addEventListener("signin", (e) => {
    clearTimeout(signoutTimerID);
    clearTimeout(warningTimerID);
    clearInterval(logInterval);
    console.log("signin event triggered");
    console.log(e);
    console.log(e.detail);
    console.log(e.detail.expiration);
    const sessionExpiration = e.detail.expiration;
    const cushion = 5000;
    console.log("setting warningTimer for: ", warningTime(sessionExpiration, cushion));
    warningTimerID = setTimeout((cushion) => {
        let now = new Date().getTime();
        alert(`You will be signed out in ${(sessionExpiration - now) / 1000} seconds.`);
    }, warningTime(sessionExpiration, cushion));
    console.log("setting signoutTimer for: ", signoutTime(sessionExpiration));
    signoutTimerID = setTimeout(() => {
        signout();
    }, signoutTime(sessionExpiration));
    logInterval = setInterval(() => {
        console.log("expiration: ", new Date(sessionExpiration));
        let timeRemaining = (sessionExpiration - new Date().getTime()) / 1000;
        console.log("time remaining: ", timeRemaining);
    }, 2000);
});
document.addEventListener("userSignout", (e) => {
    console.log("userSignout event triggered");
    clearTimeout(signoutTimerID);
    clearTimeout(warningTimerID);
    clearInterval(logInterval);
});
document.addEventListener("autoSignout", (e) => {
    console.log("autoSignout event triggered");
});
// auto-signout function
function signout() {
    return __awaiter(this, void 0, void 0, function* () {
        clearTimeout(signoutTimerID);
        clearTimeout(warningTimerID);
        clearInterval(logInterval);
        // dispatches custom event to trigger signout
        // the event listener for this event is in the user-signout component
        // this component will send request to server to signout
        const signoutDiv = document.querySelector("div#user-signout");
        signoutDiv === null || signoutDiv === void 0 ? void 0 : signoutDiv.dispatchEvent(new CustomEvent("autoSignout", { bubbles: true, cancelable: true }));
    });
}
const userMenuButton = document.querySelector("button#user-menu-button");
userMenuButton === null || userMenuButton === void 0 ? void 0 : userMenuButton.addEventListener("click", () => {
    const userMenu = document.querySelector("div#user-menu");
    userMenu === null || userMenu === void 0 ? void 0 : userMenu.classList.toggle("hidden");
});
