$(document).ready(function(){
  $(".login-button").click(function(){
    // call endpoint for login and redirect to events
  })

  $(".boletos-logout-button").click(function(){
    $.get("/logout", function(data, err){
      if (err == "success"){
        window.location.replace("/events");
      }
    })
  })

})
