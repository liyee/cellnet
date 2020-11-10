var ws;
var userId=$.getQueryVariable('id');
window.onload = function () {
    console.log("hello")
    ws = new WebSocket("ws://127.0.0.1:18802/echo");
    ws.binaryType = "arraybuffer";

    ws.onopen = function() {
        $.sendData({Userid: userId, Location: "init", Value: "00"});
    };

    ws.onmessage = function(evt) {
        if (evt.data instanceof ArrayBuffer ){
            let dv = new DataView(evt.data);

            // TODO 消息号验证
            let msgid = dv.getUint16(0, true)

            // 包体
            let msgBody = evt.data.slice(2)
            let decoder = new TextDecoder('utf8')
            let jsonBody = decoder.decode(msgBody)

            // 解码包体
            let msg = JSON.parse(jsonBody)
            manageData(msg)
            console.log( "Received Message: " , msg);
        }else{
            console.log("Require array buffer format")
        }
    };

    ws.onclose = function(evt) {
        console.log("Connection closed.");
    };
}