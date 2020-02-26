// 初始化
$(document).read(function () {

    // variables

    DEFAULT_COOKIE_EXPIRE_TIME = 30;
    uname = "";
    session = "";
    uid = 0;
    currentVideo = null;
    listedVideos = null;

    session = getCookie("session");
    uname = getCookie("username");

    // events

    // 登录注册页事件注册

    $("#regbtn").on('click', function (e) {

    });

    $("signinbtn").on('click', function (e) {

    });

    $("signinhref").on('click', function () {
        $("#regsubmit").hide();
        $("#signinsubmit").show();
    });
    $("registerhref").on('click', function () {
        $("#regsubmit").show();
        $("#signinsubmit").hide();
    });

    // 主页页面事件注册

    $("#uploadform").on('submit', function (e) {

    });

    $(".close").on('clock', function () {

    });

    $("#logout").on('click', function () {
        // 清空cookie
    });
});

// 异步操作，需要callback
function initPage(callback) {

}

function setCookie(cname, cvalue, exmin) {

}

function getCookie(cname) {

}

// DOM操作

function selectVideo(vid) {

}

function refreshComments(vid) {

}

function popupNotificationMsg(msg) {

}

function popupErrorMsg(msg) {

}

function htmlCommentListElement(cid, author, content) {

}

function htmlVideoListElement(vid, name, ctime) {

}

// 异步ajax方法


function registerUser(callback) {

}


function signinUser(callback) {

}

function getUserId(callback) {

}

function createVideo(vname, callback) {

}

function ListAllVideos(callback) {

}

function deleteVideo(vid, callback) {

}

function postComment(vid, content, callback) {

}

function ListAllComments(vid, callback) {

}
