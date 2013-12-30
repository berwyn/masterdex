var pokeball,
    start,
    animation,
    step,
    load,
    loadDone;

window.onload = function() {
  pokeball = document.getElementById('pokeball');
};

step = function(timestamp) {
  var delta = timestamp - start,
      deg = (360 * (delta/1000)) % 360;
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