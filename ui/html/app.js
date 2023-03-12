const maps = document.querySelectorAll(".map");

maps.forEach(map => {
    map.addEventListener('click', () => {
        map.classList.toggle('active');
    });
});


const faqs = document.querySelectorAll(".faq");

faqs.forEach(faq => {
    faq.addEventListener('click', () => {
        faq.classList.toggle('active');
    });
});