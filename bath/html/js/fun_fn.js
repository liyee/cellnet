jQuery.fn.extend({
    list_fun: function (name, data, count=0) {
        var div = $(this);
        var num = 0;
        var title_num;
        var limit = 6;
        var name_zh = {"rec":"前台","chr":"更衣室","bap":"浴池"};
        $.each( data, function(i, n){
            num = $.numFormat(i);
            title_num = num+1;
            var html='<div class="u_fra_div">\n' +
                '                  <div class="u_title"><span>'+name_zh[name]+' '+title_num+'</span></div>\n' +
                '                  <div class="u_content">\n' +
                '                    <div style="float: left; width: 48px">\n' +
                '                      <img src="images/page_1/regen/u3.svg"/>\n' +
                '                      <span id="'+name+'_p_'+num+'" style="margin-left: 4px">0</span>\n' +
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
                '                      <span id="'+name+'_w_'+num+'"  style="margin-left: 4px">0</span>\n' +
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
        if (title_num < limit){
            var html;
            if (count==title_num && name!="rec"){
                 html='<div class="u_fra_div">\n' +
                     '                  <img class="img" src="images/page_1/regen/u49.svg"/>\n' +
                    '                </div>';
            }else {
                 html='<div class="u_fra_div">\n' +
                    '                  <img class="img" src="images/page_1/regen/u48.svg"  onclick="'+name+'_add()"/>\n' +
                    '                </div>';
            }
            div.append(html);
        }
    },
});