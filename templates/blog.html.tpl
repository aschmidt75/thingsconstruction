{{define `main`}}
    <div class="row">
    </div>
    <div class="row">
        <div class="col s9">
        {{ if .VimeoID }}
            <p id="embed" video="{{.VimeoID}}">
            </p>
        {{ end }}
        {{ if .YoutubeID }}
            <p>
                <iframe width="640" height="360" src="https://www.youtube-nocookie.com/embed/{{.YoutubeID}}?rel=0"
                        frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe>
            </p>
        {{ end }}
            {{.HtmlOutput}}
            {{ if .Feature.Shariff }}
                    <p>
                <div class="shariff" data-services="[twitter,linkedin]" data-lang="en">
                </div>
                </p>
            {{ end -}}
        </div>
        <div class="col s1">
        </div>
        <div class="col s2">
            <p>
            <h5>Tags</h5>
            {{ range .TagChipData }}
                <div class="chip {{ if .Active }}tc-maincolor lighten-3{{end}}">
                    {{ .TagName }}
                </div>
            {{ end }}
            </p>
            <p>
            <h5>Recent posts</h5>
            <p>
                <a href="/blog"><span><i class="material-icons tiny">arrow_back</i></span>See all</a>
            </p>
            <ul>
                {{ range .AllPostsChrono }}
                <li>
                    <div style="margin-top: 0.75em"><a class="tc-maincolor-text text-lighten-1" href="/blog/{{.Name}}">{{ .Title }}</a></div>
                </li>
                {{ end }}
            </ul>
            </p>
            <p style="padding-top: 1em">
                This blog post is<br/>
                CC-BY-SA 4.0<br/>
                <a rel="license" target="tc-ext" href="http://creativecommons.org/licenses/by-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by-sa/4.0/88x31.png" /></a><br />
            </p>

        </div>
    </div>

{{end}}