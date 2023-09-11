document.getElementById('projectForm').addEventListener('submit', function(event) {
    event.preventDefault();

    const formData = new FormData(event.target);
    const data = {
        projectName: formData.get('projectName'),
        framework: formData.get('framework'),
        database: formData.get('database'),
        auth: formData.get('auth') === "on",
        cache: formData.get('cache') === "on",
        projectPath: formData.get('projectPath'),
        structName: formData.get('structName')
    };


    fetch('http://localhost:8080/v1/generate', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
        .then(response => {
            if (response.status === 201) {
                return response.json();
            } else {
                throw new Error('Erro ao gerar projeto.');
            }
        })
        .then(data => {
            console.log(data);
            const notification = document.getElementById('notification');
            notification.className = 'alert alert-success';
            notification.textContent = data.msg;
            notification.style.display = 'block';
        })
        .catch(error => {
            console.error('Erro na requisição:', error);
            const notification = document.getElementById('notification');
            notification.className = 'alert alert-danger';
            notification.textContent = 'Erro na requisição: ' + error.message;
            notification.style.display = 'block';
        });
});