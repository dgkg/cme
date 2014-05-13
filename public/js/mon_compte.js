$(window).load(function() {

    var $zoneMsgErreur = $('.msg-erreur');

    // Validation de l'image de profil
    $.validate({
        form : '.form-photo',
        errorMessagePosition: $zoneMsgErreur,
        validateOnBlur: false,
        showHelpOnFocus : false,
        addSuggestions : false,
        scrollToTopOnError : false,

        onValidate : function() {

            //Ne pas rafraîchir la page
            event.preventDefault();
        },

        onSuccess : function() {

            // Récupération des données saisies dans la page
            var $section = "photo";
            var $photo_profil = $('.photo-upload').get(0).files[0];

            var requete = $.ajax({
                type: "POST",
                url: "/mon-compte/update",
                contentType: false,
                processData: false,
                data: {section : $section, photo_profil : $photo_profil}
            });

            requete.done(function( msg ) {
                alert(msg);
            });

            requete.fail(function( msg ) {
                // Faire qquechose
            });

            /*
            // Envois des données en AJAX après avoir été trimmé
            var posting = $.post( "/mon-compte/update", {
                section : $section,
                photo_profil : $photo_profil,
            }, function(data, status){

                if (data != "error") {

                    afficherSucces();
                } else {

                    alert("Une erreur est survenue. Veuillez réessayer.");
                };
            });
            */
        }
    });

    // Validation du formulaire social (Facebook, Twitter et LinkedIn)
    $.validate({
        form : '.form-social',
        errorMessagePosition: $zoneMsgErreur,
        validateOnBlur: false,
        showHelpOnFocus : false,
        addSuggestions : false,
        scrollToTopOnError : false,

        onValidate : function() {

            //Ne pas rafraîchir la page
            event.preventDefault();
        },

        onSuccess : function() {

            // Récupération des données saisies dans la page
            var $section = "social";
            var $val_facebook = $('#url-facebook').val();
            var $val_twitter = $('#url-twitter').val();
            var $val_linkedin = $('#url-linkedin').val();

            // Envois des données en AJAX après avoir été trimmé
            var posting = $.post( "/mon-compte/update", {
                section : $section,
                val_facebook : $val_facebook.trim(),
                val_twitter : $val_twitter.trim(),
                val_linkedin : $val_linkedin.trim(),
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
    $('.formulaire').hide();
    $('html, body').animate({
        scrollTop: $(".msg-succes").offset().top
    }, 750);

}
