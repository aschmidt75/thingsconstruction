{{define `head`}}

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>THNGS:CONSTR - {{ .Title }}</title>
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

  <!-- Used as an example only to position the footer at the end of the page.
      You can delete these styles or move it to your custom css file -->
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

</head>


{{end}}