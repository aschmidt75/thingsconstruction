//
// mp = manage properties
//

// configure modal dialog
$("#mp_listitem_validation_modal").modal({
        dismissible: true, // Modal can be dismissed by clicking outside of the modal
        opacity: .5, // Opacity of modal background
        inDuration: 300, // Transition in duration
        outDuration: 200, // Transition out duration
        startingTop: "4%", // Starting top style attribute
        endingTop: "10%", // Ending top style attribute
    }
);
// disable enter to prevent submitting the wrong form
$("html").bind("keypress", function(e)
{
    if(e.keyCode === 13)
    {
        return false;
    }
});

$.encoder.init();

String.prototype.replaceAll = function(search, replacement) {
    var target = this;
    return target.replace(new RegExp(search, "g"), replacement);
};

// connect add(+) button of properties list to mp_list_add
document.getElementById("mp_add_btn").addEventListener("click", mp_list_add);


function mp_list_add_existing(obj) {
    // add a new row to the ul #mp_list
    var mp_list = document.getElementById("mp_list");

    // extract the id of the last item in the list. The new div gets id+1
    var last = mp_list.lastElementChild;
    var last_id = 0;
    if ( last != null) {
        var last_ids = last.id.split("_");
        last_id = parseInt(last_ids[last_ids.length-1]);
    }

    var mp_list_item = document.createElement("li");
    mp_list_item.id = "mp_listitem_"+(last_id+1);
    mp_list_item.className = "collection-item row";

    // template contains ## for each place where new id shall fit in
    var tplhtml = document.getElementById("tlp_mpl_list_item").innerHTML;
    mp_list_item.innerHTML = tplhtml
        .replaceAll("##", ""+(last_id+1));
    mp_list.appendChild(mp_list_item);

    // update toc data
    try {
        var items = document.getElementById("toc_current_info").innerText;
        document.getElementById("toc_current_info").innerText = ""+(Number(items)+1);
    } catch (e) {
        document.getElementById("toc_current_info").innerText = "-";
    }

    // activate dynamic type elements within
    $(document).ready(function(){
        $("ul.tabs").tabs();
    });
    $(document).ready(function() {
        $("select").material_select();
    });

    // attach buttons to handlers
    var btn_edit = document.getElementById("mp_listitem_"+(last_id+1)+"_btns_edit");
    if ( obj === null) {
        // no content yet, directly go to edit mode
        mp_list_click_edit_showHide(btn_edit);
    }
    btn_edit.addEventListener("click", mp_list_click_edit);

    var tab_item;
    tab_item = document.getElementById("mp_listitem_"+(last_id+1)+"_type_bool_click");
    tab_item.addEventListener("click", mp_listitem_typetab_click);
    tab_item = document.getElementById("mp_listitem_"+(last_id+1)+"_type_number_click");
    tab_item.addEventListener("click", mp_listitem_typetab_click);
    tab_item = document.getElementById("mp_listitem_"+(last_id+1)+"_type_str_click");
    tab_item.addEventListener("click", mp_listitem_typetab_click);

    document.getElementById("mp_listitem_"+(last_id+1)+"_btns_delete").addEventListener("click", mp_list_click_delete);

    if ( obj) {
        var ehn = $.encoder.encodeForHTML(obj.name);
        var ehd = $.encoder.encodeForHTML(obj.description);
        document.getElementById("mp_listitem_"+(last_id+1)+"_name").innerHTML = ehn;
        document.getElementById("mp_listitem_"+(last_id+1)+"_edit_name").value = ehn;
        document.getElementById("mp_listitem_"+(last_id+1)+"_edit_desc").value = ehd;

        var tabBool = document.getElementById("mp_listitem_"+(last_id+1)+"_type_bool");
        var tabNumber = document.getElementById("mp_listitem_"+(last_id+1)+"_type_number");
        var tabString = document.getElementById("mp_listitem_"+(last_id+1)+"_type_str");
        tabBool.className = "col";
        tabNumber.className = "col";
        tabString.className = "col";
        var tabBoolA = document.getElementById("mp_listitem_a_"+(last_id+1)+"_type_bool_click");
        var tabNumberA = document.getElementById("mp_listitem_a_"+(last_id+1)+"_type_number_click");
        var tabStringA = document.getElementById("mp_listitem_a_"+(last_id+1)+"_type_str_click");
        tabBoolA.className = "";
        tabNumberA.className = "";
        tabStringA.className = "";

        var typeStr = obj.type;
        var typeStrForm;
        if (obj.type === "Boolean") {
            tabBool.className = "col active";
            tabBoolA.className = "active";
            typeStrForm = "b";
        }
        if (obj.type === "Float") {
            var selectId = "mp_listitem_"+(last_id+1)+"_type_number_type";
            document.getElementById(selectId).value = "Float";
            $("#"+selectId).find('option[value="2"]').prop('selected', true);
            $("#"+selectId).material_select();
            tabNumber.className = "col active";
            tabNumberA.className = "active";
            typeStrForm = "f";
            if (obj.min !== undefined) {
                document.getElementById("mp_listitem_"+(last_id+1)+"_type_number_min").value = $.encoder.encodeForHTML(obj.min);
                typeStrForm += ";"+obj.min;
            }
            if (obj.max !== undefined) {
                document.getElementById("mp_listitem_"+(last_id+1)+"_type_number_max").value = $.encoder.encodeForHTML(obj.max);
                typeStrForm += ";"+obj.max;
            }
            if (obj.min !== undefined && obj.max !== undefined) {
                typeStr = typeStr + "["+$.encoder.encodeForHTML(obj.min)+".."+$.encoder.encodeForHTML(obj.max)+"]";
            }
        }
        if (obj.type === "Integer") {
            var selectId = "mp_listitem_"+(last_id+1)+"_type_number_type";
            document.getElementById(selectId).value = "Integer";
            $("#"+selectId).find('option[value="1"]').prop('selected', true);
            $("#"+selectId).material_select();
            tabNumber.className = "col active";
            tabNumberA.className = "active";
            typeStrForm = "i";
            if (obj.min !== undefined) {
                document.getElementById("mp_listitem_"+(last_id+1)+"_type_number_min").value = $.encoder.encodeForHTML(obj.min);
                typeStrForm += ";"+obj.min;
            }
            if (obj.max !== undefined) {
                document.getElementById("mp_listitem_"+(last_id+1)+"_type_number_max").value = $.encoder.encodeForHTML(obj.max);
                typeStrForm += ";"+obj.max;
            }
            if (obj.min !== undefined && obj.max !== undefined) {
                typeStr = typeStr + "["+$.encoder.encodeForHTML(obj.min)+".."+$.encoder.encodeForHTML(obj.max)+"]";
            }
        }
        if (obj.type === "String") {
            tabString.className = "col active";
            tabStringA.className = "active";
            typeStrForm = "s";
            if (obj.maxLength !== undefined) {
                var i = $.encoder.encodeForHTML(obj.maxLength);
                document.getElementById("mp_listitem_"+(last_id+1)+"_type_str_maxlength").value = i;
                typeStrForm += ";"+i;
                typeStr = typeStr + "["+i+"]";
            }
        }

        document.getElementById("mp_listitem_"+(last_id+1)+"_details").innerHTML =
            "<p>"+$.encoder.encodeForHTML(typeStr)+"<br/><i>"+$.encoder.encodeForHTML(obj.description)+"</i></p>"+
            "<input type=\"text\" name=\"mp_listitem_"+(last_id+1)+"_val\" class=\"hide\" "+$.encoder.encodeForHTMLAttribute("value",""+obj.name+";"+typeStrForm)+">"+
            "<input type=\"text\" name=\"mp_listitem_"+(last_id+1)+"_desc\" class=\"hide\" "+$.encoder.encodeForHTMLAttribute("value",""+obj.description)+">";

    }
}

// only allow inputs which would make up a valid identifier in
// code.
function mp_list_limit_name(e) {
    var k = e.key;
    if ( e.charCode === 0 && e.keyCode !== 0) return true;
    if ( k >= '0' && k <= '9') return true;
    if ( k >= 'a' && k <= 'z') return true;
    if ( k >= 'A' && k <= 'Z') return true;
    if ( k === '_' || k === '-') return true;
    e.preventDefault();
}

function mp_list_limit_int(e) {
    var k = e.key;
    if ( e.charCode === 0 && e.keyCode !== 0) return true;
    if ( k >= '0' && k <= '9') return true;
    if ( k === '-') return true;
    e.preventDefault();
}

function mp_list_limit_uint(e) {
    var k = e.key;
    if ( e.charCode === 0 && e.keyCode !== 0) return true;
    if ( k >= '0' && k <= '9') return true;
    e.preventDefault();
}

function mp_list_limit_float(e) {
    var k = e.key;
    if ( e.charCode === 0 && e.keyCode !== 0) return true;
    if ( k >= '0' && k <= '9') return true;
    if ( k === '-') return true;
    if ( k === '.') return true;
    e.preventDefault();
}

// add content of tpl_mpl_list_item to mp_list.
// Open "edit" part, connect buttons
function mp_list_add(e) {
    e.preventDefault();
    // add a new row to the ul #mp_list
    var mp_list = document.getElementById("mp_list");

    // extract the id of the last item in the list. The new div gets id+1
    var last = mp_list.lastElementChild;
    var last_id = 0;
    if ( last != null) {
        var last_ids = last.id.split("_");
        last_id = parseInt(last_ids[last_ids.length-1]);
    }

    var mp_list_item = document.createElement("li");
    mp_list_item.id = "mp_listitem_"+(last_id+1);
    mp_list_item.className = "collection-item row";
    // template contains ## for each place where new id shall fit in
    var tplhtml = document.getElementById("tlp_mpl_list_item").innerHTML;
    mp_list_item.innerHTML = tplhtml
        .replaceAll("##", ""+(last_id+1));
    mp_list.appendChild(mp_list_item);

    // activate dynamic type elements within
    $(document).ready(function(){
        $("ul.tabs").tabs();
    });
    $(document).ready(function() {
        $("select").material_select();
    });

    // attach buttons to handlers
    var btn_edit = document.getElementById("mp_listitem_"+(last_id+1)+"_btns_edit");
    mp_list_click_edit_showHide(btn_edit);   // directly go to edit mode
    btn_edit.addEventListener("click", mp_list_click_edit);

    var tab_item;
    tab_item = document.getElementById("mp_listitem_"+(last_id+1)+"_type_bool_click");
    tab_item.addEventListener("click", mp_listitem_typetab_click);
    tab_item = document.getElementById("mp_listitem_"+(last_id+1)+"_type_number_click");
    tab_item.addEventListener("click", mp_listitem_typetab_click);
    tab_item = document.getElementById("mp_listitem_"+(last_id+1)+"_type_str_click");
    tab_item.addEventListener("click", mp_listitem_typetab_click);

    // attach keyinput handlers
    document.getElementById("mp_listitem_"+(last_id+1)+"_type_str_maxlength").addEventListener("keypress", mp_list_limit_uint);
    document.getElementById("mp_listitem_"+(last_id+1)+"_type_number_min").addEventListener("keypress", mp_list_limit_float);
    document.getElementById("mp_listitem_"+(last_id+1)+"_type_number_max").addEventListener("keypress", mp_list_limit_float);
    document.getElementById("mp_listitem_"+(last_id+1)+"_edit_name").addEventListener("keypress", mp_list_limit_name);

    document.getElementById("mp_listitem_"+(last_id+1)+"_btns_delete").addEventListener("click", mp_list_click_delete);

    mp_list_update_count_in_toc();

    Materialize.updateTextFields();

}

function mp_listitem_typetab_click(e) {
    var x = e.target.attributes["href"].value;
    x = x.slice(1);
    var a = x.split("_");
    var target = "";
    for ( var i = 0; i < a.length-1; i++) {
        target += a[i]+"_";
    }
    document.getElementById(target).value = x
}

function mp_list_click_edit(e) {
    e.preventDefault();
    var t = e.target;
    mp_list_click_edit_showHide(t)
}

// checks if n is already a name in the mp list.
// except the row with id=edit_id
function mp_list_is_name_in_use_except(n, edit_id) {
    //console.log("Checking for uniqueness of: "+n+", except for id: "+edit_id);
    if ( n === "") {
        return false;
    }
    mp_list = document.getElementById("mp_list");
    mp_list_items = mp_list.children;
    for ( var i = 0; i < mp_list_items.length; i++) {
        if ( mp_list_items[i].id !== edit_id) {
            var id = ""+mp_list_items[i].id+"_name";
            var n2 = document.getElementById(id).innerText;

            if ( n !== "" && n === n2) {
                return true;
            }
        }
    }
    return false;
}

function mp_list_has_items() {
    var mp_list = document.getElementById("mp_list");
    var mp_list_items = mp_list.children;
    return mp_list_items.length > 0;
}

// disable all edit/delete buttons
function mp_list_disable_buttons() {
    var mp_list = document.getElementById("mp_list");
    var mp_list_items = mp_list.children;
    for ( var i = 0; i < mp_list_items.length; i++) {
        var id;
        id = ""+mp_list_items[i].id+"_btns_edit";
        document.getElementById(id).className += " disabled";
        id = ""+mp_list_items[i].id+"_btns_delete";
        document.getElementById(id).className += " disabled";
    }
    document.getElementById("mp_add_btn").className += " disabled";
    document.getElementById("mp_next").className += " disabled";
    document.getElementById("mp_prev").className += " disabled";
}

// removes the disabled class from an id
function mp_list_enable_button(id) {
    var cn = document.getElementById(id).className;
    document.getElementById(id).className = cn.replaceAll("disabled", "");
}

// enable all edit/delete buttons
function mp_list_enable_buttons() {
    var mp_list = document.getElementById("mp_list");
    var mp_list_items = mp_list.children;
    for ( var i = 0; i < mp_list_items.length; i++) {
        var id;
        id = ""+mp_list_items[i].id+"_btns_edit";
        mp_list_enable_button(id);
        id = ""+mp_list_items[i].id+"_btns_delete";
        mp_list_enable_button(id);
    }
    mp_list_enable_button("mp_add_btn");
    mp_list_enable_button("mp_next");
    mp_list_enable_button("mp_prev");
}

// given the node of the edit button, this method
// toggles the edit fields and show fields.
function mp_list_click_edit_showHide(t) {

    // depending on browser, t can be the <button> or the <i> holding the icon.
    if ( t.localName === "i") {
        t = t.parentElement;
    }

    var row, showPart, editPart;

    // Look whether this element is in "show" or in "edit" mode.
    // switch between two, do logic.
    var btn_text = t.children[0].innerText;
    if (btn_text === "mode_edit" || btn_text === "MODE_EDIT") {
        // Button shows edit icon, so it"s in show mode.

        // change icon of button
        t.children[0].innerText = "check"

        row = t.parentElement.parentElement.parentElement;

        // hide "show" part
        showPart = document.getElementById(""+row.id+"_show");
        showPart.className += " hide";
        // unhide "edit" part
        editPart = document.getElementById(""+row.id+"_edit");
        editPart.className = " col s8";

        // disable edit/delete buttons for all others.
        mp_list_disable_buttons();
        // enable the save button again
        mp_list_enable_button(t.id);
        // TODO: locate delete button, enable
    }
    if (btn_text === "check" || btn_text === "CHECK") {
        // Button shows save icon, so it"s in edit mode, and user wants to save

        row = t.parentElement.parentElement.parentElement;
        var newName = document.getElementById(""+row.id+"_edit_name").value;
        newName = $.encoder.encodeForHTML( $.encoder.canonicalize(newName));

        // validate fields. In case of error, color items, open modal, exit.
        // 1. name
        if (newName === "") {
            document.getElementById(""+row.id+"_edit_name").style.borderBottom = "solid 2px #ee2222";
            $("#mp_listitem_validation_modal").modal("open");
            return;
        }

        // 2. name must be unique, compare to all others
        if ( mp_list_is_name_in_use_except(newName, row.id)) {
            document.getElementById(""+row.id+"_edit_name").style.borderBottom = "solid 2px #ee2222";
            document.getElementById("mp_listitem_validation_modal_reason").innerText =
                "The name "+newName+" was already chosen for another property.";
            $("#mp_listitem_validation_modal").modal("open");
            return;
        }

        // 3. sanitize check
        try {
            f = $.encoder.canonicalize(document.getElementById(""+row.id+"_edit_name").value);
            $.encoder.encodeForHTML(f);

            e = document.getElementById(""+row.id+"_details");
            g = $.encoder.canonicalize(document.getElementById(""+row.id+"_edit_desc").value);
            $.encoder.encodeForHTML(g);

            $.encoder.encodeForHTMLAttribute("value", f);
            $.encoder.encodeForHTMLAttribute("value", g);
        } catch (e) {
            document.getElementById("mp_listitem_validation_modal_reason").innerText = e;
            $("#mp_listitem_validation_modal").modal("open");
            return;
        }

        // if we get here, all is valid. un-color, continue.
        document.getElementById(""+row.id+"_edit_name").style.borderBottom = "";
        mp_list_enable_buttons();

        // change icon of button
        t.children[0].innerText = "mode_edit";

        // save values..
        var e, f, g;

        try {

            e = document.getElementById(""+row.id+"_name");
            f = $.encoder.canonicalize(document.getElementById(""+row.id+"_edit_name").value);
            e.textContent = $.encoder.encodeForHTML(f);

            e = document.getElementById(""+row.id+"_details");
            g = $.encoder.canonicalize(document.getElementById(""+row.id+"_edit_desc").value);

            // pull out type
            var typeStr = "n/a";
            var typeStrForm = "";
            // look what tab is active, this determines the type. extract data for
            // both presentation and form.
            var te = document.getElementById(""+row.id+"_type_bool");
            console.log(te)
            if ( te != null && te.className === "col active") {
                typeStr = "Boolean";
                typeStrForm = "b";
            }
            te = document.getElementById(""+row.id+"_type_number");
            if ( te != null && te.className === "col active") {
                var k = document.getElementById(""+row.id+"_type_number_type");
                typeStr = "Number";
                if (k.selectedIndex === 0) {
                    typeStr = "Number: Integer";
                    typeStrForm = "i";
                }
                if (k.selectedIndex === 1) {
                    typeStr = "Number: Float";
                    typeStrForm = "f";
                }
                var i = document.getElementById(""+row.id+"_type_number_min").value;
                i = $.encoder.canonicalize(i);
                var j = document.getElementById(""+row.id+"_type_number_max").value;
                j = $.encoder.canonicalize(j);
                if ( i !== "" || j !== "") {
                    typeStr = typeStr + "["+i+".."+j+"]";
                    typeStrForm = typeStrForm + ";"+i+";"+j;
                }
            }
            te = document.getElementById(""+row.id+"_type_str");
            if ( te != null && te.className === "col active") {
                typeStr = "String";
                typeStrForm = "s";
                var i = document.getElementById(""+row.id+"_type_str_maxlength").value;
                i = $.encoder.canonicalize(i);
                if ( i != null && i.value !== "") {
                    typeStr = typeStr + "["+i+"]";
                    typeStrForm = "s;"+i;
                }
            }
            e.innerHTML = "<p>"+typeStr+"<br/><i>"+$.encoder.encodeForHTML(g)+"</i></p>"+
                "<input type=\"text\" name=\""+row.id+"_val\" class=\"hide\" "+$.encoder.encodeForHTMLAttribute("value", f+";"+typeStrForm)+">"+
                "<input type=\"text\" name=\""+row.id+"_desc\" class=\"hide\" "+$.encoder.encodeForHTMLAttribute("value", g)+">";
        } catch (e) {
            console.log(e);
        }

        // hide "edit" part
        showPart = document.getElementById(""+row.id+"_show");
        showPart.className = " col s8";
        // unhide "show" part
        editPart = document.getElementById(""+row.id+"_edit");
        editPart.className += " hide";

    }
}

// delete an item from mp_list
function mp_list_click_delete(e) {
    e.preventDefault();
    var t = e.target;

    // depending on browser, t can be the <button> or the <i> holding the icon.
    if ( t.localName === "i") {
        t = t.parentElement;
    }

    var data_row = t.parentNode.parentNode.parentNode;

    // remove the whole row
    data_row.parentNode.removeChild(data_row);

    // update toc data
    try {
        var items = document.getElementById("toc_current_info").innerText;
        document.getElementById("toc_current_info").innerText = ""+(Number(items)-1);
    } catch (e) {
        document.getElementById("toc_current_info").innerText = "-";
    }

}

function mp_list_update_count_in_toc() {
}




// details = page with properties, actions, events.

// configure modal dialog
$("#details_validation_modal").modal({
        dismissible: true, // Modal can be dismissed by clicking outside of the modal
        opacity: .5, // Opacity of modal background
        inDuration: 300, // Transition in duration
        outDuration: 200, // Transition out duration
        startingTop: "4%", // Starting top style attribute
        endingTop: "10%", // Ending top style attribute
    }
);
$("#details_validation_modal").modal();

document.getElementById("mp_next").addEventListener("click", submit_mpf);
document.getElementById("mp_prev").addEventListener("click", mp_to_framework);


function mp_disable_navbtns() {
    var btn1 = document.getElementById("mp_next");
    var btn2 = document.getElementById("mp_prev");
    btn1.className += " disabled";
    btn2.className += " disabled"

}
function mp_enable_navbtns() {
    var btn1 = document.getElementById("mp_next");
    var btn2 = document.getElementById("mp_prev");
    btn1.className.replaceAll("disabled", "")
    btn2.className.replaceAll("disabled", "")
}

function mp_to_framework(e) {
    e.preventDefault();

    var details_ok = true;

/*    var errors = "<ul>"
    if ( !mp_list_has_items()) {
        errors += "<li>You should at least have one property.</li>"
        details_ok = false
    }
    errors += "</ul>"
*/

    if (details_ok) {
        //
        var frm = document.getElementById("mpf");
        mp_disable_navbtns();

        $.ajax({
            type: frm.method,
            url: frm.action,
            data: $("#mpf").serialize(),
            success: function (data) {
                // redirect to next page
                var url = document.URL;
                window.location.replace(url.replace(/^(.*)\/properties.*/,"$1/framework"));
            },
            error: function (data) {
                // stay on page
                ma_enable_navbtns();

                document.getElementById("details_validation_modal_reason").textContent =
                    "An error occured while saving your data.";
                $("#details_validation_modal").modal("open");
            },
        });
    } else {
        document.getElementById("details_validation_modal_reason").innerHTML = errors;
        $("#details_validation_modal").modal("open");
    }
}

function submit_mpf(e) {
    e.preventDefault();

    var details_ok = true;

/*    var errors = "<ul>"
    if ( !mp_list_has_items()) {
        errors += "<li>You should at least have one property.</li>"
        details_ok = false
    }
    errors += "</ul>"
*/

    if (details_ok) {
        //
        var frm = document.getElementById("mpf");
        mp_disable_navbtns();

        $.ajax({
            type: frm.method,
            url: frm.action,
            data: $("#mpf").serialize(),
            success: function (data) {
                // redirect to next page
                var url = document.URL;
                window.location.replace(url.replace(/^(.*)\/properties.*/,"$1/actions"));
            },
            error: function (data) {
                // stay on page
                ma_enable_navbtns();

                document.getElementById("details_validation_modal_reason").textContent =
                    "An error occured while saving your data.";
                $("#details_validation_modal").modal("open");
            },
        });
    } else {
        document.getElementById("details_validation_modal_reason").innerHTML = errors;
        $("#details_validation_modal").modal("open");
    }
}

