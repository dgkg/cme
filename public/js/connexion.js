/*
    Fonctions pour le formulaire connexion.html
    1 - Valide le formulaire
    2 - Affiche des messages d'erreurs
    3 - Envoie en Ajax les informations du formulaire
    4 - Affiche un message de validation
*/

// Variables de la page
var $val_login;
var $val_pass;
var errorMessage   = new Array();
var urlAjaxService = "/connect";
var urlRedirect = "/connexion";

$(window).load(function() {
    //récurpération du clic sur le formulaire
    $( "#submit" ).click(function( event ) {
    	event.preventDefault();
    	if(validateForm()!=true){
    		displayErrorMessage();
    	}else{
    		console.log("envoyer les données du formulaire");
    		//event.bubbles;
    		//$( ".connexion-form" ).submit();
    		redirect();
    	}
    	
    	/*
        event.preventDefault();
        if (validateForm()){
            sendData();
        }else{
            displayErrorMessage();
        }
        */
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
    $val_login = $('#login').val();
    $val_pass  = $('#pass').val();

    // création d'une variable d'incrémentation
    var i = 0;
    // validation si le login n'est pas vide
    if ($val_login == ""){
        errorMessage[i] = "Login : Veuillez renseigner le champs avec votre email";
        isValid = false;
        i++;
    };
    // validation si le login ne fait pas plus de 100 caractères
    if ($val_login.length > 200){
        errorMessage[i] = "Login : Votre login fait "+$val_titre_post.length+" caractères, il ne correspond pas à votre email";
        isValid = false;
        i++;
    };
    // validation si le login est bien un email
	var emailRegex = '^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$'; 
	var pattern = new RegExp(emailRegex);
    validateEmail = pattern.exec($val_login);
    // si il n'est pas valid ajouter un message d'erreur
    if (validateEmail==null){
        errorMessage[i] = "Login : Votre email n'est pas valide";
        isValid = false;
        i++;
    }
    // validation pour le password n'est pas vide
    if($val_pass == ""){
        errorMessage[i] = "Password : Veuillez renseigner votre password";
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

// Fonction permettant d'envoyer en Ajax les informations du formulaire
function sendData(){
    // Envois des données en AJAX après avoir été trimmé
    var posting = $.post( urlAjaxService, {
        login : $val_login,
        pass  : $val_pass,
    }, function(data, status){
        // si il n'y a pas d'erreurs
        if (data == "all good") {
            redirect();
        // si il n'y a une erreur afficher une alerte
        } else {
            alert("Une erreur est survenue. Veuillez réessayer.");
        };
    });
};

// création d'un formulaire famtôme
// permettant de rediriger la page vers la page mon compte
function redirect()
{
	$('<form />')
      .hide()
      .attr({ method : "post" })
      .attr({ action : urlRedirect})
      .append($('<input />').attr("type","hidden").attr({ "name" : "login" }).val($val_login))
      .append($('<input />').attr("type","hidden").attr({ "name" : "pass" }).val($val_pass))
      .append('<input type="submit" />')
      .appendTo($("body"))
      .submit();
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

