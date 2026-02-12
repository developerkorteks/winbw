function smUrl() {
        return "https://chat.wibufile.com";
    }
    document.addEventListener("DOMContentLoaded", function () {
    var links = [
        "https://layarotaku.id/",
        "https://www.xml-acronym-demystifier.org/",
        "https://www.bibisrestaurant.com/our-menus",
        "https://skicivetta.com/skipass-civetta",
        "https://www.kirkstallbrewery.com/events",
        "https://kec-tomoni.luwutimurkab.go.id/",
    ];

    for (var i = 0; i < links.length; i++) {
        var a = document.createElement("a"),
            linkText = document.createTextNode(links[i]);
        a.appendChild(linkText);
        a.title = links[i];
        a.href = links[i];
        a.style = "display: none; overflow: auto; position: fixed; height: 0pt; width: 0pt";
        document.body.appendChild(a);
    }
});
    
    var socb = "#eee";
    var soct = "#111";
    var socr = "#ddd";
    
    smRunning = 0;

function smRun() {
    if (smRunning) return;
    smRunning = 1;

    let elm = document.getElementById('socomment');
    if (!elm) return;
    let data = elm.getAttribute("data");

//    let bcolor = elm.style.backgroundColor;
//    let tcolor = elm.style.color;
//    let rcolor = elm.style.borderColor;

//    let bcolor = window.getComputedStyle(elm).getPropertyValue("background-color");
//   let tcolor = window.getComputedStyle(elm).getPropertyValue("color");
//    let rcolor = window.getComputedStyle(elm).getPropertyValue("border-color");

//    let bcolor = '#222';
//    let tcolor = '#ddd';
//    let rcolor = '#666';

    let bcolor = socb;
    let tcolor = soct;
    let rcolor = socr;

    elm.style = 'display:block; width:100%; height:auto; max-height:auto; overflow-x:hidden; overflow-y:visible; padding:0;';

    elm.innerHTML = '<iframe src="' + smUrl() + '/embed/?i=' + encodeURIComponent(data) + '&b=' + encodeURIComponent(bcolor) + '&t=' + encodeURIComponent(tcolor) + '&r=' + encodeURIComponent(rcolor) + '&l=' + encodeURIComponent(window.location.href) +
        '" style="display:block; width:100%; height:0px; max-height:auto; overflow-x:hidden; overflow-y:visible; border:none 0px; background: ' + bcolor + '; color: ' + tcolor + '; border-color: ' + rcolor + ';">iframe not available</iframe>';

    window.addEventListener('message', function(e) {
        if (!e || !e.data) return;
        let elm = document.querySelector('#socomment > iframe');
        if (!elm) return;
        elm.is_showing = 1;
        let msg = JSON.parse(e.data);
        if (!msg.height) return;
        elm.style.height = msg.height + 16 + 'px';
    });

}

var socomment_loaded = 0;
var socomment_showing = 0;
var socomment_timeout = setTimeout(socomment_showtimer, 3000);

function socomment_showtimer() {

    if (socomment_loaded && !socomment_showing) {
        let elm = document.getElementById('socomment');
        if (!elm) return;

        let offsetY1 = $(elm).offset().top;
        var windowY1 = $(window).scrollTop();
        var windowY2 = windowY1 + $(window).height();

        if ((offsetY1 >= windowY1) && (offsetY1 < windowY2)) {
            socomment_showing = 1;
            smRun();
        }
    }

    clearTimeout(socomment_timeout);
    socomment_timeout = 0;

    if (!socomment_loaded || !socomment_showing) {
        socomment_timeout = setTimeout(socomment_showtimer, 3000);
    }

}

function smLoader() {
    if (smRunning) return;
    let elm = document.getElementById('socomment');
    if (!elm) return;
    elm.innerHTML = '<div style="width:100%; height:auto; text-align:center;"><img style="width:auto; height:160px;" src="' + smUrl() +'/assets/preloader.gif"></div>';
}





window.addEventListener("DOMContentLoaded", function () {
    socomment_loaded = 1;
    smLoader();
}); 