(() => {
    setTimeout(() => {
    console.log("dispatching signin event")
    document.dispatchEvent(new CustomEvent("signin", {
        detail: {
            expiration: new Date().getTime() + 1000 * 60 * 2
        }
    }));
}, 1000);
})();			
