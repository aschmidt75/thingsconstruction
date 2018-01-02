{{define `main`}}
<div id="main" class="no-container">
    <div class="row">
        <div class="col s12">
            <h3>An error occurred...</h3>
        {{ if .Message }}
            <div class="card red darken-2">
                <div class="card-content white-text">
                {{ .Message }}
                </div>
            </div>
        {{ else }}
            <div class="card red darken-2">
                <div class="card-content white-text">
                No message has been delivered. We're sorry!
                </div>
            </div>
        {{ end }}
        </div>
    </div>
</div>
{{end}}