
let buttonDescription = document.getElementById("buttonDescription");
 

buttonDescription.addEventListener("click", () => {
    console.log("test")
    if(d1.style.display != "none"){
      d1.style.display = "none";
    } else {
      d1.style.display = "block";
    }
  })