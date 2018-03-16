{{define `main`}}
<div id="main" class="container">
  <!-- ct == create thing -->
  <div class="row" id="ct">
    <div class="col s12">
        <p></p>
        <h4>Create a new Thing Description</h4>
        <p style="padding-bottom: 2em">
            To start a new WoT Thing Description, choose a name, a type and describe your device.<br>
            <span class="tc-maincolor-text" id="btn_more">... more</span>
            <span id="sp_more" class="hide">
                The name of your device will appear within the Thing Description and Thing API when querying
                the device. You can also choose a type for it. Default type is "Thing" which is
                a blank template. Choosing other types will automatically populate properties and actions in
                the upcoming dialogs with defaults. For an overview of thing types, click on the <i class="material-icons tiny">info_outline</i>
                button.
                <span class="tc-maincolor-text" id="btn_less">... less</span>
            </span>
        </p>
    {{ if .Message }}
        <div class="card red darken-2">
            <div class="card-content white-text">
            {{ .Message }}
            </div>
        </div>
    {{ end }}
        <br/>
        <form id="ctf" name="ctf" method="POST" action="/app{{ if .ThingId }}/{{ .ThingId}}{{end}}">
            <div class="row">
                <div class="input-field col s7">
                    <input placeholder="some fancy name for your shiny device (required)" type="text" name="ctf_name" value="{{ .CtfName }}" id="ctf_name" size="20" />
                    <label for="ctf_name">Name</label>
                </div>
                <div class="input-field col s4">
                    <select id="ctf_type" {{- if eq .AllowTypeSelection false }} disabled {{ end -}}name="ctf_type">
                        <option value="thing" selected>Thing</option>
                        <option value="onOffSwitch">OnOffSwitch</option>
                        <option value="multilevelSwitch">MultilevelSwitch</option>
                        <option value="binarySensor">BinarySensor</option>
                        <option value="multilevelSensor">MultilevelSensor</option>
                        <option value="smartPlug">SmartPlug</option>
                        <option value="onOffLight">OnOffLight</option>
                        <option value="dimmableLight">DimmableLight</option>
                        <option value="onOffColorLight">OnOffColorLight</option>
                        <option value="dimmableColorLight">DimmableColorLight</option>
                    </select>
                    <label>Select a type.</label>
                </div>
                <div class="input-field col s1">
                    <i id="type_info" class="material-icons">info_outline</i>
                </div>
            </div>
            <div class="row">
                <div class="input-field col s12">
                    <textarea placeholder="Give it a description. This is optional" class="materialize-textarea" name="ctf_desc" id="ctf_desc" size="20">{{ .CtfDesc }}</textarea>
                    <label for="ctf_desc">Description</label>
                </div>
            </div>
            <div class="row" >
                <a class="waves-effect waves-light tc-maincolor btn right" id="ct_next"><i class="material-icons right">send</i>Next</a>
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
<div id="ct_type_info" class="modal">
    <div class="modal-content">
        <h4>Type Info</h4>
            Choosing a type other than "Thing" will add default properties and/or actions to the model.
        <table>
            <tr>
                <td>Thing</td><td>A generic thing without any presets.</td>
            </tr>
            <tr>
                <td>OnOffSwitch</td><td>Aa generic device type for actuators with a simple on/off state.</td>
        </tr>
            <tr>
                <td>MultilevelSwitch</td><td>A switch with multiple levels, i.e. a dimmer.</td>
        </tr>
            <tr>
                <td>BinarySensor</td><td>Binary Sensor is a generic type for sensors with a simple on/off state<./td>
        </tr>
            <tr>
                <td>MultilevelSensor</td><td>A generic multi level sensor with a value expressed as a percentage.</td>
        </tr>
            <tr>
                <td>SmartPlug</td><td>Has a simple on/off state and measures power consumption.</td>
        </tr>
            <tr>
                <td>OnOffLight</td><td>A light that can be turned on and off.</td>
        </tr>
            <tr>
                <td>DimmableLight</td><td>A light with controllable brightness, expressed as a percentage</td>
        </tr>
            <tr>
                <td>OnOffColorLight</td><td>A light that can be turned on and off, including a color setting.</td>
        </tr>
            <tr>
                <td>DimmableColorLight</td><td>Same, with controllable brightness.</td>
            </tr>
        </table>
    </div>
    <div class="modal-footer">
        <a href="#!" class="modal-action modal-close waves-effect grey white-text btn-flat">Dismiss</a>
    </div>
</div>
{{end}}