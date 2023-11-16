// Get references to the form elements
var x = document.getElementById("login");
var y = document.getElementById("register");
var z = document.getElementById("btn");

// Function to slide the forms for login and registration
function register() {
  x.style.left = "-400px";
  y.style.left = "50px";
  z.style.left = "110px";
}

function login() {
  x.style.left = "50px";
  y.style.left = "450px";
  z.style.left = "0px";
}

// Get reference to the login modal
var loginModal = document.getElementById("login-form");

// Close the login modal if the user clicks outside it
window.onclick = function (event) {
  if (event.target == loginModal) {
    loginModal.style.display = "none";
  }
};
