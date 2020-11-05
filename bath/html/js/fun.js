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

    cost: function (num, type='add') {
        if (type=='sub'){
            if (earnings>=num) {
                earnings -= num;
            }
        }else {
            earnings += num;
        }
        //return earnings;
        $('#earnings').html($.format(earnings));
    },

    freshList: function (recs, chrs, baps, saus, spys) {
        $('#rec_div').html('');
        $('#chr_div').html('');
        $('#bap_div').html('');
        $('#sau_div').html('');
        $('#spy_div').html('');

        $('#rec_div').list_fun('rec', recs);
        $('#chr_div').list_fun('chr', chrs, Object.keys(recs).length);
        $('#bap_div').list_fun('bap', baps, Object.keys(chrs).length);
        $('#sau_div').list_fun('sau', saus, Object.keys(saus).length);
        $('#spy_div').list_fun('spy', spys, Object.keys(spys).length);
    },

    startWorker: function (name="bath", refresh, duration, number=0) {
        if(typeof(Worker) !== "undefined") {
            if(typeof(w[name][number]) == "undefined") {
                w[name][number] = new Worker("js/workers.js");
                w[name][number].postMessage({"duration": duration})
            }
            w[name][number].onmessage = function(event) {
                eval(refresh+'("'+name+'","'+refresh+'","'+event.data+'",'+number+')');
            };
        } else {
            document.getElementById("result").innerHTML = "Sorry! No Web Worker support.";
        }
    },

    stopWorker: function (name="bath",number=0) {
        w[name][number].terminate();
        w[name][number] = undefined;
    },

    listWorker: function () {
        $.startWorker("bath", "bathRefresh", duration);//店铺
        $.each( recs, function(i, n){
            $.startWorker("rec",  "recRefresh", n.duration, i);//前台
        });

        $.each( chrs, function(i, n){
            $.startWorker("chr",  "chrRefresh", n.duration, i);//更衣间
        });

        $.each( baps, function(i, n){
            $.startWorker("bap",  "bapRefresh", n.duration, i);//浴池
        });

        $.each( saus, function(i, n){
            $.startWorker("sau",  "sauRefresh", n.duration, i);//桑拿
        });

        $.each( spys, function(i, n){
            $.startWorker("spy",  "spyRefresh", n.duration, i);//SPY
        });
    },

    allot: function (data, number, name='rec', type=0) {
        var allotVal = {"num1":0, "num2":0};
        if(data[name+'_p_num']-data[name+'_p_limit']>0){
            allotVal.num1 = data[name+'_p_limit'];
            allotVal.num2 = data[name+'_p_num']-data[name+'_p_limit'];
        }else {
            allotVal.num1 = data[name+'_p_num'];
            allotVal.num2 = 0;
        }

        if (type==1){
            return allotVal;
        }else {
            $('#'+name+'_p_'+number).html($.numFormat(allotVal.num1));
            $('#'+name+'_w_'+number).html($.numFormat(allotVal.num2));
        }
    },

    nextFresh: function (r, number, name, nextData={}, nextName='', end=false) {
        var addNum = 0;
        r[name+'_p_num'] -= 1;

        if (end==false){
            $.each(nextData, function(i, n){
                if (n[nextName+'_p_num']<n[nextName+'_limit']){
                    n[nextName+'_p_num'] += 1;
                    $.allot(n, i, nextName);
                }
            });
        }

        addNum = (r[name+'_limit']-r[name+'_p_num'])>wait_num?wait_num:(r[name+'_limit']-r[name+'_p_num']);
        wait_num -= addNum;
        r.rec_p_num += addNum;
        $('#wait_num').html($.numFormat(wait_num));
        $.allot(r, number, name);
    },

    sendDate: function (msgBody={}, msgid= 1234) {
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


});