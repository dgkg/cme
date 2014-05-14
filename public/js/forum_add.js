$(window).load(function() {

    $('.post-contenu').ckeditor();

    var editeur        = $('.post-contenu');
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
            var $val_titre_post     = $('#post-nom').val();
            var $val_categorie_post = $('#post-cat').val();
            var $val_resolu_post    = $('#post-etat').val();
            var $val_contenu_post   = $('#post-contenu').val();
            var $val_user_id        = $('#user-id').val();
            var $val_forum_id       = $('#forum-id').val();
            // Envois des données en AJAX après avoir été trimmé
            var posting = $.post( "/forum/nouveau/submitform", {
                titre_post      : $val_titre_post.trim(),
                categorie_post  : $val_categorie_post.trim(),
                resolu_post     : $val_resolu_post,
                contenu_post    : $val_contenu_post.trim(),
                user_id         : $val_user_id.trim(),
                forum_id        : $val_forum_id,
            }, function(data, status){
                // si il n'y a pas d'erreurs récupérer l'id du forum créé
                if (data != "error") {
                    var dataRecue = data.split(":::", 2);
                    var idForum   = dataRecue[0];
                    afficherSucces(idForum);
                // si il n'y a une erreur afficher une alerte
                } else {
                    alert("Une erreur est survenue. Veuillez réessayer.");
                };
            });
        }
    });
});

function afficherSucces(idForum) {
    // creation du lien et animation
    $("#forum-id-created").attr("href","/forum/post/"+idForum);
    $('.nouv-post-form').hide();
    $('.msg-succes').fadeIn();
    $('html, body').animate({scrollTop: $(".msg-succes").offset().top}, 750);
}
