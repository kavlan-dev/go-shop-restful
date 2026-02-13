// Products management for Go Shop

let currentPage = 1;
let limit = 12;
let totalProducts = 0;
let categories = [];

// Initialize products page
async function initProductsPage() {
  // Check if user is admin and show admin actions
  if (isAdmin()) {
    document.getElementById("admin-actions").classList.remove("d-none");
    document
      .getElementById("add-product-btn")
      .addEventListener("click", showAddProductModal);
  }

  // Load products
  await loadProducts();

  // Set up event listeners
  document
    .getElementById("search-input")
    .addEventListener("input", debounce(loadProducts, 500));
  document
    .getElementById("category-filter")
    .addEventListener("change", loadProducts);
  document.getElementById("sort-by").addEventListener("change", loadProducts);
  document
    .getElementById("save-product-btn")
    .addEventListener("click", addProduct);
}

// Load products with filters
async function loadProducts() {
  const search = document.getElementById("search-input").value;
  const category = document.getElementById("category-filter").value;
  const sort = document.getElementById("sort-by").value;

  try {
    // Show loading
    document.getElementById("loading").classList.remove("d-none");
    document.getElementById("products-container").classList.add("d-none");

    // Build query parameters
    const params = new URLSearchParams({
      limit: limit,
      offset: (currentPage - 1) * limit,
    });

    if (search) params.append("search", search);
    if (category) params.append("category", category);
    if (sort) {
      const [field, order] = sort.split("-");
      params.append("sort_by", field);
      params.append("sort_order", order);
    }

    // Fetch products
    const response = await fetch(
      `${API_BASE_URL}/products?${params.toString()}`,
      {
        headers: getAuthHeader(),
      },
    );

    if (!response.ok) {
      throw new Error("Failed to load products");
    }

    const responseData = await response.json();

    // Handle both array and object responses
    let products = [];
    let total = 0;

    if (Array.isArray(responseData)) {
      // Direct array response
      products = responseData;
      total = responseData.length;
    } else if (responseData.products) {
      // Object with products field
      products = responseData.products;
      total = responseData.total || products.length;
    } else {
      // Unexpected format
      throw new Error("Неожиданный формат ответа от сервера");
    }

    // Extract categories
    if (categories.length === 0 && products.length > 0) {
      categories = [...new Set(products.map((p) => p.category))];
      populateCategoryFilter();
    }

    // Update total products
    totalProducts = total;

    // Render products
    renderProducts(products);
    renderPagination();
  } catch (error) {
    console.error("Error loading products:", error);
    let errorMessage = "Не удалось загрузить продукты";
    if (error.message) {
      errorMessage += ": " + error.message;
    }
    if (error.message.includes("Failed to fetch")) {
      errorMessage =
        "Не удалось подключиться к серверу. Пожалуйста, проверьте ваше интернет-соединение.";
    }
    showError(errorMessage);
  } finally {
    document.getElementById("loading").classList.add("d-none");
    document.getElementById("products-container").classList.remove("d-none");
  }
}

// Render products
function renderProducts(products) {
  const container = document.getElementById("products-container");
  container.innerHTML = "";

  if (products.length === 0) {
    container.innerHTML = `
            <div class="col-12 text-center">
                <div class="alert alert-info">
                    <h4 class="alert-heading">Продукты не найдены</h4>
                    <p>Попробуйте изменить фильтры или критерии поиска.</p>
                </div>
            </div>
        `;
    return;
  }

  products.forEach((product) => {
    const productCard = document.createElement("div");
    productCard.className = "col-md-4 col-lg-3 mb-4";
    productCard.innerHTML = `
            <div class="card product-card h-100">
                <div class="card-body d-flex flex-column">
                    <h5 class="card-title">${escapeHtml(product.title)}</h5>
                    <p class="card-text flex-grow-1">${escapeHtml(product.description || "Нет описания")}</p>
                    <div class="mb-2">
                        <span class="badge bg-info">${escapeHtml(product.category)}</span>
                        <span class="badge bg-success ms-2">В наличии: ${product.stock}</span>
                    </div>
                    <div class="d-flex justify-content-between align-items-center mt-3">
                        <h5 class="mb-0">${formatPrice(product.price)}</h5>
                        <button class="btn btn-sm btn-primary add-to-cart-btn" data-product-id="${product.ID}">
                            <i class="bi bi-cart-plus"></i> В корзину
                        </button>
                    </div>
                    ${
                      isAdmin()
                        ? `
                        <div class="mt-2 d-flex gap-2">
                            <button class="btn btn-sm btn-outline-primary edit-product-btn" data-product-id="${product.ID}" data-bs-toggle="modal" data-bs-target="#editProductModal">
                                <i class="bi bi-pencil"></i> Редактировать
                            </button>
                            <button class="btn btn-sm btn-outline-danger delete-product-btn" data-product-id="${product.ID}">
                                <i class="bi bi-trash"></i> Удалить
                            </button>
                        </div>
                    `
                        : ""
                    }
                </div>
            </div>
        `;

    container.appendChild(productCard);
  });

  // Add event listeners for buttons
  document.querySelectorAll(".add-to-cart-btn").forEach((btn) => {
    btn.addEventListener("click", function () {
      const productId = this.getAttribute("data-product-id");
      addToCart(productId);
    });
  });

  document.querySelectorAll(".delete-product-btn").forEach((btn) => {
    btn.addEventListener("click", function () {
      const productId = this.getAttribute("data-product-id");
      deleteProduct(productId);
    });
  });

  document.querySelectorAll(".edit-product-btn").forEach((btn) => {
    btn.addEventListener("click", function () {
      const productId = this.getAttribute("data-product-id");
      loadProductForEditing(productId);
    });
  });
}

// Load product data for editing
async function loadProductForEditing(productId) {
  try {
    const response = await fetch(`${API_BASE_URL}/products/${productId}`, {
      headers: getAuthHeader(),
    });

    if (!response.ok) {
      throw new Error("Failed to load product data");
    }

    const product = await response.json();

    // Fill edit form
    document.getElementById("edit-product-id").value = product.ID;
    document.getElementById("edit-product-title").value = product.title;
    document.getElementById("edit-product-description").value =
      product.description || "";
    document.getElementById("edit-product-price").value = product.price;
    document.getElementById("edit-product-category").value = product.category;
    document.getElementById("edit-product-stock").value = product.stock;
  } catch (error) {
    console.error("Error loading product for editing:", error);
    showError(
      "Не удалось загрузить данные продукта для редактирования: " +
        error.message,
    );
  }
}

// Update product
async function updateProduct() {
  const productId = document.getElementById("edit-product-id").value;
  const title = document.getElementById("edit-product-title").value;
  const description = document.getElementById("edit-product-description").value;
  const price = document.getElementById("edit-product-price").value;
  const category = document.getElementById("edit-product-category").value;
  const stock = document.getElementById("edit-product-stock").value;

  if (!title || !description || !price || !category || !stock) {
    showError("Пожалуйста, заполните все поля");
    return;
  }

  const updateBtn = document.getElementById("update-product-btn");
  const spinner = showLoading(updateBtn);

  try {
    const response = await fetch(
      `${API_BASE_URL}/admin/products/${productId}`,
      {
        method: "PUT",
        headers: {
          ...getAuthHeader(),
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          title,
          description,
          price: parseFloat(price),
          category,
          stock: parseInt(stock),
        }),
      },
    );

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || "Failed to update product");
    }

    const modal = bootstrap.Modal.getInstance(
      document.getElementById("editProductModal"),
    );
    modal.hide();

    showSuccess("Продукт успешно обновлен!");
    await loadProducts();
  } catch (error) {
    console.error("Error updating product:", error);
    showError("Не удалось обновить продукт: " + error.message);
  } finally {
    hideLoading(updateBtn, spinner);
  }
}

// Add event listener for update button
document.addEventListener("DOMContentLoaded", function () {
  const updateBtn = document.getElementById("update-product-btn");
  if (updateBtn) {
    updateBtn.addEventListener("click", updateProduct);
  }
});

// Add product to cart
async function addToCart(productId) {
  if (!isAuthenticated()) {
    showError("Пожалуйста, войдите в систему, чтобы добавить товар в корзину");
    return;
  }

  try {
    const response = await fetch(`${API_BASE_URL}/cart/${productId}`, {
      method: "POST",
      headers: getAuthHeader(),
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || "Failed to add to cart");
    }

    showSuccess("Товар добавлен в корзину!");
    await getCartCount();
  } catch (error) {
    console.error("Error adding to cart:", error);
    showError("Не удалось добавить товар в корзину: " + error.message);
  }
}

// Delete product (admin only)
async function deleteProduct(productId) {
  if (!confirm("Вы уверены, что хотите удалить этот продукт?")) {
    return;
  }

  try {
    const response = await fetch(
      `${API_BASE_URL}/admin/products/${productId}`,
      {
        method: "DELETE",
        headers: getAuthHeader(),
      },
    );

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || "Failed to delete product");
    }

    showSuccess("Продукт успешно удален!");
    await loadProducts();
  } catch (error) {
    console.error("Error deleting product:", error);
    showError("Не удалось удалить продукт: " + error.message);
  }
}

// Show add product modal
function showAddProductModal() {
  const modal = new bootstrap.Modal(document.getElementById("addProductModal"));
  modal.show();
}

// Add new product
async function addProduct() {
  const title = document.getElementById("product-title").value;
  const description = document.getElementById("product-description").value;
  const price = document.getElementById("product-price").value;
  const category = document.getElementById("product-category").value;
  const stock = document.getElementById("product-stock").value;

  if (!title || !description || !price || !category || !stock) {
    showError("Пожалуйста, заполните все поля");
    return;
  }

  const saveBtn = document.getElementById("save-product-btn");
  const spinner = showLoading(saveBtn);

  try {
    const response = await fetch(`${API_BASE_URL}/admin/products`, {
      method: "POST",
      headers: {
        ...getAuthHeader(),
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        title,
        description,
        price: parseFloat(price),
        category,
        stock: parseInt(stock),
      }),
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || "Failed to add product");
    }

    const modal = bootstrap.Modal.getInstance(
      document.getElementById("addProductModal"),
    );
    modal.hide();

    // Clear form
    document.getElementById("add-product-form").reset();

    showSuccess("Продукт успешно добавлен!");
    await loadProducts();
  } catch (error) {
    console.error("Error adding product:", error);
    showError("Не удалось добавить продукт: " + error.message);
  } finally {
    hideLoading(saveBtn, spinner);
  }
}

// Populate category filter
function populateCategoryFilter() {
  const filter = document.getElementById("category-filter");
  categories.forEach((category) => {
    const option = document.createElement("option");
    option.value = category;
    option.textContent = category;
    filter.appendChild(option);
  });
}

// Render pagination
function renderPagination() {
  const totalPages = Math.ceil(totalProducts / limit);
  const pagination = document.getElementById("pagination");

  if (totalPages <= 1) {
    pagination.innerHTML = "";
    return;
  }

  let html = '<nav><ul class="pagination">';

  // Previous button
  if (currentPage > 1) {
    html += `<li class="page-item"><button class="page-link" data-page="${currentPage - 1}">Предыдущая</button></li>`;
  } else {
    html +=
      '<li class="page-item disabled"><span class="page-link">Предыдущая</span></li>';
  }

  // Page numbers
  const startPage = Math.max(1, currentPage - 2);
  const endPage = Math.min(totalPages, currentPage + 2);

  if (startPage > 1) {
    html += `<li class="page-item"><button class="page-link" data-page="1">1</button></li>`;
    if (startPage > 2) {
      html +=
        '<li class="page-item disabled"><span class="page-link">...</span></li>';
    }
  }

  for (let i = startPage; i <= endPage; i++) {
    if (i === currentPage) {
      html += `<li class="page-item active"><span class="page-link">${i}</span></li>`;
    } else {
      html += `<li class="page-item"><button class="page-link" data-page="${i}">${i}</button></li>`;
    }
  }

  if (endPage < totalPages) {
    if (endPage < totalPages - 1) {
      html +=
        '<li class="page-item disabled"><span class="page-link">...</span></li>';
    }
    html += `<li class="page-item"><button class="page-link" data-page="${totalPages}">${totalProducts}</button></li>`;
  }

  // Next button
  if (currentPage < totalPages) {
    html += `<li class="page-item"><button class="page-link" data-page="${currentPage + 1}">Следующая</button></li>`;
  } else {
    html +=
      '<li class="page-item disabled"><span class="page-link">Следующая</span></li>';
  }

  html += "</ul></nav>";
  pagination.innerHTML = html;

  // Add event listeners
  document.querySelectorAll(".page-link[data-page]").forEach((btn) => {
    btn.addEventListener("click", function () {
      currentPage = parseInt(this.getAttribute("data-page"));
      loadProducts();
    });
  });
}

// Debounce function
function debounce(func, wait) {
  let timeout;
  return function () {
    const context = this,
      args = arguments;
    clearTimeout(timeout);
    timeout = setTimeout(() => {
      timeout = null;
      func.apply(context, args);
    }, wait);
  };
}

// Escape HTML
function escapeHtml(text) {
  const div = document.createElement("div");
  div.textContent = text;
  return div.innerHTML;
}

// Initialize when DOM is loaded
document.addEventListener("DOMContentLoaded", initProductsPage);
