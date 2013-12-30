var pokeball,
    start,
    animation,
    step,
    load,
    loadDone;

window.onload = function() {
  pokeball = document.getElementById('pokeball');
  if(window.location.pathname === "/") {
    $('footer').addClass('home-footer');
  }
};

step = function(timestamp) {
  var delta = timestamp - start,
      deg = (360 * (delta/1000)) % 360;
  pokeball.style.transform = "rotate("+deg+"deg)";
  pokeball.style.mozTransform = "rotate("+deg+"deg)";
  pokeball.style.webkitTransform = "rotate("+deg+"deg)";
  animation = requestAnimationFrame(step);
}

load = function() {
  start = new Date();
  step(start);
}

loadDone = function() {
  cancelAnimationFrame(animation);
}