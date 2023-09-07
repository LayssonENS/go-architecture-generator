document.getElementById('projectForm').addEventListener('submit', function(event) {
    event.preventDefault();

    const formData = new FormData(event.target);
    const data = {
        projectName: formData.get('projectName'),
        framework: formData.get('framework'),
        database: formData.get('database'),
        auth: formData.get('auth') === "on",  // Retorna true se estiver marcado, caso contrário, false
        cache: formData.get('cache') === "on",  // Retorna true se estiver marcado, caso contrário, false
        jsonData: formData.get('jsonData')
    };

    fetch('http://localhost:8080/v1/generate', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            // Aqui você pode processar a resposta, por exemplo, mostrar uma mensagem de sucesso.
        })
        .catch(error => {
            console.error('Erro na requisição:', error);
        });
});
