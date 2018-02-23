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
    <div class="description">
    Тут привет из {{.Hi}}<br>
    </div>
    <hr>
  </header>
  <div class="container-fluid">
    {{range .Tovar}}
    <div class="row">
    <div class="col-sm-7">{{.Наименование}}</div><div class="col-sm-2">{{.Артикул}}</div><div class="col-sm-3">{{.Изготовитель.Наименование}}</div>
    </div>
    {{end}}
  </div>
  <footer>
    <hr>
  </footer>
  <script src="/static/js/reload.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
</body>
</html>
