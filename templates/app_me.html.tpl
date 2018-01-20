{{define `main`}}
<div id="main" class="no-container">
    <div class="row">
        <div class="col s12">
        {{ if .Message }}
            <div class="card red darken-2">
                <div class="card-content white-text">
                {{ .Message }}
                </div>
            </div>
        {{ else }}
            <!-- me == manage events -->
            <div class="row" id="me">
                <div class="col s9">
                    <h4>Events</h4>
                    <div>
                        <p>
                            A WoT thing can have events, such as "ButtonPressed". Create them here
                            by clicking on the
                            <i class="tiny material-icons">playlist_add</i>
                            button below. Each event has a name and an optional description.
                        </p>
                    </div>
                    <div>
                        <p>
                            <div class="preloader-wrapper small active" id="progress">
                                <div class="spinner-layer spinner-red-only">
                                    <div class="circle-clipper left">
                                        <div class="circle"></div>
                                    </div><div class="gap-patch">
                                    <div class="circle"></div>
                                </div><div class="circle-clipper right">
                                    <div class="circle"></div>
                                </div>
                                </div>
                            </div>
                            <form id="mef" method="post" action="/app/{{.ThingId}}/events">
                                <ul id="me_list" class="collection">
                                </ul>
                                <input type="text" class="hide" id="mef_id" name="mefid" value="{{.ThingId}}">
                            </form>
                        </p>
                    </div>
                    <div style="padding-bottom: 4em">
                        <button id="me_add_btn" class="btn btn-floating deep-orange tooltipped left" data-delay="100" data-tooltip="New event">
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
                        <li>Actions{{ if .TocInfo.num_actions }}&nbsp;(<span>{{ index .TocInfo "num_actions" }}</span>){{end}}</li>
                        <li><strong class="deep-orange-text">Events</strong>&nbsp;(<span id="toc_current_info">0</span>)</li>
                        <li>Generate!</li>
                    </ul>
                    </p>
                </div>
                <div class="row">
                    <div class="col s9">
                        <button id="me_prev" class="btn-large deep-orange tooltipped left" data-delay="100" data-tooltip="Discard changes, go to previous step">
                            <i class="material-icons left">navigate_before</i>Actions

                        </button>
                        <button id="me_next" class="btn-large deep-orange tooltipped right" data-delay="100" data-tooltip="Save changes, go to next step">
                            Generate
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
    <div id="me_listitem_validation_modal" class="modal">
        <div class="modal-content">
            <h4>Action Validation</h4>
            <p>Please give each property entry a unique, non-empty name.</p>
            <p id="me_listitem_validation_modal_reason"></p>
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
    <div id="tpl_mef_list_item" class="hide">
        <div class="col s8" id="me_listitem_##_show">
          <span class="title">
            <strong id="me_listitem_##_name">Name</strong>
          </span>
            <span id="me_listitem_##_details">
            <p>Content
              <br> Description
            </p>
          </span>
        </div>
        <div class="col s8 hide" id="me_listitem_##_edit">
            <div class="input-field">
                <input type="text" placeholder="i.e. buttonPressed" id="me_listitem_##_edit_name" class="input-field validate" value="">
                <label for="mp_listitem_##_edit_name">Event Name</label>
            </div>
            <span id="me_listitem_##_details">
                <div class="input-field">
                  <input type="text" placeholder="some fancy one-line description" id="me_listitem_##_edit_desc" class="input-field" value="">
                  <label for="me_listitem_##_edit_desc">Description</label>
                </div>
            </span>
        </div>
        <div id="me_listitem_##_btns">
            <div class="right col secondary-content">
                <button id="me_listitem_##_btns_edit" class="btn-flat waves-effect waves-deep-orange">
                    <i class="material-icons large">mode_edit</i>
                </button>
                <button id="me_listitem_##_btns_delete" class="waves-effect waves-red btn-flat" style="padding: 0 0rem; width:2em; background-color:white; color: red">
                    <i class="material-icons">delete</i>
                </button>
            </div>
        </div>
    </div>
</div>

{{end}}