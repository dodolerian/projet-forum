let user = document.getElementById("user");

let popUp = document.getElementById("popUpUser");

user.addEventListener("click", () => {
    if (popUp.style.display == "flex") {
        popUp.style.display = "none";
    } else {
        popUp.style.display = "flex";
    }
})

let exitPost = document.getElementById("exitPost");

exitPost.addEventListener("click", () => {
  popUp.style.display = "none";
})