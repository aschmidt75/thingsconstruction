{{define `head`}}

<head>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta charset="UTF-8">
    <META HTTP-EQUIV="Content-Type" CONTENT="text/html; charset=utf-8">
    <meta name="description" content="Embedded IoT Code Generator for constrained devices, according to W3C's Web Of Things (WoT) Thing Description">
    <meta name="keywords" content="embedded constrained devices microcontrollers IoT Arduino Framework HTTP JSON Generator Internet of Things W3C WOT Web Of Things Thing Description">
{{ if eq .Robots false -}}
    <meta name="robots" content="noindex">
    <meta name="googlebot" content="noindex">
{{ end -}}
    <link rel="author" href="twitter.com/aschmidt75">
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>THNG:STRUCTION - {{ .Title }}</title>

    <!-- Material Icon CDN -->
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <!-- Path of the 0.9 materialize.min.css file -->
    <link rel="stylesheet" href="/css/materialize.min.css" media="screen,projection">

    <!-- Custom css file path -->
    <link rel="stylesheet" href="/css/style.css">
    <link rel="stylesheet"
        href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/styles/default.min.css">

    <link rel="stylesheet" href="/css/fontawesome.css">
    <link href="/css/fa-brands.css" rel="stylesheet">

{{ if .Feature.Shariff }}
    <link rel="stylesheet" href="/css/shariff.min.css">
{{ end -}}

    <link rel="stylesheet" href="/css/style2.css">

    {{ if .Feature.Analytics }}
    <!-- Global site tag (gtag.js) - Google Analytics -->
    <script type="text/javascript" async src="https://www.googletagmanager.com/gtag/js?id=UA-113732834-1"></script>
    <script type="text/javascript" src="/js/tcga.js"></script>
    {{ end }}
    <!-- cookie consent -->
    <link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/cookieconsent2/3.0.3/cookieconsent.min.css" />
    <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/cookieconsent2/3.0.3/cookieconsent.min.js"></script>
    <script type="text/javascript" src="/js/cookieconsent.js"></script>
</head>


{{end}}