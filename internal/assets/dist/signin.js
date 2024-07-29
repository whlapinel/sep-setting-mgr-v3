"use strict";
(() => {
    const script = document.currentScript;
    const div = script === null || script === void 0 ? void 0 : script.closest("div");
    setTimeout(() => {
        console.log("dispatching signin event");
        console.log("div", div);
        const expy = div === null || div === void 0 ? void 0 : div.getAttribute("data-expiration");
        const id = div === null || div === void 0 ? void 0 : div.id;
        console.log("id", id);
        console.log("expy", expy);
        document.dispatchEvent(new CustomEvent("signin", {
            detail: {
                expiration: Number(expy)
            }
        }));
    }, 1000);
})();
