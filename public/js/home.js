/*
    Fonctions pour la page d'accueil index.html
    1 - Envoie en Ajax les informations du formulaire sur la catégorie sélectionnée
    2 - Affiche la liste des pojets de la catégorie sélectionnée
*/


var $val_idCat;
var urlAjaxService = "/ajax/userProject/getByIdCat/";


$(window).load(function() {
    
    //récurpération du clic sur le formulaire
    $( ".cat-select" ).click(function( event ) {
    	event.preventDefault();
    	// réinitialise urlAjaxService
    	urlAjaxService = "/ajax/userProject/getByIdCat/";
        // on récupère l'id
        $val_idCat = $(this).attr("href");
        // on valide que l'id est bien un nombre
        isValid    = validate();
        if(isValid){
        	sendData();
        }
    });
});

// Fonction permettant de valider l'id
function validate(){
    // initialisation de isValid à true
    var isValid = true;
    // validation si l'id est bien un nombre
    var numberRegex = '^[0-9]+'; 
    var pattern     = new RegExp(numberRegex);
    validateNumber  = pattern.exec($val_idCat);
    // si il n'est pas valid ajouter un message d'erreur
    console.log( validateNumber );
    if (validateNumber==null){
    	isValid = false;
    }

    return isValid
}


// Fonction permettant d'envoyer en Ajax les informations du formulaire
function sendData(){
	// construction de l'url à appeler
    urlAjaxService = urlAjaxService + $val_idCat;
    console.log( urlAjaxService );
    // Envois des données en AJAX
    var posting = $.post( urlAjaxService, {}, function(data, status){
        // si il n'y a pas d'erreurs récupérer l'id du forum créé
        if (data != "error") {
            displayResult(data);
        // si il n'y a une erreur afficher une alerte
        } else {
            alert("Une erreur est survenue. Veuillez réessayer.");
        };
    });
};

// permet d'afficher les résultats
function displayResult(data){
	$('.projets').children().hide();
	console.log("*** displayResult ***");
	//console.log( data );
	var ups = jQuery.parseJSON( data );
	for (var i = ups.length - 1; i >= 0; i--) {
		up = ups[i];
		console.log(up.Title);
		var Projet =
		"<div class='img-projet'>"+
            "<img class='img-du-projet' src='"+up.ImageUrl+"'>"+
            "<div class='over-projet'>"+
                "<h3>"+up.Title+"</h3>"+
                "<p>"+up.Description+"</p>"+
                "<a href='"+up.Url+"'>Voir le projet</a>"+
            "</div>"+
        "</div>";

	    // Affichage dans la page du commentaire tout juste envoyé
	    $('.projets').append(Projet);
	};
}