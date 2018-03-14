//    ThingsConstruction, a code generator for WoT-based models
//    Copyright (C) 2017  @aschmidt75
//
//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU Affero General Public License as published
//    by the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU Affero General Public License for more details.
//
//    You should have received a copy of the GNU Affero General Public License
//    along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// CF = Choose Framework
//

function Target(id,shortDesc,desc,tags,codegeninfo) {
    this.id = id;
    this.shortDesc = shortDesc;
    this.desc = desc;
    this.tags = tags;
    this.selected = false;
    this.deps = [];
    this.codegeninfo = codegeninfo;
    // TODO: Add version
}

function TargetDep(name,url,license,copyright,info) {
    this.name = name;
    this.url = url;
    this.license = license;
    this.copyright = copyright;
    this.info = info;
}

var targets = [];

var modulesJson = "";
$.ajax({
    type: "GET",
    url: "/modules/data",
    timeout: 5000,
    async: false,
    success: function (data) {
        modulesJson += data;
    },
    error: function (data) {
        console.log(data);
        alert("I'm sorry, an error occurred when reading the modules list!");
    },
    complete: function (data) {
        if (modulesJson != null && modulesJson.length > 0) {
            var obj = JSON.parse(modulesJson);
            if (obj != null) {/// construct targets from it
                for ( var i = 0; i < obj.Targets.length; i++) {
                    var ot = obj.Targets[i];
                    console.log(ot);
                    var t = new Target(ot.Id, ot.ShortDesc, ot.Desc, ot.Tags);
                    for ( var j = 0; j < ot.Dependencies.length; j++) {
                        var oj = ot.Dependencies[j];
                        var d = new TargetDep(oj.Name, oj.URL, oj.License, oj.Copyright, oj.Info);
                        t.deps.push(d);
                    }
                    targets.push(t);
                }
            } else {
                alert("I'm sorry, an error occurred when parsing the modules list!");
            }
        }
    }
});


// given a name of a target tag (i.e. "framework:arduino"), this
// function creates a new chip and places it under nodeId
// unless it already exists
function cf_has_chip_under(target_tag, nodeId) {
    var x = document.getElementById(nodeId).childNodes;
    for ( var i = 0; i < x.length; i++) {
        if (x[i].name == target_tag) {
            return true;
        }
    }
    return false;
}
function cf_insert_here(n1,n2) {
    var component_re = /(\w+):(\w+)/
    if (n1.match(component_re)) {
        n1 = n1.replace(component_re, "$2")
    }
    if (n2.match(component_re)) {
        n2 = n2.replace(component_re, "$2")
    }
    return ( n1 < n2);
}
function cf_add_chip_to(target_tag, nodeId) {
    if ( cf_has_chip_under(target_tag, nodeId)) {
        return;
    }
    var chip = document.createElement('a');
    chip.className = "chip hoverable";
    chip.id = "tag_"+target_tag;
    chip.name = target_tag;
    var component_re = /(\w+):(\w+)/
    if (target_tag.match(component_re)) {
        var img = "";

        if (target_tag.startsWith("framework:")) {
            img = '<i class="tiny material-icons">code</i>&nbsp;';
        }
        if (target_tag.startsWith("proto:")) {
            img = '<i class="tiny material-icons">blur_on</i>&nbsp;';
        }
        if (target_tag.startsWith("app:")) {
            img = '<i class="tiny material-icons">apps</i>&nbsp;';
        }
        if (target_tag.startsWith("conn:")) {
            img = '<i class="tiny material-icons">settings_ethernet</i>&nbsp;';
        }
        if (target_tag.startsWith("mcu:")) {
            img = '<i class="tiny material-icons">memory</i>&nbsp;';
        }
        target_tag = target_tag.replace(component_re, "$2")
        chip.innerHTML = ""+img+target_tag;
    } else {
        chip.innerText = target_tag;
    }
    chip.href = '#';

    var x = document.getElementById(nodeId);
    if (x.childElementCount == 0) {
        document.getElementById(nodeId).appendChild(chip);
    } else {
        x = x.children;
        // insert sorted
        var b_inserted = false;
        for ( var i = 0; i < x.length; i++) {
            if ( cf_insert_here(chip.name,x[i].name)) {
                document.getElementById(nodeId).insertBefore(chip, x[i]);
                b_inserted = true;
                break;
            }
        }
        if ( !b_inserted) {
            document.getElementById(nodeId).appendChild(chip);
        }
    }
    chip.addEventListener('click',cf_select_chip)
    return chip
}

for (var i = 0; i < targets.length; i++) {
    var t = targets[i];
    for ( var j = 0; j < t.tags.length; j++) {
        cf_add_chip_to(t.tags[j], 'cf_targets_available');
    }
}

// Moves a chip from "available" to "selected"
function cf_select_chip(e) {
    e.preventDefault();
    var tag = e.target;
    if ( !tag.className.startsWith("chip")) {
        tag = tag.parentNode
    }
    // move to selected
    tag.parentElement.removeChild(tag);
    var chip2 = cf_add_chip_to(tag.name, 'cf_targets_selected')
    chip2.removeEventListener('click', cf_select_chip);
    tag.removeEventListener('click', cf_select_chip);
    chip2.addEventListener('click', cf_unselect_chip);

    cf_lookup_matching_targets();
}

// Moves a chip from "selected" to "available"
function cf_unselect_chip(e) {
    e.preventDefault();
    var tag = e.target;
    if ( !tag.className.startsWith("chip")) {
        tag = tag.parentNode
    }
    // move to available
    tag.parentElement.removeChild(tag);
    var chip2 = cf_add_chip_to(tag.name, 'cf_targets_available')
    tag.removeEventListener('click', cf_unselect_chip);

    cf_lookup_matching_targets();
}


// returns true if all elements of arr_part are
// also in arr_whole
function is_array_in_array(arr_part, arr_whole) {
    for ( var i = 0; i < arr_part.length; i++) {
        if ( (arr_part[i] != undefined) && (arr_whole.indexOf(arr_part[i]) < 0)) {
            return false;
        }
    }
    return true;
}

// takes all chips from "selected", uses the tags
// to look up all matching targets. Creates a <li>
// element under cf_targets_matching to display
// available generators.
function cf_lookup_matching_targets() {
    // away with the old stuff
    var cf_targets_matching = document.getElementById('cf_targets_matching');
    while (cf_targets_matching.firstChild) {
        cf_targets_matching.removeChild(cf_targets_matching.firstChild);
    }
    // in comes the new stuff. Compose an array of all selected chips
    var cf_targets_selected = document.getElementById('cf_targets_selected');
    var arr_selected = new Array(0);
    for ( var i = 0; i < cf_targets_selected.childNodes.length; i++) {
        var s = cf_targets_selected.childNodes[i];
        if (s != undefined && s.name != undefined) {
            arr_selected.push(s.name);
        }
    }
    // walk through targets, check against arr
    var num_found = 0;
    for ( var i = 0; i < targets.length; i++) {
        var t = targets[i]
        if ( is_array_in_array(arr_selected, t.tags)) {
            // match
            num_found++;
            var li = document.createElement("li");
            var depsStr = "<ul class=\"browser-default\">"
            for (var d = 0; d < t.deps.length; d++) {
                var dep = t.deps[d]
                depsStr += "<li><a target=\"tcext\" href=\""+dep.url+"\"><b>"+dep.name+"</a></b>: "+dep.copyright+" - "+dep.license+"</li>"
            }
            depsStr += "</ul>"
            li.innerHTML = `<div class="collapsible-header">
 <i class="material-icons">keyboard_arrow_right</i>
 ${t.shortDesc}
 <br><br><span class="tc-maincolor-text">...more</span>
 <span class="right col s3">
  <button style="padding-left: 5px" class="tc-maincolor darken-3 btn-floating tooltipped waves-effect waves-light" data-tooltip="Show details" type="" id="details_${t.id}">
   <i class="material-icons">description</i>
  </button>
  <button style="padding-left: 5px" class="tc-maincolor darken-3 btn-floating tooltipped waves-effect waves-light" data-tooltip="Use this generator" type="" id="go_${t.id}">
   <i class="material-icons">arrow_forward</i>
  </button>
 </span>
</div>
<div class="collapsible-body">
 <span>${t.desc}</span><br/>
 <span><b>License Information: </b>
${t.codegeninfo}</span><br/>
 <span><b>Dependencies/Library usages</b>
${depsStr}
</span>
</div>`

            cf_targets_matching.appendChild(li);

            var b = document.getElementById("go_"+t.id);
            b.addEventListener('click', cff_generator_clicked);

            b = document.getElementById("details_"+t.id);
            b.addEventListener('click', cff_generator_details_clicked);
        }
    }
    document.getElementById('cf_targets_num_showing').innerText = ""+num_found;
    if ( num_found == 0) {
        var div = document.createElement("li");
        div.className = ""
        div.style = "padding: 1em"
        div.innerHTML = `I'm sorry, i don't have a code generator on board for the aspects you selected.
        Do you like to have a code generator that is still missing? Please let us know on the <a class="tc-maincolor-text" href="/feedback">feedback</a> page
        or right here: <form id="cff_feedback" action="/feedback/q" method="POST">
        <div class="row">
        <div class="input-field col s8">
        <input id="cff_feedback_what" name="cff_feedback_what" type="text" placeholder="I'm missing ...">
        </div>
        <div class="input-field col s1">
        <button id="cff_feedback_submit" type="submit" class="waves-effect waves-light tc-maincolor btn-floating tiny"> <i class="material-icons">send</i></button>
        </div>
        </div>
        </form>
        `;

        cf_targets_matching.appendChild(div);

        // enable button
        var cff_submit_btn = document.getElementById('cff_feedback_submit');
        cff_submit_btn.addEventListener('click', cff_feedback_submit);

    }
}

function cff_feedback_submit(e) {
    e.preventDefault();

    // disable button click event
    var cff_submit_btn = document.getElementById('cff_feedback_submit');
    cff_submit_btn.removeEventListener('click', cff_feedback_submit);

    // submit
    var frm = $('#cff_feedback');
    var response = "Uh oh, i'm not sure what happened. Please try again."
    $.ajax({
        type: frm.attr('method'),
        url: frm.attr('action'),
        data: frm.serialize(),
        async: true,
        success: function (data) {
            response = data;
        },
        error: function (data) {
            response = 'An error occured: '+data;
        },
        complete: function() {
            // remove form
            var cf_targets_matching = document.getElementById('cf_targets_matching');
            while (cf_targets_matching.firstChild) {
                cf_targets_matching.removeChild(cf_targets_matching.firstChild);
            }

            var div = document.createElement("li");
            div.className = ""
            div.style = "padding: 1em"
            div.innerHTML = response+
                "<br/>As for now, you might want to try out a different framework?"+
                "<br/>For updates on our service, please follow us on Twitter or GitHub.";
            cf_targets_matching.appendChild(div);
        },
    });

}

function cff_generator_clicked(e) {
    e.preventDefault()

    var d = document.getElementById('cf_selection_form');

    var sel = document.getElementById('cf_selection');
    sel.value = e.target.id.replace(/^go_/,"");

/*    var id = document.getElementById('cf_id');
    id.value = next_url.replace(/^.*\/app\/(.*)\/properties$/,"$1");
*/
    d.submit();
}

function cff_generator_details_clicked(e) {
    e.preventDefault()

    var win = window.open("/module/"+e.target.id.replace(/^details_/,""), '_tf_module');
    win.focus();
}

// run initially, produces an empty list if nothing is selected.
cf_lookup_matching_targets();


// connect vote buttons (max.20)
for( var i = 0; i < 20; i++) {
    var n = "vote_"+i+"-down";
    var b = document.getElementById(n);
    if ( b !== null && b !== undefined)
        b.addEventListener('click', cf_vote_down);
    n = "vote_"+i+"-up";
    b = document.getElementById(n);
    if ( b !== null && b !== undefined)
        b.addEventListener('click', cf_vote_up);
}
document.getElementById('vote_submit').addEventListener('click', cf_vote_submit);


function cf_vote_submit(e) {
    e.preventDefault();

    var s = document.getElementById('vote_submit');
    s.removeEventListener('click', cff_feedback_submit);
    s.className += " disabled";

    // disconnect/gray out buttons
    for( var i = 0; i < 20; i++) {
        var n = "vote_"+i+"-down";
        var b = document.getElementById(n);
        if ( b !== null && b !== undefined) {
            b.removeEventListener('click', cf_vote_down);
        }
        n = "vote_"+i+"-up";
        b = document.getElementById(n);
        if ( b !== null && b !== undefined) {
            b.removeEventListener('click', cf_vote_up);
        }

    }

    var frm = $('#cf_vote_form');
    $.ajax({
        type: frm.attr('method'),
        url: frm.attr('action'),
        data: frm.serialize(),
        async: true,
        complete: function() {
            var f = document.getElementById('cf_vote_form');

            f.innerHTML = "<span class=\"tc-maincolor-text\">Thank you for your feedback!</span>";
        }
    });
}

function cf_vote_up(e) {
    e.preventDefault();

    var id_ = e.target.id.split("-");
    if ( id_ != null && id_.length === 2) {
        var p = id_[0];

        var input_elem = document.getElementById(""+p+"-input");
        var v = parseInt(input_elem.value);
        if ( v >= 5) {
            return;
        }
        v++;
        input_elem.value = ""+v;

        for ( var j = 1; j < 6; j++) {
            document.getElementById(""+p+"-"+j).className = "col s1"+((j==v)?" tc-maincolor":"")
        }
    }
}

function cf_vote_down(e) {
    e.preventDefault();

    var id_ = e.target.id.split("-");
    if ( id_ != null && id_.length === 2) {
        var p = id_[0];

        var input_elem = document.getElementById(""+p+"-input");
        var v = parseInt(input_elem.value);
        if ( v <= 1) {
            return;
        }
        v--;
        input_elem.value = ""+v;

        for ( var j = 1; j < 6; j++) {
            document.getElementById(""+p+"-"+j).className = "col s1"+((j==v)?" tc-maincolor":"")
        }
    }
}

function more_activate(e) {
    document.getElementById('btn_more').removeEventListener('click', more_activate);
    document.getElementById('btn_less').addEventListener('click', more_deactivate);
    document.getElementById('sp_more').className = "";
    document.getElementById('btn_more').className += " hide";
}

function more_deactivate(e) {
    document.getElementById('btn_more').addEventListener('click', more_activate);
    document.getElementById('btn_less').removeEventListener('click', more_deactivate);
    document.getElementById('sp_more').className += " hide";
    document.getElementById('btn_more').className += "tc-maincolor-text";
}

function interest_activate(e) {
    document.getElementById('btn_interest').removeEventListener('click', more_activate);
    document.getElementById('span_btn_interest').innerHTML = "";
    document.getElementById('span_interest').className = "";
}

var btn_more = document.getElementById('btn_more');
if ( btn_more != null) {
    btn_more.addEventListener('click', more_activate);
}

var btn_intr = document.getElementById('btn_interest');
if ( btn_intr != null) {
    btn_intr.addEventListener('click', interest_activate);
}
