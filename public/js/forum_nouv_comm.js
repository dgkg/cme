/*
    Fonctions pour le formulaire forum_post.html
    
    1 - Valide le formulaire
    2 - Affiche des messages d'erreurs
    3 - Envoie en Ajax les informations du formulaire
    4 - Affiche un message de validation
*/

var urlAjaxServiceAdd = "/ajax/forum-post/add";
var urlAjaxServiceDel = "/ajax/forum-post/del";

$(window).load(function() {
    // récurpération du clic sur le formulaire
    // permet d'ajouter un commentaire
    $( "#submit" ).click(function( event ) {
        event.preventDefault();
        sendData();
    });
    // récurpération du clic sur le formulaire
    // permet de supprimer un commentaire
    $( ".del-forum-post" ).click(function( event ) {
        event.preventDefault();
        deletData($(this));
    });
});


// Fonction permettant d'envoyer en Ajax les informations du formulaire
function sendData(){
    // Récupération des données saisies dans la page
    var $val_commentaire = $('#nouv-comm-contenu').val();
    var $val_post_id     = $('#nouv-comm-id').val();
    var $val_auteur_id   = $('#nouv-comm-auteur').val();
    // Envois des données en AJAX après avoir été trimmé
    var posting = $.post( urlAjaxServiceAdd, {
        val_commentaire : $val_commentaire.trim(),
        val_post_id     : $val_post_id,
        val_auteur_id   : $val_auteur_id
    }, function(data, status){
        // check si il n'y a pas d'erreurs
        if (data != "error" && data != "") {
            displaySucces(data);
        } else {
            alert("Une erreur est survenue. Veuillez réessayer.");
        };
    });
};

// Fonction permettant de supprimer le post
function deletData(obj){
    var $idCommentToDelete = obj.attr("id");
    // Envois des données en AJAX après avoir été trimmé
    var posting = $.post( urlAjaxServiceDel, {
        id_commentaire : $idCommentToDelete,
    }, function(data, status){
        // check si il n'y a pas d'erreurs
        if (data != "error" && data != "") {
            displayDelete(obj);
        } else {
            alert("Une erreur est survenue. Veuillez réessayer.");
        };
    });
}

// Fonction permettant de supprimer à l'affichage le post
function displayDelete(obj) {
    obj.parent().fadeOut();
}

// affiche un message après l'envoie du formulaire si tout c'est déroulé correctement
function displaySucces(data) {
    // Séparation en deux du string reçu
    var dataRecue    = data.split(":::", 4);
    var auteur       = dataRecue[0];
    var dateCreation = dataRecue[1];
    var commentaire  = dataRecue[2];
    var idCommentaire  = dataRecue[3];
    // Création du nouveau commentaire avec les infos
    var nouvCommentaire =
    "<div class='comm'>" +
        "<img id='"+idCommentaire+"' class='del-forum-post' src='/img/ui/circle_delete.png'>" +
        "<p class='comm-texte'>" + commentaire + "</p>" +
        "<p class='comm-auteur'>" + auteur + "</p>" +
        "<p class='comm-date'>" + dateCreation + "</p>" +
    "</div>"
    // Affichage dans la page du commentaire tout juste envoyé
    $('.comms').append(nouvCommentaire);
    //$( ".comm:last-child" ).fadeIn( "slow" );
    $( ".comm:last-child" ).hide();
    $( ".comm:last-child" ).fadeIn();
    $('#nouv-comm-contenu').val('');
};


