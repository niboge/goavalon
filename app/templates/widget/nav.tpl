{include file="widget/head.tpl"}

<ul class="nav nav-pills nav-justified">
    <li class="dropdown">
        <a class="dropdown-toggle" data-toggle="dropdown">
        tokyo时区测试:{php}echo date('Y-m-d H:i:s');{/php} &nbsp;&nbsp;&nbsp;我 <span class="caret"></span>
        </a>
        <ul class="dropdown-menu">
            <li><a href="/user/home">个人中心</a></li>
            <li><a href="/user/login">登录</a></li>
            <li><a href="/user/logout">登出</a></li>
            <li class="divider"></li>
            <li><a href="/user/addbug">提交Bug</a></li>
        </ul>
    </li>
</ul>

<ul class="nav nav-pills nav-justified">
    <li><a href="/">首页</a></li>
    <li {*class="active">*}><a href="/room/index">桌游</a></li>
    <li><a href="/doc/ss">SS服务</a></li>
    {*todo跳转写的这么艰辛*}
    <li><a href=javascript:window.location.href='http://'+_HOST >空间站</a></li>
</ul><br><br><br>