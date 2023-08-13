/**
 * 弹出式提示框，默认1.2秒自动消失
 * @param message 提示信息
 * @param style 提示样式，有alert-success、alert-danger、alert-warning、alert-info
 * @param time 消失时间
 */
var prompt = function (message, style, time)
{
    style = (style === undefined) ? 'alert-success' : style;
    time = (time === undefined) ? 1200 : time;
    $('<div id="promptModal">')
        .appendTo('body')
        .addClass('alert '+ style)
        .css({"display":"block",
            "z-index":99999,
            "left":($(document.body).outerWidth(true) - 120) / 2,
            "top":10,
            "position": "absolute",
            "padding": "10px",
            "border-radius": "5px"})
        .html(message)
        .show()
        .delay(time)
        .fadeOut(10,function(){
            $('#promptModal').remove();
        });
};

// 成功提示
var success_prompt = function(message, time)
{
    prompt(message, 'alert-success', time);
};

// 失败提示
var fail_prompt = function(message, time)
{
    prompt(message, 'alert-danger', time);
};

// 提醒
var warning_prompt = function(message, time)
{
    prompt(message, 'alert-warning', time);
};

// 信息提示
var info_prompt = function(message, time)
{
    prompt(message, 'alert-info', time);
};

// 信息提示
var alert_prompt = function(message, time)
{
    prompt(message, 'alert-pormpt', time);
};

function checkTelephone(telephone) {
    var reg = /^[1][3,4,5,7,8][0-9]{9}$/;
    if (!reg.test(telephone)) {
        return false;
    } else {
        return true;
    }
}

function setTime() {
    var timeNum = $("#timeNum").val();
    if (timeNum <= 1) {
        $("#timeNum").val(60);
        $("#sendMsg").text("发送验证码");
        $("#sendMsg").attr("class","btn btn-primary");
        $("#sendMsg").attr("disabled",false);
    } else {
        timeNum--
        $("#timeNum").val(timeNum)
        $("#sendMsg").text(timeNum + "s");
        $("#sendMsg").attr("class","btn btn-secondary");
        $("#sendMsg").attr("disabled",true);
        setTimeout("setTime()", 1000)
    }
}