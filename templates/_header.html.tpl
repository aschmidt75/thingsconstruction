{{define `header`}}
<header>
  <div class="navbar-fixed">
    <nav class="deep-orange darken-1">
      <div class="nav-wrapper">
        <div class="container">
          <a href="/index.html" class="brand-logo"><img style="height:40px; margin-top: 10px" src="/img/p0small.png"></a>
          <ul class="right hide-on-med-and-down">
            <li>
              <a href="/about.html">About</a>
            </li>
            {{ if .BlogFeature }}
            <li>
              <a href="/blog">Blog</a>
            </li>
            {{ end }}
            {{ if .AppFeature }}
            <li>
                <a class="waves-effect waves-light btn grey accent-2" href="/app"><i class="material-icons left">launch</i>Get Started</a>
            </li>
            {{ end }}
          </ul>
        </div>
      </div>
    </nav>
  </div>
</header>
{{end}}