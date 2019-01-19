$(document).ready(function(){
  $(".login-button").click(function(){
    // call endpoint for login and redirect to events
  })

  $(".boletos-logout-button").click(function(){
    console.log("aquui");
    $.get("/logout", function(data, err){
      console.log(data);
      console.log(err);
      if (err == "success"){
        window.location.replace("/events");
      }
    })
  })
})
