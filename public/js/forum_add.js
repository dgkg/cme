$(window).load(function() {

    //$('form').submit(alert("Bonjour"));

    /* $('form').submit(function(e) {
        console.log($data);
    }); */

    //log.console(contenu);

    $('.post-contenu').ckeditor();

    var editeur = $('.post-contenu');

    var $zoneMsgErreur = $('.msg-erreur');

    $.validate({
        errorMessagePosition: $zoneMsgErreur,
        validateOnBlur: false,
    });
});
