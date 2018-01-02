{{define `script`}}
<script src="/js/ma.js"></script>
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
            alert("An error occured while querying session data")
        },
        complete: function() {
            var obj = JSON.parse(propertiesJson)
            for (var i = 0; i < obj.length; i++) {
                var prop = obj[i]
                ma_list_add_existing(prop)
            }
        },
    });

</script>
{{end}}