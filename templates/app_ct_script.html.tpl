{{define `script`}}
<script src="/js/ct.js"></script>
<script>
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
</script>
{{end}}