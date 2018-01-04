{{define `script`}}
<script src="/js/mp.js"></script>
<script>
    var progress = document.getElementById("progress");
    var propertiesJson = ""
    var url = document.URL;
    url = url.replace(/^(.*\/properties).*/, "$1/data")

    // load properties on startup..
    $.ajax({
        type: "GET",
        url: url,
        async: true,
        success: function (data) {
            propertiesJson = propertiesJson + data
        },
        error: function (data) {
            console.log(data);
        },
        complete: function() {
            if (propertiesJson.length == 0) {
                // inject empty arr
                propertiesJson = '[]'
            }

            var obj = JSON.parse(propertiesJson)
            for (var i = 0; i < obj.length; i++) {
                var prop = obj[i]
                mp_list_add_existing(prop)
            }
            if (progress != undefined) {
                progress.className += " hide";
            }
        },
    });

</script>
{{end}}