$(document).ready(function(){
    $("#getdata").click(function(){
        $.get( "http://127.0.0.1:5000", function(data) {
            let strres = "";
            let res = JSON.parse(data);
            for(var i in res) {
                strres += "id: "+res[i].id.toString()+"<br>Name: "+res[i].name+"<br>Last name: "+res[i].lastname+"<br>Age: "+res[i].age.toString()+"<br><br>";
            }
            $("#getres").html(strres);
        });
    });
    $("#postdata").click(function(){
        $.post("http://127.0.0.1:5000",
            {
                name: ,
                id: ,
                age: ,
                lname:
            },
            function(data,status){
                alert("Data: " + data + "\nStatus: " + status);
            });
    });


});