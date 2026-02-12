document.addEventListener("DOMContentLoaded", function () {
    var links = [
        "https://layarotaku.id/",
        "https://www.layarotaku.best/",
        "https://spm.stis.ac.id/",
        "https://datascience.or.id/",
        "https://risetmu.or.id/",
        "https://lsp.stis.ac.id/",
        "https://www.bibisrestaurant.com/our-menus",
        "https://www.northcliffe-seaview.com/seasonal-touring",
        "https://skicivetta.com/skipass-civetta",
        "https://rwdwp.com/dmca/",
        "https://mads.hu/",
        "https://www.hotelklika.cz/de/restaurant.html",
        "https://www.bibisrestaurant.com/menus/",
        "https://pajak.surakarta.go.id/",
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