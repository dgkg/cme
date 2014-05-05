$(window).load(function() {

    //$('form').submit(alert("Bonjour"));

    /* $('form').submit(function(e) {
        console.log($data);
    }); */

    //log.console(contenu);

    $('.nouv-post-form').submit(function(e){
          e.preventDefault();
    });

    $('.post-contenu').ckeditor();

    var editeur = $('.post-contenu');

    var $zoneMsgErreur = $('.msg-erreur');

    $.validate({
        errorMessagePosition: $zoneMsgErreur,
        validateOnBlur: false,
        onSuccess : function() {
            //afficherSucces();
        }
    });

    $( ".nouv-post-form" ).submit(function( event ) {

        // Get some values from elements on the page:
        var $form = $( this );

        // Send the data using post
        var posting = $.post( "/submitForm", $form.serialize() );

        // Put the results in a div
        posting.done(function( data ) {
            afficherSucces();
        });
    });
});

function afficherSucces() {
    $('.msg-succes').show();
}
