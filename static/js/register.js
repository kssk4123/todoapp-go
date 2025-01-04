document.getElementById("register-form").addEventListener("submit", async function(e) {
  e.preventDefault();

  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;

  const response = await fetch("/api/register", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username, password }),
  });

  const result = await response.json();
  if (!result.message) {
    alert("この名前では登録できません");
  } else {
    alert(result.message);
  }

  if (response.ok) {
    setTimeout(() => {
      window.location.href = "/home";  // メッセージ後に/homeへリダイレクト
    }, 1000);  // 少し待ってからリダイレクト（1秒など）
  }
});

