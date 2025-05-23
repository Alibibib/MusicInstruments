document.addEventListener("DOMContentLoaded", () => {
    const usersTableBody = document.querySelector("#users-table tbody");
    const editModal = document.getElementById("editModal");
    const closeModalBtn = document.getElementById("closeModal");
    const editRoleForm = document.getElementById("editRoleForm");
    const editUserIdInput = document.getElementById("editUserId");
    const roleSelect = document.getElementById("roleSelect");

    // Функция загрузки всех пользователей с сервера
    async function loadUsers() {
        try {
            const res = await fetch("/users", {
                headers: { "Authorization": "Bearer " + localStorage.getItem("token") }
            });
            if (!res.ok) throw new Error("Ошибка загрузки пользователей");
            const users = await res.json();

            usersTableBody.innerHTML = "";

            users.forEach(user => {
                // Добавляем строку с пользователем
                const tr = document.createElement("tr");
                tr.innerHTML = `
          <td>${user.ID || user.id || user.id}</td>
          <td>${user.Username || user.username}</td>
          <td>${user.Email || user.email}</td>
          <td>${user.Role ? user.Role.Name || user.Role.name : "user"}</td>
          <td>
            <button class="edit-btn" data-id="${user.ID || user.id}" data-role="${user.Role ? user.Role.Name || user.Role.name : "user"}">Редактировать</button>
            <button class="delete-btn" data-id="${user.ID || user.id}">Удалить</button>
          </td>
        `;
                usersTableBody.appendChild(tr);
            });

            // Назначаем обработчики на кнопки
            document.querySelectorAll(".edit-btn").forEach(btn => {
                btn.addEventListener("click", openEditModal);
            });
            document.querySelectorAll(".delete-btn").forEach(btn => {
                btn.addEventListener("click", deleteUser);
            });

        } catch (error) {
            alert(error.message);
        }
    }

    // Открыть модальное окно редактирования
    function openEditModal(event) {
        const btn = event.target;
        const userId = btn.dataset.id;
        const role = btn.dataset.role;

        editUserIdInput.value = userId;
        roleSelect.value = role.toLowerCase();

        editModal.style.display = "block";
    }

    // Закрыть модальное окно
    closeModalBtn.onclick = () => {
        editModal.style.display = "none";
    };
    window.onclick = event => {
        if (event.target === editModal) {
            editModal.style.display = "none";
        }
    };

    // Обработка формы редактирования роли
    editRoleForm.addEventListener("submit", async e => {
        e.preventDefault();
        const userId = editUserIdInput.value;
        const newRole = roleSelect.value;

        // Здесь мы можем отправить PUT-запрос на обновление роли пользователя
        // Предполагается, что бекенд позволяет обновлять роль через PUT /users/:id с передачей role
        try {
            const res = await fetch(`/users/${userId}`, {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": "Bearer " + localStorage.getItem("token"),
                },
                body: JSON.stringify({ role: newRole }),
            });

            if (!res.ok) {
                const err = await res.json();
                throw new Error(err.error || "Ошибка при обновлении роли");
            }

            alert("Роль успешно обновлена");
            editModal.style.display = "none";
            loadUsers();
        } catch (error) {
            alert(error.message);
        }
    });

    // Удаление пользователя
    async function deleteUser(event) {
        const userId = event.target.dataset.id;
        if (!confirm("Вы уверены, что хотите удалить пользователя?")) return;

        try {
            const res = await fetch(`/users/${userId}`, {
                method: "DELETE",
                headers: {
                    "Authorization": "Bearer " + localStorage.getItem("token"),
                },
            });

            if (res.status !== 204) {
                const err = await res.json();
                throw new Error(err.error || "Ошибка при удалении пользователя");
            }

            alert("Пользователь удалён");
            loadUsers();
        } catch (error) {
            alert(error.message);
        }
    }

    // Загрузка данных при старте
    loadUsers();
});
