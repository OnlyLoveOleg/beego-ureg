<div class="container" style="padding-top: 80px" id="content">
  <div class="row">
    <form class="form-horizontal" method="POST">
      <div class="form-group{{if .Errors.First}} has-error has-feedback{{end}}">
        <label for="First" class="col-sm-2 control-label">First name</label>
        <div class="col-sm-8">
          <input type="text" name="First" id="First" class="form-control" value="{{.User.First}}">
          {{if .Errors.First}}<span class="help-block">{{.Errors.First}}</span>{{end}}
        </div>
      </div>

      <div class="form-group{{if .Errors.Last}} has-error has-feedback{{end}}">
        <label for="Last" class="col-sm-2 control-label">Last name</label>
        <div class="col-sm-8">
          <input type="text" name="Last" class="form-control" id="Last" value="{{.User.Last}}">
          {{if .Errors.Last}}<span class="help-block">{{.Errors.Last}}</span>{{end}}
        </div>
      </div>

      <div class="form-group{{if .Errors.Email}} has-error has-feedback{{end}}">
        <label for="Email" class="col-sm-2 control-label">Email address</label>
        <div class="col-sm-8">
          <input type="email" name="Email" class="form-control" id="Email" value="{{.User.Email}}">
          {{if .Errors.Email}}<span class="help-block">{{.Errors.Email}}</span>{{end}}
        </div>
      </div>

      <div class="form-group{{if .Errors.Phone1}} has-error has-feedback{{end}}">
        <label for="Phone1" class="col-sm-2 control-label">Main phone</label>
        <div class="col-sm-8">
          <input type="tel" name="Phone1" class="form-control" id="Phone1" value="{{.User.Phone1}}" placeholder="Example: 647-123-4567">
          {{if .Errors.Phone1}}<span class="help-block">{{.Errors.Phone1}}</span>{{end}}
        </div>
      </div>

      <div class="form-group{{if .Errors.Phone2}} has-error has-feedback{{end}}">
        <label for="Phone2" class="col-sm-2 control-label">Secondary phone</label>
        <div class="col-sm-8">
          <input type="tel" name="Phone2" class="form-control" id="Phone2" value="{{.User.Phone2}}" placeholder="Example: 647-123-4567">
          {{if .Errors.Phone2}}<span class="help-block">{{.Errors.Phone2}}</span>{{end}}
        </div>
      </div>

      <div class="form-group{{if .Errors.Address}} has-error has-feedback{{end}}">
        <label for="Address" class="col-sm-2 control-label">Address</label>
        <div class="col-sm-8">
          <input type="text" name="Address" class="form-control" id="Address" value="{{.User.Address}}">
          {{if .Errors.Address}}<span class="help-block">{{.Errors.Address}}</span>{{end}}
        </div>
      </div>

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
  <hr>
  <div class="panel panel-danger">
    <div class="panel-heading">
      <h3 class="panel-title">Danger zone</h3>
    </div>
    <div class="panel-body">
      <p class="text-center"><a class="btn btn-danger" href="//{{.base_url}}/accounts/delete">Remove account</a></p>
    </div>
  </div>
</div><!-- end container -->
