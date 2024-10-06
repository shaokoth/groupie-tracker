// Function to handle the click event and return the id of the clicked artist
function getClickedArtist(event) {
    
    const clickedElement = event.target; // The element that was clicked
    const artistId = clickedElement.id;  // Retrieve the id of the clicked element
    alert(`Clicked artist id = ${artistId}`);               // Output the id (you can return or do anything with it)
    return artistId;
}

// Attach event listeners to all elements with class 'artist'
document.querySelectorAll('.artist').forEach(function(artist) {
    artist.addEventListener('click', getClickedArtistId);
});
