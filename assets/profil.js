/* Description: Script JS pour la page profil.html */

let showDescription = document.getElementById("buttonDescription");

let description = document.getElementById("description");

showDescription.addEventListener("click", () => {
  if (description.style.display == "flex") {
    description.style.display = "none";
  } else {
    description.style.display = "flex";
  }
});

let exitDescription = document.getElementById("exitDescription");

exitDescription.addEventListener("click", () => {
  description.style.display = "none";
});

/* Add post: Script JS pour la page profil.html */

let buttonPost = document.getElementById("buttonPost");

let addPost = document.getElementById("addPost");

buttonPost.addEventListener("click", () => {
  console.log("test");
  if (addPost.style.display == "flex") {
    addPost.style.display = "none";
  } else {
    addPost.style.display = "flex";
  }
});

let exitPost = document.getElementById("exitPost");

exitPost.addEventListener("click", () => {
  addPost.style.display = "none";
});
