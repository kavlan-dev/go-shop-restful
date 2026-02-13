// Cart management for Go Shop

// Initialize cart page
async function initCartPage() {
  if (!isAuthenticated()) {
    document.getElementById("cart-content").innerHTML = `
            <div class="alert alert-warning">
                <h4 class="alert-heading">Требуется авторизация</h4>
                <p>Пожалуйста, <a href="/login.html" class="alert-link">войдите</a> в систему, чтобы просмотреть вашу корзину.</p>
            </div>
        `;
    return;
  }

  await loadCart();
}

// Load cart
async function loadCart() {
  try {
    const response = await fetch(`${API_BASE_URL}/cart`, {
      headers: getAuthHeader(),
    });

    if (!response.ok) {
      throw new Error("Failed to load cart");
    }

    const cart = await response.json();
    renderCart(cart);
  } catch (error) {
    console.error("Error loading cart:", error);
    document.getElementById("cart-content").innerHTML = `
            <div class="alert alert-danger">
                <h4 class="alert-heading">Ошибка</h4>
                <p>Не удалось загрузить корзину: ${error.message}</p>
            </div>
        `;
  }
}

// Render cart
function renderCart(cart) {
  if (!cart || !cart.items || cart.items.length === 0) {
    document.getElementById("cart-content").innerHTML = `
            <div class="alert alert-info">
                <h4 class="alert-heading">Ваша корзина пуста</h4>
                <p>Перейдите в <a href="/products.html" class="alert-link">каталог продуктов</a>, чтобы добавить товары.</p>
            </div>
        `;
    return;
  }

  let html = `
        <div class="table-responsive">
            <table class="table table-dark table-hover">
                <thead>
                    <tr>
                        <th>Продукт</th>
                        <th>Цена</th>
                        <th>Количество</th>
                        <th>Итого</th>
                        <th>Действия</th>
                    </tr>
                </thead>
                <tbody>
    `;

  let total = 0;

  cart.items.forEach((item) => {
    const itemTotal = item.price * item.quantity;
    total += itemTotal;

    html += `
            <tr>
                <td>
                    <strong>${escapeHtml(item.product.title)}</strong>
                    <br>
                    <small class="text-muted">${escapeHtml(item.product.category)}</small>
                </td>
                <td>${formatPrice(item.price)}</td>
                <td>
                    <div class="input-group" style="width: 120px;">
                        <button class="btn btn-sm btn-outline-secondary decrease-quantity" data-item-id="${item.ID}">-</button>
                        <input type="text" class="form-control text-center quantity-input" value="${item.quantity}" readonly>
                        <button class="btn btn-sm btn-outline-secondary increase-quantity" data-item-id="${item.ID}">+</button>
                    </div>
                </td>
                <td>${formatPrice(itemTotal)}</td>
                <td>
                    <button class="btn btn-sm btn-outline-danger remove-from-cart" data-item-id="${item.ID}">
                        <i class="bi bi-trash"></i> Удалить
                    </button>
                </td>
            </tr>
        `;
  });

  html += `
                </tbody>
                <tfoot>
                    <tr>
                        <th colspan="3" class="text-end">Итого:</th>
                        <th colspan="2">${formatPrice(total)}</th>
                    </tr>
                </tfoot>
            </table>
        </div>

        <div class="d-flex justify-content-between mt-4">
            <button id="clear-cart-btn" class="btn btn-outline-danger">
                <i class="bi bi-cart-x"></i> Очистить корзину
            </button>
            <button id="checkout-btn" class="btn btn-success">
                <i class="bi bi-credit-card"></i> Оформить заказ
            </button>
        </div>
    `;

  document.getElementById("cart-content").innerHTML = html;

  // Add event listeners
  document.querySelectorAll(".remove-from-cart").forEach((btn) => {
    btn.addEventListener("click", function () {
      const itemId = this.getAttribute("data-item-id");
      removeFromCart(itemId);
    });
  });

  document.querySelectorAll(".increase-quantity").forEach((btn) => {
    btn.addEventListener("click", function () {
      const itemId = this.getAttribute("data-item-id");
      updateQuantity(itemId, 1);
    });
  });

  document.querySelectorAll(".decrease-quantity").forEach((btn) => {
    btn.addEventListener("click", function () {
      const itemId = this.getAttribute("data-item-id");
      updateQuantity(itemId, -1);
    });
  });

  document
    .getElementById("clear-cart-btn")
    .addEventListener("click", clearCart);
  document.getElementById("checkout-btn").addEventListener("click", checkout);
}

// Remove item from cart
async function removeFromCart(itemId) {
  try {
    const response = await fetch(`${API_BASE_URL}/cart/items/${itemId}`, {
      method: "DELETE",
      headers: getAuthHeader(),
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || "Failed to remove item from cart");
    }

    showSuccess("Товар удален из корзины!");
    await loadCart();
    await getCartCount();
  } catch (error) {
    console.error("Error removing from cart:", error);
    showError("Не удалось удалить товар из корзины: " + error.message);
  }
}

// Update quantity
async function updateQuantity(itemId, change) {
  try {
    // In a real implementation, you would have an API endpoint to update quantity
    // For now, we'll simulate it by removing and re-adding
    showError("Изменение количества пока не реализовано в API");
  } catch (error) {
    console.error("Error updating quantity:", error);
    showError("Не удалось обновить количество: " + error.message);
  }
}

// Clear cart
async function clearCart() {
  if (!confirm("Вы уверены, что хотите очистить корзину?")) {
    return;
  }

  try {
    const response = await fetch(`${API_BASE_URL}/cart`, {
      method: "DELETE",
      headers: getAuthHeader(),
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || "Failed to clear cart");
    }

    showSuccess("Корзина очищена!");
    await loadCart();
    await getCartCount();
  } catch (error) {
    console.error("Error clearing cart:", error);
    showError("Не удалось очистить корзину: " + error.message);
  }
}

// Checkout
function checkout() {
  showError("Функция оформления заказа пока не реализована");
}

// Escape HTML
function escapeHtml(text) {
  const div = document.createElement("div");
  div.textContent = text;
  return div.innerHTML;
}

// Initialize when DOM is loaded
document.addEventListener("DOMContentLoaded", initCartPage);
