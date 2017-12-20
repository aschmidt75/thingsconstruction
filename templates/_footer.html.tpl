{{define `footer`}}
  <footer class="page-footer deep-orange darken-1">
    <div class="container">
      <div class="row">
        <div class="col s12 m8 l8">
          <h5 class="white-text">Web Of Things</h5>
            <p class="grey-text text-lighten-4">
                Read more about <a class="grey-text text-lighten-1" target="tc-wot-ext" href="https://www.w3.org/WoT/"><i class="tiny material-icons">link</i>WoT</a> and the
                <a class="grey-text text-lighten-1" target="tc-wot-ext" href="https://iot.mozilla.org/wot/"><i class="tiny material-icons">link</i>Web Thing Description</a></p>
        </div>
        <div class="col s12 m4 l4 right-align">
            <ul>
                    <a target="tc-ext" style="padding-right: 10px" href="https://github.com/aschmidt75/thingsconstruction"><i class="white-text fab fa-github fa-2x"></i></a>
                    <a target="tc-ext" style="padding-right: 10px" href="https://twitter.com/aschmidt75"><i class="white-text fab fa-twitter fa-2x"></i></a>
            </ul>
        </div>
      </div>
    </div>
    <div class="footer-copyright">
      <div class="container">
        © 2017 @aschmidt75
        <a class="grey-text text-lighten-4 right" href="/imprint.html">Imprint</a>
      </div>
    </div>
  </footer>
  <!-- The jQuery file path -->
  <script src="/js/jquery-3.2.1.min.js"></script>
  <script src="/js/jquery.form.js"></script>
  <!-- Path of the materialize.min.js file -->
  <script src="/js/materialize.min.js"></script>
  <script src="/js/main.js"></script>
{{ template `script` . }}

{{end}}