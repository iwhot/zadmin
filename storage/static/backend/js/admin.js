;!function (win) {
    "use strict";
    var doc = document
        , admin = function () {
        this.v = '2.2'; //版本号
    };

    admin.prototype.init = function () {

    };


    win.admin = new admin();

}(window);


layui.use(['layer','element','jquery'],function() {
    layer = layui.layer;
    element = layui.element;
    $ = layui.jquery;

    // 打开页面初始
    admin.init();
});