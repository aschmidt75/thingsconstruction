// CT = Create Thing
//
$("document").ready(function(){
    $('select').material_select();

    $('#ct_type_info').modal({
        dismissible: true
    });

    // configure modal dialog
    $('#ct_form_validation_modal').modal({
            dismissible: true, // Modal can be dismissed by clicking outside of the modal
            opacity: .5, // Opacity of modal background
            inDuration: 300, // Transition in duration
            outDuration: 200, // Transition out duration
            startingTop: '4%', // Starting top style attribute
            endingTop: '10%', // Ending top style attribute
        }
    );
    //$('#ct_form_validation_modal').modal();

    document.getElementById('ct_next').addEventListener('click', ct_next);

    var t = document.getElementById('type_info');
    console.log(t);
    t.addEventListener('click', ct_type_info_modal);
    console.log(t);

});

// validate form, post to backend
function ct_next(e) {
    e.preventDefault();

    var n = document.getElementById('ctf_name').value;
    if ( n == "") {
        document.getElementById('ctf_name').style.borderBottom = 'solid 2px #ee2222'
        $('#ct_form_validation_modal').modal('open');
        return;
    } else {
        document.getElementById('ctf_name').style.borderBottom = 'solid 1px #000'
    }

    // is valid, submit
    var form = document.getElementById('ctf');
    form.submit();
}


function ct_type_info_modal(e) {
    console.log(e);
    e.preventDefault();

    $('#ct_type_info').modal('open');

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

document.getElementById('btn_more').addEventListener('click', more_activate);

