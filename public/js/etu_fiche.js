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

});