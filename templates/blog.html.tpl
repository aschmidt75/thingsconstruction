{{define `main`}}
    <div class="row">
    </div>
    <div class="row">
        <div class="col s9">
            {{.HtmlOutput}}
        </div>
        <div class="col s3">
            <p>
            <h5>Tags</h5>
            {{ range .TagChipData }}
                <div class="chip {{ if .Active }}deep-orange lighten-3{{end}}">
                    {{ .TagName }}
                </div>
            {{ end }}
            </p>
            <p>
            <h5>Recent posts</h5>
            <ul>
                {{ range .AllPostsChrono }}
                <li>
                    <div><a class="deep-orange-text text-lighten-1" href="/blog/{{.Name}}">{{ .Title }}</a></div>
                </li>
                {{ end }}
            </ul>
            </p>
            <p>
                <a href="/blog"><span><i class="material-icons tiny">arrow_back</i></span>See all</a>
            </p>
            <p style="padding-top: 1em">
                <a rel="license" target="tc-ext" href="http://creativecommons.org/licenses/by-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by-sa/4.0/88x31.png" /></a><br />
            </p>

        </div>
    </div>

{{end}}