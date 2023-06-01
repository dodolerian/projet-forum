let btncomment = document.getElementById("comment");

let comment = document.getElementById("addComment");

btncomment.addEventListener("click", () => {
    if (comment.style.display == "flex") {
        comment.style.display = "none";
    } else {
        comment.style.display = "flex";
    }
})