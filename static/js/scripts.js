const searchForm = document.getElementById('search-form');
const searchBar = document.getElementById('searchBar');
const resultsDiv = document.getElementById('results');

// Fetch the artist data from the API
let artistsData = [];
let locationsData = [];

async function fetchArtists() {
    try {
        const response = await fetch('/api/proxy?endpoint=artists');
        artistsData = await response.json();
    } catch (error) {
        console.error('Error fetching artist data:', error);
    }
}

async function fetchLocations() {
    try {
        const response = await fetch('/api/proxy?endpoint=locations');
        locationsData = await response.json();
    } catch (error) {
        console.error('Error fetching locations data:', error);
    }
}

// Call the functions to load data initially
fetchArtists();
fetchLocations();

// Function to create a search result item
function createSearchResultItem(category, name, description, id) {
    const resultDiv = document.createElement('div');
    resultDiv.className = 'search-result-item';
    resultDiv.dataset.id = id;
    
    const categorySpan = document.createElement('span');
    categorySpan.className = 'search-result-category';
    categorySpan.textContent = category;
    
    const nameSpan = document.createElement('span');
    nameSpan.textContent = `: ${name}`;
    
    const descriptionDiv = document.createElement('div');
    descriptionDiv.className = 'search-result-description';
    descriptionDiv.textContent = description;
    
    resultDiv.appendChild(categorySpan);
    resultDiv.appendChild(nameSpan);
    resultDiv.appendChild(descriptionDiv);
    
    return resultDiv;
}

// Function to filter and display search results
searchBar.addEventListener('input', function () {
    const query = searchBar.value.toLowerCase();
    resultsDiv.innerHTML = '';  // Clear previous results
    
    if (query === '') {
        resultsDiv.style.display = 'none';
        return;
    }
    
    const matchedResults = [];
    
    // Search through each artist's data
    artistsData.forEach(artist => {
        const { id, name, members, creationDate, firstAlbum } = artist;
        
        // Check if the artist name matches the query
        if (name.toLowerCase().includes(query)) {
            matchedResults.push(createSearchResultItem('Artist', name, 'Band name', id));
        }
        
        // Check if any of the members match the query
        members.forEach(member => {
            if (member.toLowerCase().includes(query)) {
                matchedResults.push(createSearchResultItem('Member', member, `Member of ${name}`, id));
            }
        });
        
        // Check if the creation date matches the query
        if (creationDate.toString().includes(query)) {
            matchedResults.push(createSearchResultItem('Creation Date', creationDate, `Formation year of ${name}`, id));
        }
        
        // Check if the first album matches the query
        if (firstAlbum.toLowerCase().includes(query)) {
            matchedResults.push(createSearchResultItem('First Album', firstAlbum, `Debut album of ${name}`, id));
        }
    });

    // Search through each location's data
    locationsData.index.forEach(locationEntry => {
        const { id, locations } = locationEntry;

        // Check if any of the locations match the query
        locations.forEach(location => {
            if (location.toLowerCase().includes(query)) {
                matchedResults.push(createSearchResultItem('Location', location, 'Concert location', id));
            }
        });
    });
    
    // Display the matched results or show "No results found"
    if (matchedResults.length > 0) {
        matchedResults.forEach(result => {
            resultsDiv.appendChild(result);
        });
        resultsDiv.style.display = 'block';
    } else {
        const noResultsDiv = document.createElement('div');
        noResultsDiv.textContent = 'No results found';
        noResultsDiv.className = 'no-results';
        resultsDiv.appendChild(noResultsDiv);
        resultsDiv.style.display = 'block';
    }
});

// Handle clicking on search results
resultsDiv.addEventListener('click', function(event) {
    const clickedItem = event.target.closest('.search-result-item');
    
    if (clickedItem) {
        // Extract the text content of the search result item
        const query = clickedItem.querySelector('span:nth-child(2)').textContent.trim().substring(2); // Remove ": " from the query string
        window.location.href = `/search?query=${encodeURIComponent(query)}`;
    }
});

// Hide results when clicking outside the search bar
document.addEventListener('click', function(event) {
    if (!searchBar.contains(event.target) && !resultsDiv.contains(event.target)) {
        resultsDiv.style.display = 'none';
    }
});

// Prevent form submission
searchForm.addEventListener('submit', function(event) {
    event.preventDefault();
});

// search icon functionality (for search submissions)
function submitSearch() {
    const form = document.getElementById('search-form');
    const searchInput = document.getElementById('searchBar');

    if (searchInput.value.trim() !== '') {
        form.submit();
    } else {
        alert('Please enter a search term.'); // Optional alert
    }
}

function checkEnter(event) {
    if (event.key === 'Enter') {
        event.preventDefault(); // Prevent default form submission
        submitSearch();
    }
}

// HUmburger button functionality for small screens
document.addEventListener('DOMContentLoaded', () => {
    const hamburger = document.querySelector('.hamburger');
    const navLinks = document.querySelector('.nav-links');
    const overlay = document.querySelector('.overlay');

    function toggleMenu() {
        hamburger.classList.toggle('active');
        navLinks.classList.toggle('active');
        overlay.classList.toggle('active');
        document.body.style.overflow = navLinks.classList.contains('active') ? 'hidden' : '';
    }

    hamburger.addEventListener('click', toggleMenu);
    overlay.addEventListener('click', toggleMenu);

    // Close menu when clicking on a link
    navLinks.querySelectorAll('a').forEach(link => {
        link.addEventListener('click', () => {
            if (navLinks.classList.contains('active')) {
                toggleMenu();
            }
        });
    });

    // Close menu on escape key
    document.addEventListener('keydown', (e) => {
        if (e.key === 'Escape' && navLinks.classList.contains('active')) {
            toggleMenu();
        }
    });

    // Handle resize events
    let resizeTimer;
    window.addEventListener('resize', () => {
        clearTimeout(resizeTimer);
        resizeTimer = setTimeout(() => {
            if (window.innerWidth > 768 && navLinks.classList.contains('active')) {
                toggleMenu();
            }
        }, 250);
    });
});
