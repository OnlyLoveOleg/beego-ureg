$(document).ready(function() {  
  //you have to use keyup, because keydown will not catch the currently entered value
  $('#Password').keyup(function() { 
 
    // set variables
    var pswd = $(this).val();
    var lengthPass = false 
    var letterPass = false 
    var capitalPass = false 
    var numberPass = false 

    //validate the length
    if ( pswd.length < 8 ) {
      $('#length').removeClass('text-success').addClass('text-danger');
      $('#length span').removeClass('glyphicon-ok').addClass('glyphicon-remove');
    } else {
      $('#length').removeClass('text-danger').addClass('text-success');
      $('#length span').removeClass('glyphicon-remove').addClass('glyphicon-ok');
      lengthPass = true;
    }
    
    //validate lowercase letter
    if ( pswd.match(/[a-z]/) ) {
      $('#letter').removeClass('text-danger').addClass('text-success');
      $('#letter span').removeClass('glyphicon-remove').addClass('glyphicon-ok');
      letterPass = true;
    } else {
      $('#letter').removeClass('text-success').addClass('text-danger');
      $('#letter span').removeClass('glyphicon-ok').addClass('glyphicon-remove');
    }
    
    //validate uppercase letter
    if ( pswd.match(/[A-Z]/) ) {
      $('#capital').removeClass('text-danger').addClass('text-success');
      $('#capital span').removeClass('glyphicon-remove').addClass('glyphicon-ok');
      capitalPass = true;
    } else {
      $('#capital').removeClass('text-success').addClass('text-danger');
      $('#capital span').removeClass('glyphicon-ok').addClass('glyphicon-remove');
    }

    //validate number
    if ( pswd.match(/\d/) ) {
      $('#number').removeClass('text-danger').addClass('text-success');
      $('#number span').removeClass('glyphicon-remove').addClass('glyphicon-ok');
      numberPass = true
    } else {
      $('#number').removeClass('text-success').addClass('text-danger');
      $('#number span').removeClass('glyphicon-ok').addClass('glyphicon-remove');
    }

    //update panel color
    if ( numberPass && capitalPass && letterPass && lengthPass ) {
      $('#pswd_info').removeClass('panel-danger').addClass('panel-success');
    } else {
      $('#pswd_info').removeClass('panel-success').addClass('panel-danger');
    }
    
    }).focus(function() {
      //$('#pswd_info').show();
      $('#pswd_info').removeClass('hidden');
    }).blur(function() {
      $('#pswd_info').addClass('hidden');
    }); 
});
