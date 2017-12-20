{{define `script`}}
<script>
    function Target(id,shortDesc,desc,tags) {
        this.id = id;
        this.shortDesc = shortDesc;
        this.desc = desc;
        this.tags = tags;
        this.selected = false;
    };

    var targets = [];
    {{ range .AppGenTargets }}
    targets.push(new Target('{{.Id}}', '{{.ShortDesc}}', '{{.Desc}}', [{{ range .Tags }}'{{.}}', {{end}} ]))
    {{ end }}
</script>
<script src="/js/cf.js"></script>
{{end}}