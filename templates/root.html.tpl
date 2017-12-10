{{define `root`}}
<!DOCTYPE html>
<html>
{{ template `head` . }}
<body>
{{ template `header` . }}
<main>
    <div id="main" class="container">
    {{ template `main` . }}
    </div>
</main>
{{ template `footer` . }}
</body>
</html>
{{end}}