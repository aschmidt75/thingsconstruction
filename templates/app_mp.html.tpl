{{define `main`}}
<div id="main" class="cantainer">
    <div class="row">
        <div class="col s12">
        {{ if .Msg }}
            <div class="card red darken-2">
                <div class="card-content white-text">
                {{ .Msg }}
                </div>
            </div>
        {{ else }}
            <!-- mp == manage properties -->
            <div class="row section scrollspy" id="mp">
                <div class="col s9">
                    <h4>Properties</h4>
                    <div>
                        <p>
                            A WoT thing can have properties, such as temperature, led state and so on. Create them here
                            by clicking on the
                            <i class="tiny material-icons">playlist_add</i>
                            button below. Each property has a name, a type and an optional description.
                        </p>
                    </div>
                    <div>
                        <p>
                        <form id="mpf" method="post" action="/app/{{.ThingId}}/properties">
                            <ul id="mp_list" class="collection">
                            </ul>
                            <input type="text" class="hide" id="mpf_id" name="mpfid" value="{{.ThingId}}">

                    </form>
                        </p>
                    </div>
                    <div style="padding-bottom: 4em">
                        <button id="mp_add_btn" class="btn btn-floating deep-orange tooltipped left" data-delay="100" data-tooltip="New property">
                            <i class="material-icons">playlist_add</i>
                        </button>
                    </div>
                </div>
                <div class="row section scrollspy" id="details_next_row" >
                    <div class="col s6 offset-s6">
                        <button id="details_next" class="btn-large deep-orange tooltipped right-align" data-delay="100" data-tooltip="Save changes, go to next step">
                            Next step
                            <i class="material-icons right">navigate_next</i>
                        </button>
                    </div>
                </div>
            </div>

        {{ end }}

            <!--{{.ThingId}}-->
        </div>
    </div>


<!-- hidden, tools -->
<!-- Modal Structure -->
<div id="mp_listitem_validation_modal" class="modal">
    <div class="modal-content">
        <h4>Property Validation</h4>
        <p>Please give each property entry a unique, non-empty name.</p>
        <p id="mp_listitem_validation_modal_reason"></p>
    </div>
    <div class="modal-footer">
        <a href="#!" class="modal-action modal-close grey white-text waves-effect btn-flat">Got it!</a>
    </div>
</div>
<div id="details_validation_modal" class="modal">
    <div class="modal-content">
        <h4>Details Validation</h4>
        <p>There are some things missing. Please correct them first.</p>
        <p id="details_validation_modal_reason"></p>
    </div>
    <div class="modal-footer">
        <a href="#!" class="modal-action modal-close grey white-text waves-effect btn-flat">Got it!</a>
    </div>
</div>
<!-- template for Property list item -->
<div id="tlp_mpl_list_item" class="hide">
    <div class="col s8" id="mp_listitem_##_show">
      <span class="title">
        <strong id="mp_listitem_##_name">Name</strong>
      </span>
        <span id="mp_listitem_##_details">
        <p>Content
          <br> Description
        </p>
      </span>
    </div>
    <div class="col s8 hide" id="mp_listitem_##_edit">
        <div class="input-field">
            <input type="text" placeholder="i.e. temperature" id="mp_listitem_##_edit_name" class="input-field validate" value=""></input>
            <label for="mp_listitem_##_edit_name">Property Name</label>
        </div>
        <span id="mp_listitem_##_details">
        <div class="input-field">
          <input type="text" placeholder="some fancy one-line description" id="mp_listitem_##_edit_desc" class="input-field" value="">
          <label for="mp_listitem_##_edit_desc">Description</label>
        </div>
        <!-- type -->
        <div class="row">
          <div class="col s12">
           <input type="text" id="mp_listitem_##_type_" class="input-field hide" value="mp_listitem_##_type_str">
           <ul class="tabs deep-orange lighten-5 tabs-fixed-width">
              <li id="mp_listitem_##_type_bool_click" class="tab col">
                <a href="#mp_listitem_##_type_bool">Bool</a>
              </li>
              <li  id="mp_listitem_##_type_number_click" class="tab col">
                <a href="#mp_listitem_##_type_number">Number</a>
              </li>
              <li id="mp_listitem_##_type_str_click" class="tab col">
                <a class="active" href="#mp_listitem_##_type_str">String</a>
              </li>
            </ul>
          </div>
          <div id="mp_listitem_##_type_bool" class="col">
            <p>Property is a simple boolean. On or off. No need to specify more.</p>
          </div>
          <div id="mp_listitem_##_type_number" class="col">
              <p>Property is a number. You can specify type and min/max values here.</p>
            <div class="row s8">
              <div class="input-field col s4">
                <select id="mp_listitem_##_type_number_type" >
                  <option value="1">Integer</option>
                  <option value="2">Float</option>
                </select>
                <label>Type</label>
              </div>
              <div class="input-field col s2">
                <input type="text" id="mp_listitem_##_type_number_min" class="input-field" value="">
                <label for="mp_listitem_##_type_number_min">Min</label>
              </div>
              <div class="input-field col s2">
                <input type="text" id="mp_listitem_##_type_number_max" class="input-field" value="">
                <label for="mp_listitem_##_type_number_max">Max</label>
              </div>
            </div>
          </div>
          <div id="mp_listitem_##_type_str" class="col">
            <p>Property is a string. You can specify a maximum length here (optional).</p>
            <div class="input-field">
              <input type="text" id="mp_listitem_##_type_str_maxlength" class="input-field" value=""></input>
              <label for="mp_listitem_##_type_str_maxlength">Maximum length</label>
            </div>
          </div>
        </div>

        </span>
    </div>
    <div id="mp_listitem_##_btns">
        <div class="right col secondary-content">
            <button id="mp_listitem_##_btns_edit" class="btn-flat waves-effect waves-deep-orange">
                <i class="material-icons large">mode_edit</i>
            </button>
            <button id="mp_listitem_##_btns_delete" class="waves-effect waves-red btn-flat" style="padding: 0 0rem; width:2em; background-color:white; color: red">
                <i class="material-icons">delete</i>
            </button>
        </div>
    </div>
</div>

{{end}}