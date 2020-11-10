$.extend({
    getQueryVariable: function (variable) {
        var query = window.location.search.substring(1);
        var vars = query.split("&");
        for (var i=0;i<vars.length;i++) {
            var pair = vars[i].split("=");
            if(pair[0] == variable){return pair[1];}
        }
        return(false);
    },

    format: function (num) {
        return (num+ '').replace(/(\d{1,3})(?=(\d{3})+(?:$|\.))/g,'$1,');
    },

    numFormat: function (val) {
        var num;
        num = isNaN(Number(val))?0:Number(val);
        return num>0?num:0;
    },

    startWorker: function (faes, number=0, sign="p") {
        if(typeof(Worker) !== "undefined") {
            if(typeof(w[faes['name']][sign][number]) == "undefined") {
                w[faes['name']][sign][number] = new Worker("js/workers.js");
                w[faes['name']][sign][number].postMessage({"speed": faes['list'][number][sign]['speed']})
            }
            w[faes['name']][sign][number].onmessage = function(event) {
                eval('refresh("'+faes['name']+'","'+number+'","'+sign+'","'+event.data+'")');
            };
        } else {
            document.getElementById("result").innerHTML = "Sorry! No Web Worker support.";
        }
    },

    stopWorker: function (name="bath",number=0, sign="p") {
        w[name][sign][number].terminate();
        w[name][sign][number] = undefined;
    },

    listWorker: function () {
        $.startWorker(baths,0, 'p');//店铺
        $.each( recs.list, function(i, n){
            $.startWorker(recs,  i, 'p');//前台
            $.startWorker(recs,  i, 'w');//前台
        });

        $.each( chrs.list, function(i, n){
            $.startWorker(chrs,  i, 'p');//前台
            $.startWorker(chrs,  i, 'w');//前台
            // $.startWorker("chrs",  "chrRefresh", n.duration, i);//更衣间
        });

        $.each( baps.list, function(i, n){
            $.startWorker(baps,  i, 'p');//前台
            $.startWorker(baps,  i, 'w');//前台
            // $.startWorker("baps",  "bapRefresh", n.duration, i);//浴池
        });

        $.each( saus.list, function(i, n){
            $.startWorker(saus,  i, 'p');//前台
            $.startWorker(saus,  i, 'w');//前台
            // $.startWorker("saus",  "sauRefresh", n.duration, i);//桑拿
        });

        $.each( spys.list, function(i, n){
            $.startWorker(spys,  i, 'p');//前台
            $.startWorker(spys,  i, 'w');//前台
            // $.startWorker("spys",  "spyRefresh", n.duration, i);//SPY
        });
    },

    sendData: function (msgBody={}, msgid= 1234) {
        console.log("Connection open ...");

        // 消息体编码
        // 注意：需要对字符串做url编码， 否则中文乱码。该问题仅限于json传输模式
        // cellnet接收时，必须使用wsjson编码处理
        let msgData = JSON.stringify(msgBody)

        let encoder = new TextEncoder('utf8')
        let jsonBody= encoder.encode( msgData)

        // TODO 实现消息ID与消息体绑定

        let pkt = new ArrayBuffer( 2+ jsonBody.length)
        let dv = new DataView(pkt)

        // 写入消息号
        dv.setUint16(0, msgid, true)

        // 这里body使用的是Json编码
        for(let i = 0;i <jsonBody.length;i++){
            dv.setUint8(2+i, jsonBody[i])
        }

        // 发送
        ws.send(pkt);
    },

    build: function (key="") {
        $.sendData({"Userid":userid,"Location":"hincrby", "Key":key+"_num", "Value":"1"});
    },

    buildNew: function (name) {
        var rs = eval(name+"s");
        //var property = eval(name+"_property");
        if (earnings>costs[name]){
            $.build(name);
            $.cost(costs[name], "sub");
            //rs['list'][rs['num']] = Object.assign({}, property);
            rs['list'][rs['num']] = JSON.parse(JSON.stringify(eval(name+"_property")));
            rs['num'] = $.numFormat(rs['num'])+1;
            $.freshList(recs, chrs, baps, saus, spys);
            //$.startWorker(name,  name+"Refresh", property.duration, Object.keys(rs).length-1);//前台
        }else {
            alert("金额不足！")
        }
    },
    //用户初始化
    init: function (msg) {
        var data=JSON.parse(msg.Value)
        switch (msg.Location) {
            case "init":
                $('#wait_num').html(baths.wait_num);
                $('#level').html(data.level)
                earnings = $.numFormat(data.earnings);
                $('#earnings').html($.format(data.earnings));
                $.buildInit(data);
                $.freshList(recs, chrs, baps, saus, spys);
                $.listWorker();
                break;
            default:
        }
    }

});