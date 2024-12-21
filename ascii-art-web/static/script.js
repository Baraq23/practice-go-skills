const body = document.querySelector("body");
const container = document.querySelector(".container");
const themeToggler = document.querySelector(".theme-modes");

// Toggle dark mode functionality
function toggleTheme() {
  body.classList.toggle("dark-mode");
  container.classList.toggle("dark-mode");
  themeToggler.classList.toggle("active");
}

// Add evennt listener for theme toggling
themeToggler.addEventListener("click", toggleTheme);
