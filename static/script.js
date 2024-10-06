function getClickedArtist() {

    document.addEventListener('DOMContentLoaded', function () {
        const artistElements = document.querySelectorAll('.artist');
      
        // Loop through each element and add a click event listener
        artistElements.forEach(function(artist) {
            artist.addEventListener('click', function() {
                const artistId = artist.getAttribute('id');
    
                     //alert(`Artist ID: ${artistId}`);

             //  Send the artistId to the server using fetch
                if (artistId) {
                    fetch(`/details`, {
                        method: 'POST',  
                        headers: {
                            'Content-Type': 'application/json',
                        },
                        body: JSON.stringify({ artistId: artistId })
                    })
                    .then(response => response.json()) // assuming JSON response
                    .then(data => {
                        console.log('Response from server:', data);
                    })
                    .catch((error) => {
                        console.error('Error:', error);
                    });
                } else {
                    alert('Artist ID not found');
                }
            });
        });
    });
}

getClickedArtist()

