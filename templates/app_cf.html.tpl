{{define `main`}}
<div id="main" class="container">
    <div class="row">
        <div class="col s12">
            <!-- cf == choose framework -->
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
    </div>
</div>
{{end}}