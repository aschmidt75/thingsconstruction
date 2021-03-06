{{define `main`}}
    <div class="row">
    </div>
    <div class="row">
        <div id="bp_overview" class="col s10">
            <h4><span id="bp_title">Blog posts</span><span id="bp_count" class="badge" data-badge-caption=""></span></h4>
            <p>
            <br>
            {{ range .AllPostsChrono }}
                <div id="bp: {{.Name}}" class="blogpost row ">
                    <div class="col s3 left hide-on-small-only show-on-medium-and-up" >
                    {{ range .Tags }}
                        <div class="chip">
                        {{ . }}
                        </div>
                    {{end}}
                    <p>{{ .DateFormatted }}
                    </p>


                    </div>
                    <div class="col s12 m9 right">
                        <h5 style="margin-top: 0px;"><a class="tc-maincolor-text text-lighten-1" href="/blog/{{.Name}}">{{ .Title }}</a>
                    {{ if .HasVideo }}
                        <i class="material-icons">ondemand_video</i>
                    {{ end }}
                        </h5>

                    {{.Abstract}}
                    </div>

                </div>

            {{ end }}
            </p>
        </div>
        <div id="bp_all_tags" class="col s2">
            <h5>Filter by tags</h5>
            {{ range .TagChipData }}
                <div class="chip hoverable tc-filter">
                    {{ .TagName }}
                </div>
            {{ end }}
            <p style="padding-top: 5em">
                All blog posts are<br/>
                CC-BY-SA 4.0<br/>
                <a rel="license" target="tc-ext" href="http://creativecommons.org/licenses/by-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by-sa/4.0/88x31.png" /></a><br />
            </p>
        </div>
    </div>

{{end}}
