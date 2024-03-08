document
  .getElementById("productForm")
  .addEventListener("submit", function (event) {
    event.preventDefault();

    var category = document.getElementById("category").value;
    var name = document.getElementById("name").value;
    var price = document.getElementById("price").value;
    var description = document.getElementById("description").value;
    var color = document.getElementById("color").value;

    fetch("/admin/catalog/edit/add", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        category: category,
        name: name,
        price: price,
        description: description,
        color: color,
      }),
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.id < 0) {
          console.log("Ошибка при добавлении товара");
        } else {
          console.log("OK!");
        }
      })
      .catch((error) => console.error("Ошибка", error));
  });
