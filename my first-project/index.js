const saveEl = document.getElementById("save-el")
const countEl = document.getElementById("count-el")
const totalEl = document.getElementById("total-el")
let count = 0
let total = 0;

// console.log(saveEl)

function Registered() {
    count += 1
    countEl.textContent = count
}

function save() {
   let countStr = count + " + "
    saveEl.textContent += countStr
    countEl.textContent = 0
    totalEl.textContent += 
    total += count
    count = 0
    console.log(total)
}

console.log("Let's count people on the subway") 

let errorParagraph = document.getElementById("error")
console.log(errorParagraph)

function purchase() {
    console.log("Got a rich Customer")
    errorParagraph.textContent = "Something went wrong, please try again"
}

document.write("Today is" + Date());

