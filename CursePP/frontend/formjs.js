const token = JSON.parse(sessionStorage.getItem("JWT"));
console.log(token); // Здесь ваши данные
if ((token==="")|| (token===null)){
  logout();
}
//const url='http://localhost:8080/data';
//import * as d3 from 'd3';
$.support.cors = true;
//$('.taskInB').click(selectcInside);
var globalCurrentTask=0;
var globalCurrentTaskFull='0';
//$('#MyButton').click(mypost);
$('#add').click(selectByParentIdToFill);
$('#back').click(selectBack);
$('#all').click(selectToFill);
$('#tree').click(selectTree);
$('#logout').click(logout);
$('#entertask').click(enterTask);
$('#closetaskcr').click(closetaskcr);
$('#closetaskch').click(closetaskch);
$('#closetaskaddW').click(closetaskaddW);
$('#Reports').click(showReportsMenu);
$('#closerepots').click(closerepots);
$('#DaysLeft').click(DaysLeft);
$('#DoneFor').click(DoneFor);
function DaysLeft(){
  var url = "http://localhost:8080/report/left";
  let days= document.getElementById("RepDays").value;
  
  var data = { pi: Number(days) };
    
  fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': token.tokenStr,
    },
    body: JSON.stringify(data),
    
  })
  .then(response => response.json())
  .then(data => {DaysLeft2(data);})
  .catch(error => console.error(error));
}

function DaysLeft2(data){
  let par=document.getElementById("ReportsDialog");

  document.querySelectorAll('#RepDialTable').forEach(e => e.remove());
  let table = document.createElement('table');
  table.id="RepDialTable";
  table.className="greyGridTable";
  let header = document.createElement('tr');

  let hId = document.createElement('th');
  hId.innerHTML="ID";
  let hTask = document.createElement('th');
  hTask.innerHTML="Задача";
  let hMark = document.createElement('th');
  hMark.innerHTML="Отметки";
  let hStatus = document.createElement('th');
  hStatus.innerHTML="Статус";
  let hDH = document.createElement('th');
  hDH.innerHTML="Дней дано";
  let hDL = document.createElement('th');
  hDL.innerHTML="Дней осталось";
  header.appendChild(hId);
  header.appendChild(hTask);
  header.appendChild(hMark);
  header.appendChild(hStatus);
  header.appendChild(hDH);
  header.appendChild(hDL);
  table.appendChild(header);
  data.forEach(function(entry) {
    let row = document.createElement('tr');

    let rId = document.createElement('td');
    rId.innerHTML=entry.id;
    let rTask = document.createElement('td');
    rTask.innerHTML=entry.task;
    let rMark = document.createElement('td');
    rMark.innerHTML=entry.mark;
    let rStatus = document.createElement('td');
    rStatus.innerHTML=entry.status;
    let rDH = document.createElement('td');
    rDH.innerHTML=entry.daysHad;
    let rDL = document.createElement('td');
    rDL.innerHTML=entry.daysLeft;
    row.appendChild(rId);
    row.appendChild(rTask);
    row.appendChild(rMark);
    row.appendChild(rStatus);
    row.appendChild(rDH);
    row.appendChild(rDL);
    table.appendChild(row);
  });
  par.appendChild(table);
}
function DoneFor(){
  var url = "http://localhost:8080/report/forLast";
  let days= document.getElementById("RepDays").value;
  
  var data = { pi: Number(days) };
    
  fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': token.tokenStr,
    },
    body: JSON.stringify(data),
    
  })
  .then(response => response.json())
  .then(data => {DoneFor2(data);})
  .catch(error => console.error(error));

}
function DoneFor2(data){
  let par=document.getElementById("ReportsDialog");

  document.querySelectorAll('#RepDialTable').forEach(e => e.remove());
  let table = document.createElement('table');
  table.id="RepDialTable";
  table.className="greyGridTable";
  let header = document.createElement('tr');


  let hId = document.createElement('th');
  hId.innerHTML="ID";
  let hTask = document.createElement('th');
  hTask.innerHTML="Задача";
  let hMark = document.createElement('th');
  hMark.innerHTML="Отметки";
  let hStatus = document.createElement('th');
  hStatus.innerHTML="Статус";

  let hDL = document.createElement('th');
  hDL.innerHTML="Выполнено столько дней назад";
  header.appendChild(hId);
  header.appendChild(hTask);
  header.appendChild(hMark);
  header.appendChild(hStatus);

  header.appendChild(hDL);
  table.appendChild(header);
  data.forEach(function(entry) {
    let row = document.createElement('tr');

    let rId = document.createElement('td');
    rId.innerHTML=entry.id;
    let rTask = document.createElement('td');
    rTask.innerHTML=entry.task;
    let rMark = document.createElement('td');
    rMark.innerHTML=entry.mark;
    let rStatus = document.createElement('td');
    rStatus.innerHTML=entry.status;

    let rDL = document.createElement('td');
    rDL.innerHTML=entry.daysPassed;
    row.appendChild(rId);
    row.appendChild(rTask);
    row.appendChild(rMark);
    row.appendChild(rStatus);

    row.appendChild(rDL);
    table.appendChild(row);
  });
  par.appendChild(table);
}
function showReportsMenu(){
  let dialog = document.getElementById('crDialog');
  hide(dialog);
  dialog = document.getElementById('chDialog');
  hide(dialog);
  dialog = document.getElementById('addWDialog');
  hide(dialog);
    dialog = document.getElementById('ReportsDialog');
    show2(dialog);

}
function logout(){

  sessionStorage.removeItem("JWT");
  window.open("./login.html", "_self");
}
function maxDepth(arr) {
  let maxValue = Number.MIN_VALUE;
  for (let i = 0; i < arr.length; i++) {
      if (arr[i].depth > maxValue) {
          maxValue = arr[i].depth;
      }
  }
  return maxValue;
}
function show (elem) {  /* added argument */
  elem.style.display="block"; /* changed variable to argument */
}
function show2 (elem) {  /* added argument */
  elem.style.display="flex"; /* changed variable to argument */
}
function hide (elem) { /* added argument */
  elem.style.display="none";  /* changed variable to argument */
}
function makeTree(tree){
  //let s=document.getElementById("nav");
  let s=document.getElementById("nav");
  show(s);
  //s.visibility="visible";
  var indTek=globalCurrentTask;
if (s) {
  
  let width = window.innerWidth;
let height = window.innerHeight;

  s.setAttribute('width',width*0.8);
  s.setAttribute('height',height*0.8);
  s.setAttribute('viewBox',`0 0 ${width*0.8} ${height*0.8}`);


  window.addEventListener('resize', () => {
    let width = window.innerWidth;
    let height = window.innerHeight;
    
      s.setAttribute('width',width*0.8);
      s.setAttribute('height',height*0.8);
      s.setAttribute('viewBox',`0 0 ${width*0.8} ${height*0.8}`);
  });


};
document.addEventListener("click", function(event) {
  // Проверяем, был ли клик вне элемента nav
  if (!s.contains(event.target)) {
      // Код, который нужно выполнить при клике вне блока nav
      hide(s);
      s.innerHTML = ` <defs>
    <linearGradient id="blackred">
      <stop offset="0%" stop-color="black" />
      <stop offset="100%" stop-color="black" />
    </linearGradient>
  </defs>`;
      //console.log("Клик вне блока nav");
      // Здесь можно, например, скрыть меню или выполнить другие действия
  }
});
  //let levels=maxDepth(tree);
  let nodes=[];
  let links =[];
  tree.forEach(function(el,ind){
    if (ind==0){
      nodes.push({id: el.parentId, name: `Задача #${el.parentId}`, status: el.status});
      nodes.push({id: el.id, name: `Задача #${el.id}`, status: el.status});
      links .push({source: el.parentId, target: el.id});
    } else{
      nodes.push({id: el.id, name: `Задача #${el.id}`, status: el.status});
    links .push({source: el.parentId, target: el.id});
    };
    
   
  });
  var data ={nodes: nodes,links: links};
  //console.log(data)
  
  var svg = d3.select("svg"),
      width = +svg.attr("width"),
      height = +svg.attr("height");
  
  var simulation = d3.forceSimulation()
      .force("link", d3.forceLink().id(function(d) { return d.id; }))
      .force("charge", d3.forceManyBody())
      .force("center", d3.forceCenter(width / 2, height / 2));
  
  var link = svg.selectAll(".link")
      .data(data.links)
      .enter().append("line")
      .attr("class", "link");
      
  
  var node = svg.selectAll(".node")
      .data(data.nodes)
      .enter().append("circle")
      .attr("class", "node")
      .attr("r", 5)
      .attr("fill", function(d) {
        // Установите цвет в зависимости от значения d.status
        
         if (d.id == indTek ) {
          
          return 'grey'; // Цвет для неактивного статуса
        }else if (d.id === 0) {
          return 'fuchsia'; // Цвет для активного статуса
        }
         else if (d.status === 'done') {
          return 'green'; // Цвет для активного статуса
        } else if (d.status === 'in process') {
          return 'yellow'; // Цвет для неактивного статуса
        }else if ((d.status === 'inactive' ) || (d.status === 'waiting' )) {
          return 'red'; // Цвет для неактивного статуса
        }
         else {
          return 'black'; // Цвет по умолчанию
        }
      })
      .on("mouseover", function(d) {
        //const x = event.clientX; // Получение координаты X
        //const y = event.clientY; // Получение координаты Y
        //console.log("Координаты мыши (X, Y):", x, y);
        //console.log("Наведен курсор на элемент: ", d);
        // настройте свой код обработки событий mouseover здесь
        let dialogTask=document.getElementById("dialogTask");
        //let cont=document.getElementById("tasksContainer");
       // var sizes = document.getElementById('tasksContainer').offsetHeight;
        dialogTask.innerHTML="Задача #" + d.id
        dialogTask.style.top=`${d.y+30}px`
        dialogTask.style.left=`${d.x+15}px`
        //dialogTask.style.objectPosition= d.x.toString()+" "+d.y.toString()
        show(dialogTask);
      })
      .on("mouseout", function(d) {
        //console.log("Курсор ушел с элемента: ", d);
        // настройте свой код обработки событий mouseout здесь
        hide(dialogTask);
      })
      .on("click", function(d) {
        hide(dialogTask);
       selectcInside(d.id);
       //let s=document.getElementById("nav"); // Получаем SVG элемент

        // Устанавливаем значение свойства innerHTML равным пустой строке
        s.innerHTML = ` <defs>
    <linearGradient id="blackred">
      <stop offset="0%" stop-color="black" />
      <stop offset="100%" stop-color="black" />
    </linearGradient>
  </defs>`;
       // s.style.visibility="hidden";
       hide(s);
      });
  
  simulation
      .nodes(data.nodes)
      .on("tick", ticked);
  
  simulation.force("link")
      .links(data.links);
  
  function ticked() {
    link
        .attr("x1", function(d) { return d.source.x; })
        .attr("y1", function(d) { return d.source.y; })
        .attr("x2", function(d) { return d.target.x; })
        .attr("y2", function(d) { return d.target.y; });
  
    node
        .attr("cx", function(d) { return d.x; })
        .attr("cy", function(d) { return d.y; });
  }









}


function selectTree(){
    let ind=0;
    var url = "http://localhost:8080/children";
    
    
    fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': token.tokenStr,
      },
      
    })
    .then(response => response.json())
    .then(data => {makeTree(data);})
    .catch(error => console.error(error));
}
function selectChildrenTree(id){
   let ind=globalCurrentTask;
   var url = "http://localhost:8080/children";
   var data = { id: ind };
   
   fetch(url, {
     method: 'POST',
     headers: {
       'Content-Type': 'application/json',
       'Authorization': token.tokenStr,
     },
     body: JSON.stringify(data),
   })
   .then(response => response.json())
   .then(data => {console.log(data)})
   .catch(error => console.error(error));
}
function selectParentTree(id){
    let ind=globalCurrentTask;
    var url = "http://localhost:8080/parents";
    var data = { id: ind };
    
    fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': token.tokenStr,
      },
      body: JSON.stringify(data),
    })
    .then(response => response.json())
    .then(data => {console.log(data)})
    .catch(error => console.error(error));
}
function selectcInside(ind){
    if ( isNaN(ind) ) {ind=$(this).parent().attr('id'); };
  
    console.log(ind);
    ind=Number(ind);
    if ( isNaN(ind)) {
        
    }else{
        globalCurrentTask=ind;
        globalCurrentTaskFull+='/'+ind.toString()
    };
    
    


    //
    document.getElementById('path').innerHTML=globalCurrentTaskFull;

    var url = "http://localhost:8080/selectByParent";
    var data = { pi: ind };
    
    fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': token.tokenStr,
      },
      body: JSON.stringify(data),
    })
    .then(response => response.json())
    .then(data => {add(data)})
    .catch(error => console.error(error));
}
function selectBack(){
    let massPathStr=globalCurrentTaskFull.split('/');
    var massPathInt = massPathStr.map(function (x) {
        return parseInt(x, 10);
      });
    let ind=Number(massPathInt[massPathInt.length - 2]);
    console.log(ind);
    globalCurrentTask=ind;
    globalCurrentTaskFull='';
    massPathInt.forEach(function(elem, idx, array) {
        if (idx === array.length - 2){globalCurrentTaskFull+=elem.toString();}
        else if(idx === array.length - 1){}
        else{globalCurrentTaskFull+=elem.toString()+'/';}
    
});
document.getElementById('path').innerHTML=globalCurrentTaskFull;

    var url = "http://localhost:8080/selectByParent";
    var data = { pi: ind };
    
    fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': token.tokenStr,
      },
      body: JSON.stringify(data),
    })
    .then(response => response.json())
    .then(data => {add(data)})
    .catch(error => console.error(error));
}
function selectToFill(){
  url= 'http://localhost:8080/selectAll';
  fetch(url, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': token.tokenStr,
    },
    body: JSON.stringify(data),
  })
  .then(response => response.json())
  .then(data => {add(data)})
  .catch(error => console.error(error));
  
    globalCurrentTask=0;
    globalCurrentTaskFull='0';
    document.getElementById('path').innerHTML='';
}
function selectByParentIdToFill(){
    globalCurrentTask=0;
    globalCurrentTaskFull='0';





    document.getElementById('path').innerHTML=globalCurrentTaskFull;

var url = "http://localhost:8080/selectByParent";
var data = { pi: 0 };

fetch(url, {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
    'Authorization': token.tokenStr,
  },
  body: JSON.stringify(data),
})
.then(response => response.json())
.then(data => {add(data)})
.catch(error => console.error(error));


 /*
    $.ajax({
        type: 'post',
        url: 'http://localhost:8080/selectByParent',
        
        
        
        ossDomain: true,
        
        contentType: 'application/json; charset=utf-8',
        dataType: 'application/json',
        data: {name: "Value1", occupation: "Value2"},
        success: function(responseData, textStatus, jqXHR) {
            //alert('GET finished');
            //let view_data = responseData.view_data;
            //console.log(view_data); //Shows the correct piece of information
            add(responseData); // Pass data to a function
           // return responseData;
                },
        error: function (responseData, textStatus, errorThrown) {
            //alert( textStatus);
        }
    });*/
}


function add(data){
  
    document.querySelectorAll('.taskInstance').forEach(e => e.remove()); //очистка старых значений
    data.forEach(function(entry) {
       // console.log(entry);
        let div = document.createElement('div');
div.className = "taskInstance";

let table = document.createElement('table');
table.className = "taskInside";
//Dates
let headerDate = document.createElement('tr');
let Dbeg = document.createElement('th');
let Dend = document.createElement('th');
Dbeg.innerHTML=('Дата нач.');
Dend.innerHTML=('Дата кон.');
headerDate.appendChild(Dbeg);
headerDate.appendChild(Dend);
table.appendChild(headerDate);
//Dates Data
let dataDate = document.createElement('tr');
Dbeg = document.createElement('td');
Dend = document.createElement('td');
Dbeg.innerHTML=entry.TaskF.dateBeg.slice(0,10);
Dend.innerHTML=entry.TaskF.dateEnd.slice(0,10);
dataDate.appendChild(Dbeg);
dataDate.appendChild(Dend);
table.appendChild(dataDate);
//Workers
let headerResp = document.createElement('tr');
let resp = document.createElement('th');
resp.setAttribute('colspan','2');
resp.innerHTML=('Ответственные');
headerResp.appendChild(resp);
table.appendChild(headerResp);
//Workers Data
let dataResp = document.createElement('tr');
if (entry.WorkersF.length!=0) {

//ЦИКЛ СЮДА НУЖЕН БУДЕТ
resp = document.createElement('td');
resp.setAttribute('colspan','2');

entry.WorkersF.forEach(function(worker){
    resp.innerHTML+=worker.login+ "<br>";
});
} else {
    resp = document.createElement('td');
resp.setAttribute('colspan','2');
resp.innerHTML='Не назначенно';
}



dataResp.appendChild(resp);
table.appendChild(dataResp);
//
//Task
let headerTask = document.createElement('tr');
let task = document.createElement('th');
task.setAttribute('colspan','2');
task.innerHTML=('Задача');
headerTask.appendChild(task);
table.appendChild(headerTask);
//Task Data
let dataTask = document.createElement('tr');
//ЦИКЛ СЮДА НУЖЕН БУДЕТ
task = document.createElement('td');
task.setAttribute('colspan','2');
task.innerHTML=entry.TaskF.task;
dataTask.appendChild(task);
table.appendChild(dataTask);
//
//Marks
let headerMarks = document.createElement('tr');
let marks = document.createElement('th');
marks.setAttribute('colspan','2');
marks.innerHTML=('Метки');
headerMarks.appendChild(marks);
table.appendChild(headerMarks);
//Marks Data
//Task Data
let dataMarks = document.createElement('tr');
//ЦИКЛ СЮДА НУЖЕН БУДЕТ?
marks = document.createElement('td');



marks.innerHTML=entry.TaskF.mark;
dataMarks.appendChild(marks);
marks = document.createElement('td');
marks.innerHTML=entry.TaskF.status;
dataMarks.appendChild(marks);
table.appendChild(dataMarks);
//
//End
let taskid= document.createElement('p');
taskid.innerHTML=entry.TaskF.id;
taskid.className = "taskNumP";
let button = document.createElement('button');
button.className = "taskInB";
button.textContent = 'Перейти ввнутрь';
let button2 = document.createElement('button');
button2.className = "taskChange";
button2.textContent = 'Изменить';
let button3 = document.createElement('button');
button3.className = "taskDelete";
button3.textContent = 'Удалить';
let button4 = document.createElement('button');
button4.className = "taskAddWorker";
button4.textContent = 'Добавить человека в задачу';
div.appendChild(taskid);
div.appendChild(button);
div.appendChild(button2);
div.appendChild(button3);
div.appendChild(button4);
div.appendChild(table);
div.id=entry.TaskF.id;


$('.tasksContainer').append(div);
    });



    $('.taskInB').click(selectcInside);
    $('.taskChange').click(ChangeTask);
    $('.taskDelete').click(DeleteTask);
    $('.taskAddWorker').click(taskAddWorker);
}
function taskAddWorker(){
  document.getElementById("addWWarn").innerHTML="";
  ind=$(this).parent().attr('id');
  dialog = document.getElementById('addWDialog');
  show(dialog);
  let dialog= document.getElementById("chDialog");
  hide(dialog);
  let dialogcr= document.getElementById("crDialog");
  hide(dialogcr);

  id = document.getElementById('addWId');
  id.innerHTML=ind
}

function DeleteTask(){
  ind=$(this).parent().attr('id');

  //let parent = document.getElementById(ind);
var url = "http://localhost:8080/deleteTask";
var data = {pi: Number(ind)};

fetch(url, {
method: 'DELETE',
headers: {
  'Content-Type': 'application/json',
  'Authorization': token.tokenStr,
},
body: JSON.stringify(data),
})
.then(response => response.json())
.then(data => {selectcInside(globalCurrentTask);})
.catch(error => console.error(error));
}


function ChangeTask(){
  document.getElementById("crWarn").innerHTML="";
  document.getElementById("chWarn").innerHTML="";
  let dialog= document.getElementById("chDialog");
  show(dialog);
  let dialogcr= document.getElementById("crDialog");
  hide(dialogcr);
  let dialaddw=document.getElementById("addWDialog");
  hide(dialaddw);
  let dialogr=document.getElementById("ReportsDialog");
  hide(dialogr);
  ind=$(this).parent().attr('id'); 
  let parent = document.getElementById(ind);
  let children = parent.querySelectorAll('td');
//2 строчку с ответственным по другому будет
//console.log(children[2].innerHTML)
var idText = document.getElementById("chId");
idText.innerHTML=ind;
  var dateB = document.getElementById("chdn");
  dateB.value=children[0].innerHTML;
   //dateB.value=
  var dateE = document.getElementById("chde");
  dateE.value=children[1].innerHTML;
  var Parent = document.getElementById("chparent");
  Parent.value= globalCurrentTask;
  var Task = document.getElementById("cht");
  Task.value=children[3].innerHTML;
  var Marks = document.getElementById("chm");
  Marks.value=children[4].innerHTML;
  var Status = document.getElementById("chs");
  Status.value=children[5].innerHTML;
}
//Скорость перетаскивания зависит от скорости мыши, а как найти значение не знаю
function enterTask(){
  document.getElementById("crWarn").innerHTML="";
  document.getElementById("chWarn").innerHTML="";
  let dialog = document.getElementById('crDialog');
  show(dialog);
  let dialogch = document.getElementById('chDialog');
  hide(dialogch);

  let dialaddw=document.getElementById("addWDialog");
  hide(dialaddw);

  let dialogr = document.getElementById('ReportsDialog');
    hide(dialogr);
  /*let isDragging = false;
 // let offsetX, offsetY;
  
  //dialog.addEventListener('mousedown', startDrag);
  //dialog.addEventListener('mouseup', stopDrag);

  function startDrag(e) {
    isDragging = true;
    offsetX = e.offsetX;
    offsetY = e.offsetY;
}

function stopDrag() {
    isDragging = false;
}

document.addEventListener('mousemove', drag);

function drag(e) {
    if (isDragging) {
        dialog.style.top = (e.clientY - offsetY*1) + 'px';
        dialog.style.left = (e.clientX - offsetX*1) + 'px';
    }
}*/
}


function closetaskcr(){
  let dialog = document.getElementById('crDialog');
  hide(dialog);

}

function closetaskch(){

  dialog = document.getElementById('chDialog');
  hide(dialog);
}
function closetaskaddW(){
  dialog = document.getElementById('addWDialog');
  hide(dialog);
}
function closerepots(){
  dialog = document.getElementById('ReportsDialog');
  hide(dialog);
}



