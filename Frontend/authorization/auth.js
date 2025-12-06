const API_URL = 'http://localhost:8080/'

async function login() {
    const userLogin = {
        username: document.getElementById('username').value,
        password: document.getElementById('password').value
    };
    
    if (!username || !password) {
        alert('Заполните все поля');
        return;
    }
    
    try {
        const zapros = await fetch(`${API_URL}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(userLogin)
        });

        const data = await zapros.json();

        if (zapros.ok) {
            document.getElementById('username').value = '';
            document.getElementById('password').value = '';
            console.log('Все отлично')
        }
    } catch (error) {
        console.error('Ошибка:', error);
    }
    
    console.log('Вход', { username, password });
    
    alert(`Попытка входа для пользователя: ${username}`);
}

function goToRegister() {
    window.location.href = 'register.html';
}

