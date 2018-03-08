<!DOCTYPE html>

<html>
<head>
  <title>1c outlet - каталог товаров из 1с</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
</head>

<body>
  <header>
    <h1 class="logo">Каталог товаров из 1С</h1>
    <div class="description">{{.Hi}} Счетчик вызовов равер {{.SecretCode}}</div>
    <p>Записей в таблице:{{.Count}} Ошибка:{{.Err}}</p>
    <hr>
  </header>
  {{range .Tovars}}
  <div class="container-fluid">
    <div class="row">
      <div class="col-sm-6">{{.Name}}</div><div class="col-sm-4">{{.Idc}}</div><div class="col-sm-1">2</div>
    </div>
  </div>
  {{end}}
  <footer>
    <hr>
  </footer>
  <script src="/static/js/reload.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
</body>
</html>
