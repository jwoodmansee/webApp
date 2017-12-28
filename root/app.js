$(document).ready(function () {

document.getElementById("click").onclick = function() {go()};

function go() {
    $.get("http://localhost:8080", {params: document.getElementById("params").value}, function(data, status) {

        console.log(status);
    });
}
});
