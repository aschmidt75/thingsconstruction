{{define `head`}}

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Title of the document</title>
  <!-- Material Icon CDN -->
  <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
  <!-- Path of the materialize.min.css file -->
  <link rel="stylesheet" href="/css/materialize.min.css" media="screen,projection">

    <script defer src="https://use.fontawesome.com/releases/v5.0.0/js/all.js"></script>

    <!-- Custom css file path -->
  <link rel="stylesheet" href="/css/style.css">
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
</head>


{{end}}