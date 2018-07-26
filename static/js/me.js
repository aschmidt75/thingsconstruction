//
// me = manage events
//


// configure modal dialog
$("#me_listitem_validation_modal").modal({
        dismissible: true, // Modal can be dismissed by clicking outside of the modal
        opacity: .5, // Opacity of modal background
        inDuration: 300, // Transition in duration
        outDuration: 200, // Transition out duration
        startingTop: "4%", // Starting top style attribute
        endingTop: "10%" // Ending top style attribute
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

// connect add(+) button of properties list to me_list_add
document.getElementById("me_add_btn").addEventListener("click", me_list_add);

var btn = document.getElementById("me_next");
var u = btn.attributes["customizationUrl"];
if ( u !== undefined && u.nodeValue !== "") {
    btn.innerHTML = btn.innerHTML.replaceAll(/Generate/,"Customize");
    btn.innerHTML = btn.innerHTML.replaceAll(/navigate_next/,"developer_board");
}

// adds an existing action object to the list.
function me_list_add_existing(obj) {
    // add a new row to the ul #me_list
    var me_list = document.getElementById("me_list");

    // extract the id of the last item in the list. The new div gets id+1
    var last = me_list.lastElementChild;
    var last_id = 0;
    if ( last != null) {
        var last_ids = last.id.split("_");
        last_id = parseInt(last_ids[last_ids.length-1]);
    }

    var me_list_item = document.createElement("li");
    me_list_item.id = "me_listitem_"+(last_id+1);
    me_list_item.className = "collection-item row";
    // template contains ## for each place where new id shall fit in
    var tplhtml = document.getElementById("tpl_mef_list_item").innerHTML;
    me_list_item.innerHTML = tplhtml
        .replaceAll("##", ""+(last_id+1));
    me_list.appendChild(me_list_item);

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
    var btn_edit = document.getElementById("me_listitem_"+(last_id+1)+"_btns_edit");
    if ( obj == null) {
        // no content yet, directly go to edit mode
        me_list_click_edit_showHide(btn_edit);
    }
    btn_edit.addEventListener("click", me_list_click_edit);

    document.getElementById("me_listitem_"+(last_id+1)+"_btns_delete").addEventListener("click", me_list_click_delete);

    if ( obj !== null) {
        document.getElementById("me_listitem_"+(last_id+1)+"_name").textContent = $.encoder.encodeForHTML(obj.name);
        document.getElementById("me_listitem_"+(last_id+1)+"_edit_name").value = $.encoder.encodeForHTML(obj.name);
        document.getElementById("me_listitem_"+(last_id+1)+"_edit_desc").value = $.encoder.encodeForHTML(obj.description);

        document.getElementById("me_listitem_"+(last_id+1)+"_details").innerHTML = "<p><i>"+$.encoder.encodeForHTML(obj.description)+"</i></p>"+
            "<input type=\"text\" name=\"me_listitem_"+(last_id+1)+"_val\" class=\"hide\" "+$.encoder.encodeForHTMLAttribute("value", obj.name)+">"+
            "<input type=\"text\" name=\"me_listitem_"+(last_id+1)+"_desc\" class=\"hide\" "+$.encoder.encodeForHTMLAttribute("value", obj.description)+">";

    }
}

// add content of tpl_mpl_list_item to me_list.
// Open "edit" part, connect buttons
function me_list_add(e) {
    e.preventDefault();
    me_list_add_existing(null);
}

function me_list_click_edit(e) {
    e.preventDefault();
    var t = e.target;
    me_list_click_edit_showHide(t)
}

// checks if n is already a name in the mp list.
// except the row with id=edit_id
function me_list_is_name_in_use_except(n, edit_id) {
    if ( n === "") {
        return false;
    }
    var me_list = document.getElementById("me_list");
    var me_list_items = me_list.children;
    for ( var i = 0; i < me_list_items.length; i++) {
        if ( me_list_items[i].id !== edit_id) {
            var id = ""+me_list_items[i].id+"_name";
            var n2 = document.getElementById(id).innerText;

            if ( n !== "" && n === n2) {
                return true;
            }
        }
    }
    return false;
}

function me_list_has_items() {
    var me_list = document.getElementById("me_list");
    var me_list_items = me_list.children;
    return me_list_items.length > 0;
}

// removes the disabled class from an id
function me_list_enable_button(id) {
    var cn = document.getElementById(id).className;
    document.getElementById(id).className = cn.replaceAll("disabled", "");
}

// enable all edit/delete buttons
function me_list_enable_buttons() {
    var me_list = document.getElementById("me_list");
    var me_list_items = me_list.children;
    for ( var i = 0; i < me_list_items.length; i++) {
        var id;
        id = ""+me_list_items[i].id+"_btns_edit";
        me_list_enable_button(id);
        id = ""+me_list_items[i].id+"_btns_delete";
        me_list_enable_button(id);
    }
    me_list_enable_button("me_add_btn");
    me_list_enable_button("me_next");
    me_list_enable_button("me_prev");
}

// disable all edit/delete buttons
function me_list_disable_buttons() {
    var me_list = document.getElementById("me_list");
    var me_list_items = me_list.children;
    for ( var i = 0; i < me_list_items.length; i++) {
        var id;
        id = ""+me_list_items[i].id+"_btns_edit";
        document.getElementById(id).className += " disabled";
        id = ""+me_list_items[i].id+"_btns_delete";
        document.getElementById(id).className += " disabled";
    }
    document.getElementById("me_add_btn").className += " disabled";
    document.getElementById("me_next").className += " disabled";
    document.getElementById("me_prev").className += " disabled";
}


// only allow inputs which would make up a valid identifier in
// code.
function me_list_limit_name(e) {
    var k = e.key;
    if ( e.charCode === 0 && e.keyCode !== 0) return true;
    if ( k >= '0' && k <= '9') return;
    if ( k >= 'a' && k <= 'z') return;
    if ( k >= 'A' && k <= 'Z') return;
    if ( k === '_' || k === '-') return;
    e.preventDefault();
}

// given the node of the edit button, this method
// toggles the edit fields and show fields.
function me_list_click_edit_showHide(t) {

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
        me_list_disable_buttons();
        // enable the save button again
        me_list_enable_button(t.id);

        //
        document.getElementById(""+row.id+"_edit_name").addEventListener("keypress", me_list_limit_name)

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
            $("#me_listitem_validation_modal").modal("open");
            return;
        };

        // 2. name must be unique, compare to all others
        if ( me_list_is_name_in_use_except(newName, row.id)) {
            document.getElementById(""+row.id+"_edit_name").style.borderBottom = "solid 2px #ee2222";
            document.getElementById("me_listitem_validation_modal_reason").innerText =
                "The name "+newName+" was already chosen for another property.";
            $("#me_listitem_validation_modal").modal("open");
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
            document.getElementById("me_listitem_validation_modal_reason").innerText = e;
            $("#me_listitem_validation_modal").modal("open");
            return;
        }

        // if we get here, all is valid. un-color, continue.
        document.getElementById(""+row.id+"_edit_name").style.borderBottom = "";
        me_list_enable_buttons();

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
            e.innerHTML = "<p><i>"+$.encoder.encodeForHTML(g)+"</i></p>"+
                "<input type=\"text\" name=\""+row.id+"_val\" class=\"hide\" "+$.encoder.encodeForHTMLAttribute("value", f)+">"+
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

// delete an item from me_list
function me_list_click_delete(e) {
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

document.getElementById("me_next").addEventListener("click", me_submit);
document.getElementById("me_prev").addEventListener("click", me_to_actions);

function me_disable_navbtns() {
    var btn1 = document.getElementById("me_next");
    var btn2 = document.getElementById("me_prev");
    btn1.className += " disabled";
    btn2.className += " disabled"

}
function me_enable_navbtns() {
    var btn1 = document.getElementById("me_next");
    var btn2 = document.getElementById("me_prev");
    btn1.className = btn1.className.replaceAll("disabled", "");
    btn2.className = btn2.className.replaceAll("disabled", "");
}

function me_to_actions(e) {
    e.preventDefault();

    var details_ok = true;
    if (details_ok) {
        //
        var frm = document.getElementById("mef");
        me_disable_navbtns();

        $.ajax({
            type: frm.method,
            url: frm.action,
            data: $("#mef").serialize(),
            success: function (data) {
                // redirect to next page
                var url = document.URL;
                window.location.replace(url.replace(/^(.*)\/events.*/,"$1/actions"));
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

function me_submit(e) {
    e.preventDefault();

    var details_ok = true;
/*
    var errors = "<ul>";
    if ( !me_list_has_items()) {
        errors += "<li>You should at least have one property.</li>";
        details_ok = false
    }
    errors += "</ul>";
*/
    if (details_ok) {
        //
        var frm = document.getElementById("mef");
        me_disable_navbtns();

        $.ajax({
            type: frm.method,
            url: frm.action,
            data: $("#mef").serialize(),
            success: function (data) {
                // redirect to next page
                var url = document.URL;

                var btn = document.getElementById("me_next");
                var u = btn.attributes["customizationUrl"];
                if ( u !== undefined && u.nodeValue !== "") {
                    var s = u.nodeValue;
                    console.log("redirecting to "+s);
                    window.location.replace(url.replace(/^(.*)(\/app.*)/,"$1"+s.toString()));
                } else {
                    // forward to generate page
                    window.location.replace(url.replace(/^(.*)\/events.*/,"$1/generate"));
                }
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











var progress = document.getElementById("progress");
var eventsJson = "";
var url = document.URL;
url = url.replace(/^(.*\/events).*/, "$1/data")
$.ajax({
    type: "GET",
    url: url,
    async: true,
    success: function (data) {
        eventsJson = eventsJson + data
    },
    error: function (data) {
        console.log(data);
    },
    complete: function() {
        var obj = JSON.parse(eventsJson);
        if (obj !== null) {
            for (var key in obj) {
                if (obj.hasOwnProperty(key)) {
                    prop = obj[key]
                    prop.name = key
                    me_list_add_existing(prop)
                }
            }
        }
        progress.className += " hide";

    }
});



