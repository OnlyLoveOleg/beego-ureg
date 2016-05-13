<div class="container" id="content">
  <div class="row">
    <h1>Reset password</h1>
    <br>
    <form class="form-horizontal" method="POST">
      <div class="form-group{{if .Errors.password}} has-error has-feedback{{end}}">
        <label for="inputPassword" class="col-sm-2 control-label">New password</label>
        <div class="col-sm-8">
          <input type="password" name="password" id="inputPassword" class="form-control">
          {{if .Errors.password}}
          <span class="help-block">{{.Errors.password}}</span>
          {{else}}
          <span class="help-block">New password (must be at least 8 characters)</span>
          {{end}}
        </div>
      </div>

      <div class="form-group{{if .Errors.password2}} has-error has-feedback{{end}}">
        <label for="inputPassword2" class="col-sm-2 control-label">Confirm new password</label>
        <div class="col-sm-8">
          <input type="password" name="password2" id="inputPassword2" class="form-control">
          {{if .Errors.password2}}<span class="help-block">{{.Errors.password2}}</span>{{end}}
        </div>
      </div>

      <div class="form-group">
        <div class="col-sm-offset-2 col-sm-8">
          <button type="submit" class="btn btn-primary" value="Reset password">Reset password</button>
        </div>
      </div>
    </form>
</div><!-- end main container -->
