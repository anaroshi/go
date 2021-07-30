"use strict";

$(document).ready(function () {
  
  if (!window.WebSocket) {
    alert("No WebSocket!");
    return
  }
  
  let connect = "";
  let ws = "";

  let addmessage = function(data) {
    $chatlog.prepend("<div><span>"+data+"</span></div>");
  }

  // window.location.host 현재 윈도의 호스트를 가져옴
  connect = function() {
    ws = new WebSocket("ws://"+window.location.host+"/ws");
    ws.onopen = function(e) {
      console.log("onopen",arguments);
    }
    ws.onclose = function(e) {
      console.log("onclose",arguments);
    }
    ws.onmessage = function(e) {
      addmessage(e.data);
    }
  }

  connect();

  var $chatlog = $('#chat-log');
  var $chatmsg = $('#chat-msg');

  var isBlank = function(string) {
    return string == null || string.trim() === "";
  };

  var userName;
  while(isBlank(userName)) {
    userName = prompt("What's your name?");
    if (!isBlank(userName)) {
      $('#user-name').html('<b>'+userName+'</b>');
    }
  }

  $('#input-form').on('submit', function(e) {
    if (ws.readyState === ws.OPEN) {
      ws.send(JSON.stringify({
        type: "msg",
        data: $chatmsg.val()
      }))
    }
    $chatmsg.val("");
    $chatmsg.focus();
    return false;
  });  

});