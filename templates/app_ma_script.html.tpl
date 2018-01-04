{{define `script`}}
<script src="/js/ma.js"></script>
<script>
    var progress = document.getElementById("progress");
    var actionsJson = ""
    var url = document.URL;
    url = url.replace(/^(.*\/actions).*/, "$1/data")

    $.ajax({
        type: "GET",
        url: url,
        async: true,
        success: function (data) {
            actionsJson = actionsJson + data
        },
        error: function (data) {
            console.log(data);
            progress.className += " hide";
        },
        complete: function() {
            if (actionsJson.length == 0) {
                // inject empty arr
                actionsJson = '[]'
            }

            var obj = JSON.parse(actionsJson)
            for (var i = 0; i < obj.length; i++) {
                var prop = obj[i]
                ma_list_add_existing(prop)
            }
            progress.className += " hide";
        },
    });

</script>
{{end}}