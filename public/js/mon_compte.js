$(window).load(function() {

    // Initialisation des handlers vers chaque fonction associée

    // Sauvegarde de la photo de profil
    $('.form-photo').submit( function(event) {
        event.preventDefault();
        savePhotoProfil();
    });

    // Sauvegarde de l'image de cover
    $('.form-cover').submit( function(event) {
        event.preventDefault();
        savePhotoCover();
    })

    // Sauvegarde de la bio
    $('.form-bio').submit( function(event) {
        event.preventDefault();
        saveBio();
    });

    // Sauvegarde de les réseaux sociaux
    $('.form-social').submit( function(event) {
        event.preventDefault();
        saveSocial();
    });

    // Sauvegarde de l'année de graduation
    $('.form-gradu').submit( function(event) {
        event.preventDefault();
        saveGraduation();
    });

    // Sauvegarde du nom d'utilisateur
    $('.form-compte').submit( function(event) {
        event.preventDefault();
        saveNomUtilisateur();
    });

    // Suppression du compte
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

    $('#photo-upload').fileValidator({
        onValidation: function(files){
            $('#photo-upload').css("background-color","none");
            $('#submitPhoto').removeAttr("disabled");
        },
        onInvalid:    function(validationType, file) {
            $('#photo-upload').css("background-color","red");
            $('#submitPhoto').attr("disabled", true);
        },
        maxSize:      '2m',
        type:         'image'
    });
});

function afficherSucces(section) {
    var msgSection = "msg-" + section;
    $('.' + msgSection).fadeIn('fast');
}

// Envoi de l'image de profil
function savePhotoProfil() {

    // Initialisation des variables de base
    var idUser = $('#id-user').val();
    var imgUser = $('#photo-upload').get(0).files[0];
    var strSection = "savePhoto";

    // Création de l'objet formulaire
    var formData = new FormData();

    // Ajout d'une entrée dans l'objet formulaire
    formData.append('photo-upload', imgUser, imgUser.name);

    // Création de la requête
    var xhr = new XMLHttpRequest();

    xhr.open('POST', '/mon-compte/avatar', true);

    // Vérification de l'envoi du formulaire
    xhr.onload = function () {
        if (xhr.status === 200) {
            afficherSucces("photo");
        } else {
            alert('Une erreur est survenue. Veuillez réessayer.');
        }
    };

    // Envoi du formulaire et de son contenu
    xhr.send(formData);
}

// Envoi de l'image de cover
function savePhotoCover() {

    // Initialisation des variables de base
    var idUser = $('#id-user').val();
    var imgCover = $('#cover-upload').get(0).files[0];
    var strSection = "savePhoto";

    // Création de l'objet formulaire
    var formData = new FormData();

    // Ajout d'une entrée dans l'objet formulaire
    formData.append('cover-upload', imgCover, imgCover.name);

    // Création de la requête
    var xhr = new XMLHttpRequest();

    xhr.open('POST', '/mon-compte/cover', true);

    // Vérification de l'envoi du formulaire
    xhr.onload = function () {
        if (xhr.status === 200) {
            afficherSucces("cover");
        } else {
            alert('Une erreur est survenue. Veuillez réessayer.');
        }
    };

    // Envoi du formulaire et de son contenu
    xhr.send(formData);
}

function saveBio() {
    var idUser = $('#id-user').val();
    var strBiographie = $('#bio-text').val();
    var strSection = "saveBio";
    var requete = $.post('/ajax/mon-compte/update', { idUser : idUser, section : strSection, biographie : strBiographie }, function(data, status){
        if (data != "error") {
            afficherSucces("bio");
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
    var requete = $.post('/ajax/mon-compte/update', { idUser : idUser, section : strSection, facebook : strFb, twitter : strTw, linkedin : strLi }, function(data, status){
        if (data != "error") {
            afficherSucces("social");
        } else {
            alert("Une erreur est survenue. Veuillez réessayer.");
        };
    });
}

function saveGraduation() {
    var idUser = $('#id-user').val();
    var intGraduation = $('#graduation').val();
    var strSection = "saveGraduation";
    var requete = $.post('/ajax/mon-compte/update', { idUser : idUser, section : strSection, graduation : intGraduation }, function(data, status){
        if (data != "error") {
            afficherSucces("gradu");
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
    var requete = $.post('/ajax/mon-compte/update', { idUser : idUser, section : strSection, prenom : strPrenom, nom : strNom }, function(data, status){
        if (data != "error") {
            afficherSucces("compte");
        } else {
            alert("Une erreur est survenue. Veuillez réessayer.");
        };
    });
}

function supprimerCompte() {
    var idUser = $('#id-user').val();
    var strSection = "supprimerCompte";
    var requete = $.post('/ajax/mon-compte/update', { idUser : idUser, section : strSection}, function(data, status){
        if (data != "error") {
            afficherSucces();
        } else {
            alert("Une erreur est survenue. Veuillez réessayer.");
        };
    });
}
