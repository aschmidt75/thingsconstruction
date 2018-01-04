{{define `script`}}
<script src="/js/me.js"></script>
<script>
    var progress = document.getElementById("progress");
    var eventsJson = ""
    var url = document.URL;
    url = url.replace(/^(.*\/events).*/, "$1/data")
    $.ajax({
        type: "GET",
        url: url,
        async: true,
        success: function (data) {
            eventsJson = eventsJson + data
        },
        error: function (data) {
            console.log(data);
        },
        complete: function() {
            if (eventsJson.length == 0) {
                // inject empty arr
                eventsJson = '[]'
            }
            var obj = JSON.parse(eventsJson)
            for (var i = 0; i < obj.length; i++) {
                var prop = obj[i]
                me_list_add_existing(prop)
            }
            progress.className += " hide";

        },
    });

</script>
{{end}}