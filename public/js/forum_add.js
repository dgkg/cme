$(window).load(function() {

    //$('form').submit(alert("Bonjour"));

    /* $('form').submit(function(e) {
        console.log($data);
    }); */

    //log.console(contenu);

    /* $('#submit').click(function(e){
          // do whatever actions you need here
          e.preventDefault();
    }); */

    $('.post-contenu').ckeditor();

    var editeur = $('.post-contenu');

    var $zoneMsgErreur = $('.msg-erreur');

    $.validate({
        errorMessagePosition: $zoneMsgErreur,
        validateOnBlur: false,
        onSuccess : function() {
            afficherSucces();
        }
    });
});

function afficherSucces() {
    $('.msg-succes').show();
}
