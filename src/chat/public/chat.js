$(document).ready(function () {
  if(!window.EventSource) {
    alert("No EventSource!")
    return
  }

  let $chatlog = $('#chat-log')
  let $chatmsg = $('#chat-msg')

  let isBlank = function(string) {
    return string == null || string.trim() === "";    
  }

  let username;
  while(isBlank(username)) {
    username = prompt("What's your name?")
    if (!isBlank(username)) {
      $('#user-name').html('<b>'+username+'</b>')
    }
  }

  $('#input-form').on('submit', function () {
    $.post("/messages", {
      msg: $chatmsg.val(),
      name: username    
    });
    $chatmsg.val('');
    $chatmsg.focus();
    return false;
  });

  let addMessage = function(data) {
    let text = "";
    if (!isBlank(data.name)) {
      text = '<strong>'+ data.name+': </strong>';      
    }
    text += data.msg;
    $chatlog.prepend('<div><span>'+text+'</span></div>');
  };

  // addMessage({
  //   name : 'test',
  //   msg:'hello'
  // })

  // addMessage({    
  //   msg:'hello2'
  // })

  let es = new EventSource('/stream');
  es.onopen =function(e) {
    $.post("users/", {
      name: username
    });
  }

  es.onmessage = function(e) {
    let msg = JSON.parse(e.data);
    addMessage(msg);
  }
  
  window.onbeforeunload = function() {
    $.ajax({
      type: "DELETE",
      url: "/users?username="+username
    });
    es.close();
  }

});