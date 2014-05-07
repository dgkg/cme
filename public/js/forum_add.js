$(window).load(function() {

    $('.post-contenu').ckeditor();

    var editeur = $('.post-contenu');

    var $zoneMsgErreur = $('.msg-erreur');

    $.validate({
        errorMessagePosition: $zoneMsgErreur,
        validateOnBlur: false,

        onValidate : function() {

            //Ne pas rafraîchir la page
            event.preventDefault();
        },

        onSuccess : function() {

            // Récupération des données saisies dans la page
            var $val_titre_post = $('#post-nom').val();
            var $val_categorie_post = $('#post-cat').val();
            var $val_resolu_post = $('#post-etat').val();
            var $val_contenu_post = $('#post-contenu').val();

            // Envois des données en AJAX après avoir été trimmé
            var posting = $.post( "/forum/nouveau/submitform", {
                titre_post : $val_titre_post.trim(),
                categorie_post : $val_categorie_post.trim(),
                resolu_post : $val_resolu_post.trim(),
                contenu_post: $val_contenu_post.trim()
            }, function(data, status){

                if (data != "error") {
                    afficherSucces();

                } else {
                    alert("Une erreur est survenue. Veuillez réessayer.");
                };
            });
        }
    });
});

function afficherSucces() {
    $('.msg-succes').show();
    $('.nouv-post-form').hide();
    $('html, body').animate({
        scrollTop: $(".msg-succes").offset().top
    }, 750);

}
