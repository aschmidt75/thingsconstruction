{{define `head`}}

<head>
    <meta charset="UTF-8">
    <META HTTP-EQUIV="Content-Type" CONTENT="text/html; charset=utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="description" content="Embedded IoT Code Generator for constrained devices, according to W3C's Web Of Things (WoT) Thing Description">
    <meta name="keywords" content="embedded constrained devices microcontrollers IoT Arduino Framework HTTP JSON Generator Internet of Things W3C WOT Web Of Things Thing Description">
    <link rel="author" href="twitter.com/aschmidt75">
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>THNG:STRUCTION - {{ .Title }}</title>

    <!-- Material Icon CDN -->
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <!-- Path of the 0.9 materialize.min.css file -->
    <link rel="stylesheet" href="/css/materialize.min.css" media="screen,projection">
    <!-- 1.0 Compiled and minified CSS
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0-alpha.3/css/materialize.min.css">-->

    <!-- Custom css file path -->
    <link rel="stylesheet" href="/css/style.css">
    <link rel="stylesheet"
        href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/styles/default.min.css">
    <script defer src="https://use.fontawesome.com/releases/v5.0.0/js/all.js"></script>

    <style>
    body {
      display: flex;
      min-height: 100vh;
      flex-direction: column;
    }

    main {
      flex: 1 0 auto;
    }

    li a.active {
      background: rgba(0,0,0,0.2);
    }
    </style>


    <!-- Global site tag (gtag.js) - Google Analytics -->
    <script async src="https://www.googletagmanager.com/gtag/js?id=UA-113732834-1"></script>
    <script>
      window.dataLayer = window.dataLayer || [];
      function gtag(){dataLayer.push(arguments);}
      gtag('js', new Date());

      gtag('config', 'UA-113732834-1', { 'anonymize_ip': true });
    </script>

    <!-- cookie consent -->
    <link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/cookieconsent2/3.0.3/cookieconsent.min.css" />
    <script src="https://cdnjs.cloudflare.com/ajax/libs/cookieconsent2/3.0.3/cookieconsent.min.js"></script>
    <script>
        window.addEventListener("load", function(){
            window.cookieconsent.initialise({
                "palette": {
                    "popup": {
                        "background": "#000"
                    },
                    "button": {
                        "background": "#fff"
                    }
                },
                "content": {
                    "href": "https://thngstruction.online/imprint.html"
                }
            })});
    </script>

</head>


{{end}}