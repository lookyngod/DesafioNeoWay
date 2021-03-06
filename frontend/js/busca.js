$('#valida').on('submit', insereCPFCNPJ);


//FUNÇÃO PARA INSERIR CPF/CNPJ
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
            ).then((result) => {
                if (result.isConfirmed) {
                    location.reload();
                }
            });
            setTimeout(function() {
                location.reload();
            }, 2000);
        } else {
            Swal.fire(
                'Ops...',
                data["error"],
                'warning',
            );
        }
    }).fail(function(data) {
        Swal.fire(
            'Ops...',
            data["responseJSON"]["error"],
            'warning',
        );
    });
}

//JQUERY PARA DELETAR CPFCNPJ

function deleteCPFCNPJ(event) {
    event.preventDefault();

    $.ajax({
        url: "http://localhost:5000/api/cpfcnpj",
        method: "DELETE",
        data: JSON.stringify({
            cpfcnpj: $("#cpfcnpj").val(),
        }, )
    }).done(function(data) {
        if (data == "CPF/CNPJ deletado com sucesso!") {
            Swal.fire(
                'Sucesso!',
                data,
                'success'
            );
            setTimeout(function() {
                location.reload();
            }, 2000);
        } else {
            Swal.fire(
                'Ops...',
                data["error"],
                'warning',
            );
        }
    }).fail(function(data) {
        Swal.fire(
            'Ops...',
            data["responseJSON"]["error"],
            'warning',
        );
    });
}

//JQUERY PARA BUSCAR OS DADOS

function buscaDados() {
    $(document).ready(function() {
        $.ajax({
            url: "http://localhost:5000/api/getall",
            method: "GET",
        }).done(function() {}).fail(function() {
            Swal.fire(
                'Ops...',
                'Falha na busca',
                'error'
            );
        });
    });

}