$(document).ready(function(){

  var selectedSeatId = "";

  $("#inlineFormCustomSelect").change(function() {
    selectedSeatId = $(this).children(":selected").attr("id").toString();
  });

  var first_time = false;
  $("#inlineFormCustomSelect").on('click', function() {
    if (!first_time) {
      first_time = true;
      var eventId = $(".hidden-event-id").text()
      $.get("/api/seats/" + eventId, function(data, err){
        if (err == "success"){
          console.log(data);
          var seats = data["data"];
          for (var i = 0; i < seats.length; i++){
            seat = seats[i];
            console.log(seat);
            $("#inlineFormCustomSelect").append("" +
              "<option class='boletos-asiento-seleccionado' id="+ seat.id + "> " + seat.number + "|" +
              seat.zoneName + "|" + seat.zonePrice + "</option>"
            );
          }
        }
      });
    }
  });

  $("#boletos-comprar-boton").on('click', function() {
    var eventId = $(".hidden-event-id").text();
    var userId = 1;
    $.ajax({
      type: "POST",
      url: "/api/purchase",
      data: {
        "user-ID": userId.toString(),
        "event-ID": eventId.toString(),
        "seat-ID": selectedSeatId.toString()
      },
      success: function(data){
        window.location.replace("/profile");
      },
      error: function(XMLHttpRequest, textStatus, errorThrown){
        if (XMLHttpRequest.status == 404){
          alert("El asiento que intentas comprar ya ha sido comprado! Recarga la pagina!");
          setTimeout(function(){ location.reload(); }, 3000);
        } else {
          alert("Algo salio mal! Recarga la pagina!");
          setTimeout(function(){ location.reload(); }, 3000);
        }
      },
      dataType: "json"
    });
  });

})
