const API_URL = 'http://localhost:8080/'

async function loginUser() {
    const userLogin = {
        login: document.getElementById('login').value,
        password: document.getElementById('password').value
    };
    
    if (!login || !password) {
        alert('Заполните все поля');
        return;
    }
    
    try {
        const zapros = await fetch(`${API_URL}auth`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(userLogin)
        });

        const data = await zapros.json();

        if (zapros.ok) {
            document.getElementById('login').value = '';
            document.getElementById('password').value = '';
            console.log('Все отлично')
        }
    } catch (error) {
        console.error('Ошибка:', error);
    }
    
    console.log('Вход', { login, password });
    
    alert(`Попытка входа для пользователя: ${login}`);
}

function goToRegister() {
    window.location.href = 'register.html';
}

