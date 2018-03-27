function more_activate(e) {
    document.getElementById('btn_more').removeEventListener('click', more_activate);
    document.getElementById('btn_less').addEventListener('click', more_deactivate);
    document.getElementById('sp_more').className = "";
    document.getElementById('btn_more').className += " hide";
}

function more_deactivate(e) {
    document.getElementById('btn_more').addEventListener('click', more_activate);
    document.getElementById('btn_less').removeEventListener('click', more_deactivate);
    document.getElementById('sp_more').className += " hide";
    document.getElementById('btn_more').className += "tc-maincolor-text";
}

var btn_more = document.getElementById('btn_more');
if ( btn_more != null) {
    btn_more.addEventListener('click', more_activate);
}

document.getElementById("gen_accept_cb").addEventListener("click", gen_accept_cb);

function gen_accept_cb(e) {
    var cb = e.target;
    var d = document.getElementById("gen_go_div");
    var t = document.getElementById("tpl_gen_go_div");

    if (cb.checked === true) {
        var dataJson = "";
        var url = "/app/"+ThingId+"/generate/accept";
        $.ajax({
            type: "POST",
            url: url,
            async: true,
            success: function (data) {
                dataJson = dataJson + data
            },
            error: function (data) {
            },
            complete: function () {
                var obj = JSON.parse(dataJson);

                d.innerHTML = t.innerHTML;

                var tkn = document.getElementById("gen_go_token");
                tkn.value = obj.Token;

                console.log(tkn);

                cb.disabled = " disabled"
            }
        })


    } else {
        d.innerHTML = "";
    }
}


$(document).ready(function(){
    $('.tabs').tabs();
});

$.encoder.init();

var ThingId = document.getElementById('span_thing_id').innerText;

var dataJson = "";
var url = "/app/"+ThingId+"/generate/data";
$.ajax({
    type: "GET",
    url: url,
    async: true,
    success: function (data) {
        dataJson = dataJson + data
    },
    error: function (data) {
    },
    complete: function() {
        if (dataJson.length === 0) {
        } else {
            var obj = JSON.parse(dataJson);
            console.log(obj);
            var div;

            var tdesc = "";
            if (obj.wtd.description !== null) {
                tdesc = "<p>"+$.encoder.encodeForHTML(obj.wtd.description)+"</p>";
            }

            div = document.getElementById("gen_review_thing");
            div.innerHTML = "<div class=\"col s12\"><p>Name: <strong>"+
                $.encoder.encodeForHTML(obj.wtd.name)+
                "</strong></p>"+
                tdesc+
                "<a class=\"tc-maincolor-text\" href=\"/app/"+ThingId+"\"><i class=\"material-icons\">edit</i></a></div>"
            "<a class=\"tc-maincolor-text\" href=\"/app/"+ThingId+"\"><i class=\"material-icons\">edit</i></a></div>"

            var cf = obj.target;
            if ( cf != null && cf !== undefined ) {
                var depsStr = "<b>Dependencies/Library usages</b>";
                if ( cf.Dependencies != null) {
                    depsStr += "<ul class=\"browser-default\">";
                    for (var i = 0; i < cf.Dependencies.length; i++) {
                        var d = cf.Dependencies[i];
                        depsStr += "<li>"+d.Name+", "+d.Copyright+", "+d.License+"<br/>"+
                            "<a href=\""+d.URL+"\">"+d.URL+"</a>"+
                            "</li>"
                    }
                    depsStr += "</ul>"
                }
                div = document.getElementById("gen_review_framework");
                div.innerHTML = "<div class=\"col s12\"><p>Name: <strong>"+
                    $.encoder.encodeForHTML(cf.ShortDesc)+
                    "</strong></p><p>"+
                    $.encoder.encodeForHTML(cf.Desc)+
                    "</p><p><b>License Information</b><br>"+
                    $.encoder.encodeForHTML(cf.CodeGenInfo)+
                    "</p><p>"+
                    depsStr+
                    "</p><a class=\"tc-maincolor-text\" href=\"/app/"+ThingId+"/framework\"><i class=\"material-icons\">edit</i></a>"+
                    "<a target=\"_tf_module\" class=\"tc-maincolor-text\" href=\"/module/"+$.encoder.encodeForHTML(cf.Id)+"\"><i class=\"material-icons\">description</i></a>"+
                    "</div>"

            }

            var tgt;

            tgt = obj.wtd.properties;
            if (tgt !== null) {
                var f = false;
                var t = "<table class=\"responsive-table\"><thead><tr><th>Name</th><th>Type</th><th>Description</th></thead><tbody>";
                for (var key in tgt) {
                    if (tgt.hasOwnProperty(key)) {
                        p = tgt[key];
                        p.name = key;
                        t += "<tr><td><strong>"+ $.encoder.encodeForHTML(p.name)+"</strong></td>";
                        t += "<td>"+ $.encoder.encodeForHTML(p.type)+"</td>";
                        t += "<td>"+ $.encoder.encodeForHTML(p.description)+"</td>";
                        t += "</tr>";
                        f = true;
                    }
                }
                t += "</tbody></table>";
                div = document.getElementById("gen_review_properties");
                if (f) {
                    div.innerHTML =
                        "<div class=\"col s11\">"+
                        t+ "</div>"+
                        "<div class=\"col s1\"><p><a class=\"tc-maincolor-text\" href=\"/app/"+ThingId+"/properties\"><i class=\"material-icons\">edit</i></a></p></div>"+
                        "";
                } else {
                    div.innerHTML = "<div class=\"col s1\"><p><a class=\"tc-maincolor-text\" href=\"/app/"+ThingId+"/properties\"><i class=\"material-icons\">edit</i></a></p></div>";
                }
            }
            tgt = obj.wtd.actions;
            if (tgt !== null) {
                var f = false;
                var t = "<table class=\"responsive-table\"><thead><tr><th>Name</th><th>Description</th></thead><tbody>";
                for (var key in tgt) {
                    if (tgt.hasOwnProperty(key)) {
                        p = tgt[key];
                        p.name = key;
                        t += "<tr><td><strong>"+ $.encoder.encodeForHTML(p.name)+"</strong></td>";
                        t += "<td>"+ $.encoder.encodeForHTML(p.description)+"</td>";
                        t += "</tr>";
                        f = true
                    }
                }
                t += "</tbody></table>";
                div = document.getElementById("gen_review_actions");
                if (f) {
                    div.innerHTML =
                        "<div class=\"col s11\">"+
                        t+ "</div>"+
                        "<div class=\"col s1\"><p><a class=\"tc-maincolor-text\" href=\"/app/"+ThingId+"/actions\"><i class=\"material-icons\">edit</i></a></p></div>"+
                        "";
                } else {
                    div.innerHTML = "<div class=\"col s1\"><p><a class=\"tc-maincolor-text\" href=\"/app/"+ThingId+"/actions\"><i class=\"material-icons\">edit</i></a></p></div>";
                }
            }
            tgt = obj.wtd.events;
            if (tgt !== null) {
                var f = false;
                var t = "<table class=\"responsive-table\"><thead><tr><th>Name</th><th>Description</th></thead><tbody>";
                for (var key in tgt) {
                    if (tgt.hasOwnProperty(key)) {
                        p = tgt[key];
                        p.name = key;
                        f = true;
                        t += "<tr><td><strong>" + $.encoder.encodeForHTML(p.name) + "</strong></td>";
                        t += "<td>" + $.encoder.encodeForHTML(p.description) + "</td>";
                        t += "</tr>";
                    }
                }
                t += "</tbody></table>";
                div = document.getElementById("gen_review_events");
                if (f === true) {
                    div.innerHTML =
                        "<div class=\"col s11\">"+
                        t+ "</div>"+
                        "<div class=\"col s1\"><p><a class=\"tc-maincolor-text\" href=\"/app/"+ThingId+"/events\"><i class=\"material-icons\">edit</i></a></p></div>"+
                        "";
                } else {
                    div.innerHTML = "<div class=\"col s1\"><p><a class=\"tc-maincolor-text\" href=\"/app/"+ThingId+"/events\"><i class=\"material-icons\">edit</i></a></p></div>";
                }
            }
        }

    }
});

