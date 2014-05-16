/*
    Fonctions pour le formulaire inscription.html
    1 - Valide le formulaire
    2 - Affiche des messages d'erreurs
    3 - Envoie en Ajax les informations du formulaire
    4 - Affiche un message de validation
*/

// Variables de la page
var $val_prenom;
var $val_nom;
var $val_email;
var $val_pass;
var $val_pass2;
var errorMessage   = new Array();
var urlRedirect = "/inscription";

$(window).load(function() {
    //récurpération du clic sur le formulaire
    $( "#submit" ).click(function( event ) {
        event.preventDefault();
        if(validateForm()!=true){
            displayErrorMessage();
        }else{
            //console.log("envoyer les données du formulaire");
            sendData();
        }
    });
});

/*
 match username
 var usernameRegex = '/^[a-z0-9_-]{3,16}$/'; 
 match password
 var passwordRegex = '/^[a-z0-9_-]{6,18}$/'; 
 Fonction permettant de valides les champs du formulaire forum_add.html
*/
function validateForm(){
    // initialisation de isValid à true
    var isValid = true;
    // Récupération des données saisies dans la page
    $val_prenom = $('#prenom').val();
    $val_nom = $('#nom').val();
    $val_email = $('#email').val();
    $val_pass = $('#pass').val();
    $val_pass2 = $('#pass2').val();

    // création d'une variable d'incrémentation
    var i = 0;
    // vide le tableau des anciens messages
    errorMessage = [];
    // validation si le prénom n'est pas vide
    if ($val_prenom == ""){
        errorMessage[i] = "Prénom : Veuillez renseigner le champs avec votre prénom";
        isValid = false;
        i++;
    };
    // validation si le prénom ne fait pas plus de 200 caractères
    if ($val_prenom.length > 200){
        errorMessage[i] = "Prénom : Votre prénom est trop long, il fait " + $val_titre_post.length + " caractères.";
        isValid = false;
        i++;
    };
    // validation si le nom n'est pas vide
    if ($val_nom == ""){
        errorMessage[i] = "Nom : Veuillez renseigner le champs avec votre nom";
        isValid = false;
        i++;
    };
    // validation si le nom ne fait pas plus de 200 caractères
    if ($val_nom.length > 200){
        errorMessage[i] = "Nom : Votre nom est trop long, il fait " + $val_titre_post.length + " caractères.";
        isValid = false;
        i++;
    };

    // validation si le login est bien un email
    var emailRegex = '^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$'; 
    var pattern = new RegExp(emailRegex);
    validateEmail = pattern.exec($val_email);
    // si il n'est pas valid ajouter un message d'erreur
    if (validateEmail==null){
        errorMessage[i] = "Email : Votre email n'est pas valide";
        isValid = false;
        i++;
    }
    // validation si l'email n'est pas vide
    if ($val_email == ""){
        errorMessage[i] = "Email : Veuillez renseigner le champs avec votre email";
        isValid = false;
        i++;
    };
    // validation si l'email ne fait pas plus de 200 caractères
    if ($val_email.length > 200){
        errorMessage[i] = "Email : Votre email est trop long, il fait " + $val_titre_post.length + " caractères.";
        isValid = false;
        i++;
    };
    // validation pour le password n'est pas vide
    if($val_pass == ""){
        errorMessage[i] = "Password : Veuillez renseigner votre password";
        isValid = false;
        i++;
    };
    // validation si l'email ne fait pas plus de 10 caractères
    if ($val_pass.length > 10){
        errorMessage[i] = "Password : Veuillez prendre un password de maximum 10 caractères alapha numériques";
        isValid = false;
        i++;
    };
    // validation si l'email ne fait pas plus de 10 caractères
    if ($val_pass != $val_pass2 ){
        errorMessage[i] = "Password : Veuillez confirmer votre password";
        isValid = false;
        i++;
    };

    // retour de la variable si valide ou non
    return isValid;
};

// Fonction pour afficher la listes des erreurs
function displayErrorMessage(){
    var nouvCommentaire = "";

    for (var i = 0; i < errorMessage.length; i++) {
        nouvCommentaire += "<p class=\".button-warning\" >" + errorMessage[i]; + "</p>"
    };
    
    // Affichage dans la page du commentaire tout juste envoyé
    $('.msg-erreur').append(nouvCommentaire);
    $('html, body').animate({scrollTop: $(".body").offset().top}, 800);
};

// création d'un formulaire famtôme
// permettant de rediriger la page vers la page mon compte
function sendData()
{
    /*
    console.log("$val_prenom : "+$val_prenom);
    console.log("$val_nom : "+$val_nom);
    console.log("$val_email : "+$val_email);
    console.log("$val_pass : "+$val_pass);
    */
    
    $('<form />')
      .hide()
      .attr({ method : "post" })
      .attr({ action : urlRedirect})
      .append($('<input />').attr("type","hidden").attr({ "name" : "prenom" }).val($val_prenom))
      .append($('<input />').attr("type","hidden").attr({ "name" : "nom" }).val($val_nom))
      .append($('<input />').attr("type","hidden").attr({ "name" : "email" }).val($val_email))
      .append($('<input />').attr("type","hidden").attr({ "name" : "pass" }).val($val_pass))
      .append('<input type="submit" />')
      .appendTo($("body"))
      .submit();
    
};






