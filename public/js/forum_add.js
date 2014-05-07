$(window).load(function() {

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

        //Ne pas rafraîchir la page
        event.preventDefault();

        // Récupération des données saisies dans la page
        var $val_titre_post = $('#post-nom').val();
        var $val_categorie_post = $('#post-cat').val();
        var $val_resolu_post = $('#post-etat').val();
        var $val_contenu_post = $('#post-contenu').val();

        // Envois des données en AJAX
        var posting = $.post( "/forum/nouveau/submitform", {
            titre_post : $val_titre_post,
            categorie_post : $val_categorie_post,
            resolu_post : $val_resolu_post,
            contenu_post: $val_contenu_post
        }, function(data, status){

            if (data != "error") {
                afficherSucces();
            } else {
                alert("Une erreur est survenue. Veuillez réessayer.");
            };
        });
    });
});

function afficherSucces() {
    $('.msg-succes').show();
}
