$(window).load(function() {

    moment.lang('fr', {
        months : "janvier_février_mars_avril_mai_juin_juillet_août_septembre_octobre_novembre_décembre".split("_"),
        monthsShort : "janv._févr._mars_avr._mai_juin_juil._août_sept._oct._nov._déc.".split("_"),
        weekdays : "dimanche_lundi_mardi_mercredi_jeudi_vendredi_samedi".split("_"),
        weekdaysShort : "dim._lun._mar._mer._jeu._ven._sam.".split("_"),
        weekdaysMin : "Di_Lu_Ma_Me_Je_Ve_Sa".split("_"),
        longDateFormat : {
            LT : "HH:mm",
            L : "DD/MM/YYYY",
            LL : "D MMMM YYYY",
            LLL : "D MMMM YYYY LT",
            LLLL : "dddd D MMMM YYYY LT"
        },
        calendar : {
            sameDay: "[Aujourd'hui à] LT",
            nextDay: '[Demain à] LT',
            nextWeek: 'dddd [à] LT',
            lastDay: '[Hier à] LT',
            lastWeek: 'dddd [dernier à] LT',
            sameElse: 'L'
        },
        relativeTime : {
            future : "dans %s",
            past : "il y a %s",
            s : "quelques secondes",
            m : "une minute",
            mm : "%d minutes",
            h : "une heure",
            hh : "%d heures",
            d : "un jour",
            dd : "%d jours",
            M : "un mois",
            MM : "%d mois",
            y : "une année",
            yy : "%d années"
        },
        ordinal : function (number) {
            return number + (number === 1 ? 'er' : '');
        },
        week : {
            dow : 1, // Monday is the first day of the week.
            doy : 4  // The week that contains Jan 4th is the first week of the year.
        }
    });

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
                        "<p class='comm-date'>" + moment().format('Do MMMM YYYY à h:mm:ss A'); + "</p>" +
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
