toggleSideNav();
setTimeout(() => {
    toggleSideNav();
}, 1000);

function toggleSideNav() {
    const sideNav = document.querySelector("#side-nav");
    const div1 = document.querySelector("#side-nav-div-1");
    const div2 = document.querySelector("#side-nav-div-2");
    const subDiv = document.querySelector("#side-nav-subdiv");
    const nav = document.querySelector("#side-nav nav");
    sideNav?.classList.toggle("w-60");
    sideNav?.classList.toggle("w-fit");
    div1?.classList.toggle("w-0");
    div1?.classList.toggle("w-full");
    div2?.classList.toggle("w-0");
    div2?.classList.toggle("w-fit");
    nav?.classList.toggle("hidden");
    nav?.classList.toggle("flex");
    subDiv?.classList.toggle("opacity-100");
    subDiv?.classList.toggle("opacity-0");
}

