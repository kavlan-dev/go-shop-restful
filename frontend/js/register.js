// Register page functionality

document.addEventListener("DOMContentLoaded", function () {
  const registerForm = document.getElementById("register-form");

  if (registerForm) {
    registerForm.addEventListener("submit", async function (e) {
      e.preventDefault();

      const username = document.getElementById("username").value.trim();
      const email = document.getElementById("email").value.trim();
      const password = document.getElementById("password").value.trim();
      const confirmPassword = document
        .getElementById("confirm-password")
        .value.trim();
      const registerBtn = document.getElementById("register-btn");

      if (!username || !email || !password || !confirmPassword) {
        showError("Пожалуйста, заполните все поля");
        return;
      }

      // Validate username length
      if (username.length < 3) {
        showError("Имя пользователя должно содержать не менее 3 символов");
        return;
      }

      // Validate email format
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      if (!emailRegex.test(email)) {
        showError("Пожалуйста, введите корректный email");
        return;
      }

      // Validate password length
      if (password.length < 6) {
        showError("Пароль должен содержать не менее 6 символов");
        return;
      }

      if (password !== confirmPassword) {
        showError("Пароли не совпадают");
        return;
      }

      const spinner = showLoading(registerBtn);

      try {
        await register(username, password, email);
        showSuccess(
          "Регистрация прошла успешно! Теперь вы можете войти в систему.",
        );

        // Redirect to login page
        setTimeout(() => {
          window.location.href = "/login.html";
        }, 2000);
      } catch (error) {
        showError("Ошибка регистрации: " + error.message);
      } finally {
        hideLoading(registerBtn, spinner);
      }
    });
  }
});
