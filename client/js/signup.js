document
  .getElementById("usernameForm")
  .addEventListener("submit", function (event) {
    event.preventDefault();

    var username = document.getElementById("name").value;
    var password = document.getElementById("password").value;

    fetch("/auth/signup", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username: username, password: password }),
    })
      .then((response) => response.json())
      .then((data) => {
        window.location.href = "login";
      })
      .catch((error) => console.error("Ошибка", error));
  });
