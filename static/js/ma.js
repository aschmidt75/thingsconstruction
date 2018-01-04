//
// ma = manage actions
//


// configure modal dialog
$("#ma_listitem_validation_modal").modal({
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

String.prototype.replaceAll = function(search, replacement) {
    var target = this;
    return target.replace(new RegExp(search, "g"), replacement);
};

// connect add(+) button of properties list to ma_list_add
document.getElementById("ma_add_btn").addEventListener("click", ma_list_add);

// adds an existing action object to the list.
function ma_list_add_existing(obj) {
    // add a new row to the ul #ma_list
    var ma_list = document.getElementById("ma_list");

    // extract the id of the last item in the list. The new div gets id+1
    var last = ma_list.lastElementChild;
    var last_id = 0;
    if ( last != null) {
        var last_ids = last.id.split("_");
        last_id = parseInt(last_ids[last_ids.length-1]);
    }

    var ma_list_item = document.createElement("li");
    ma_list_item.id = "ma_listitem_"+(last_id+1);
    ma_list_item.className = "collection-item row";
    // template contains ## for each place where new id shall fit in
    var tplhtml = document.getElementById("tpl_maf_list_item").innerHTML;
    ma_list_item.innerHTML = tplhtml
        .replaceAll("##", ""+(last_id+1));
    ma_list.appendChild(ma_list_item);

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
    var btn_edit = document.getElementById("ma_listitem_"+(last_id+1)+"_btns_edit");
    if ( obj == null) {
        // no content yet, directly go to edit mode
        ma_list_click_edit_showHide(btn_edit);
    }
    btn_edit.addEventListener("click", ma_list_click_edit);

    document.getElementById("ma_listitem_"+(last_id+1)+"_btns_delete").addEventListener("click", ma_list_click_delete);

    if ( obj != null) {
        document.getElementById("ma_listitem_"+(last_id+1)+"_name").innerHTML = obj.Name;
        document.getElementById("ma_listitem_"+(last_id+1)+"_edit_name").value = obj.Name;
        document.getElementById("ma_listitem_"+(last_id+1)+"_edit_desc").value = obj.Description;

        document.getElementById("ma_listitem_"+(last_id+1)+"_details").innerHTML = "<p><i>"+obj.Description+"</i></p>"+
            "<input type=\"text\" name=\"ma_listitem_"+(last_id+1)+"_val\" class=\"hide\" value=\""+obj.Name+"\">"+
            "<input type=\"text\" name=\"ma_listitem_"+(last_id+1)+"_desc\" class=\"hide\" value=\""+obj.Description+"\">";

    }
}

// add content of tpl_mpl_list_item to ma_list.
// Open "edit" part, connect buttons
function ma_list_add(e) {
    e.preventDefault();
    ma_list_add_existing(null);
}

function ma_list_click_edit(e) {
    e.preventDefault();
    var t = e.target;
    ma_list_click_edit_showHide(t)
}

// checks if n is already a name in the mp list.
// except the row with id=edit_id
function ma_list_is_name_in_use_except(n, edit_id) {
    if ( n === "") {
        return false;
    }
    var ma_list = document.getElementById("ma_list");
    var ma_list_items = ma_list.children;
    for ( var i = 0; i < ma_list_items.length; i++) {
        if ( ma_list_items[i].id !== edit_id) {
            var id = ""+ma_list_items[i].id+"_name";
            var n2 = document.getElementById(id).innerText;

            if ( n !== "" && n === n2) {
                return true;
            }
        }
    }
    return false;
}

function ma_list_has_items() {
    var ma_list = document.getElementById("ma_list");
    var ma_list_items = ma_list.children;
    return ma_list_items.length > 0;
}

// disable all edit/delete buttons
function ma_list_disable_buttons() {
    var ma_list = document.getElementById("ma_list");
    var ma_list_items = ma_list.children;
    for ( var i = 0; i < ma_list_items.length; i++) {
        var id;
        id = ""+ma_list_items[i].id+"_btns_edit";
        document.getElementById(id).className += " disabled";
        id = ""+ma_list_items[i].id+"_btns_delete";
        document.getElementById(id).className += " disabled";
    }
    document.getElementById("ma_add_btn").className += " disabled";
    document.getElementById("ma_next").className += " disabled";
    document.getElementById("ma_prev").className += " disabled";
}

// removes the disabled class from an id
function ma_list_enable_button(id) {
    var cn = document.getElementById(id).className;
    var cns = cn.split(" ");
    cn = "";
    for ( var j = 0; j < cns.length; j++) {
        if (cns[j] !== "disabled") {
            cn += cns[j]+" ";
        }
    }
    document.getElementById(id).className = cn;
}

// enable all edit/delete buttons
function ma_list_enable_buttons() {
    var ma_list = document.getElementById("ma_list");
    var ma_list_items = ma_list.children;
    for ( var i = 0; i < ma_list_items.length; i++) {
        var id;
        id = ""+ma_list_items[i].id+"_btns_edit";
        ma_list_enable_button(id);
        id = ""+ma_list_items[i].id+"_btns_delete";
        ma_list_enable_button(id);
    }
    ma_list_enable_button("ma_add_btn");
    ma_list_enable_button("ma_next");
    ma_list_enable_button("ma_prev");
}

// given the node of the edit button, this method
// toggles the edit fields and show fields.
function ma_list_click_edit_showHide(t) {

    // depending on browser, t can be the <button> or the <i> holding the icon.
    if ( t.localName === "i") {
        t = t.parentElement;
    }

    // Look whether this element is in "show" or in "edit" mode.
    // switch between two, do logic.
    var btn_text = t.children[0].innerText;
    var row, showPart, editPart;
    if (btn_text === "mode_edit" || btn_text === "MODE_EDIT") {
        // Button shows edit icon, so it"s in show mode.

        // change icon of button
        t.children[0].innerText = "check";

        row = t.parentElement.parentElement.parentElement;

        // hide "show" part
        showPart = document.getElementById(""+row.id+"_show");
        showPart.className += " hide";
        // unhide "edit" part
        editPart = document.getElementById(""+row.id+"_edit");
        editPart.className = " col s8";

        // disable edit/delete buttons for all others.
        ma_list_disable_buttons();
        // enable the save button again
        ma_list_enable_button(t.id);
    }
    if (btn_text === "check" || btn_text === "CHECK") {
        // Button shows save icon, so it"s in edit mode, and user wants to save

        row = t.parentElement.parentElement.parentElement;
        var newName = document.getElementById(""+row.id+"_edit_name").value;

        // validate fields. In case of error, color items, open modal, exit.
        // 1. name
        if (newName === "") {
            document.getElementById(""+row.id+"_edit_name").style.borderBottom = "solid 2px #ee2222";
            $("#ma_listitem_validation_modal").modal("open");
            return;
        }

        // 2. name must be unique, compare to all others
        if ( ma_list_is_name_in_use_except(newName, row.id)) {
            document.getElementById(""+row.id+"_edit_name").style.borderBottom = "solid 2px #ee2222";
            document.getElementById("ma_listitem_validation_modal_reason").innerText =
                "The name "+newName+" was already chosen for another property.";
            $("#ma_listitem_validation_modal").modal("open");
            return;
        }

        // if we get here, all is valid. un-color, continue.
        document.getElementById(""+row.id+"_edit_name").style.borderBottom = "";
        ma_list_enable_buttons();

        // change icon of button
        t.children[0].innerText = "mode_edit";

        // save values..
        var e, f, g;

        e = document.getElementById(""+row.id+"_name");
        f = document.getElementById(""+row.id+"_edit_name");
        e.innerText =  f.value;
        e = document.getElementById(""+row.id+"_details");
        g = document.getElementById(""+row.id+"_edit_desc");
        e.innerHTML = "<p><i>"+g.value+"</i></p>"+
            "<input type=\"text\" name=\""+row.id+"_val\" class=\"hide\" value=\""+f.value+"\">"+
            "<input type=\"text\" name=\""+row.id+"_desc\" class=\"hide\" value=\""+g.value+"\">";

        // hide "edit" part
        showPart = document.getElementById(""+row.id+"_show");
        showPart.className = " col s8";
        // unhide "show" part
        editPart = document.getElementById(""+row.id+"_edit");
        editPart.className += " hide";

    }
}

// delete an item from ma_list
function ma_list_click_delete(e) {
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

// configure modal dialog
$("#details_validation_modal").modal({
        dismissible: true, // Modal can be dismissed by clicking outside of the modal
        opacity: .5, // Opacity of modal background
        inDuration: 300, // Transition in duration
        outDuration: 200, // Transition out duration
        startingTop: "4%", // Starting top style attribute
        endingTop: "10%" // Ending top style attribute
    }
);

document.getElementById("ma_next").addEventListener("click", ma_submit);
document.getElementById("ma_prev").addEventListener("click", ma_to_properties);

function ma_disable_navbtns() {
    var btn1 = document.getElementById("ma_next");
    var btn2 = document.getElementById("ma_prev");
    btn1.className += " disabled";
    btn2.className += " disabled"

}
function ma_enable_navbtns() {
    var btn1 = document.getElementById("ma_next");
    var btn2 = document.getElementById("ma_prev");
    btn1.className.replaceAll("disabled", "");
    btn2.className.replaceAll("disabled", "");
}

function ma_to_properties(e) {
    e.preventDefault();

    var details_ok = true;

    /*    var errors = "<ul>";
        if ( !ma_list_has_items()) {
            errors += "<li>You should at least have one property.</li>";
            details_ok = false
        }
        errors += "</ul>";
    */
    if (details_ok) {
        //
        var frm = document.getElementById("maf");
        ma_disable_navbtns();

        $.ajax({
            type: frm.method,
            url: frm.action,
            async: true,
            success: function (data) {
            },
            error: function (data) {
                //
                console.log(data);

                // stay on page
                ma_enable_navbtns();
            },
            complete: function() {
                // redirect to next page
                var url = document.URL;
                window.location.replace(url.replace(/^(.*)\/actions.*/,"$1/properties"));
            }
        });

    } else {
        document.getElementById("details_validation_modal_reason").innerHTML = errors;
        $("#details_validation_modal").modal("open");
    }
}

function ma_submit(e) {
    e.preventDefault();

    var details_ok = true;

/*    var errors = "<ul>";
    if ( !ma_list_has_items()) {
        errors += "<li>You should at least have one property.</li>";
        details_ok = false
    }
    errors += "</ul>";
*/
    if (details_ok) {
        //
        var frm = document.getElementById("maf");
        ma_disable_navbtns();

        $.ajax({
            type: frm.method,
            url: frm.action,
            async: true,
            success: function (data) {
            },
            error: function (data) {
                //
                console.log(data);

                // stay on page
                ma_enable_navbtns();
            },
            complete: function() {
                // redirect to next page
                var url = document.URL;
                window.location.replace(url.replace(/^(.*)\/actions.*/,"$1/events"));
            }
        });

    } else {
        document.getElementById("details_validation_modal_reason").innerHTML = errors;
        $("#details_validation_modal").modal("open");
    }
}
