(() => {
    const script = document.currentScript;
    const div = script?.closest("div");
    setTimeout(() => {
        console.log("dispatching signin event")
        console.log("div", div);
        const expy = div?.getAttribute("data-expiration");
        const id = div?.id;
        console.log("id", id);
        console.log("expy", expy);
        document.dispatchEvent(new CustomEvent("signin", {
            detail: {
                expiration: Number(expy)
            }
        }));
    }, 1000);
})();

