{{define `main`}}
<div class="row">
</div>
<div class="row">
    <div class="col s12">
        <h4 class="center">Turn your embedded device into a Web Thing!</h4>
        <div class="row">
            <div class="col s4" style="height:20em">
                <div class="card horizontal" style="height:100%">
                    <div class="card-stacked">
                        <div class="card-content">
                            <p>
                                This service let's you create communication models
                                for your embedded device similar to the
                                Web Things API, and turn them into source code for
                                embedded devices, ready to run.
                            </p>
                            <p></p>
                        </div>
                    </div>
                </div>
            </div>
            <div class="col s4" style="height:20em">
                <div class="card horizontal" style="height:100%">
                    <div class="card-stacked">
                        <div class="card-content">
                            <p>
                                Models cover device <strong>properties</strong> - such as
                                temperature or LED state -, <strong>actions</strong> and
                                <strong>events</strong> - i.e. "Flash LED" or "Temperature low".
                            </p>
                            <p></p>
                        </div>
                    </div>
                </div>
            </div>
            <div class="col s4" style="height:20em">
                <div class="card horizontal" style="height:100%">
                    <div class="card-stacked">
                        <div class="card-content">
                            <p>
                                Code generators produce API code, so you don't have to.
                                Examples for APIs are HTTP+JSON REST, MQTT or LPWAN interface.

                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    {{ if .Feature.App  }}
    <div class="row">
        <p></p>
        <div class="carousel-fixed-item center">
                <a class="waves-effect waves-light btn blue-grey darken-3 white-text" href="/app"><i class="material-icons left">settings</i>Get Started</a></p>
        </div>
    </div>
    {{end}}
        <div>

        </div>

    </div>
</div>
{{end}}