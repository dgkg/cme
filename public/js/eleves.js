$(window).load(function() {

	var margeGrille = $('.wrapper-eleves').width() * 0.04;

	$('.wrapper-eleves').isotope({
		masonry: {
			gutter: margeGrille,
			margin: margeGrille,
			//columnWidth: 100,
		},
		itemSelector: '.eleve',
	});

	$('.anneeGraduation').click(function(e) {
		e.preventDefault();
		var annee = "." + $(this).attr("data-annee");
		//alert(annee);
		$('.wrapper-eleves').isotope({ filter: annee });
	});

	/* $('.eleve').click(function(e) {
		var profilSelectionne = $(this).attr("data-userId");
		alert(profilSelectionne);
		window.location.href = "/antoine-lord";
	}); */

	// Effet hover sur les projets en accueil
	/*
	$('.etudiant-photo').contenthover({
	    overlay_background: '#000',
	    overlay_opacity: 0.9,
	    data_selector: '.hover-infos',
	});
	*/

});
