$(window).load(function() {

    var $zoneMsgErreur = $('.msg-erreur');

    $.validate({
        modules: 'security',
        errorMessagePosition: $zoneMsgErreur,
        validateOnBlur: false,
        scrollToTopOnError : false
    });
});
