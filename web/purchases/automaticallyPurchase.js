let baseUrlPurchases = "http://localhost:8080/purchases/?filer_all=false";
const baseUrlCreatePurchase = "http://localhost:8080/purchases/";

document.addEventListener('DOMContentLoaded', () => {
    const userInfo = document.getElementById('user-info');
    const userMail = document.createElement('p');
    userMail.textContent = localStorage.getItem('user-mail');
    userMail.classList.add('green-color', 'bold-words', 'user-mail');
    userInfo.appendChild(userMail);
    getFoods()
});
document.getElementById("foodType").onchange = () => {
    getFoods()
}

const searchBar = document.getElementById("search-bar");

searchBar.addEventListener("submit", (e) => {
    e.preventDefault();
    getFoods();
});

let aproximation
let type

function getFoods() {
    createUrl()
    makeRequest(
        baseUrlPurchases,
        "GET",
        "",
        "application/json",
        true,
        showMinimumList,
        failedGet
    )
   // baseUrlPurchases = "http://localhost:8080/purchases/?filer_all=false";
}

function createUrl() {
    baseUrlPurchases = "http://localhost:8080/purchases/?filer_all=false";
    aproximation = document.getElementById('searchInput').value;
    type = document.getElementById('foodType').value
    if (aproximation != "") {
        baseUrlPurchases = `${baseUrlPurchases}&filter_aproximation=${aproximation}`;
    }
    if (type != "all") {
        baseUrlPurchases = `${baseUrlPurchases}&filter_type=${type}`;
    }
}


function showMinimumList(data) {
    const foodTable = document.getElementById('dynamic-food-table');
    if (data.message) {
        showAlert('There are not any products with current quantity below the minimum quantity. We will redirect you to home')
        document.getElementById("alert-button").addEventListener(('click'), () => {
            window.location.href = "../home/home.html"
        });
        
    }
    foodTable.innerHTML = ""
    data.forEach(food => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td class="food-name">${food.name}</td>
            <td>U$D ${food.unit_price}</td>
            <td>${food.current_quantity}</td>
            <td>${food.minimum_quantity}</td>
            <td class="quantity-to-buy">${food.minimum_quantity - food.current_quantity}</td>
          `;

        foodTable.appendChild(row);
    });
}

function failedGet(response) {
    console.log("Failed to get foods:", response);
}

document.getElementById('btnAutomatically').addEventListener('click', automaticallyPurchase);

function automaticallyPurchase() {
    const rows = document.querySelectorAll('#dynamic-food-table tr');
    const purchases = [];

    rows.forEach(row => {
        const foodName = row.querySelector('.food-name').textContent;
        const quantityToBuy = parseInt(row.querySelector('.quantity-to-buy').textContent, 10);

        if (quantityToBuy > 0) {
            purchases.push({
                name: foodName,
                quantity: quantityToBuy
            });
        }
    });

    const options = {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            Accept: "application/json",
            Authorization: `Bearer ${localStorage.getItem("authToken")}`,
        },
        body: JSON.stringify({ purchases })
    };

    fetch(baseUrlCreatePurchase, options)
        .then(response => {
            if (!response.ok) {
                throw new Error('Error en la respuesta de la API');
            }
            return response.json();
        })
        .then(data => {
            showAlert('Purchase completed successfully, Thank you!')
            document.getElementById("alert-button").addEventListener(('click'), () => {
                window.location.href = "../home/home.html"
            })
            console.log('Respuesta de la API:', data);
        })
        .catch(error => {
            showAlert('Failed to make purchase, please try again later.')
            document.getElementById("alert-button").addEventListener(('click'), () => {
                window.location.href = "../home/home.html"
            })
            console.error('Error al realizar la compra automática:', error);
            console.log('Failed to make purchase:', response);
           

        });
}

