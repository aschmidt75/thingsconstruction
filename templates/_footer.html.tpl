{{define `footer`}}
  <footer class="page-footer tc-maincolor darken-1">
    <div class="container">
      <div class="row">
        <div class="col s12 m8 l8">
          <h5 class="white-text">Web Of Things</h5>
            <p class="grey-text text-lighten-4">
                Read more about <a class="grey-text text-lighten-1" target="tc-wot-ext" href="https://www.w3.org/WoT/"><i class="tiny material-icons">exit_to_app</i>WoT</a> and the
                <a class="grey-text text-lighten-1" target="tc-wot-ext" href="https://iot.mozilla.org/wot/"><i class="tiny material-icons">exit_to_app</i>Web Thing API</a></p>
        </div>
        <div class="col s12 m4 l4 right-align">
            <ul>
                {{ if .Feature.GitHub }}
                    <a target="tc-ext" style="padding-right: 10px" href="{{ .GitHubUrl }}"><i class="white-text fab fa-github fa-2x"></i></a>
                {{ end }}
                {{ if .Feature.Twitter }}
                    <a target="tc-ext" style="padding-right: 10px" href="{{ .TwitterUrl }}"><i class="white-text fab fa-twitter fa-2x"></i></a>
                {{ end }}
                {{ if .Feature.LinkedIn }}
                    <a target="tc-ext" style="padding-right: 10px" href="{{ .LinkedInUrl }}"><i class="white-text fab fa-linkedin fa-2x"></i></a>
                {{ end }}
            </ul>
        </div>
      </div>
    </div>
    <div class="footer-copyright">
        <div class="container"><span style="font-size: 0.8em"><!--by A.Schmidt&nbsp;|&nbsp;-->{{ .CopyrightLine }}</span>
          <div class="right">
              <a class="grey-text text-lighten-4" href="/privacy.html">Privacy Policy</a> &nbsp;|&nbsp;
              <a class="grey-text text-lighten-4" href="/imprint.html">Imprint</a> &nbsp;|
              <a class="grey-text text-lighten-4" href="/terms.html">Terms</a>
          </div>
            <div style="font-size: 0.5em">
            {{ .Notices }}
            </div>
        </div>
    </div>
  </footer>
  <script src="/js/jquery-3.2.1.min.js"></script>
  <script src="/js/jquery.form.js"></script>
  <script src="/js/class.min.js"></script>
  <script src="/js/jquery.jquery-encoder.min.js"></script>
  <script src="/js/materialize.min.js"></script>
  <!-- 1.0 alpha Compiled and minified JavaScript
  <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0-alpha.3/js/materialize.min.js"></script>
          -->
  {{ if .Feature.Shariff }}
  <script src="/js/shariff.min.js"></script>
  {{ end -}}

  <script src="/js/main.js"></script>
{{ template `script` . }}

{{end}}