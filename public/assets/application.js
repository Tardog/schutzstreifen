const pointsUrl = '/points';
const hazardMap = L.map('main-map').setView([51.163375, 10.447683], 7);

// Fetch a JSON representation of geographical points to create markers
fetch(pointsUrl).then(response => {
  if (!response.ok) {
    throw new Error('Failed to fetch marker data.');
  }
  return response.json();
}).then(jsonData => {
  for (let point of jsonData) {
    let hazard = L.marker([point.lat, point.lon], {
      title: point.label,
      alt: point.label,
      riseOnHover: true,
    }).addTo(hazardMap);

    hazard.bindPopup(point.description);
  }
}).catch(error => {
  console.error('Error during fetch operation:', error.message);
});

L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
  attribution:
    '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
}).addTo(hazardMap);

// TODO: Create an interactive map during hazard creation/editing, allowing the user to pick coordinates by clicking
hazardMap.on('click', event => {
  console.log(event.latlng);
});
