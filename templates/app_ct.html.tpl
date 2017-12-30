{{define `main`}}
<div id="main" class="container">
  <!-- ct == create thing -->
  <div class="row" id="ct">
    <div class="col s12">
        <p></p>
        <h4>Create a new Thing Description</h4>
        <p style="padding-bottom: 2em">
            To start a new WoT Thing Description, start with a Name and Description.
        </p>
    {{ if .Msg }}
        <div class="card red darken-2">
            <div class="card-content white-text">
            {{ .Msg }}
            </div>
        </div>
    {{ end }}
        <br/>
        <form id="ctf" name="ctf" method="POST" action="/app">
            <div class="row">
                <div class="input-field col s8">
                    <input placeholder="some fancy name for your shiny device (required)" type="text" name="ctf_name" value="{{ .CtfName }}" id="ctf_name" size="20" />
                    <label for="ctf_name">Name</label>
                </div>
                <div class="input-field col s4">
                    <select id="ctf_type" name="ctf_type">
                        <option value="thing" selected>Thing</option>
                        <!--
                        <option value="onoffthing">OnOffThing</option>
                        <option value="binarysensor">BinarySensor</option>
                        -->
                    </select>
                    <label>Select a type.</label>
                </div>
            </div>
            <div class="row">
                <div class="input-field col s12">
                    <textarea placeholder="Give it a description. This is optional" class="materialize-textarea" name="ctf_desc" id="ctf_desc" size="20">{{ .CtfDesc }}</textarea>
                    <label for="ctf_desc">Description</label>
                </div>
            </div>
            <div class="row" >
                <a class="waves-effect waves-light deep-orange btn right" id="ct_next"><i class="material-icons right">send</i>Next</a>
            </div>
        </form>
    </div>
  </div>
</div>
<!-- Modal Structure -->
<div id="ct_form_validation_modal" class="modal">
    <div class="modal-content">
        <h4>Thing Description Validation</h4>
        <p>Please enter a name.</p>
        <p id="ct_form_validation_modal_reason"></p>
    </div>
    <div class="modal-footer">
        <a href="#!" class="modal-action modal-close waves-effect grey white-text btn-flat">Got it!</a>
    </div>
</div>
{{end}}