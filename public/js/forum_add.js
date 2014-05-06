$(window).load(function() {

    //$('form').submit(alert("Bonjour"));

    /* $('form').submit(function(e) {
        console.log($data);
    }); */

    //log.console(contenu);

    /* $('.nouv-post-form').submit(function(e){
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

    $( ".nouv-post-form" ).submit(function( event ) {

        //Ne pas rafraîchir la page
        event.preventDefault();

        // Get some values from elements on the page:
        var $form = $( this );

        // Send the data using post
        var posting = $.post( "/submitForm", $form.serialize() );

        // Afficher la confirmation
        posting.done(function( data ) {
            alert("La page /submitForm n'existe pas.");
            afficherSucces();
        });
    });
});

function afficherSucces() {
    $('.msg-succes').show();
}
