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
            var obj = JSON.parse(eventsJson)
            console.log(obj)
            if (obj !== null) {
                for (var key in obj) {
                    if (obj.hasOwnProperty(key)) {
                        prop = obj[key]
                        prop.name = key
                        me_list_add_existing(prop)
                    }
                }
            }
            progress.className += " hide";

        },
    });

</script>
{{end}}