<div class="container" id="content">
<h1>Please register</h1>

<form method="POST">
  <div class="form-group{{if .Errors.First}} has-error has-feedback{{end}}">
    <label for="First">First name <small class="text-danger">(required)</small></label>
    <input type="text" class="form-control" id="First" name="First" {{if .Form.First}}value="{{.Form.First}}"{{end}}>
    {{if .Errors.First}}<span class="help-block">{{.Errors.First}}</span>{{end}}
  </div>

  <div class="form-group{{if .Errors.Last}} has-error has-feedback{{end}}">
    <label for="Last">Last name <small class="text-danger">(required)</small></label>
    <input type="text" class="form-control" id="Last" name="Last" {{if .Form.Last}}value="{{.Form.Last}}"{{end}}>
    {{if .Errors.Last}}<span class="help-block">{{.Errors.Last}}</span>{{end}}
  </div>

  <div class="form-group{{if .Errors.Email}} has-error has-feedback{{end}}">
    <label for="Email">Email address <small class="text-danger">(required)</small></label>
    <input type="email" class="form-control" id="Email" name="Email" {{if .Form.Email}}value="{{.Form.Email}}"{{end}} placeholder="user@example.org">
    {{if .Errors.Email}}<span class="help-block">{{.Errors.Email}}</span>{{end}}
  </div>

  <div class="form-group{{if .Errors.Password}} has-error has-feedback{{end}}">
    <label for="Password">Create a password <small class="text-danger">(required)</small></label>
    <input type="password" class="form-control" id="Password" name="Password">
    {{if .Errors.Password}}<span class="help-block">{{.Errors.Password}}</span>{{end}}
  </div>

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
  </div><!--end password helper-->

  <div class="form-group{{if .Errors.ConfirmPassword}} has-error has-feedback{{end}}">
    <label for="ConfirmPassword">Confirm your password <small class="text-danger">(required)</small></label>
    <input type="password" name="ConfirmPassword" class="form-control" id="ConfirmPassword">
    {{if .Errors.ConfirmPassword}}<span class="help-block">{{.Errors.ConfirmPassword}}</span>{{end}}
  </div>

  <div class="form-group{{if .Errors.Captcha}} has-error has-feedback{{end}}">
    {{create_captcha}}
    <input type="text" name="captcha" class="form-control" id="InputCaptcha">
    {{if .Errors.Captcha}}<span class="help-block">{{.Errors.Captcha}}</span>{{end}}
  </div>

  <button type="submit" class="btn btn-primary btn-lg">Sign Up</button>
</form>
</div>
