$(document).ready(function(){
  var first_time = false;
  $("#inlineFormCustomSelect").on('click', function() {
    if (!first_time) {
      first_time = true;
      var eventId = $(".hidden-event-id").text()
      $.get("/api/seats/" + eventId, function(data, err){
        if (err == "success"){
          var seats = data["data"];
          for (var i = 0; i < seats.length; i++){
            seat = seats[i];
            $("#inlineFormCustomSelect").append("" +
              "<option id="+ seat.id + "> " + seat.number + "|" +
              seat.zoneName + "|" + seat.zonePrice + "</option>"
            );
          }
        }
      });
    }
  });
})
