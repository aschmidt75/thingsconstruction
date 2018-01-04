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
// CT = Create Thing
//
$("document").ready(function(){
    $('select').material_select();
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
$('#ct_form_validation_modal').modal();

document.getElementById('ct_next').addEventListener('click', ct_next);

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
