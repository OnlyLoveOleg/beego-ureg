<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="Beego User Registration Example">
    <meta name="author" content="Tal Lannder">
    <link rel="icon" href="//{{.base_url}}/static/img/favicon.ico">

    <title>{{.Title}}</title>
    <!-- Bootstrap core CSS -->
    <link rel="stylesheet" href="//{{.base_url}}/static/bootstrap-3.3.5-dist/css/bootstrap.min.css">
    <!-- Custom Global styles for this template -->
    <link rel="stylesheet" href="//{{.base_url}}/static/css/globalStyles.css">
  </head>

  <body>

  <nav class="navbar navbar-default navbar-fixed-top">
    <div class="container-fluid">
      <div class="navbar-header">
        <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
          <span class="sr-only">Toggle navigation</span><span class="icon-bar"></span><span class="icon-bar"></span><span class="icon-bar"></span>
        </button><a class="navbar-brand" href="/"><img src="//{{.base_url}}/static/img/logo.png" alt="My Logo"></a>
      </div><!-- end navbar head -->
      <div id="navbar" class="navbar-collapse collapse">
        <ul class="nav navbar-nav pull-right">
          {{if .Session}}
          <li><p class="navbar-text">Welcome, {{.Session.first}}</p></li>
          <li><a href="//{{.base_url}}/accounts/logout"><span class="glyphicon glyphicon-off"></span></a></li>
          <li><a href="//{{.base_url}}/accounts/profile"><span class="glyphicon glyphicon-user"></span></a></li>
          {{else}}
          <form class="navbar-form navbar-left" role="signin" action="//{{.base_url}}/accounts/signin" method="POST">
            <div class="form-group">
              <input type="email" name="email" class="form-control" placeholder="email">
            </div>
            <div class="form-group">
              <input type="password" name="password" class="form-control" placeholder="password">
            </div>
            <button type="submit" class="btn btn-primary">Log in</button>
          </form>
          <li><p class="navbar-btn"><a href="//{{.base_url}}/accounts/signup" class="btn btn-success" role="button">Sign up</a></p></li>
          {{end}}
        </ul><!-- end nav pull-right -->
      </div><!-- end nav-collapse -->
    </div><!-- end container -->
  </nav><!-- end navbar -->

  {{if .flash}}
  <div class="container">
    <div class="row">
      {{if .flash.error}}
      <div class="alert alert-danger" role="alert">{{.flash.error}}</div>
      {{end}}
      {{if .flash.warning}}
      <div class="alert alert-warning" role="alert">{{.flash.warning}}</div>
      {{end}}
      {{if .flash.notice}}
      <div class="alert alert-success" role="alert">{{.flash.notice}}</div>
      {{end}}
    </div>
  </div>
  {{end}}

    <div class="container animated fadeInLeft" id="main">
      <div class="row" id="bigCallout">
         <div class="col-12">
          <div class="jumbotron">
            <h3><em>Beego User Registration Example</em></h3>
          </div><!-- end jumbotron -->
        </div><!-- end col-12 -->
      </div><!-- end bigCallout -->
    </div><!-- end main container -->
  </body>

  <footer>
    <div class="container-fluid">
      <div class="row">
        <div class="col-xs-12">
          <h6>&copy; 2015 Some Inc.</h6>
        </div><!-- end col-sm-2 -->
      </div><!-- end row -->
    </div><!-- end container -->
  </footer>

  <script src="//{{.base_url}}/static/jquery/jquery-2.1.4.min.js"></script>
  <script src="//{{.base_url}}/static/bootstrap-3.3.5-dist/js/bootstrap.min.js"></script>
</html>
