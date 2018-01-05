{include file="widget/nav.tpl"}

<div class="container">
    <div class="jumbotron">
        <h1>阿瓦隆    狼人杀</h1>
        <p>This is a good game  It includes  ...</p>
        <a class="btn btn-primary btn-large" href="/doc/gamerule">Learn more</a>
    </div>
    <div class="row">
        {foreach from=$data item='item'}
            <div class="panel panel-info">
                <div class="panel-heading">
                    <h3 class="panel-title">房间名: {$item.name} &nbsp;&nbsp; 房间id:{$item.id}</h3>
                </div>
                <div class="panel-body pb">
                    <h3>房间名: {$item.name}</h3>
                    <h6>创建人: {$item.creator_name}</h6>
                    <h6>公告: {$item.notice}</h6>
                    <h6>游戏类型: {$item.game_type}</h6>
                </div>
            </div>
            {foreach from=$item key='key' item='vv'}
                {*{$key}*}
            {/foreach}
        {/foreach}

    </div>
</div>
