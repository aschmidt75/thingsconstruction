{{define `header`}}
<header>
  <div class="navbar-fixed">
    <nav class="deep-orange darken-1">
      <div class="nav-wrapper">
        <div class="container">
          <a href="/index.html" class="brand-logo"><img style="height:40px; margin-top: 10px" src="/img/logosmall.png"></a>
          <a href="#" data-activates="mobile" class="button-collapse"><i class="material-icons">menu</i></a>
          <ul class="right hide-on-med-and-down">
            {{ if .Feature.Blog }}
            <li>
                <a href="/blog/about.md">About</a>
            </li>
            <li>
              <a {{ if .InBlog }}class="active"{{end}} href="/blog">Blog</a>
            </li>
            {{ end }}
            {{ if .Feature.App }}
            <li>
                <a {{ if .InApp }}class="active"{{end}} href="/app" title="Get Started!"></i><i class="material-icons">settings</i></a>
            </li>
            {{ end }}
            {{ if .Feature.Contact }}
              <li>
                  <a {{ if .InContact }}class="active"{{end}} href="/feedback" title="Your feedback is welcome!"><i class="material-icons">comment</i></a>
              </li>
            {{ end }}

          </ul>
        </div>
      </div>
    </nav>
  </div>
    <!-- https://github.com/Dogfalo/materialize/issues/3982 -->
    <ul class="side-nav" id="mobile">
        <li>
            <a href="/about.html">About</a>
        </li>
    {{ if .Feature.Blog }}
        <li>
            <a href="/blog">Blog</a>
        </li>
    {{ end }}
    {{ if .Feature.App }}
        <li>
            <a class="active" href="#"></i>App</a>
        </li>
    {{ end }}
    </ul>
</header>
{{end}}