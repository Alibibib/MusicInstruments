const API_URL = 'http://localhost:8080'; // или адрес твоего бэкенда

async function registerUser(data) {
    try {
        localStorage.removeItem('token'); // очистка старого токена
        const response = await fetch(`${API_URL}/register`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data),
        });
        const result = await response.json();
        if (response.ok) {
            return { success: true };
        } else {
            return { success: false, error: result.error || 'Ошибка регистрации' };
        }
    } catch (e) {
        return { success: false, error: e.message };
    }
}

async function loginUser(data) {
    try {
        const response = await fetch(`${API_URL}/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data),
        });
        const result = await response.json();
        if (response.ok) {
            return { success: true, token: result.token };
        } else {
            return { success: false, error: result.error || 'Ошибка входа' };
        }
    } catch (e) {
        return { success: false, error: e.message };
    }
}

// Обработчик формы логина
document.getElementById('loginForm')?.addEventListener('submit', async (e) => {
    e.preventDefault();

    const email = e.target.email.value;
    const password = e.target.password.value;

    const result = await loginUser({ email, password });

    if (result.success) {
        localStorage.setItem('token', result.token);
        window.location.href = '/'; // переход на главную страницу
    } else {
        alert(result.error);
    }
});


async function renderHeader() {
    const token = localStorage.getItem('token');
    if (!token) return;

    try {
        const res = await fetch(`${API_URL}/me`, {
            headers: { Authorization: `Bearer ${token}` }
        });
        const data = await res.json();

        document.getElementById('guest-links').style.display = 'none';
        document.getElementById('user-links').style.display = 'block';

        if (data.role === 'staff') {
            document.getElementById('staff-links').style.display = 'block';
        }

        if (data.role === 'admin') {
            document.getElementById('staff-links').style.display = 'block';
            document.getElementById('admin-links').style.display = 'block';
        }
    } catch (err) {
        console.error('Ошибка при получении роли:', err);
    }
}

window.addEventListener('DOMContentLoaded', renderHeader);



document.getElementById('logout-btn')?.addEventListener('click', () => {
    localStorage.removeItem('token');
    location.href = '/login';
});

document.addEventListener('DOMContentLoaded', () => {
    document.querySelectorAll('.add-to-cart-btn').forEach(button => {
        button.addEventListener('click', async () => {
            const id = button.dataset.id;
            const token = localStorage.getItem('token');

            if (!token) {
                alert('Пожалуйста, войдите в систему чтобы добавить товар в корзину.');
                return;
            }


        });
    });
});


document.addEventListener('DOMContentLoaded', () => {
    const searchBtn = document.getElementById('search-btn');
    const searchInput = document.getElementById('search-input');

    if (searchBtn && searchInput) {
        searchBtn.addEventListener('click', () => {
            const query = searchInput.value.trim().toLowerCase();

            document.querySelectorAll('.card').forEach(card => {
                const title = card.querySelector('.card-title')?.textContent.toLowerCase() || '';
                const description = card.querySelector('.card-text')?.textContent.toLowerCase() || '';

                if (title.includes(query) || description.includes(query)) {
                    card.parentElement.style.display = 'block';
                } else {
                    card.parentElement.style.display = 'none';
                }
            });
        });

        // Поиск по Enter
        searchInput.addEventListener('keypress', (e) => {
            if (e.key === 'Enter') {
                searchBtn.click();
            }
        });
    }
});


async function loadCategories() {
    try {
        const res = await fetch(`${API_URL}/categories`);
        const categories = await res.json();
        const select = document.getElementById('category-select');

        categories.forEach(category => {
            const option = document.createElement('option');
            option.value = category;
            option.textContent = category;
            select.appendChild(option);
        });
    } catch (err) {
        console.error("Ошибка при загрузке категорий:", err);
    }
}

document.getElementById('category-select')?.addEventListener('change', function () {
    const selectedCategory = this.value;
    const params = new URLSearchParams(window.location.search);
    if (selectedCategory) {
        params.set('category', selectedCategory);
    } else {
        params.delete('category');
    }
    window.location.search = params.toString(); // обновляет URL и страницу
});

window.addEventListener('DOMContentLoaded', () => {
    renderHeader();
    loadCategories(); // ← вот это важно
});
