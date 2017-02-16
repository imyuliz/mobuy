/*加载年份*/
function years(obj, Cyear) {
    var len = 134; // select长度,即option数量
    var selObj = document.getElementById(obj);
    var selIndex = len - 1;
    var newOpt; // OPTION对象

    // 循环添加OPION元素到年份select对象中
    for (i = 1; i <= len; i++) {
        if (selObj.options.length != len) { // 如果目标对象下拉框升度不等于设定的长度则执行以下代码
            newOpt = window.document.createElement("OPTION"); // 新建一个OPTION对象
            newOpt.text = Cyear - len + i; // 设置OPTION对象的内容
            newOpt.value = Cyear - len + i; // 设置OPTION对象的值
            selObj.options.add(newOpt, i); // 添加到目标对象中
        }

    }
}

function months(){
    var month = document.getElementById("month");
    month.length = 0;  
    for (i = 1; i < 13; i++) {  
        month.options.add(new Option(i, i));  
    }
    
}


function change_date(){  
   // var birthday = document.birthday;  
    var year  = document.getElementById("year");  
    var month = document.getElementById("month");  
    var day   = document.getElementById("date");  
    vYear  = parseInt(year.value);  
    vMonth = parseInt(month.value);  
    date.length=0;  
      
    //根据年月获取天数  
    var max = (new Date(vYear,vMonth, 0)).getDate();  
    for (var i=1; i <= max; i++) {  
        date.options.add(new Option(i, i));  
    }  
}
