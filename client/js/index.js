document.getElementById("login").addEventListener("click", function (event) {
  event.preventDefault();

  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/auth/login", true);
  xhr.setRequestHeader("Content-Type", "application/json");
  xhr.onload = function () {
    if (xhr.status == 200) {
      window.location.href = "/auth/login";
    } else {
      console.log("Error ", xhr.status);
    }
  };
  xhr.send();
});
