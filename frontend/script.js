const API = "http://localhost:8080";
let editingTaskId = null;

// ----------------- Utilities -----------------
function getToken() {
  return localStorage.getItem("token");
}

function showToast(msg) {
  const t = document.getElementById("toast");
  t.innerText = msg;
  t.style.display = "block";
  setTimeout(() => t.style.display = "none", 3000);
}

function logout() {
  localStorage.clear();
  window.location = "login.html";
}

// ----------------- Load Tasks -----------------
async function loadTasks() {
  try {
    const res = await fetch(`${API}/tasks`, {
      headers: { "Authorization": "Bearer " + getToken() }
    });

    if (!res.ok) throw new Error("Failed");

    const tasks = await res.json();
    renderTasks(tasks);
  } catch (err) {
    showToast("Failed to load tasks");
    console.error(err);
  }
}

// ----------------- Render Tasks -----------------
function renderTasks(tasks) {
  const container = document.getElementById("tasks");
  container.innerHTML = "";

  let completed = 0;

  tasks.forEach(task => {
    if (task.status === "completed") completed++;

    const card = document.createElement("div");
    card.className = "task-card";

    card.innerHTML = `
      <div class="task-header">
        <div class="task-title">${task.title}</div>
        <span class="status-badge ${task.status}">
          ${task.status.toUpperCase()}
        </span>
      </div>
      <div class="task-desc">${task.description}</div>
      <div class="task-actions">
        <button class="status-btn ${task.status}">
          ${task.status === "pending" ? "Mark Completed" : "Mark Pending"}
        </button>
        <button class="edit-btn">Edit</button>
        <button class="delete-btn">Delete</button>
      </div>
    `;

    container.appendChild(card);

    // ---- STATUS BUTTON (ONLY STATUS CHANGE) ----
    card.querySelector(".status-btn")
      .addEventListener("click", () => toggleStatus(task.id, task.status));

    // ---- EDIT BUTTON (ONLY TITLE + DESCRIPTION) ----
    card.querySelector(".edit-btn")
      .addEventListener("click", () =>
        openEdit(task.id, task.title, task.description)
      );

    // ---- DELETE ----
    card.querySelector(".delete-btn")
      .addEventListener("click", () => deleteTask(task.id));
  });

  document.getElementById("totalCount").innerText = tasks.length;
  document.getElementById("completedCount").innerText = completed;
  document.getElementById("pendingCount").innerText = tasks.length - completed;
}

// ----------------- Add Task -----------------
async function createTask() {
  const title = document.getElementById("title").value.trim();
  const description = document.getElementById("desc").value.trim();

  if (!title || !description)
    return showToast("Enter title & description");

  try {
    await fetch(`${API}/tasks`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + getToken()
      },
      body: JSON.stringify({ title, description })
    });

    document.getElementById("title").value = "";
    document.getElementById("desc").value = "";

    showToast("Task added!");
    loadTasks();
  } catch (err) {
    showToast("Failed to add task");
    console.error(err);
  }
}

// ----------------- TOGGLE STATUS ONLY -----------------
async function toggleStatus(id, currentStatus) {
  const newStatus =
    currentStatus === "pending" ? "completed" : "pending";

  try {
    await fetch(`${API}/tasks/${id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + getToken()
      },
      body: JSON.stringify({
        status: newStatus   // ONLY STATUS CHANGES
      })
    });

    showToast(`Task marked ${newStatus.toUpperCase()}`);
    loadTasks();
  } catch (err) {
    showToast("Failed to update status");
    console.error(err);
  }
}

// ----------------- EDIT (ONLY TITLE & DESCRIPTION) -----------------
function openEdit(id, title, description) {
  editingTaskId = id;
  document.getElementById("editTitle").value = title;
  document.getElementById("editDesc").value = description;
  document.getElementById("editModal").style.display = "flex";
}

async function saveEdit() {
  const title = document.getElementById("editTitle").value.trim();
  const description = document.getElementById("editDesc").value.trim();

  if (!title || !description)
    return showToast("Enter title & description");

  try {
    await fetch(`${API}/tasks/${editingTaskId}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + getToken()
      },
      body: JSON.stringify({
        title,
        description   // STATUS NOT SENT HERE
      })
    });

    document.getElementById("editModal").style.display = "none";
    showToast("Task updated!");
    loadTasks();
  } catch (err) {
    showToast("Failed to update task");
    console.error(err);
  }
}

// ----------------- DELETE -----------------
async function deleteTask(id) {
  if (!confirm("Delete this task?")) return;

  try {
    await fetch(`${API}/tasks/${id}`, {
      method: "DELETE",
      headers: { "Authorization": "Bearer " + getToken() }
    });

    showToast("Task deleted!");
    loadTasks();
  } catch (err) {
    showToast("Failed to delete task");
    console.error(err);
  }
}

// ----------------- Close Modal -----------------
window.onclick = function (event) {
  const modal = document.getElementById("editModal");
  if (event.target === modal)
    modal.style.display = "none";
};