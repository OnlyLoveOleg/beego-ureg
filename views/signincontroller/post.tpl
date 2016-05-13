<div class="container" id="content">
  <div class="row">
    <h1>Please login</h1>
    <form method="POST">
      <div class="form-group{{if .Errors.email}} has-error has-feedback{{end}}">
        <label for="InputEmail">Email address</label>
        <input type="email" class="form-control" id="InputEmail" name="email" {{if .Email}}value="{{.Email}}"{{else}}placeholder="Please enter your email"{{end}}>
        {{if .Errors.email}}<span class="help-block">{{.Errors.email}}</span>{{end}}
      </div>
      <div class="form-group{{if .Errors.password}} has-error has-feedback{{end}}">
        <label for="InputPassword">Password</label>
        <input type="password" name="password" class="form-control" id="InputPassword" placeholder="Password">
        {{if .Errors.password}}<span class="help-block">{{.Errors.password}}</span>{{end}}
      </div>
      <button type="submit" class="btn btn-primary btn-lg">Log In</button>
    </form>
    <p><a href="//{{.base_url}}/accounts/forgotpassword">Forgot password?</a></p>
  </div><!-- end row -->
  <hr>
  <div class="row">
    <h4 class="text-center">Or</h4>
  </div><!-- end row -->
  <div class="row">
    <p class="text-center"><a href="//{{.base_url}}/accounts/signup" class="btn btn-success">Create new account</a></p>
  </div><!-- end row -->
</div>
