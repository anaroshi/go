"use strict";

$(document).ready(function () {

  $(function() {
    if(!window.EventSource) {
      alert("No EventSource!")
      return
    }

    let $chatLog = $('#chat-log');
    let $chatMsg = $('#chat-msg');

    let isBlank = function(string) {
      return string == null || string.trim() === "";
    };

    let userName = "";
    while(isBlank(userName)) {
      userName = prompt("What's your name?")
      if(!isBlank(userName)) {
        $('#user-name').html('<b>'+userName+'</b>');
      }
    }

    $('#input-form').on('submit', function(e) {
      $.post("/messages", {
          msg: $chatMsg.val(),
          name: userName
      });
      $chatMsg.val("");
      $chatMsg.focus();
      return false; // 다른 페이지로 넘어가는 것을 막음
    });

    let addMsg = function(data) {
      let text = "";
      if (!isBlank(data.name)) {
        text = '<strong>'+ data.name +':</strong>';
      }
      text += data.msg;
      $chatLog.prepend('<div><span>'+ text +'</span></div>');
    };

    let es = new EventSource('/stream')  
    es.onopen = function(e) {
      $.post('users/', {
        name: userName
      });
    }
    es.onmessage = function(e) {
      let msg = JSON.parse(e.data)
      addMsg(msg)
    };

    // 윈도우가 닫히기 전에 호출됨
    window.onbeforeunload = function() {
      $.ajax({        
        url: "/users?userName=" + userName,
        type: "Delete"
      });
      es.close();
    };
  })
});

