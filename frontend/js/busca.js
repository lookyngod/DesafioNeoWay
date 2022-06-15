$('#valida').on('submit', insereCPFCNPJ);

function insereCPFCNPJ(event) {
    event.preventDefault();

    $.ajax({
        url: "http://localhost:5000/api/cpfcnpj",
        method: "POST",
        data: JSON.stringify({
            cpfcnpj: $("#cpfcnpj").val(),
        }, )
    }).done(function(data) {
        if (data == "CPF/CNPJ inserido com sucesso!") {
            Swal.fire(
                'Sucesso!',
                data,
                'success'
            );
            setTimeout(function() {
                location.reload();
            }, 3000);
        } else {
            Swal.fire(
                'Ops...',
                data["error"],
                'warning',
            );
        }
    }).fail(function(data) {
        console.log(data);
        Swal.fire(
            'Ops...',
            data["responseJSON"]["error"],
            'warning',
        );
    });
}


function buscaDados() {

    $(document).ready(function() {
        $.ajax({
            url: "http://localhost:5000/api/getall",
            method: "GET",
        }).done(function() {
            console.log("é gol do rony");
        }).fail(function() {
            console.log("é gol do rony");
            Swal.fire(
                'Ops...',
                'Usuário ou senha incorretos!',
                'error'
            );
        });
    });

}