var i = 0;
var elem = document.getElementById('username');
//var txt = string.concat('Wake up, ', elem, "The Matrix has you... Follow the white rabbit. Knock, knock, ", elem)
var txt = "wake up, Neo"
var speed = 50; /* The speed/duration of the effect in milliseconds */

function typeWriter() {
    if (i < txt.length) {
        elem.innerHTML += txt.charAt(i);
        i++;
        setTimeout(typeWriter, speed);
    }
}