{{define `main`}}
    <div class="row">
    </div>
    <div class="row">
        <div id="bp_overview" class="col s9">
            <h4><span id="bp_title">Blog posts</span><span id="bp_count" class="badge" data-badge-caption=""></span></h4>
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
                        <h5><a class="tc-maincolor-text text-lighten-1 truncate" href="/blog/{{.Name}}">{{ .Title }}</a></h5>
                        {{.Abstract}}
                    </div>

                </div>
            {{ end }}
            </p>
        </div>
        <div id="bp_all_tags" class="col s3">
            <h5>Filter by tags</h5>
            {{ range .TagChipData }}
                <div class="chip">
                    {{ .TagName }}
                </div>
            {{ end }}
            <p style="padding-top: 5em">
                <a rel="license" target="tc-ext" href="http://creativecommons.org/licenses/by-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by-sa/4.0/88x31.png" /></a><br />
            </p>
        </div>
    </div>

{{end}}