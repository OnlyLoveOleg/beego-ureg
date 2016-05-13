package models

import (
  "regexp"
  "github.com/astaxie/beego/validation"
)


type FormAccountSignUp struct {
  First           string `valid:"Required;MinSize(2);Alpha"`
  Last            string `valid:"Required;MinSize(2);Alpha"`
  Email           string `valid:"Email"`
  Password        string `valid:"MinSize(8);MaxSize(36)"`
  ConfirmPassword string `valid:"Required"`
}


type FormAccountUpdate struct {
  First           string `valid:"Required;MinSize(2);Alpha"`
  Last            string `valid:"Required;MinSize(2);Alpha"`
  Email           string `valid:"Email"`
  Phone1          string
  Phone2          string
  Address         string
  CurrentPassword string `valid:"Required"`
}


type FormPasswordUpdate struct {
  NewPassword              string `valid:"MinSize(8);MaxSize(36)"`
  NewPasswordConfirmation  string `valid:"Required"`
  CurrentPassword          string `valid:"Required"`
}


func (form *FormPasswordUpdate) Valid(v *validation.Validation) {
  // Check passwords match
  if form.NewPassword != form.NewPasswordConfirmation {
    v.SetError("NewPassword", "Passwords don't match")
    v.SetError("NewPasswordConfirmation", "Passwords don't match")
  }

  // Check password contain capital letter
  r, _ := regexp.Compile(`[A-Z]`)

  if !r.MatchString(form.NewPassword) {
    v.SetError("NewPassword", "Password must contain at least one capital letter")
  }

  // Check password contain lowercase letter
  r, _ = regexp.Compile(`[a-z]`)

  if !r.MatchString(form.NewPassword) {
    v.SetError("NewPassword", "Password must contain at least one lower case letter")
  }

  // Check password contain number
  r, _ = regexp.Compile(`[0-9]`)

  if !r.MatchString(form.NewPassword) {
    v.SetError("NewPassword", "Password must contain at least one number")
  }

}//end FormPasswordUpdate.valid func


func (form *FormAccountSignUp) Valid(v *validation.Validation) {

  // Check passwords match
  if form.Password != form.ConfirmPassword {
    v.SetError("Password", "Passwords don't match")
    v.SetError("ConfirmPassword", "Passwords don't match")
  }

  // Check password contain capital letter
  r, _ := regexp.Compile(`[A-Z]`)

  if !r.MatchString(form.Password) {
    v.SetError("Password", "Password must contain at least one capital letter")
  }

  // Check password contain lowercase letter
  r, _ = regexp.Compile(`[a-z]`)

  if !r.MatchString(form.Password) {
    v.SetError("Password", "Password must contain at least one lower case letter")
  }

  // Check password contain number
  r, _ = regexp.Compile(`[0-9]`)

  if !r.MatchString(form.Password) {
    v.SetError("Password", "Password must contain at least one number")
  }

}//end FormAccountSignUp validation



func (form *FormAccountUpdate) Valid(v *validation.Validation) {

  valid := validation.Validation{}

  if form.Phone1 != "" {
    matched, err := regexp.MatchString("^\\d{3}\\-\\d{3}\\-\\d{4}$", form.Phone1)

    if err != nil {
      v.SetError("Phone1", "Something wierd is going on here.")
    } else if !matched {
      v.SetError("Phone1", "Phone should be something like: 647-123-4567")
    }
  } // end Phone1 validation if block

  if form.Phone2 != "" {
    matched, err := regexp.MatchString("^\\d{3}\\-\\d{3}\\-\\d{4}$", form.Phone2)

    if err != nil {
      v.SetError("Phone2", "Something wierd is going on here.")
    } else if !matched {
      v.SetError("Phone2", "Phone should be something like: 647-123-4567")
    }
  } // end Phone2 validation if block

  for _, err := range valid.Errors {
    v.SetError(err.Key, err.Message)
  }

}//end FormAccountUpdate validation
