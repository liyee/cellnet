var i = 0;

function timedCount() {
    i -= 1;
    i = i<10?"0"+i:i;
    postMessage(i);
    setTimeout("timedCount()",1000);
}

self.onmessage = function (e) {
    i = e.data.speed +1;
    timedCount();
};