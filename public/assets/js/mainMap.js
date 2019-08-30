const pointsUrl = '/points';
const hazardMap = L.map('main-map', {
    center: [51.163375, 10.447683],
    zoom: 7,
});

osmTileLayer.addTo(hazardMap);

// Fetch a JSON representation of geographical points to create markers
fetch(pointsUrl).then((response) => {
    if (!response.ok) {
        throw new Error('Failed to fetch marker data.');
    }
    return response.json();
}).then((jsonData) => {
    for (let point of jsonData) {
        let hazard = L.marker([point.lat, point.lon], {
            title: point.label,
            alt: point.label,
            riseOnHover: true,
        }).addTo(hazardMap);

        let createdDate = new Date(point.created_at);

        hazard.bindPopup(`
      <h3 class="popup-heading">${point.label}</h3>
      <p class="popup-hazard-type">${point.hazard_type.label}</p>
      <p class="popup-description">${point.description}</p>
      <p class="popup-author"><small>Submitted by ${point.user.name} on ${createdDate.toLocaleString()}</small></p>
    `);
    }
}).catch((error) => {
    console.error('Error during fetch operation:', error.message);
});
