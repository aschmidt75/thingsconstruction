{{define `script`}}
<script src="/js/me.js"></script>
<script>
    var propertiesJson = ""
    // load properties on startup..
    $.ajax({
        type: "GET",
        url: document.URL+"/data",
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
            console.log(obj)
            for (var i = 0; i < obj.length; i++) {
                var prop = obj[i]
                me_list_add_existing(prop)
            }
        },
    });

</script>
{{end}}