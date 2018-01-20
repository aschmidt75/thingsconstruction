{{define `script`}}
<script>

    $(document).ready(function() {
        $('select').material_select();
    });

    $("#select-archive").on('change', function() {
        var v = $(this)[0].value;

        var l = document.getElementById("select-archive-link");
        if ( v === "") {
            l.className = "hide";
            l.href = "";

        } else {
            l.className = "";
            l.href = "/app/{{ .ThingId }}/result/asset-archive/"+v;
        }
    });

    $('#url_copy_btn').on('click', function(e) {
        var i = document.getElementById("url_copy_input");
        i.value = document.getElementById("url_copy_a").textContent;
        i.select();
        document.execCommand("copy");
        //M.toast({html: 'Copied.'});
    })
</script>
{{end}}