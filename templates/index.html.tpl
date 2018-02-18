{{define `main`}}
<div class="row">
    <div class="col s12">
        <div>
            <p>

            </p>
        </div>

        <div class="carousel carousel-slider center" data-indicators="true">
        {{ if .Feature.App  }}
            <div class="carousel-fixed-item center">
                <a class="waves-effect waves-light btn teal waves-teal" href="/app"><i class="material-icons left">settings</i>Get Started</a>
            </div>
        {{end}}
            <div class="carousel-item grey lighten-3 black-text" href="#one!">
                <h2>Support for embedded device targets</h2>
                <p>Code generators are available for different embedded development boards</p>
                <img src="/img/index_carousel_module.png">
            </div>
            <div class="carousel-item grey lighten-4 black-text" href="#two!">
                <h2>Uses W3C's Web of Things Descriptions</h2>
                <p>Model device behaviour including properties, actions and events.</p>
                <img src="/img/index_carousel_wot.png">
            </div>
            <div class="carousel-item grey lighten-5 black-text" href="#three!">
                <h2>Generates code for a complete Web Of Things API</h2>
                <p>Generate and download your code as a skeleton, ready to run.</p>
                <img src="/img/index_carousel_download.png">
            </div>
            <div class="carousel-item grey lighten-4 black-text" href="#four!">
                <h2>Run code in your favourite development environment</h2>
                <p>Paste code into your favourite embedded IDE, adapt & flash</p>
                <img src="/img/index_carousel_code.png">
            </div>
        </div>


    </div>
</div>
{{end}}