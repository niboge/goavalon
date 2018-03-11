{{ define "room.tpl" }}
{{template "nav.tpl" }}

<div class="container">
    <div class="jumbotron">
        <h2>阿瓦隆    狼人杀</h2>
        <p>This is a good game  It includes  ...</p>
        <a class="btn btn-primary btn-large" href="/doc/rule">Learn more</a>
    </div>
    <div class="row">
        {{range $k, $room := .data}}
            <div class="panel panel-info">
                <div class="panel-heading">
                    <h3 class="panel-title">房间名: {{$room.Name}} &nbsp;&nbsp; 房间id:{{$k}}</h3>
                </div>
                <div class="panel-body pb">
                    <h3>房间名: {{$room.Name}}</h3>
                    <h6>创建人: {{$room.Owner}}</h6>
                    <h6>公告: {{$room.Notice}}</h6>
                    <h6>游戏类型: {{$room.Type}}</h6>
                    <li><a class="btn btn-primary btn-large" href=/room/in{{$room.Name}} >Learn more</a></li>
                </div>
            </div>
        {{end}}

    </div>
</div>
{{ end }}