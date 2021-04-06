;!function (win) {
    "use strict";
    var doc = document
        , admin = function () {
        this.v = '2.2'; //版本号
    };

    admin.prototype.init = function () {

    };

    /**
     * 引入编辑器
     * @param editorDivID
     * @param txtObj
     * @constructor
     */
    admin.prototype.WangEditor = function (editorDivID,txtObj){
        var E = window.wangEditor;
        var editor = new E(editorDivID);
        var $text1 = $(txtObj)
        // 默认情况下，显示所有菜单
        editor.config.menus = [
            'head',
            'bold',
            'fontSize',
            'fontName',
            'italic',
            'underline',
            'strikeThrough',
            'indent',
            'lineHeight',
            'foreColor',
            'backColor',
            'link',
            'list',
            'todo',
            'justify',
            'quote',
            'emoticon',
            'image',
            'video',
            'table',
            'code',
            'splitLine',
            'undo',
            'redo',
        ]

        editor.config.onchange = function (html) {
            // 第二步，监控变化，同步更新到 textarea
            $text1.val(html)
        }
        editor.create()
        // 第一步，初始化 textarea 的值
        $text1.val(editor.txt.html())
    }

    win.admin = new admin();

}(window);


layui.use(['layer','element','jquery'],function() {
    var layer = layui.layer;
    var element = layui.element;
    var $ = layui.jquery;

    // 打开页面初始
    admin.init();
});