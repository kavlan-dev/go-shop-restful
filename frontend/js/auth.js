// Authentication functions for Go Shop

const API_BASE_URL = "/api";

// Check if user is authenticated
function isAuthenticated() {
  return localStorage.getItem("token") !== null;
}

// Get current user info from JWT
function getCurrentUser() {
  const token = localStorage.getItem("token");
  if (!token) return null;

  try {
    const payload = JSON.parse(atob(token.split(".")[1]));
    return {
      userId: payload.user_id,
      username: payload.username,
      role: payload.role,
      email: payload.email,
    };
  } catch (e) {
    console.error("Error parsing token:", e);
    return null;
  }
}

// Update UI based on authentication status
function updateAuthUI() {
  const user = getCurrentUser();
  const authSection = document.getElementById("auth-section");
  const userSection = document.getElementById("user-section");

  if (user) {
    if (authSection) authSection.classList.add("d-none");
    if (userSection) userSection.classList.remove("d-none");

    const usernameSpan = document.getElementById("current-username");
    if (usernameSpan) {
      usernameSpan.textContent = `${user.username} (${user.role})`;
    }
  } else {
    if (authSection) authSection.classList.remove("d-none");
    if (userSection) userSection.classList.add("d-none");
  }
}

// Login function
async function login(username, password) {
  try {
    const response = await fetch(`${API_BASE_URL}/auth/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username, password }),
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || "Login failed");
    }

    const data = await response.json();
    localStorage.setItem("token", data.token);

    updateAuthUI();
    return data;
  } catch (error) {
    console.error("Login error:", error);
    throw error;
  }
}

// Register function
async function register(username, password, email) {
  try {
    const response = await fetch(`${API_BASE_URL}/auth/register`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username, password, email }),
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || "Registration failed");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Registration error:", error);
    throw error;
  }
}

// Logout function
function logout() {
  localStorage.removeItem("token");
  updateAuthUI();
  window.location.href = "/index.html";
}

// Get auth header for API requests
function getAuthHeader() {
  const token = localStorage.getItem("token");
  return token ? { Authorization: `Bearer ${token}` } : {};
}

// Check if user is admin
function isAdmin() {
  const user = getCurrentUser();
  return user && user.role === "admin";
}

// Initialize auth on page load
document.addEventListener("DOMContentLoaded", function () {
  updateAuthUI();

  // Add logout event listener
  const logoutBtn = document.getElementById("logout-btn");
  if (logoutBtn) {
    logoutBtn.addEventListener("click", logout);
  }
});
