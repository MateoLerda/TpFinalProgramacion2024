const baseUrl = "http://localhost:8080/foods/?filter_all=true";
const baseUrlCreatePurchase = "http://localhost:8080/purchases/";


document.addEventListener('DOMContentLoaded', getFoodsInPurchases);

function getFoodsInPurchases() {
    const userInfo = document.getElementById('user-info');
    const userMail = document.createElement('p');
    userMail.textContent = localStorage.getItem('user-mail');
    userMail.classList.add('green-color', 'bold-words', 'user-mail');
    userInfo.appendChild(userMail);
    makeRequest(
        baseUrl,
        "GET",
        "",
        "application/json",
        true,
        showFoods,
        failedGet
    );

}

function failedGet(response) {
    console.log("Failed to get foods:", response);
}

function showFoods(data) {
    if (data.message) {
        showAlert('Not found any products to buy.');
        document.getElementById("alert-button").addEventListener(('click'), () => {
            window.location.href= "../foods/foods.html"
        })        
        return
    }
    const foodTable = document.getElementById('dynamic-food-table');
    data.forEach(food => {
        const row = document.createElement('tr');
        row.setAttribute('food-code', food._id);
        row.innerHTML = `
            <td id= "name">${food.name}</td>
            <td id="price">U$D ${food.unit_price}</td>
            <td>${food.current_quantity}</td>
            <td>${food.minimum_quantity}</td>
            <td>
                <button class="decrement"><i class="fa-solid fa-minus"></i></button>
                <input type="number" value="0" size="2" readonly>
                <button class="increment"><i class="fa-solid fa-plus"></i></button>
            </td>
        `;
        const decrementButton = row.querySelector('.decrement');
        const input = row.querySelector('input');
        input.id = 'quantityToBuy';
        const incrementButton = row.querySelector('.increment');
        incrementButton.addEventListener('click', () => {
            input.value = parseInt(input.value) + 1;
        });

        decrementButton.addEventListener('click', () => {
            if (parseInt(input.value) > 0) {
                input.value = parseInt(input.value) - 1;
            }
        });

        foodTable.appendChild(row);
    });
}

document.getElementById('btnManually').addEventListener('click', manuallyPurchase);

function manuallyPurchase() {
    const foodTable = document.getElementById('dynamic-food-table');
    const rows = foodTable.getElementsByTagName('tr');
    const selectedFoods = [];

    for (let row of rows) {
        const input = row.querySelector('input');
        const quantity = parseInt(input.value);
        const name = row.querySelector('#name').textContent;
        const price = row.querySelector('#price').textContent;
        const priceValue = parseFloat(price.replace('U$D ', ''));
        if (quantity > 0) {
            const foodCode = row.getAttribute('food-code');
            selectedFoods.push({ foodCode, name, quantity, priceValue });
        }
    }

    if (selectedFoods.length == 0) {
      showAlert("You have to select a food to buy.");
      return;
    }
    let TotalCost = 0;
    const purchaseDto = {

        foods: selectedFoods.map(food => {
            TotalCost += food.priceValue * food.quantity;
            return {
                _id: food.foodCode,
                name: food.name,
                quantity: food.quantity,
            };
        }),
        total_cost: TotalCost
    };
    makeRequest(
        baseUrlCreatePurchase,
        "POST",
        purchaseDto,
        "application/json",
        true,
        successPurchase,
        failedPurchase
    );

}

function successPurchase() {
    showAlert('Purchase completed successfully, Thank you!')
    document.getElementById("alert-button").addEventListener(('click'), () => {
        window.location.href= "../home/home.html"
    })
}

function failedPurchase(response,responseBody) {
    console.log("Failed to create purchase:", response);        
    showAlert(responseBody.error);
}
