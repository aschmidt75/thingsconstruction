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
            var obj = JSON.parse(actionsJson);
            if (obj !== null) {
                for (var key in obj) {
                    if (obj.hasOwnProperty(key)) {
                        action = obj[key];
                        action.name = key;
                        ma_list_add_existing(action)
                    }
                }
            }
            progress.className += " hide";
        },
    });

</script>
{{end}}