$(window).load(function() {

	// Création d'un tableau de 40 chiffres (0 à 40)
	var arrProjets = new Array();

	for (var i = 40; i >= 1; i--) {
		arrProjets.push(i);
	};

	// Mélanger les chiffres au hasard
	Shuffle(arrProjets);

	// Initialisation de base
	var strProjetsAAjouter = '';
	var viewportWidth = $(window).width();
	var viewportHeight = $(window).height();

	$(".mosaique").height(viewportHeight);

	// Ajuster le nombre de projets à afficher selon la taille de l'écran
	if(viewportWidth > 1000) {
		for (var i = 22; i >= 1; i--) {
			var strTableau = arrProjets.length - 1;
			strProjetsAAjouter  += '<a href="#"><img src="img/projets/mosaique/prj_' + arrProjets[strTableau] + '.jpg" alt="Image de projet"></a>';
			arrProjets.pop();
		}
	} else {
		if (viewportWidth > 768 && viewportWidth < 1000 ) {
			for (var i = 16; i >= 1; i--) {
				var strTableau = arrProjets.length - 1;
				strProjetsAAjouter  += '<a href="#"><img src="img/projets/mosaique/prj_' + arrProjets[strTableau] + '.jpg" alt="Image de projet"></a>';
				arrProjets.pop();
			}
		} else {
			for (var i = 7; i >= 1; i--) {
				var strTableau = arrProjets.length - 1;
				strProjetsAAjouter  += '<a href="#"><img src="img/projets/mosaique/prj_' + arrProjets[strTableau] + '.jpg" alt="Image de projet"></a>';
				arrProjets.pop();
			}
		}
	}

	// Ajout des projets sélectionnés au HTML de l'accueil
	$('.mosaique').html(strProjetsAAjouter);

	// Initialisation de la mosaique
	initierMosaique();

	// Scroll vers le contenu
	$(".btn").click(function() {
	    $('html, body').animate({
	        scrollTop: $(".zone-membre").offset().top
	    }, 750);
	});

	// Effet hover sur les projets en accueil
	$('.img-du-projet').contenthover({
	    overlay_background: '#000',
	    overlay_opacity: 0.9,
	    data_selector: '.over-projet',
	});

	var margeGrille = $('.projets').width() * 0.04;

	$('.projets').isotope({
		masonry: {
			gutter: margeGrille,
			margin: margeGrille,
		},
		itemSelector: '.img-projet',
	})

});

function Shuffle(o) {
	for(var j, x, i = o.length; i; j = parseInt(Math.random() * i), x = o[--i], o[i] = o[j], o[j] = x);
	return o;
};

function initierMosaique() {

	var viewportWidth = $(window).width();
	var viewportHeight = $(window).height();

	if(viewportWidth > 1000) {
		var hauteurImgMosaique = viewportHeight/3;
	} else {
		var hauteurImgMosaique = viewportHeight/4;
	}

	$(".mosaique").justifiedGallery({

		rowHeight: hauteurImgMosaique,
		maxRowHeight: hauteurImgMosaique,
		fixedHeight: viewportHeight,
		captions: false,
		margins: 0,
		lastRow: 'justify',
		sizeRangeSuffixes: {
			'lt100':'',
			'lt240':'',
			'lt320':'',
			'lt500':'',
			'lt640':'',
			'lt1024':''},
	});

	$(".mosaique").justifiedGallery().on('jg.complete', function (e) {

		if( $(".mosaique").css('height') != viewportHeight ) {
			$(".mosaique").height(viewportHeight);
		}
	});

}
