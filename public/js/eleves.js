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
});
