// theme.js
 document.addEventListener('DOMContentLoaded', function() {
      const btn = document.getElementById('theme-toggle');
      const body = document.body;
      let dark = false;
      btn.addEventListener('click', function() {
        dark = !dark;
        body.classList.toggle('dark-mode', dark);
        btn.textContent = dark ? '☀️ Tema Claro' : '🌙 Tema Oscuro';
      });
    });