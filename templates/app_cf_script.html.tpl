{{define `script`}}
<script>
    function Target(id,shortDesc,desc,tags,codegeninfo) {
        this.id = id;
        this.shortDesc = shortDesc;
        this.desc = desc;
        this.tags = tags;
        this.selected = false;
        this.deps = []
        this.codegeninfo = codegeninfo
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
{{end}}