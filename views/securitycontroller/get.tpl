<div class="container" style="padding-top: 80px" id="content">
  <div class="row">
    <form class="form-horizontal" method="POST">

      <div class="form-group{{if .Errors.NewPassword}} has-error has-feedback{{end}}">
        <label for="NewPassword" class="col-sm-2 control-label">New password</label>
        <div class="col-sm-8">
          <input type="password" name="NewPassword" class="form-control" id="NewPassword">
          {{if .Errors.NewPassword}}<span class="help-block">{{.Errors.NewPassword}}</span>{{end}}

          <br>

          <div class="panel panel-danger hidden" id="pswd_info">
            <div class="panel-heading">
              <h4 class="panel-title">Password must meet the following requirements:</h4>
            </div>
            <div class="panel-body">
              <ul>
                <li id="letter" class="text-danger">
                  <span class="glyphicon glyphicon-remove" aria-hidden="true"></span>
                  At least <strong>one lowercase letter</strong>
                </li>
                <li id="capital" class="text-danger">
                  <span class="glyphicon glyphicon-remove" aria-hidden="true"></span>
                  At least <strong>one capital letter</strong>
                </li>
                <li id="number" class="text-danger">
                  <span class="glyphicon glyphicon-remove" aria-hidden="true"></span>
                  At least <strong>one number</strong>
                </li>
                <li id="length" class="text-danger">
                  <span class="glyphicon glyphicon-remove" aria-hidden="true"></span>
                  Be at least <strong>8 characters</strong>
                </li>
              </ul>
            </div>
          </div><!--end panel password helper-->

        </div>
      </div><!--end form-group-->

      <div class="form-group{{if .Errors.NewPasswordConfirmation}} has-error has-feedback{{end}}">
        <label for="NewPasswordConfirmation" class="col-sm-2 control-label">Confirm new password</label>
        <div class="col-sm-8">
          <input type="password" name="NewPasswordConfirmation" class="form-control" id="NewPasswordConfirmation">
          {{if .Errors.NewPasswordConfirmation}}<span class="help-block">{{.Errors.NewPasswordConfirmation}}</span>{{end}}
        </div>
      </div>

      <br>
      <h4 class="text-center">Still you?</h4>
      <br>

      <div class="form-group{{if .Errors.CurrentPassword}} has-error has-feedback{{end}}">
        <label for="CurrentPassword" class="col-sm-2 control-label">Current password</label>
        <div class="col-sm-8">
          <input type="password" name="CurrentPassword" class="form-control" id="CurrentPassword">
          {{if .Errors.CurrentPassword}}<span class="help-block">{{.Errors.CurrentPassword}}</span>{{end}}
        </div>
      </div>

      <div class="form-group">
        <div class="col-sm-offset-2 col-sm-8">
          <button type="submit" class="btn btn-primary" value="Update">Update</button>
        </div>
      </div>
    </form>
  </div>
</div><!-- end container -->
