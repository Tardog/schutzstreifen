const locationMap = L.map('location-select-map', {
    center: [51.163375, 10.447683],
    zoom: 3,
});

mapboxLayer.addTo(locationMap);

locationMap.on('click', (event) => {
    const coordinates = event.latlng;
    document.getElementById('hazard-Lat').value = coordinates.lat;
    document.getElementById('hazard-Lon').value = coordinates.lng;
});
