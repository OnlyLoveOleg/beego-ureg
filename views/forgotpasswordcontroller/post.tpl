<div class="container" id="content">
  <div class="row">
    <h1>Reset password</h1>
    <br>

    <form method="POST">
      <div class="form-group{{if .Errors.email}} has-error has-feedback{{end}}">
        <label for="InputEmail">Email address</label>
        <input type="email" class="form-control" id="InputEmail" name="email" placeholder="Enter email here">
        {{if .Errors.email}}<span class="help-block">{{.Errors.email}}</span>{{end}}
      </div>

      <div class="form-group{{if .Errors.Captcha}} has-error has-feedback{{end}}">
        {{create_captcha}}
        <input type="text" name="captcha" class="form-control" id="InputCaptcha">
        {{if .Errors.Captcha}}<span class="help-block">{{.Errors.Captcha}}</span>{{end}}
      </div>

      <button type="submit" class="btn btn-primary">Request password reset</button>
    </form>
  </div>
</div>
