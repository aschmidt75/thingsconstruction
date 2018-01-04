{{define `main`}}
<div id="main" class="no-container">
    <div class="row">
        <div class="col s12">
            <!-- cf == choose framework -->
            <div class="col s9">
                <div class="row">
                    <p><br></p>
                    <h4>Choose Code Generation Module</h4>
                    <p>
                        The list below shows all available generators. You can narrow down the list of
                        generators by selecting implementation aspects from the stack below.
                    </p>
                    <div class="row">
                        <div class="col s2 deep-orange-text darken-1">
                            Available
                        </div>
                        <div id="cf_targets_available" class="col s8 teal-text">
                        </div>
                    </div>
                    <div class="row">
                        <div class="col s2 deep-orange-text darken-1">
                            Filter by ...
                        </div>
                        <div id="cf_targets_selected" class="col s8 teal-text">
                        </div>
                    </div>

                </div>
                <div class="row">
                    <h4>Available generators<span class="new badge grey darken-1" id="cf_targets_num_showing" data-badge-caption="">0</span></h4>
                    <ul class="collapsible" data-collapsible="accordion" id="cf_targets_matching">
                    </ul>
                </div>
                <div class="hide">
                    <form id="cf_selection_form" name="cf_selection_form" method="POST" action="">
                        <input type="text" id="cf_selection" name="cfs">
                        <input type="text" id="cf_id" name="cfid" value="">
                    </form>
                </div>
            </div>
            <div class="col s2 offset-s1 hide-on-med-and-down">
                <p>
                <ul class="section table-of-contents">
                    <li>Create</li>
                    <li><strong class="deep-orange-text">Framework</strong>{{ if .TocInfo.framework }}&nbsp;(<span>{{ index .TocInfo "framework" }}</span>){{end}}</li>
                    <li>Properties{{ if .TocInfo.num_properties }}&nbsp;(<span>{{ index .TocInfo "num_properties" }}</span>){{end}}</li>
                    <li>Actions{{ if .TocInfo.num_events }}&nbsp;(<span id="toc_current_info">{{ index .TocInfo "num_actions" }}</span>){{end}}</li>
                    <li>Events{{ if .TocInfo.num_events }}&nbsp;(<span>{{ index .TocInfo "num_events" }}</span>){{end}}</li>
                    <li>Generate!</li>
                </ul>
                </p>
            </div>
        </div>
    </div>
</div>
{{end}}