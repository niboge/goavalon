{{ define "head.tpl" }}

<meta http-equiv="cache-control" content="max-age=3600">
<meta name="viewport" content="width=device-width, initial-scale=1.0" xmlns="http://www.w3.org/1999/html">
<link rel="stylesheet" href="http://cdn.static.runoob.com/libs/bootstrap/3.3.7/css/bootstrap.min.css">
<script type="text/javascript"  src="http://cdn.static.runoob.com/libs/jquery/2.1.1/jquery.min.js"></script>
<script type="text/javascript"  src="http://cdn.static.runoob.com/libs/bootstrap/3.3.7/js/bootstrap.min.js"> </script>

<title>Shiki</title>



<script type="text/javascript">
    var _HOST = 'shikii.cc';
    var _WEB_HOST = 'www.shikii.cc';
    if( window.location.host.indexOf(':') > 0) {
        _HOST = 'haibo.com:8080';
        _WEB_HOST = 'haibo.com';
    }


    //    function(){
//        $('#dropdown-menu li').click(function(){
//            alert('sdfsdf');
//            this.addClass("active").siblings().removeClass("active");
//        });
    //    };
</script>

{{end}}