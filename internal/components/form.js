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
    console.log(e.target)
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
