    $(document).ready(function(){
        // the "href" attribute of the modal trigger must specify the modal ID that wants to be triggered
        $('.modal').modal();
    });

var ThingId = document.getElementById('span_thing_id').innerText;

function view_element(e) {
    $('#view_modal').modal('open');

    var permaLink = e.target.getAttribute("linkid");
    var target = document.getElementById("view_modal_content");
    var viewUrl = "/app/"+ThingId+"/result/assetview/"+permaLink;

    var content = "";

    $.ajax({
        type: "GET",
        url: viewUrl,
        async: true,
        success: function (data) {
            content = content + data
        },
        error: function (data) {
            console.log(data);
            target.innerHTML = "<p>An error occured while fetching the document. Please try again later.</p>";
        },
        complete: function() {
            target.innerHTML = content;
            (function(){
                new Clipboard('#copycode');
            })();
        },
    });

}

// connect all view links
function getAllElementsWithAttribute(attribute)
{
    var matchingElements = [];
    var allElements = document.getElementsByTagName('*');
    for (var i = 0, n = allElements.length; i < n; i++)
    {
        if (allElements[i].getAttribute(attribute) !== null)
        {
            matchingElements.push(allElements[i]);
        }
    }
    return matchingElements;
}
var allViewLinks = getAllElementsWithAttribute('linkid');
for ( var i = 0; i < allViewLinks.length; i++) {
    allViewLinks[i].addEventListener('click', view_element);
}

$(document).ready(function() {
    $('select').material_select();
});

$("#select-archive").on('change', function() {
    var v = $(this)[0].value;

    var l = document.getElementById("select-archive-link");
    if ( v === "") {
        l.className = "hide";
        l.href = "";

    } else {
        l.className = "";
        l.href = "/app/"+ThingId+"/result/asset-archive/"+v;
    }
});

$('#url_copy_btn').on('click', function(e) {
    var i = document.getElementById("url_copy_input");
    i.value = document.getElementById("url_copy_a").textContent;
    i.select();
    document.execCommand("copy");
    //M.toast({html: 'Copied.'});
})

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

var btn_more = document.getElementById('btn_more');
if ( btn_more != null) {
    btn_more.addEventListener('click', more_activate);
}

function delete_data(e) {
    var url = "/app/"+ThingId;

    $.ajax({
        type: "DELETE",
        url: url,
        async: true,
        success: function (data) {
            content = content + data
        },
        error: function (data) {
            console.log(data);
            target.innerHTML = "<p>An error occured while fetching the document. Please try again later.</p>";
        },
        complete: function() {
            var url = document.URL;
            console.log(url);
            url = url.replace(/^(.*\/app).*/, "$1");

            console.log(url);
            window.location.replace(url);
        }
    });
}

var btn_delete = document.getElementById('btn-delete-data');
if ( btn_delete != null) {
    btn_delete.addEventListener('click', delete_data);
}
