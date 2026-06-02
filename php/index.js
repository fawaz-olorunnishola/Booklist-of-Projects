const menuIcon = document.querySelector(".menu-icon")
const sidebar = document.querySelector(".sidebar")
const container = document.querySelector(".container")
// const settingsMenu = document.querySelector(".settings-menu")
const darkBtn = document.getElementById("dark-btn")

// function settingsMenuToggle () {
//     settingsMenu.classList.toggle("settings-menu-height");
// }

darkBtn.onclick = function () {
    darkBtn.classList.toggle("dark-btn-on");
    document.body.classList.toggle("dark-theme");

    if(localStorage.getItem("theme") == "light") {
        localStorage.setItem("theme", "dark");
    } else {
        localStorage.setItem("theme", "light");
    } 
}

if(localStorage.getItem("theme") == "light"){
    darkBtn.classList.remove("dark-btn-on");
    document.body.classList.remove("dark-theme");

} else if(localStorage.getItem("theme") == "dark") {
    darkBtn.classList.add("dark-btn-on");
    document.body.classList.add("dark-theme");
} else{
    localStorage.setItem("theme", "light");
}

menuIcon.onclick = function() {
    sidebar.classList.toggle("small-sidebar")
    container.classList.toggle("large-container")
}