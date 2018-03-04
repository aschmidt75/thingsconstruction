{{define `script`}}
<script>
    function Target(id,shortDesc,desc,tags,codegeninfo) {
        this.id = id;
        this.shortDesc = shortDesc;
        this.desc = desc;
        this.tags = tags;
        this.selected = false;
        this.deps = [];
        this.codegeninfo = codegeninfo;
        // TODO: Add version
    }

    function TargetDep(name,url,license,copyright,info) {
        this.name = name;
        this.url = url;
        this.license = license;
        this.copyright = copyright;
        this.info = info;
    }

    var targets = [];
    {{ range .AppGenTargets.Targets }}
    t = new Target('{{.Id}}', '{{.ShortDesc}}', '{{.Desc}}', [{{ range .Tags }}'{{.}}', {{end}} ], '{{ .CodeGenInfo}}');
    {{ range .Dependencies }}
    d = new TargetDep('{{.Name}}','{{.URL}}','{{.License}}','{{.Copyright}}','{{.Info}}');
    t.deps.push(d);
    {{ end }}
    targets.push(t);
    {{ end }}

    var id = document.getElementById('cf_id');
    id.value = "{{ .AppPageData.ThingId }}";
    var d = document.getElementById('cf_selection_form');
    d.action = "/app/{{ .AppPageData.ThingId }}/framework";

</script>
<script src="/js/cf.js"></script>
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