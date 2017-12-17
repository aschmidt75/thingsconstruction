{{define `header`}}
  <header>
    <nav class="deep-orange darken-1">
      <div class="nav-wrapper">
        <div class="container">
          <a href="/index.html" class="brand-logo">ThingsConstruction</a>
          <ul class="right hide-on-med-and-down">
            <li>
              <a href="/about.html">About</a>
            </li>
            {{ if .BlogFeature }}
            <li>
              <a href="/blog">Blog</a>
            </li>
            {{ end }}
          </ul>
        </div>
      </div>
    </nav>
  </header>
{{end}}