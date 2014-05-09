$(window).load(function() {

    var $zoneMsgErreur = $('.msg-erreur');

    $.validate({
        modules: 'security',
        errorMessagePosition: $zoneMsgErreur,
        validateOnBlur: false,
        scrollToTopOnError : false,

        onValidate : function() {

            //Ne pas rafraîchir la page
            $('.msg-erreur').empty();
            event.preventDefault();
        },

        onSuccess : function() {

            // Récupération des données saisies dans la page
            var $val_insc_prenom = $('#insc-prenom').val().trim();
            var $val_insc_nom = $('#insc-nom').val().trim();
            var $val_insc_email = $('#insc-email').val().trim();
            var $val_insc_mdp = $('#insc-mdp').val().trim();

            // Envois des données en AJAX après avoir été trimmé
            var posting = $.post( "/inscription/submitform", {
                prenom : $val_insc_prenom,
                nom : $val_insc_nom,
                email : $val_insc_email,
                mdp: $val_insc_mdp
            }, function(data, status){

                if (data != "error") {
                    //afficherSucces();
                    alert("Réussi!")

                } else {
                    alert("Une erreur est survenue. Veuillez réessayer.");
                };
            });
        }
    });
});
