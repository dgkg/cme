$(window).load(function() {

	// Effet hover sur les projets en accueil
	$('.img-du-projet').contenthover({
	    overlay_background: '#000',
	    overlay_opacity: 0.9,
	    data_selector: '.over-projet',
	});

	var margeGrille = $('.etu-projet-liste').width() * 0.04;

	$('.etu-projet-liste').isotope({
		masonry: {
			gutter: margeGrille,
			margin: margeGrille,
		},
		itemSelector: '.img-projet',
	})

	$('.ajout-projet').click(function(e){
		e.preventDefault();
		var nouvProjet = $('#nouv-projet')

		nouvProjet.css('display', 'block');
		$('.etu-projet-liste').prepend( nouvProjet );
		$('.etu-projet-liste').isotope( 'prepended', nouvProjet );
		nouvProjet.css('display', 'block'); // Je sais que cette ligne est dédoublée (voir 3 lignes plus haut) mais elle est nécessaire au bon fonctionnement de l'effet.
	})

});
