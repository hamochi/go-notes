window.onload = (event) => {
	document.getElementsByClassName("burger")[0].getElementsByClassName("icon")[0].addEventListener("click", toggleMenu)
}

function toggleMenu(event) {
	let menu = document.getElementsByClassName("container")[0].getElementsByClassName("left")[0];
	if (menu.style.display == "none") {
		menu.style.display = "block !important";
	} else {
		menu.style.display = "none"
	}
}
