{{define `main`}}
<div class="row">
    <div class="col s12">
        <h3>Welcome!</h3>

    {{ if .AppFeature  }}
            <a class="waves-effect waves-light btn grey accent-2" href="/app"><i class="material-icons left">launch</i>Get Started</a>
    {{end}}

    </div>
</div>
{{end}}