$(window).load(function() {

    $.validate({

        onValidate : function() {

            //Ne pas rafraîchir la page
            event.preventDefault();
        },

        onSuccess : function() {

            // Récupération des données saisies dans la page
            var $val_commentaire = $('#nouv-comm-contenu').val();
            var $val_post_id = $('#nouv-comm-id').val();
            var $val_auteur_id = $('#nouv-comm-auteur').val();

            // Envois des données en AJAX après avoir été trimmé
            var posting = $.post( "/forum/post/nouvcomm", {
                val_commentaire : $val_commentaire.trim(),
                val_post_id : $val_post_id,
                val_auteur_id : $val_auteur_id

            }, function(data, status){

                if (data != "error" && data != "") {

                    // Séparation en deux du string reçu
                    var dataRecue = data.split(":::", 2);

                    var auteur = dataRecue[0];
                    var commentaire = dataRecue[1];

                    // Création du nouveau commentaire avec les infos
                    var nouvCommentaire =

                    "<div class='comm'>" +
                        "<p class='comm-texte'>" + commentaire + "</p>" +
                        "<p class='comm-auteur'>" + auteur + "</p>" +
                        "<p class='comm-date'>" + "" + "</p>" +
                    "</div>"

                    // Affichage dans la page du commentaire tout juste envoyé
                    $('.comms').append(nouvCommentaire);

                } else {
                    alert("Une erreur est survenue. Veuillez réessayer.");
                };
            });
        }
    });
});
