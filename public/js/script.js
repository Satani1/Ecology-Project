ymaps.ready(init);

let map;

function init() {
    map = new ymaps.Map("map", {
        center: [55.76, 37.64],
        zoom: 10
    });

    var addMarkerButton = document.getElementById('add-marker-btn');
    var markerForm = document.getElementById('marker-form');
    var markerFormData = document.getElementById('marker-form-data');

    addMarkerButton.addEventListener('click', function() {
        markerForm.style.display = 'block';
    });

    markerFormData.addEventListener('submit', function(e) {
        e.preventDefault();
        var name = document.getElementById('marker-name').value;
        var description = document.getElementById('marker-description').value;
        var address = document.getElementById('marker-address').value;
        ymaps.geocode(address).then(function(res) {
            var coords = res.geoObjects.get(0).geometry.getCoordinates();
            var placemark = new ymaps.Placemark(coords, {
                hintContent: name,
                balloonContentHeader: name,
                balloonContentBody: description
            });
            map.geoObjects.add(placemark);
            markerForm.style.display = 'none';
        });
    });
}


// $.ajax({
//     url: "/get_markers",
//     method: "GET",
//     dataType: "json",
//     success: function(data) {
//         // Добавляем на карту каждую метку из списка
//         $.each(data, function(index, marker) {
//             var placemark = new ymaps.Placemark([marker.lat, marker.lng], {
//                 hintContent: marker.name,
//                 balloonContent: marker.description
//             });
//             myMap.geoObjects.add(placemark);
//         });
//     },
//     error: function(jqXHR, textStatus, errorThrown) {
//         console.log("Error:", textStatus, errorThrown);
//     }
// });

