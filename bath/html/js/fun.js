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

    cost: function (earnings, num, type='add') {
        if (type=='sub'){
            if (earnings>=num) {
                earnings -= num;
            }
        }else {
            earnings = earnings+num;
        }
        //return earnings;
        $('#earnings').html($.format(earnings));
    },

    freshList: function (recs, chrs, baps) {
        $('#rec_div').html('');
        $('#chr_div').html('');
        $('#bap_div').html('');

        $('#rec_div').list_fun('rec', recs);
        $('#chr_div').list_fun('chr', chrs, Object.keys(recs).length);
        $('#bap_div').list_fun('bap', baps, Object.keys(chrs).length);
    },

    startWorker: function (name="bath", refresh, duration, number=0) {
        if(typeof(Worker) !== "undefined") {
            if(typeof(w[name][number]) == "undefined") {
                w[name][number] = new Worker("js/workers.js");
                w[name][number].postMessage({"duration":duration})
            }
            w[name][number].onmessage = function(event) {
                eval(refresh+'("'+name+'","'+refresh+'","'+event.data+'",'+number+')');
            };
        } else {
            document.getElementById("result").innerHTML = "Sorry! No Web Worker support.";
        }
    },

    stopWorker: function (name="bath",number=0) {
        w[name][0].terminate();
        w[name][0] = undefined;
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
            $.startWorker("bap",  "chrRefresh", n.duration, i);//浴池
        });
    }

});