{{define `script`}}
<script>


    $(document).ready(function(){
        $('.tabs').tabs();
    });

    $.encoder.init();

    var dataJson = "";
    var url = "/app/{{.ThingId}}/generate/data";
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

                div = document.getElementById("gen_review_thing");
                div.innerHTML = "<div class=\"col s12\"><p>Name: <strong>"+
                        $.encoder.encodeForHTML(obj.wtd.Name)+
                        "</strong></p><p>"+
                        $.encoder.encodeForHTML(obj.wtd.Description)+
                        "</p><a class=\"deep-orange-text\" href=\"/app/{{.ThingId}}\"><i class=\"material-icons\">edit</i></a></div>"

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
                            "</p><a class=\"deep-orange-text\" href=\"/app/{{.ThingId}}/framework\"><i class=\"material-icons\">edit</i></a></div>"

                }

                var tgt;

                tgt = obj.wtd.Properties;
                if (tgt != null && tgt.length > 0) {
                    var t = "<table class=\"responsive-table\"><thead><tr><th>Name</th><th>Type</th><th>Description</th></thead><tbody>";
                    for (var i = 0; i < tgt.length; i++) {
                        var p = tgt[i];
                        t += "<tr><td><strong>"+ $.encoder.encodeForHTML(p.Name)+"</strong></td>";
                        t += "<td>"+ $.encoder.encodeForHTML(p.Type)+"</td>";
                        t += "<td>"+ $.encoder.encodeForHTML(p.Description)+"</td>";
                        t += "</tr>";
                    }
                    t += "</tbody></table>";
                    div = document.getElementById("gen_review_properties");
                    div.innerHTML =
                            "<div class=\"col s11\">"+
                            t+ "</div>"+
                            "<div class=\"col s1\"><p><a class=\"deep-orange-text\" href=\"/app/{{.ThingId}}/properties\"><i class=\"material-icons\">edit</i></a></p></div>"+
                            "";
                } else {
                    div = document.getElementById("gen_review_properties");
                    div.innerHTML = "<div class=\"col s1\"><p><a class=\"deep-orange-text\" href=\"/app/{{.ThingId}}/properties\"><i class=\"material-icons\">edit</i></a></p></div>";
                }
                tgt = obj.wtd.Actions;
                if (tgt != null && tgt.length > 0) {
                    var t = "<table class=\"responsive-table\"><thead><tr><th>Name</th><th>Description</th></thead><tbody>";
                    for (var i = 0; i < tgt.length; i++) {
                        var p = tgt[i];
                        t += "<tr><td><strong>"+ $.encoder.encodeForHTML(p.Name)+"</strong></td>";
                        t += "<td>"+ $.encoder.encodeForHTML(p.Description)+"</td>";
                        t += "</tr>";
                    }
                    t += "</tbody></table>";
                    div = document.getElementById("gen_review_actions");
                    div.innerHTML =
                            "<div class=\"col s11\">"+
                            t+ "</div>"+
                            "<div class=\"col s1\"><p><a class=\"deep-orange-text\" href=\"/app/{{.ThingId}}/actions\"><i class=\"material-icons\">edit</i></a></p></div>"+
                            "";
                } else {
                    div = document.getElementById("gen_review_actions");
                    div.innerHTML = "<div class=\"col s1\"><p><a class=\"deep-orange-text\" href=\"/app/{{.ThingId}}/actions\"><i class=\"material-icons\">edit</i></a></p></div>";
                }
                tgt = obj.wtd.Events;
                if (tgt != null && tgt.length > 0) {
                    var t = "<table class=\"responsive-table\"><thead><tr><th>Name</th><th>Description</th></thead><tbody>";
                    for (var i = 0; i < tgt.length; i++) {
                        var p = tgt[i];
                        t += "<tr><td><strong>"+ $.encoder.encodeForHTML(p.Name)+"</strong></td>";
                        t += "<td>"+ $.encoder.encodeForHTML(p.Description)+"</td>";
                        t += "</tr>";
                    }
                    t += "</tbody></table>";
                    div = document.getElementById("gen_review_events");
                    div.innerHTML =
                            "<div class=\"col s11\">"+
                            t+ "</div>"+
                            "<div class=\"col s1\"><p><a class=\"deep-orange-text\" href=\"/app/{{.ThingId}}/events\"><i class=\"material-icons\">edit</i></a></p></div>"+
                            "";
                } else {
                    div = document.getElementById("gen_review_events");
                    div.innerHTML = "<div class=\"col s1\"><p><a class=\"deep-orange-text\" href=\"/app/{{.ThingId}}/events\"><i class=\"material-icons\">edit</i></a></p></div>";
                }
            }

        }
    });

    document.getElementById("gen_accept_cb").addEventListener("click", gen_accept_cb);

    function gen_accept_cb(e) {
        var cb = e.target;
        var d = document.getElementById("gen_go_div");
        var t = document.getElementById("tpl_gen_go_div");

        if (cb.checked === true) {
            var dataJson = "";
            var url = "/app/{{.ThingId}}/generate/accept";
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
                    console.log(obj);

                    d.innerHTML = t.innerHTML;

                    var tkn = document.getElementById("gen_go_token");
                    tkn.value = obj.Token;

                    cb.disabled = " disabled"
                }
            })


        } else {
            d.innerHTML = "";
        }
    }
</script>
{{end}}