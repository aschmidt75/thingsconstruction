{{define `main`}}
    <div class="row">
    </div>
    <div class="row">
        <div id="bp_overview" class="col s9">
            <h4><span id="bp_title">All blog posts</span><span id="bp_count" class="badge new" data-badge-caption=""></span></h4>
            <p>
            <br>
            {{ range .AllPostsChrono }}
                <div id="bp: {{.Name}}" class="blogpost row">
                    <div class="col s4 left" >
                    {{ range .Tags }}
                        <div class="chip">
                        {{ . }}
                        </div>
                    {{end}}
                    <p>{{ .Date.Format "Mon, Jan 2 2006" }}</p>


                    </div>
                    <div class="col s8 right">
                        <h5><a class="deep-orange-text text-lighten-1 truncate" href="/blog/{{.Name}}">{{ .Title }}</a></h5>
                        {{.Abstract}}
                    </div>

                </div>
            {{ end }}
            </p>
        </div>
        <div id="bp_all_tags" class="col s3">
            <h5>Filter by tags</h5>
            {{ range .TagChipData }}
                <div class="chip bp_selected deep-orange lighten-3">
                    {{ .TagName }}
                </div>
            {{ end }}
        </div>
    </div>

{{end}}