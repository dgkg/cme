$(window).load(function() {

    $('.form-bio').submit( function(event) {
        event.preventDefault();
        saveBio();
    });

    $('.form-social').submit( function(event) {
        event.preventDefault();
        saveSocial();
    });

    $('.form-gradu').submit( function(event) {
        event.preventDefault();
        saveGraduation();
    });

    $('.form-compte').submit( function(event) {
        event.preventDefault();
        saveNomUtilisateur();
    });

    $('#supp-compte').click( function(event) {
        event.preventDefault();
        $('#supp-compte').hide();
        $('#msg-confirm').append( "<p class='msg'>Êtes-vous absolument certain de vouloir supprimer définitivement votre compte?</p>");
        $('#msg-confirm').append( "<input type='submit' class='pure-button button-error' value='Je confirme, je souhaite supprimer mon compte.'>" );
    });

    $('.form-supp').submit( function(event) {
        event.preventDefault();
        supprimerCompte();
    });
});

function afficherSucces() {
    $('.msg-succes').show();
    //$('.formulaire').hide();
    $('html, body').animate({
        scrollTop: $(".msg-succes").offset().top
    }, 500);
}

/*
function savePhotoProfil() {
    //    @todo : Envoyer la photo de profil en AJAX

    var idUser = $('#id-user').val();
    var imgPhoto = $('#bio-text').val();
    var strSection = "savePhoto";
    var requete = $.post('/mon-compte/update', { idUser : idUser, section : strSection, biographie : strBiographie }, function(data, status){
        if (data != "error") {
            afficherSucces();
        } else {
            alert("Une erreur est survenue. Veuillez réessayer.");
        };
    });

} */

function saveBio() {
    var idUser = $('#id-user').val();
    var strBiographie = $('#bio-text').val();
    var strSection = "saveBio";
    var requete = $.post('/mon-compte/update', { idUser : idUser, section : strSection, biographie : strBiographie }, function(data, status){
        if (data != "error") {
            afficherSucces();
        } else {
            alert("Une erreur est survenue. Veuillez réessayer.");
        };
    });
}

function saveSocial() {
    var idUser = $('#id-user').val();
    var strFb = $('#url-facebook').val();
    var strTw = $('#url-twitter').val();
    var strLi = $('#url-linkedin').val();
    var strSection = "saveSocial";
    var requete = $.post('/mon-compte/update', { idUser : idUser, section : strSection, facebook : strFb, twitter : strTw, linkedin : strLi }, function(data, status){
        if (data != "error") {
            afficherSucces();
        } else {
            alert("Une erreur est survenue. Veuillez réessayer.");
        };
    });
}

function saveGraduation() {
    var idUser = $('#id-user').val();
    var intGraduation = $('#graduation').val();
    var strSection = "saveGraduation";
    var requete = $.post('/mon-compte/update', { idUser : idUser, section : strSection, graduation : intGraduation }, function(data, status){
        if (data != "error") {
            afficherSucces();
        } else {
            alert("Une erreur est survenue. Veuillez réessayer.");
        };
    });
}

function saveNomUtilisateur() {
    var idUser = $('#id-user').val();
    var strPrenom = $('#prenom-utilisateur').val();
    var strNom = $('#nom-utilisateur').val();
    var strSection = "saveNomUtilisateur";
    var requete = $.post('/mon-compte/update', { idUser : idUser, section : strSection, prenom : strPrenom, nom : strNom }, function(data, status){
        if (data != "error") {
            afficherSucces();
        } else {
            alert("Une erreur est survenue. Veuillez réessayer.");
        };
    });
}

function supprimerCompte() {
    var idUser = $('#id-user').val();
    var strSection = "supprimerCompte";
    var requete = $.post('/mon-compte/update', { idUser : idUser, section : strSection}, function(data, status){
        if (data != "error") {
            afficherSucces();
        } else {
            alert("Une erreur est survenue. Veuillez réessayer.");
        };
    });
}
