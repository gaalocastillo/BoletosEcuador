$(document).ready(function(){
  $.get("/api/user_tickets/", function(data, err){
    if (err == "success"){
      var seats = data["data"];
      for (var i = 0; i < seats.length; i++){
        seat = seats[i];
        $("#boletos-tickets-comprados").append("" +
          "<li class='list-group-item'>" +
            "Evento: " + seat.eventName + "<br>" +
            "Numbero de Asiento: " + seat.ticketNumber + "<br>" +
            "Zona: " + seat.zoneName + "<br>" +
            "Precio: " + seat.zonePrice + "<br>" +
          "</li>"
        );
      }
    }
  });
})
