<!DOCTYPE html>

<html>
<head>
  <title>{{if .Gruppa}}Группа товаров: {{.Gruppa.Name}} {{else}}Каталог товаров из 1С{{end}}</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
  <link rel="stylesheet" href="static/css/catalog.css">
</head>

<body>
  <header class="header">
    <h1 class="logo"><a href="/catalog">Каталог товаров из 1С.</a> Товаров: {{.Count}} </h1>
  </header>
  <div class="container-fluid row">
    <div class="col-sm-9">
      {{range .Tovars}}
      <div class="row bgcol{{.}}">
        <div class="col-6"><a href="/tovar?tovid={{.Id}}">{{.Tname}}</a></div><div class="col-2">{{.Articul}}</div><div class="col-4">{{.Gname}}</div>
      </div>
      {{end}}
    </div>
    <div class="col-sm-3">
      <div class="row">
        <div class="col-12"><b><a href="/catalog?grpid={{.Gruppa.Id}}">{{.Gruppa.Name}}</a></b> (Групп: {{.Countg}})</div>
      </div>
      {{range .Grupps}}
      <div class="row">
        <div class="col">
          <a href="/catalog?grpid={{.Id}}">{{.Id}}.{{.Name}}</a>
        </div>
      </div>
      {{end}}
    </div>    
  </div>

  <footer class="footer">
    <div>Контакты магазина</div>
  </footer>

  <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js" integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" crossorigin="anonymous"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>

</body>
</html>
