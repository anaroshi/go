$(document).ready(function () {

  $('.error-page').hide(0);

  $('.login-button , .no-access').click(function(){
    let id = $("#id").val();
    let pwd = $("#pwd").val();
    alert (id+" : "+pwd);    
    // $('.login').slideUp(500);
    // $('.error-page').slideDown(1000);
  });

  $('.try-again').click(function(){    
    $('.error-page').hide(0);
    $('.login').slideDown(1000);
  });

  $('.sign-up').on('click', function () {
    location.href='./join';    
  });

});


