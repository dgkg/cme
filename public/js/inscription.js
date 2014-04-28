$(window).load(function() {

    var $zoneMsgErreur = $('.msg-erreur');

    $.validate({
        errorMessagePosition: $zoneMsgErreur,
        validateOnBlur: false,
        scrollToTopOnError : false
    });
});
