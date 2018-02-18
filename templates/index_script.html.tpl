{{define `script`}}
<script type="text/javascript" language="javascript">
    $( document ).ready(function() {
        $('.carousel.carousel-slider').carousel({
            fullWidth: true,
            indicators: true,
            dist:0,
        });
    });
    autoplay();
    function autoplay() {
        $('.carousel').carousel('next');
        setTimeout(autoplay, 5000);
    }
</script>
{{end}}