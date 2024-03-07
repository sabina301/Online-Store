const addProductBtn = document.getElementById("addProduct");
addProductBtn.style.display = "none";

window.addEventListener("load", function (event) {
  event.preventDefault();
  fetch("/catalog/getrole", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.role == "admin") {
        addProductBtn.style.display = "block";
      } else {
        addProductBtn.style.display = "none";
      }
    })
    .catch((error) => console.error("Ошибка", error));
});

addProductBtn.addEventListener("click", function (event) {
  event.preventDefault();
  window.location.href = "/admin/catalog/edit";
});
