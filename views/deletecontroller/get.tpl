<div class="container" id="content">
  <h1>Delete Account</h1>

  <div class="panel panel-danger">
    <div class="panel-heading">
      <h3 class="panel-title">Caution: all related transactions and data will also be removed. Are you sure?</h3>
    </div>
    <br>
    <div class="panel-body">
      <form class="form-horizontal" method="POST">
        <div class="form-group{{if .Errors.current}} has-error has-feedback{{end}}">
          <label for="currentpassword" class="col-sm-2 control-label">Current password:</label>
          <div class="col-sm-8">
            <input type="password" name="currentpassword" id="currentpassword" class="form-control">
            {{if .Errors.currentpassword}}<span class="help-block">{{.Errors.currentpassword}}</span>{{end}}
          </div>
        </div>
        <br>
        <div class="row text-center">
          <div class="col-xs-6">
            <button type="submit" class="btn btn-danger" value="Delete">Delete account</button>
          </div>
          <div class="col-xs-6">
            <p><a class="btn btn-primary" href="//{{.base_url}}/accounts/profile">Cancel</a></p>
          </div>
        </div>
      </form>
    </div>
  </div><!-- end panel -->
</div><!-- end container -->
