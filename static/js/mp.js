//
// mp = manage properties
//

// configure modal dialog
$('#mp_listitem_validation_modal').modal({
        dismissible: true, // Modal can be dismissed by clicking outside of the modal
        opacity: .5, // Opacity of modal background
        inDuration: 300, // Transition in duration
        outDuration: 200, // Transition out duration
        startingTop: '4%', // Starting top style attribute
        endingTop: '10%', // Ending top style attribute
    }
);
$('#mp_listitem_validation_modal').modal();

String.prototype.replaceAll = function(search, replacement) {
    var target = this;
    return target.replace(new RegExp(search, 'g'), replacement);
};

// connect add(+) button of properties list to mp_list_add
document.getElementById('mp_add_btn').addEventListener('click', mp_list_add);

// add content of tpl_mpl_list_item to mp_list.
// Open "edit" part, connect buttons
function mp_list_add(e) {
    console.log(e)
    e.preventDefault();

    // add a new row to the ul #mp_list
    var mp_list = document.getElementById('mp_list');

    // extract the id of the last item in the list. The new div gets id+1
    var last = mp_list.lastElementChild;
    var last_id = 0;
    if ( last != null) {
        var last_ids = last.id.split('_');
        last_id = parseInt(last_ids[last_ids.length-1]);
    }

    var mp_list_item = document.createElement('li');
    mp_list_item.id = 'mp_listitem_'+(last_id+1);
    mp_list_item.className = "collection-item row";
    // template contains ## for each place where new id shall fit in
    var tplhtml = document.getElementById('tlp_mpl_list_item').innerHTML;
    console.log(tplhtml)
    mp_list_item.innerHTML = tplhtml
        .replaceAll("##", ''+(last_id+1));
    mp_list.appendChild(mp_list_item);

    // activate dynamic type elements within
    $(document).ready(function(){
        $('ul.tabs').tabs();
    });
    $(document).ready(function() {
        $('select').material_select();
    });

    // attach buttons to handlers
    var btn_edit = document.getElementById('mp_listitem_'+(last_id+1)+'_btns_edit');
    mp_list_click_edit_showHide(btn_edit);   // directly go to edit mode
    btn_edit.addEventListener('click', mp_list_click_edit);

    var tab_item;
    tab_item = document.getElementById('mp_listitem_'+(last_id+1)+'_type_bool_click');
    tab_item.addEventListener('click', mp_listitem_typetab_click)
    tab_item = document.getElementById('mp_listitem_'+(last_id+1)+'_type_number_click');
    tab_item.addEventListener('click', mp_listitem_typetab_click)
    tab_item = document.getElementById('mp_listitem_'+(last_id+1)+'_type_str_click');
    tab_item.addEventListener('click', mp_listitem_typetab_click)

    document.getElementById('mp_listitem_'+(last_id+1)+'_btns_delete').addEventListener('click', mp_list_click_delete);

    mp_list_update_count_in_toc();
    Materialize.updateTextFields();
}

function mp_listitem_typetab_click(e) {
    var x = e.target.attributes["href"].value;
    x = x.slice(1);
    console.log(x)
    var a = x.split('_')
    var target = "";
    for ( var i = 0; i < a.length-1; i++) {
        target += a[i]+"_";
    }
    console.log(target)
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
    console.log("Checking for uniqueness of: "+n+", except for id: "+edit_id);
    if ( n == "") {
        return false;
    }
    mp_list = document.getElementById('mp_list');
    mp_list_items = mp_list.children;
    for ( var i = 0; i < mp_list_items.length; i++) {
        if ( mp_list_items[i].id != edit_id) {
            var id = ""+mp_list_items[i].id+"_name";
            var n2 = document.getElementById(id).innerText;

            if ( n != "" && n == n2) {
                return true;
            }
        }
    }
    return false;
}

function mp_list_has_items() {
    var p_list = document.getElementById('mp_list');
    var mp_list_items = mp_list.children;
    return mp_list_items.length > 0
}

// disable all edit/delete buttons
function mp_list_disable_buttons() {
    var mp_list = document.getElementById('mp_list');
    var mp_list_items = mp_list.children;
    for ( var i = 0; i < mp_list_items.length; i++) {
        var id;
        id = ""+mp_list_items[i].id+"_btns_edit";
        document.getElementById(id).className += " disabled";
        id = ""+mp_list_items[i].id+"_btns_delete";
        document.getElementById(id).className += " disabled";
    }
    document.getElementById('mp_add_btn').className += " disabled";
    document.getElementById('details_next').className += " disabled";
}

// removes the disabled class from an id
function mp_list_enable_button(id) {
    var cn = document.getElementById(id).className;
    var cns = cn.split(" ");
    cn = "";
    for ( var j = 0; j < cns.length; j++) {
        if (cns[j] != "disabled") {
            cn += cns[j]+" ";
        }
    }
    document.getElementById(id).className = cn;
}

// enable all edit/delete buttons
function mp_list_enable_buttons() {
    var mp_list = document.getElementById('mp_list');
    var mp_list_items = mp_list.children;
    for ( var i = 0; i < mp_list_items.length; i++) {
        var id;
        id = ""+mp_list_items[i].id+"_btns_edit";
        mp_list_enable_button(id);
        id = ""+mp_list_items[i].id+"_btns_delete";
        mp_list_enable_button(id);
    }
    mp_list_enable_button('mp_add_btn')
    mp_list_enable_button('details_next')
}

// given the node of the edit button, this method
// toggles the edit fields and show fields.
function mp_list_click_edit_showHide(t) {

    // depending on browser, t can be the <button> or the <i> holding the icon.
    if ( t.localName == "i") {
        t = t.parentElement;
    }

    // Look whether this element is in "show" or in "edit" mode.
    // switch between two, do logic.
    var btn_text = t.children[0].innerText;
    if (btn_text == "mode_edit" || btn_text == "MODE_EDIT") {
        // Button shows edit icon, so it's in show mode.

        // change icon of button
        t.children[0].innerText = 'check'

        var row = t.parentElement.parentElement.parentElement;

        // hide "show" part
        var showPart = document.getElementById(''+row.id+"_show");
        showPart.className += " hide";
        // unhide "edit" part
        var editPart = document.getElementById(''+row.id+"_edit");
        editPart.className = " col s8";

        // disable edit/delete buttons for all others.
        mp_list_disable_buttons();
        // enable the save button again
        mp_list_enable_button(t.id);
        // TODO: locate delete button, enable
    }
    if (btn_text == "check" || btn_text == "CHECK") {
        // Button shows save icon, so it's in edit mode, and user wants to save

        var row = t.parentElement.parentElement.parentElement;
        var newName = document.getElementById(''+row.id+"_edit_name").value;

        // validate fields. In case of error, color items, open modal, exit.
        // 1. name
        if (newName == "") {
            document.getElementById(''+row.id+"_edit_name").style.borderBottom = 'solid 2px #ee2222'
            $('#mp_listitem_validation_modal').modal('open');
            return;
        };

        // 2. name must be unique, compare to all others
        if ( mp_list_is_name_in_use_except(newName, row.id)) {
            document.getElementById(''+row.id+"_edit_name").style.borderBottom = 'solid 2px #ee2222'
            document.getElementById('mp_listitem_validation_modal_reason').innerText =
                "The name "+newName+" was already chosen for another property."
            $('#mp_listitem_validation_modal').modal('open');
            return;
        }

        // if we get here, all is valid. un-color, continue.
        document.getElementById(''+row.id+"_edit_name").style.borderBottom = ''
        mp_list_enable_buttons();

        // change icon of button
        t.children[0].innerText = 'mode_edit'

        // save values..
        var e, f, g;

        e = document.getElementById(''+row.id+"_name");
        f = document.getElementById(''+row.id+"_edit_name");
        e.innerText =  f.value;
        e = document.getElementById(''+row.id+"_details");
        g = document.getElementById(''+row.id+"_edit_desc");
        // pull out type
        typeStr = "n/a"
        te = document.getElementById(''+row.id+"_type_bool");
        if ( te != null && te.className == "col active") {
            typeStr = "Boolean"
        }
        te = document.getElementById(''+row.id+"_type_number");
        if ( te != null && te.className == "col active") {
            var k = document.getElementById(''+row.id+"_type_number_type");
            typeStr = "Number"
            if (k.selectedIndex == 0) {
                typeStr = "Number: Integer"
            }
            if (k.selectedIndex == 1) {
                typeStr = "Number: Float"
            }
            var i = document.getElementById(''+row.id+"_type_number_min");
            var j = document.getElementById(''+row.id+"_type_number_max");
            if ( i != "" || j != "") {
                typeStr = typeStr + "["+i.value+".."+j.value+"]";
            }
        }
        te = document.getElementById(''+row.id+"_type_str");
        if ( te != null && te.className == "col active") {
            typeStr = "String"
            var i = document.getElementById(''+row.id+"_type_str_maxlength");
            if ( i != null && i.value != "") {
                typeStr = typeStr + "["+i.value+"]";
            }
        }
        e.innerHTML = '<p>'+typeStr+'<br/><i>'+g.value+'</i></p>';

        // hide "edit" part
        var showPart = document.getElementById(''+row.id+"_show");
        showPart.className = " col s8";
        // unhide "show" part
        var editPart = document.getElementById(''+row.id+"_edit");
        editPart.className += " hide";

    }
}

// delete an item from mp_list
function mp_list_click_delete(e) {
    e.preventDefault();
    var t = e.target;

    // depending on browser, t can be the <button> or the <i> holding the icon.
    if ( t.localName == "i") {
        t = t.parentElement;
    }

    var data_row = t.parentNode.parentNode.parentNode;

    // remove the whole row
    data_row.parentNode.removeChild(data_row);
    mp_list_update_count_in_toc();

}

function mp_list_update_count_in_toc() {
    var mp_list = document.getElementById('mp_list');
    console.log("mp_list="+mp_list);
    var toc_elem = document.getElementById('mp_items_toc');
    console.log("toc_elem="+toc_elem);
    toc_elem.innerText = mp_list.children.length;
}




// details = page with properties, actions, events.

// configure modal dialog
$('#details_validation_modal').modal({
        dismissible: true, // Modal can be dismissed by clicking outside of the modal
        opacity: .5, // Opacity of modal background
        inDuration: 300, // Transition in duration
        outDuration: 200, // Transition out duration
        startingTop: '4%', // Starting top style attribute
        endingTop: '10%', // Ending top style attribute
    }
);
$('#details_validation_modal').modal();

document.getElementById('details_next').addEventListener('click', details_next);

function details_next(e) {
    e.preventDefault()

    var details_ok = true

    var errors = "<ul>"
    if ( !mp_list_has_items()) {
        errors += "<li>You should at least have one property.</li>"
        details_ok = false
    }
    errors += "</ul>"

    if (details_ok) {
        //
        console.log("submit TODO")
    } else {
        document.getElementById('details_validation_modal_reason').innerHTML = errors;
        $('#details_validation_modal').modal('open');
    }
}
