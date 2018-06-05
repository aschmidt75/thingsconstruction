function autoplay() {
    $('.carousel').carousel('next');
    setTimeout(autoplay, 5000);
}

function load_blog_posts() {
    var dataJson = "";
    var url = "/blog/data";
    $.ajax({
        type: "GET",
        url: url,
        async: true,
        success: function (data) {
            dataJson = dataJson + data
        },
        error: function (data) {
        },
        complete: function () {
            if (dataJson.length === 0) {
            } else {
                var obj = JSON.parse(dataJson);
                console.log(obj);
                var div;

                var content = "";

                var n = obj.length;
                if ( n > 3) {
                    n = 3;
                }
                for (var i = 0; i < n; i++) {
                    var p = obj[i];

                    content += "<div class=\"blogpost row\">";
                    content += "<div class=\"col s3 left hide-on-small-only show-on-medium-and-up\" >";
                    for ( var j = 0; j < p.Tags.length; j++) {
                        content += "<div class=\"chip small hide-on-small-and-down \">"+p.Tags[j]+"</div>";

                    }
                    content += "<br/><span>"+p.DateElapsed+"</span></div>";
                    content += "<div class=\"col s7 m9 \">";
                    content += "<a class=\"tc-maincolor-text text-lighten-1\" href=\"/blog/"+p.Name+"\">"+p.Title+"</a>";
                    content += "</div>";
                    content += "</div>";
                }

                div = document.getElementById("index_blog_posts");
                div.innerHTML = content;
            }
        }
    });
}

$( document ).ready(function() {
    $('.carousel.carousel-slider').carousel({
        fullWidth: true,
        indicators: true,
        dist:0,
    });
    autoplay();
    load_blog_posts();
});

