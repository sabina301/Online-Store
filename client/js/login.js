document
  .getElementById("usernameForm")
  .addEventListener("submit", function (event) {
    event.preventDefault();

    var username = document.getElementById("name").value;
    var password = document.getElementById("password").value;

    fetch("/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username: username, password: password }),
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.token == "") {
        } else {
          let jwt = data.token;

          goToMainPage(jwt);
        }
      })
      .catch((error) => console.error("Ошибка", error));
  });

function goToMainPage(jwt) {
  fetch("/api/t", {
    method: "GET",
    headers: {
      Authorization: "Bearer " + jwt,
    },
  })
    .then((response) => {
      if (response.ok) {
        window.location.href = "/api/t";
      } else {
        console.error("Failed to authenticate user");
      }
    })
    .catch((error) => {
      console.error("Error:", error);
    });
}
