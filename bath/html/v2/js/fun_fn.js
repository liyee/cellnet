jQuery.fn.extend({
    list_fun: function (data) {
        var div = $(this);
        // var num = 0;
        var title_num;
        //var limit = 6;
        //var name_zh = {"rec":"前台","chr":"更衣室","bap":"浴池", "sau":"桑拿", "spy": "SPY"};
        var addDiv = "u_fra_div";
        var addClass = "img";
        if (data.name=="sau" || data.name=="spy") {
            addDiv="c_fra_div";
            addClass = "img2";
            // limit = 2;
        }

        var extents = ["sau","spy"];
        var name = data['name'];
        var prior_num = data['prior']!=null?data['prior']['num']:0;
        var num = 0;
        $.each( data['list'], function(i, n){
            num = $.numFormat(i);
            title_num = num+1;
            // var allotVal = $.allot(n, i, name, 1);
            var html='<div class="'+addDiv+'">\n' +
                '                  <div class="u_title"><span>'+data['name_zh']+' '+title_num+'</span></div>\n' +
                '                  <div class="u_content">\n' +
                '                    <div style="float: left; width: 48px">\n' +
                '                      <img src="images/page_1/regen/u3.svg"/>\n' +
                '                      <span id="'+name+'_p_'+num+'" style="margin-left: 4px">'+n['p']['num']+'</span>\n' +
                '                    </div>\n' +
                '\n' +
                '                    <div style="float: left; width: 66px">\n' +
                '                      <img src="images/page_1/regen/u6.svg"/>\n' +
                '                      00:<span id="'+name+'_pt_'+num+'">00</span>\n' +
                '                    </div>\n' +
                '                  </div>\n' +
                '                  <div class="u_content">\n' +
                '                    <div style="float: left; width: 48px">\n' +
                '                      <img src="images/page_1/regen/u4.svg"/>\n' +
                '                      <span id="'+name+'_w_'+num+'"  style="margin-left: 4px">'+n['w']['num']+'</span>\n' +
                '                    </div>\n' +
                '\n' +
                '                    <div style="float: left; width: 66px">\n' +
                '                      <img src="images/page_1/regen/u6.svg"/>\n' +
                '                      00:<span id="'+name+'_wt_'+num+'">00</span>\n' +
                '                    </div>\n' +
                '                  </div>\n' +
                '                </div>';
            div.append(html);
        });
        if (data['num'] < data['max']){
            var html;
            if (extents.includes(name)){
                html='<div class="'+addDiv+'">\n' +
                    '                  <img class="'+addClass+'" src="images/page_1/regen/u48.svg"  onclick="$.buildNew(\''+name+'\')"/>\n' +
                    '                </div>';
            }else if (prior_num>=data['num'] || name=="rec"){
                html='<div class="'+addDiv+'">\n' +
                    '                  <img class="'+addClass+'" src="images/page_1/regen/u48.svg"  onclick="$.buildNew(\''+name+'\')"/>\n' +
                    '                </div>';

            }else {
                html='<div class="'+addDiv+'">\n' +
                    '                  <img class="'+addClass+'" src="images/page_1/regen/u49.svg"/>\n' +
                    '                </div>';
            }
            div.append(html);
        }
    },
});