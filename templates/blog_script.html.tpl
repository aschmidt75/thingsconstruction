{{define `script`}}
<!-- immediately before </body> -->
<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/highlight.min.js"></script>
<script>
    // set document title to blog post title
    document.title = "{{ .PageData.Title }}"
</script>
<script>
    hljs.configure({useBR: false});

    $('div > pre').each(function(i, block) {
        hljs.highlightBlock(block);
    });
</script>

<script>
var bp_selected_tags = new Map();

// if we are on the overview page, do some filtering magic
all_posts = document.getElementById("bp_overview")
if (all_posts != null) {
    // enable handlers for all chips on the right
    all_tags = document.getElementById("bp_all_tags")
    c = all_tags.children
    for (i = 0; i < c.length; i++) {
        if (c[i].className.startsWith("chip")) {
            // add listener
            c[i].addEventListener('click', bp_tag_clicked)

        }
    }
}
bp_update_list_filter()

function bp_update_list_filter() {
    var all_on = (bp_selected_tags.size == 0)
    var total = 0
    var active = 0
    all_bps = document.getElementsByClassName("blogpost")
    for ( var i = 0; i < all_bps.length; i++) {
        var c = all_bps[i].firstElementChild
        var num_chips = 0;
        var num_matching = 0;
        for ( var j = 0; j < c.children.length; j++) {
            if (c.children[j].className.startsWith("chip")) {
                num_chips += 1
                var tagName = c.children[j].innerText.trim()
                if ( bp_selected_tags.get(tagName) != null) {
                    num_matching += 1
                } else {
                }
            }
        }
        total += 1
        if (( all_on == true) || (num_matching > 0)) {
            // all tags of this post are selected on right -> activate
            all_bps[i].className = "blogpost row"
            active += 1
        } else {
            // not all tags are selected on the right -> deactivate
            all_bps[i].className = "blogpost row hide"
        }
    }

    // update title
    if ( total == active) {
        document.getElementById("bp_title").innerHTML = "Blog posts"
        document.getElementById("bp_count").innerText = "all"
        document.getElementById("bp_count").className = "badge deep-orange white-text"

    } else {
        document.getElementById("bp_title").innerText = "Blog posts"
        if ( active > 0) {
            document.getElementById("bp_count").innerText = ""+active+"/"+total
            document.getElementById("bp_count").className = "badge deep-orange white-text tiny"
        } else {
            //
            document.getElementById("bp_count").innerText = ""
            document.getElementById("bp_count").className = ""
        }
    }
}

function bp_tag_clicked(e) {
    e.preventDefault()

    tag = e.target
    tagName = tag.innerText
    if (tag.className.startsWith("chip bp_selected")) {
        tag.className = "chip"
        bp_selected_tags.delete(tagName)
    } else {
        tag.className = "chip bp_selected deep-orange lighten-3"
        bp_selected_tags.set(tagName, tag)
    }

    // filter list according to new selection
    bp_update_list_filter()
}

</script>
{{end}}