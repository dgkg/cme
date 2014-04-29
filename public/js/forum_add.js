CKEDITOR.replace( 'post-contenu', {

    //var data = CKEDITOR.instances.post-contenu.getData();

});

$(window).load(function() {

    //$('form').submit(alert("Bonjour"));

    /* $('form').submit(function(e) {
        console.log($data);
    }); */

    var $zoneMsgErreur = $('.msg-erreur');

    $.validate({
        errorMessagePosition: $zoneMsgErreur,
        validateOnBlur: false,
    });
});
