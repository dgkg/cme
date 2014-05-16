/*
    Fonctions pour le formulaire forum_add.html
    1 - Initialise CKeditor
    2 - Valide le formulaire
    3 - Affiche des messages d'erreurs
    4 - Envoie en Ajax les informations du formulaire
    5 - Affiche un message de validation
*/

// Variables de la page
var $val_titre_post;
var $val_categorie_post;
var $val_resolu_post;
var $val_contenu_post;
var $val_user_id;
var $val_forum_id;
var errorMessage   = new Array();
var urlAjaxService = "/forum/nouveau/submitform";

$(window).load(function() {
    $('.msg-succes').hide();
    // initialisation de ckeditor
    $('.post-contenu').ckeditor();
    //récurpération du clic sur le formulaire
    $( "#submit" ).click(function( event ) {
        event.preventDefault();
        if (validateForm()){
            sendData();
        }else{
            displayErrorMessage();
        }
    });
});

// Fonction permettant de valides les champs du formulaire forum_add.html
function validateForm(){
    // initialisation de isValid à true
    var isValid = true;
    // Récupération des données saisies dans la page
    $val_titre_post     = $('#post-nom').val();
    $val_categorie_post = $('#post-cat').val();
    $val_resolu_post    = $('#post-etat').val();
    $val_contenu_post   = $('#post-contenu').val();
    $val_user_id        = $('#user-id').val();
    $val_forum_id       = $('#forum-id').val();
    // suppression des espaces sur les titre et le contenu
    $val_titre_post     = $val_titre_post.trim();
    $val_contenu_post   = $val_contenu_post.trim();

    // création d'une variable d'incrémentation
    var i = 0;
    // vide le tableau des anciens messages
    errorMessage = [];
    // validation si le titre n'est pas vide
    if ($val_titre_post == ""){
        errorMessage[i] = "Titre du sujet : Veuillez renseigner le titre";
        isValid = false;
        i++;
    };
    // validation si le titre ne fait pas plus de 100 caractères
    if ($val_titre_post.length > 100){
        errorMessage[i] = "Titre du sujet : Votre titre fait "+$val_titre_post.length+" caractères, veuillez le raccourcir à 100.";
        isValid = false;
        i++;
    };
    // validation sur la catégorie voir si elle est bien sélectionnée
    if($val_categorie_post == -1){
        errorMessage[i] = "Titre du sujet : La catégorie n a pas été renseignée";
        isValid = false;
        i++;
    };
    // validation pour le contenu voir si il n'est pas vide
    if($val_contenu_post == ""){
        errorMessage[i] = "Contenu : Veuillez renseigner le contenu de votre question";
        isValid = false;
        i++;
    };
    // validation pour le contenu voir si il n'est pas vide
    if($val_contenu_post > 10000){
        errorMessage[i] = "Contenu : Votre contenu est trop important";
        isValid = false;
        i++;
    };
    // retour de la variable si valide ou non
    return isValid;
};

// Fonction pour afficher la listes des erreurs
function displayErrorMessage(){
    var nouvCommentaire="";

    console.log(errorMessage);

    for (var i = 0; i < errorMessage.length; i++) {
        console.log(errorMessage[i] + " / "+i);
        nouvCommentaire += "<p class=\".button-warning\" >" + errorMessage[i]; + "</p>"
    };
    
    // Affichage dans la page du commentaire tout juste envoyé
    $('.msg-erreur').append(nouvCommentaire);
    $('html, body').animate({scrollTop: $(".body").offset().top}, 800);
};

// Fonction permettant d'envoyer en Ajax les informations du formulaire
function sendData(){
    // Envois des données en AJAX après avoir été trimmé
    var posting = $.post( urlAjaxService, {
        titre_post      : $val_titre_post,
        categorie_post  : $val_categorie_post,
        resolu_post     : $val_resolu_post,
        contenu_post    : $val_contenu_post,
        user_id         : $val_user_id,
        forum_id        : $val_forum_id,
    }, function(data, status){
        // si il n'y a pas d'erreurs récupérer l'id du forum créé
        if (data != "error") {
            var dataRecue = data.split(":::", 2);
            var idForum   = dataRecue[0];
            displaySucces(idForum);
        // si il n'y a une erreur afficher une alerte
        } else {
            alert("Une erreur est survenue. Veuillez réessayer.");
        };
    });
};

// affiche un message après l'envoie du formulaire si tout c'est déroulé correctement
function displaySucces(idForum) {
    // creation du lien et animation
    $('.msg-erreur').hide();
    $("#forum-id-created").attr("href","/forum/post/"+idForum);
    $('.nouv-post-form').hide();
    $('.msg-succes').fadeIn();
    $('html, body').animate({scrollTop: $(".msg-succes").offset().top}, 750);
};

