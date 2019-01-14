$(document).ready(function(){
  $("#boletos-add-ticket").click(function(){
    $(".boletos-tickets-a-comprar").append('<div class="form-row align-items-center boletos-ticket-selector">' +
                '<div class="col-auto mb-2">' +
                '  <label class="mr-sm-6" for="inlineFormCustomSelect">Selecciona tu boleto: </label>' +
                '  <select class="custom-select mr-sm-2" id="inlineFormCustomSelect">' +
                '    <option selected>Escoge...</option>' +
                '    <option value="1">1</option>' +
                '    <option value="2">2</option>' +
                '    <option value="3">3</option>' +
                '  </select>' +
                '</div>' +
                '<div class="col-auto my-1">' +
                '<button class="btn btn-secondary btn-sm">Bloquear Boleto</button>' +
                '</div>' +
              '</div>'
    );
  });

  $("#boletos-remove-ticket").click(function(){
    $(".boletos-ticket-selector").last().remove();
  })

});
