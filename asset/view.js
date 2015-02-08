var slope, pslope, Lx, Ly, leanVsPslope, stren = 1
var pslopeVsWS = 0, 
  pslopeVsWN = 0, 
  pslopeVsES = 0, 
  pslopeVsEN = 0;
var Ox = 300
var Oy = 200
var canvas = document.getElementById("example");
ctx = canvas.getContext('2d');

soc.onmessage = function (event) {
  var obj = JSON.parse(event.data)
  console.log(obj)
  slope = obj.Slope
  pslope = obj.Pslope
  Lx = Math.floor(obj.Orient.lr*10)
  Ly = Math.floor(obj.Orient.fb*10)
  leanVsPslope = Math.sign((Lx*(-pslope)) + Ly)

  var x = 100
  var y = 100

  pslopeVsWS = Math.sign( (-x*(-pslope)) - y ) // -y = x*-psl
  pslopeVsES = Math.sign( (x*(-pslope)) - y ) // -y = x*-psl
  pslopeVsWN = Math.sign( (-x*(-pslope)) + y) // -y = x*-psl
  pslopeVsEN = Math.sign( (x*(-pslope)) + y ) // -y = x*-psl
  stren = (Math.abs(obj.Orient.lr)+Math.abs(obj.Orient.fb)) / 50
}

function nearestToP (xb, yb, sl, psl) {
  var b = (yb - sl * xb)
  var xn = (-b / (sl - psl))
  var yn = (psl * xn)
  var dist = Math.abs(xb-xn) + Math.abs(yb-yn)
  return [xn, yn, b, dist * stren]
}

function draw(){

  ctx.clearRect ( 0 , 0 , canvas.width, canvas.height );
  ctx.save()

	ctx.lineWidth = 1;
	ctx.strokeStyle = 'black';
  ctx.fillStyle = 'gray';

  ctx.fillRect(100 + Ox, 100 + Oy, 2, 2)
  ctx.fillRect(-100+ Ox, -100+ Oy, 2, 2)
  ctx.fillRect(100 + Ox, -100+ Oy, 2, 2)
  ctx.fillRect(-100+ Ox, 100 + Oy, 2, 2)

  ctx.beginPath()
  ctx.moveTo(Ox, Oy);
  ctx.lineTo(200+Ox, (-200*pslope + Oy));
  ctx.moveTo(Ox, Oy);
  ctx.lineTo(-200+Ox, 200*pslope + Oy);
  ctx.stroke();
  ctx.closePath()

  ctx.lineWidth = 2;
  ctx.strokeStyle = 'red';
  ctx.beginPath()
  ctx.moveTo(Ox, Oy);
  ctx.lineTo(Lx+Ox, -Ly+Oy);
  ctx.stroke();
  ctx.closePath()

  ctx.lineWidth = 1;
  ctx.strokeStyle = 'gray';
  ctx.save()
  
  np = nearestToP(100, -100, -slope, -pslope)
  ctx.beginPath()
  ctx.moveTo(100+Ox, -100+Oy);
  ctx.lineTo(np[0]+Ox, np[1]+Oy);
  ctx.stroke();
  ctx.closePath()

  if (leanVsPslope ==  pslopeVsEN ) {
    ctx.fillRect(100+Ox, -100+Oy, 10, np[3])
  } else {
    ctx.fillRect(100+Ox, -100+Oy-np[3], 10, np[3])
  }

  np = nearestToP(100, 100, -slope, -pslope)
  ctx.beginPath()
  ctx.moveTo(100+Ox, 100+Oy);
  ctx.lineTo(np[0]+Ox, np[1]+Oy);
  ctx.stroke();
  ctx.closePath()
  
  if (leanVsPslope ==  pslopeVsES ) {
    ctx.fillRect(100+Ox, 100+Oy, 10, np[3])
  } else {
    ctx.fillRect(100+Ox, 100+Oy-np[3], 10, np[3])
  }

  np = nearestToP(-100, 100, -slope, -pslope)
  ctx.beginPath()
  ctx.moveTo(-100+Ox, 100+Oy);
  ctx.lineTo(np[0]+Ox, np[1]+Oy);
  ctx.stroke();
  ctx.closePath()

  if (leanVsPslope ==  pslopeVsWS ) {
    ctx.fillRect(-100+Ox, 100+Oy, 10, np[3])
  } else {
    ctx.fillRect(-100+Ox, 100+Oy-np[3], 10, np[3])
  }

  np = nearestToP(-100, -100, -slope, -pslope)
  ctx.beginPath()
  ctx.moveTo(-100+Ox, -100+Oy);
  ctx.lineTo(np[0]+Ox, np[1]+Oy);
  ctx.stroke();
  ctx.closePath()
  
  if (leanVsPslope ==  pslopeVsWN ) {
    ctx.fillRect(-100+Ox, -100+Oy, 10, np[3])
  } else {
    ctx.fillRect(-100+Ox, -100+Oy-np[3], 10, np[3])
  }

  ctx.restore();
  window.requestAnimationFrame(draw);

}

draw()

