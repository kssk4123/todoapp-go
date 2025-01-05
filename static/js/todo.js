document.getElementById("addtodo").addEventListener("submit", async function (e) {
    e.preventDefault(); // フォームのデフォルト動作を防止

    const taskTitle = document.getElementById("task").value;

    if (!taskTitle) {
        alert("タスクを入力してください。");
        return;
    }

    // 送信ボタンを無効化して、二重送信を防止
    const submitButton = document.querySelector("#addtodo button[type='submit']");
    submitButton.disabled = true;

    try {
        // Fetch APIでPOSTリクエストを送信
        const response = await fetch("/api/todos", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ title: taskTitle }),
        });

        if (!response.ok) {
            const errorData = await response.json();
            alert(`エラー: ${errorData.error}`);
            return;
        }

        await response.json();
        document.getElementById("task").value = ""; // 入力欄をクリア

        // タスク追加後にリストを再読み込み
        fetchTodos();
    } catch (error) {
        console.error("エラー:", error);
        alert("タスクの追加に失敗しました。");
    } finally {
        // 送信ボタンを再度有効化
        submitButton.disabled = false;
    }
});

// Todosを非同期で取得する関数
async function fetchTodos() {
    try {
        const response = await fetch(`/api/todos`);
        
        if (!response.ok) {
            throw new Error('Failed to fetch todos');
        }

        const data = await response.json();

        // Todosを表示する関数を呼び出し
        displayTodos(data.todos);
    } catch (error) {
        console.error('Error fetching todos:', error);
    }
}

// Todosを画面に表示する関数
function displayTodos(todos) {
    const todosContainer = document.getElementById('todos');
    todosContainer.innerHTML = ''; // 既存のTodosをクリア
    const todoFragment = document.createDocumentFragment();

    if (todos && todos.length > 0) {
        todos.forEach(todo => {
            const div = document.createElement('div');
            div.setAttribute("id", `todo-${todo.id}`);
            div.classList.add("todo-item");
            div.innerHTML = `
              <input type="checkbox" ${todo.is_completed ? 'checked' : ''} data-id="${todo.id}" class="checkbox" />
              <span>${todo.title}</span>
              <button onclick="deleteTodo(${todo.id})" class="delete-button">削除</button>
            `;
            todoFragment.appendChild(div);
        });
        todosContainer.appendChild(todoFragment);
    } else {
        todosContainer.innerHTML = '<p>No todos found.</p>';
    }
}

async function deleteTodo(todoId) {
  try {
    const response = await fetch(`/api/todos/${todoId}` , {
      method: `DELETE`,
    })

    if (!response.ok) {
      const errorData = await response.json();
      alert(`エラー: ${errorData.error}`);
      return;
    }

    fetchTodos();
  } catch (error) {
    console.error(`エラー:`, error);
    alert("タスクの削除に失敗しました");
  }
}

// チェックボックスが変更された時に呼ばれる関数
async function updateTodoCompletion(todoId, isCompleted) {
    try {
        const response = await fetch(`/api/todos/${todoId}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ is_completed: isCompleted }),
        });

        if (!response.ok) {
            const errorData = await response.json();
            alert(`エラー: ${errorData.error}`);
            return;
        }

        // Todoリストを再取得して表示
        fetchTodos();
    } catch (error) {
        console.error('エラー:', error);
        alert("タスクの更新に失敗しました。");
    }
}

// チェックボックスの変更を監視するイベントリスナー
document.getElementById('todos').addEventListener('change', (event) => {
    if (event.target.type === 'checkbox') {
        const todoId = event.target.dataset.id;
        const isCompleted = event.target.checked;

        // チェックボックスが変更された場合にTodoの完了状態を更新
        updateTodoCompletion(todoId, isCompleted);
    }
});

document.getElementById('delete-user').addEventListener('click', (event) => {
  if (confirm("本当にこのユーザーを削除しますか？")) {
    // ユーザーが「OK」を押した場合にAPIリクエストを送信

    fetch("/api/delete-user", {
      method: "DELETE",  // 削除リクエスト
      headers: {
        "Content-Type": "application/json",
      },
      // 必要ならばリクエストボディに追加データを送ることも可能
      // body: JSON.stringify({ user_id: 123 }) など
    })
      .then(response => {
        if (response.ok) {
          // 削除成功した場合、ユーザーに通知してログインページにリダイレクト
          alert("ユーザーは削除されました。ログイン画面にリダイレクトします。");
          window.location.href = "/login"; // ログインページにリダイレクト
        } else {
          // 削除に失敗した場合
          alert("ユーザー削除に失敗しました。");
        }
      })
      .catch(error => {
        // ネットワークエラーなどが発生した場合
        console.error("Error deleting user:", error);
        alert("ユーザー削除中にエラーが発生しました。");
      });
  }
});

// ページの読み込み後に呼び出す関数
window.onload = function() {
    fetchTodos();
};

