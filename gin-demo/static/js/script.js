document.addEventListener('DOMContentLoaded', function() {
    // 获取DOM元素
    const getUsersBtn = document.getElementById('get-users');
    const getUserBtn = document.getElementById('get-user');
    const createUserBtn = document.getElementById('create-user');
    const userIdInput = document.getElementById('user-id');
    const newIdInput = document.getElementById('new-id');
    const newUsernameInput = document.getElementById('new-username');
    const newEmailInput = document.getElementById('new-email');
    const resultElement = document.getElementById('result');

    // 显示结果的函数
    function displayResult(data) {
        resultElement.textContent = JSON.stringify(data, null, 2);
    }

    // 处理API错误的函数
    function handleError(error) {
        resultElement.textContent = `错误: ${error.message}`;
        console.error('API错误:', error);
    }

    // 获取所有用户
    getUsersBtn.addEventListener('click', async function() {
        try {
            const response = await fetch('/api/users');
            const data = await response.json();
            displayResult(data);
        } catch (error) {
            handleError(error);
        }
    });

    // 获取指定ID的用户
    getUserBtn.addEventListener('click', async function() {
        const userId = userIdInput.value.trim();
        if (!userId) {
            alert('请输入用户ID');
            return;
        }

        try {
            const response = await fetch(`/api/users/${userId}`);
            const data = await response.json();
            displayResult(data);
        } catch (error) {
            handleError(error);
        }
    });

    // 创建新用户
    createUserBtn.addEventListener('click', async function() {
        const id = newIdInput.value.trim();
        const username = newUsernameInput.value.trim();
        const email = newEmailInput.value.trim();

        if (!id || !username || !email) {
            alert('请填写所有字段');
            return;
        }

        const newUser = {
            ID: id,
            Username: username,
            Email: email
        };

        try {
            const response = await fetch('/api/users', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(newUser)
            });
            const data = await response.json();
            displayResult(data);

            // 清空输入框
            newIdInput.value = '';
            newUsernameInput.value = '';
            newEmailInput.value = '';
        } catch (error) {
            handleError(error);
        }
    });
});