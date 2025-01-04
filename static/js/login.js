document.getElementById("login-form").addEventListener("submit", async function(e) {
  e.preventDefault();

  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;

  const response = await fetch("/api/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username, password }),
  });

  await response.json();
  //const result = await response.json();
  //alert(result.message);

  if (response.ok) {
    window.location.href = "/home";  // メッセージ後に/homeへリダイレクト
  }
});
