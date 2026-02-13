// Login page functionality

document.addEventListener("DOMContentLoaded", function () {
  const loginForm = document.getElementById("login-form");

  if (loginForm) {
    loginForm.addEventListener("submit", async function (e) {
      e.preventDefault();

      const username = document.getElementById("username").value.trim();
      const password = document.getElementById("password").value.trim();
      const loginBtn = document.getElementById("login-btn");

      if (!username || !password) {
        showError("Пожалуйста, заполните все поля");
        return;
      }

      const spinner = showLoading(loginBtn);

      try {
        await login(username, password);
        showSuccess("Вы успешно вошли в систему!");

        // Redirect to previous page or home
        const redirectUrl =
          localStorage.getItem("redirectAfterLogin") || "/index.html";
        localStorage.removeItem("redirectAfterLogin");
        window.location.href = redirectUrl;
      } catch (error) {
        showError("Ошибка входа: " + error.message);
      } finally {
        hideLoading(loginBtn, spinner);
      }
    });
  }
});
