<!DOCTYPE html>
<html>
<head>
  <title>{{.Tovar.Name}} описание товара</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
  <link rel="stylesheet" href="static/css/catalog.css">
  <link rel="stylesheet" href="static/css/tovar.css">
</head>

<body>
    <header class="header">
        <h1 class="logo">Товар #{{.Tovid}}</h1>
    </header>

    <div class="container-fluid">
        <div class="row">
            <div class="col-8">
                <div class="row"><h2>{{.Tovar.Name}}</h2></div>
                <div class="row" id="tovar">{{.Tovar.Articul}}</div>
                <div class="row">{{.Tovar.Opisanie}}</div>
                
            </div>
            <div class="col-4" id="gruppa">
                <h2>Группа</h2>
            </div>
        </div>
        <div class="row">
            <div class="col-12">
                {{.Tovar.Idc}}
            </div>
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
