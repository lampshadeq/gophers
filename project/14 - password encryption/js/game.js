var shape;

function init() {
  // 0 = O, 1 = X
  switch (Math.floor(Math.random() * 2)) {
    case 0:
      shape = "images/piece_o.png";
      document.getElementById("you_img").setAttribute("src", shape);
      break;
    case 1:
      shape = "images/piece_x.png";
      document.getElementById("you_img").setAttribute("src", shape);
  }
}

function place(id) {
  var s = document.getElementById(id).style;
  
  s.backgroundImage = "url('" + shape + "')";
  s.backgroundRepeat = "no-repeat";
  s.backgroundPosition = "center center";
}