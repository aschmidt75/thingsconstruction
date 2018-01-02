{{define `main`}}
<div id="main" class="no-container">
    <div class="row">
        <div class="col s12">
        {{ if .Msg }}
            <div class="card red darken-2">
                <div class="card-content white-text">
                {{ .Msg }}
                </div>
            </div>
        {{ else }}
            <!-- ma == manage actions -->
            <div class="row" id="ma">
                <div class="col s9">
                    <h4>Action</h4>
                    <div>
                        <p>
                            A WoT thing can have actions, such as "reset". Create them here
                            by clicking on the
                            <i class="tiny material-icons">playlist_add</i>
                            button below. Each action has a name and an optional description.
                        </p>
                    </div>
                    <div>
                        <p>
                            <form id="maf" method="post" action="/app/{{.ThingId}}/actions">
                                <ul id="ma_list" class="collection">
                                </ul>
                                <input type="text" class="hide" id="maf_id" name="mafid" value="{{.ThingId}}">
                            </form>
                        </p>
                    </div>
                    <div style="padding-bottom: 4em">
                        <button id="ma_add_btn" class="btn btn-floating deep-orange tooltipped left" data-delay="100" data-tooltip="New property">
                            <i class="material-icons">playlist_add</i>
                        </button>
                    </div>

                </div>
                <div class="col s2 offset-s1 hide-on-med-and-down">
                    <p>
                    <ul class="section table-of-contents">
                        <li>Create</li>
                        <li>Framework{{ if .TocInfo.framework }}&nbsp;(<span>{{ index .TocInfo "framework" }}</span>){{end}}</li>
                        <li>Properties{{ if .TocInfo.num_properties }}&nbsp;(<span>{{ index .TocInfo "num_properties" }}</span>){{end}}</li>
                        <li><strong class="deep-orange-text">Actions</strong>&nbsp;(<span id="toc_actions_info">0</span>)</li>
                        <li>Events{{ if .TocInfo.num_events }}&nbsp;(<span>{{ index .TocInfo "num_events" }}</span>){{end}}</li>
                        <li>Generate!</li>
                    </ul>
                    </p>
                </div>

                <div class="row">
                    <div class="col s9">
                        <button id="ma_prev" class="btn-large deep-orange tooltipped left" data-delay="100" data-tooltip="Discard changes, go to previous step">
                            <i class="material-icons left">navigate_before</i>Previous step

                        </button>
                        <button id="ma_next" class="btn-large deep-orange tooltipped right" data-delay="100" data-tooltip="Save changes, go to next step">
                            Next step
                            <i class="material-icons right">navigate_next</i>
                        </button>
                    </div>
                </div>
            </div>

        {{ end }}
        </div>
    </div>


    <!-- hidden, tools -->
    <!-- Modal Structure -->
    <div id="ma_listitem_validation_modal" class="modal">
        <div class="modal-content">
            <h4>Action Validation</h4>
            <p>Please give each property entry a unique, non-empty name.</p>
            <p id="ma_listitem_validation_modal_reason"></p>
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
    <!-- template for Action list item -->
    <div id="tpl_maf_list_item" class="hide">
        <div class="col s8" id="ma_listitem_##_show">
          <span class="title">
            <strong id="ma_listitem_##_name">Name</strong>
          </span>
            <span id="ma_listitem_##_details">
            <p>Content
              <br> Description
            </p>
          </span>
        </div>
        <div class="col s8 hide" id="ma_listitem_##_edit">
            <div class="input-field">
                <input type="text" placeholder="i.e. reset" id="ma_listitem_##_edit_name" class="input-field validate" value="">
                <label for="mp_listitem_##_edit_name">Action Name</label>
            </div>
            <span id="ma_listitem_##_details">
                <div class="input-field">
                  <input type="text" placeholder="some fancy one-line description" id="ma_listitem_##_edit_desc" class="input-field" value="">
                  <label for="ma_listitem_##_edit_desc">Description</label>
                </div>
            </span>
        </div>
        <div id="ma_listitem_##_btns">
            <div class="right col secondary-content">
                <button id="ma_listitem_##_btns_edit" class="btn-flat waves-effect waves-deep-orange">
                    <i class="material-icons large">mode_edit</i>
                </button>
                <button id="ma_listitem_##_btns_delete" class="waves-effect waves-red btn-flat" style="padding: 0 0rem; width:2em; background-color:white; color: red">
                    <i class="material-icons">delete</i>
                </button>
            </div>
        </div>
    </div>
</div>

{{end}}