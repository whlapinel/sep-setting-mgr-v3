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
    console.log("signin event triggered");
    setTimeout(() => {
        console.log("test timeout");
    }, 1000);
    console.log(e);
    console.log(e.detail);
    console.log(e.detail.expiration);
    const sessionExpiration = e.detail.expiration;
    const cushion = sessionExpiration - 5000;
    console.log("setting warningTimer for: ", warningTime(sessionExpiration, cushion));
    warningTimerID = setTimeout((cushion) => {
        let now = new Date().getTime();
        alert(`You will be signed out in ${sessionExpiration - now} milliseconds.`);
    }, warningTime(sessionExpiration, cushion));
    console.log("setting signoutTimer for: ", signoutTime(sessionExpiration));
    signoutTimerID = setTimeout(() => {
        signout();
    }, signoutTime(sessionExpiration));
    clearInterval(logInterval);
    logInterval = setInterval(() => {
        console.log("expiration: ", new Date(sessionExpiration));
        let timeRemaining = (sessionExpiration - new Date().getTime()) / 1000;
        console.log("time remaining: ", timeRemaining);
    }, 2000);
});
document.addEventListener("signout", (e) => {
    console.log("signout event triggered");
});
function signout() {
    return __awaiter(this, void 0, void 0, function* () {
        console.log("signOut");
        // dispatches custom event to trigger signout
        // the event listener for this event is in the user-signout component
        // this component will send request to server to signout
        const signoutDiv = document.querySelector("div#user-signout");
        if (signoutDiv != null) {
            signoutDiv.dispatchEvent(new CustomEvent("signout", { bubbles: true, cancelable: true }));
        }
    });
}
