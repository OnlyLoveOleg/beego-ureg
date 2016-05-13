  <nav class="navbar navbar-default navbar-fixed-top">
    <div class="container-fluid">
       <div class="navbar-header">
        <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
          <span class="sr-only">Toggle navigation</span><span class="icon-bar"></span><span class="icon-bar"></span><span class="icon-bar"></span>
        </button>
        <a class="navbar-brand" href="/"><img src="//{{.base_url}}/static/img/logo.png" alt="Some Inc. Logo"></a>
      </div><!-- end navbar head -->
      {{if .Session}}
      <div id="navbar" class="navbar-collapse collapse">
        <ul class="nav navbar-nav pull-right">
          <li><p class="navbar-text">Welcome, {{.Session.firstname}}</p></li>
          <li><a href="//{{.base_url}}/accounts/signout"><span class="glyphicon glyphicon-off"></span></a></li>
          <li><a href="//{{.base_url}}/accounts/profile"><span class="glyphicon glyphicon-user"></span></a></li>
        </ul><!-- end nav pull-right -->
      </div><!-- end nav-collapse -->
      {{end}}
    </div><!-- end container -->
  </nav><!-- end navbar -->
  <div class="container" id="content">
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
  {{if .Session}}
  <div class="container">
    <ul class="nav nav-tabs nav-justified">
      <li role="presentation"{{if .ProfileActive}} class="active"{{end}}>
        <a href="//{{.base_url}}/accounts/profile">
          <span class="glyphicon glyphicon-user" aria-hidden="true"></span> Profile
        </a>
      </li>
      <li role="presentation"{{if .SecurityActive}} class="active"{{end}}>
        <a href="//{{.base_url}}/accounts/security">
          <span class="glyphicon glyphicon-lock" aria-hidden="true"></span> Security
        </a>
      </li>
    </ul>
  </div>
  {{end}}
