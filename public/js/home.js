/*
    Fonctions pour la page d'accueil index.html
    1 - Envoie en Ajax les informations du formulaire sur la catégorie sélectionnée
    2 - Affiche la liste des pojets de la catégorie sélectionnée
*/


var $val_idCat=0;
var urlAjaxService = "/ajax/userProject/getByIdCat/";
var catSelected = false;
var page = 0;
var isMore = false;

$(window).load(function() {

    //récurpération du clic sur le formulaire
    $( ".cat-select" ).click( function( event ) {
    	event.preventDefault();
    	// réinitialise urlAjaxService
    	urlAjaxService = "/ajax/userProject/getByIdCat/";
        // réinitialise is more
        isMore = false;
        // on récupère l'id
        $val_idCat = $(this).attr("data-idCat");
        // on valide que l'id est bien un nombre
        isValid    = validateNumber($val_idCat);
        if(isValid){
        	//sendData();
            filterImgProjets($val_idCat);
        }
    });

    //récurpération du clic sur bouton "Afficher plus de projets"
    $( ".btn-plus-de-projets" ).click(function( event ) {
        event.preventDefault();
        // réinitialise urlAjaxService
        urlAjaxService = "/ajax/userProject/getByIdCat/";
        // réinitialise is more
        isMore = false;
        isValid = validateNumber(page);
        if(isValid){
            page = page + 1;
            if ($val_idCat == 0){
                urlAjaxService = "/ajax/userProject/getMore/"+ page;
            }else{
                urlAjaxService = "/ajax/userProject/getByIdCat/"+$val_idCat+"/"+page;
            }
            sendData();
        }else{
            console.log("le numero de pages n'est pas un nombre"+page);
        }
    });
});

// Fonction permettant de valider l'id
function validateNumber(number){
    // initialisation de isValid à true
    var isValid = true;
    // validation si l'id est bien un nombre
    var numberRegex = '^[0-9]+'; 
    var pattern     = new RegExp(numberRegex);
    validateNumb  = pattern.exec(number);
    // si il n'est pas valid ajouter un message d'erreur
    console.log( validateNumb );
    if (validateNumb==null){
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
        if (data != "error" && data != "null") {
            displayResult(data, $val_idCat);
        // si il n'y a une erreur afficher une alerte
        } else {
            console.log("Une erreur est survenue. Veuillez réessayer.");
        };
    });
};

// Filtre les projets en page d'accueil selon
// le ID de categorie recue
function filterImgProjets(val_idCat) {

    var catFilter = "." + val_idCat;

    $('.projets').isotope({    
        filter: catFilter
    });
}

// Ajoute des projets à la page d'accueil
function addImgProjets() {
    
}

// permet d'afficher les résultats
function displayResult(data){

    /*
    
    if (isMore == false){
        //$('.projets').children().remove();
        //$('.img-projet').remove();
        //console.log("isMore == false > remove projects")
    }
    
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

    
    // Effet hover sur les projets en accueil
    $('.img-du-projet').contenthover({
        overlay_background: '#000',
        overlay_opacity: 0.9,
        data_selector: '.over-projet',
    }); 

    //var margeGrille = $('.projets').width() * 0.04;

    //$('.projets').isotope('destroy');

    var $container = $('.projets').imagesLoaded( function() {
        console.log("les images sont loadés : imagesLoaded appelé ! dans home.js");
        
        $container.isotope({
            masonry: {
                gutter: margeGrille,
                margin: margeGrille,
            },
            itemSelector: '.img-projet',
        });

        $('.projets').css('height', 'auto');
    });

    */

    /* $('.projets').isotope({
        masonry: {
        gutter: margeGrille,
        margin: margeGrille,
        },
        itemSelector: '.img-projet',
    }); */

    //initUserProjets();

}




