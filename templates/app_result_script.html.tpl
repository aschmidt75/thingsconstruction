{{define `script`}}
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

        var content = ""

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
{{end}}