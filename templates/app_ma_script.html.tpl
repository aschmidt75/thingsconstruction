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
        }
    });

</script>
<script>
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
</script>

{{end}}