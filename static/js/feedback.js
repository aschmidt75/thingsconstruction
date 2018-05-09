$("document").ready(function(){
    $('select').material_select();
});

document.getElementById('feedback_accept_cb').addEventListener('click', feedback_accept_cb);

function feedback_accept_cb(e) {
    var cb = e.target;

    var d = document.getElementById('fbf_send');
    var x = d.className;
    if (cb.checked === true) {
        d.className = x.replace(new RegExp(" disabled", "g"), "");
        d.addEventListener('click', fbf_do_send);
    } else {
        d.className += " disabled";
        d.removeEventListener('click', fbf_do_send);
    }
}

function fbf_do_send(e) {
    e.preventDefault();

    // disable button click event
    var cff_submit_btn = document.getElementById('fbf_send');
    cff_submit_btn.removeEventListener('click', fbf_do_send);

    // submit
    var frm = $('#fbf');
    var response = "Uh oh, i'm not sure what happened. Please try again.";
    $.ajax({
        type: frm.attr('method'),
        url: frm.attr('action'),
        data: frm.serialize(),
        async: true,
        success: function (data) {
            response = data;
        },
        error: function (data) {
            response = 'I\'m sorry, there was an error on our side :-/ Please try again later.';
        },
        complete: function() {
            // remove form
            var fbf = document.getElementById('fbf');
            while (fbf.firstChild) {
                fbf.removeChild(fbf.firstChild);
            }

            var div2 = document.createElement("div");
            div2.className = "card-panel grey darken-1";
            div2.innerHTML = "<span class=\"white-text\">"+response+"</span>";
            fbf.appendChild(div2);
        }
    });


}
