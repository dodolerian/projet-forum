
function commentPop(div) {
    const comment = div.getElementsByTagName('div')[0]; // get the fisrt div element
    if (comment.style.display == "flex") {
        comment.style.display = "none";
    } else {
        comment.style.display = "flex";
    }
}