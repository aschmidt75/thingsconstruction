{{define `script`}}
<script src="https://cdnjs.cloudflare.com/ajax/libs/clipboard.js/1.7.1/clipboard.min.js"></script>
<script>
    $(document).ready(function(){
        // the "href" attribute of the modal trigger must specify the modal ID that wants to be triggered
        $('.modal').modal();
    });
    {{ if .Files }}
{{ range .Files }}
    document.getElementById("view_{{ .Permalink }}").addEventListener('click', view_element);
{{ end }}
{{ end }}

    function view_element(e) {
        $('#view_modal').modal('open');

        var permaLink = e.target.getAttribute("linkid");
        var target = document.getElementById("view_modal_content");
        var viewUrl = "/app/{{.ThingId}}/result/assetview/"+permaLink;

        var content = "";

        $.ajax({
            type: "GET",
            url: viewUrl,
            async: true,
            success: function (data) {
                content = content + data
            },
            error: function (data) {
                console.log(data);
                target.innerHTML = "<p>An error occured while fetching the document. Please try again later.</p>";
            },
            complete: function() {
                target.innerHTML = content;
                (function(){
                    new Clipboard('#copycode');
                })();
            },
        });

    }

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