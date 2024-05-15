const submitBtn = document.querySelector("dialog button[type='submit']");
const inputs = document.querySelectorAll("input");
const body = document.querySelector("body");

body.addEventListener("htmx:afterRequest", (e)=> {
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


body.addEventListener("click", (e)=> {
    console.log("body clicked!")
    console.log(e.target.id === "open-dialog")
    if (e.target.id === "open-dialog") {
        console.log("showButton clicked!")
        document.querySelector("dialog").showModal();
    }
})

body.addEventListener("click", (e)=> {
    const closeButton = document.querySelector("dialog button#cancel");
    if (e.target === closeButton) {
        console.log("closeButton clicked!")        
        e.preventDefault();		
        document.querySelector("dialog").close();
    }
})
