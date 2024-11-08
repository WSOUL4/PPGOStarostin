

$('#reg').click(ShowRegistration);
$('#backToLog').click(backToLog);
//$('#log').click(login);reg
function ShowRegistration(){
    let l = document.getElementById("loginDiv");
    let r = document.getElementById("regDiv");
    hide(l);show(r);
    
}
function backToLog(){
    let l = document.getElementById("loginDiv");
    let r = document.getElementById("regDiv");
    hide(r);show(l);
}
function show (elem) {  /* added argument */
    elem.style.display="flex"; /* changed variable to argument */
  }
  function hide (elem) { /* added argument */
    elem.style.display="none";  /* changed variable to argument */
  }


