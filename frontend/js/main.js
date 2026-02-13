// Main application functions for Go Shop

// Show error message
function showError(message) {
  const errorDiv = document.createElement("div");
  errorDiv.className = "error-message";
  errorDiv.textContent = message;

  const main = document.querySelector("main");
  if (main) {
    main.prepend(errorDiv);
    setTimeout(() => errorDiv.remove(), 5000);
  }
}

// Show success message
function showSuccess(message) {
  const successDiv = document.createElement("div");
  successDiv.className = "success-message";
  successDiv.textContent = message;

  const main = document.querySelector("main");
  if (main) {
    main.prepend(successDiv);
    setTimeout(() => successDiv.remove(), 5000);
  }
}

// Show loading spinner
function showLoading(element) {
  const spinner = document.createElement("span");
  spinner.className = "loading-spinner ms-2";
  spinner.setAttribute("role", "status");
  spinner.setAttribute("aria-hidden", "true");

  if (element) {
    element.appendChild(spinner);
    element.disabled = true;
  }

  return spinner;
}

// Hide loading spinner
function hideLoading(element, spinner) {
  if (spinner && spinner.parentNode) {
    spinner.parentNode.removeChild(spinner);
  }
  if (element) {
    element.disabled = false;
  }
}

// Format price
function formatPrice(price) {
  return new Intl.NumberFormat("ru-RU", {
    style: "currency",
    currency: "RUB",
  }).format(price);
}

// Get cart count
async function getCartCount() {
  try {
    const response = await fetch(`${API_BASE_URL}/cart`, {
      headers: getAuthHeader(),
    });

    if (response.ok) {
      const cart = await response.json();
      const count = cart.items
        ? cart.items.reduce((sum, item) => sum + item.quantity, 0)
        : 0;

      const cartCountElement = document.getElementById("cart-count");
      if (cartCountElement) {
        cartCountElement.textContent = count;
      }

      return count;
    }
  } catch (error) {
    console.error("Error getting cart count:", error);
  }
  return 0;
}

// Initialize main functions
document.addEventListener("DOMContentLoaded", function () {
  // Update cart count if user is authenticated
  if (isAuthenticated()) {
    getCartCount();
  }
});
