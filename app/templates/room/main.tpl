{{define "room/main.tpl"}}
<button type="button" onclick="window.history.go(-1)" class="btn btn-primary">Back</button>
<button type="button" href="#nav" data-toggle="collapse" class="btn btn-success">Info</button>
<div id="nav" class="collapse"> {{template "nav.tpl" }} </div>

<script type="text/javascript"> var roomName = {{.data.Name}}  </script>
<script type="text/javascript"  src="/public/room/room.js"> </script>
  
<div class="container-fluid" style="background-color: #001090;">
<div class="row-fluid" style="color: #FFF; margin-top: 1%; margin-bottom: 1%;">
  <table width=100%>
  <tr>
  <td width=6%>
      <a href="/login">
  <i class="glyphicon glyphicon-home" style="font-size:23px; color:#FFF;"></i>
  </a>
  </td>
  <td class="text-center" style="color: #FFF; font-size:23px" >
     {{.data.Name}}
  </td>
  </tr>
  </table>
  
</div>
</div>

<!-- 游戏配置 -->
<div class="container-fluid" style="margin-top: 40px;">
	<div class="row-fluid">
      <a href="#form_cfg" data-toggle="collapse" class="btn btn-primary" style="margin-top:-10%;margin-bottom:-15%">修改配置</a><hr />

      <form id="form_cfg" class="collapse" method="post" action="/room/InceptionSpace">
        <li>游戏类型</li>
        <select name="type" class="selectpicker" data-style="btn-info">
        	<option value=1>狼人杀</option>
          <option value=2>阿瓦隆</option>
        </select>
        <li>本局主题</li>
        <input name="notice" class="form-control" required autofocus value={{.data.Notice}}>
        <li style="color:#AA1010">普通狼个数</li>
        <input name="wolf" class="form-control"
        placeholder="3">
        <li style="color:red">白狼王个数</li>
        <input name="wolf_white" class="form-control"
        placeholder="0">
        <li style="color:#0010A0">普村个数</li>
        <input name="wolf_white" class="form-control"
        placeholder="3">
        <li style="color:#0010FF">神职</li>
        <input type="checkbox" name="check1" value="check1" checked="checked"/>预言家
        <input type="checkbox" name="check2" value="check2" checked="checked"/>女巫 
        <input type="checkbox" name="check3" value="check3" checked="checked"/>猎人
        <input type="checkbox" name="check4" value="check4" />白吃
        <button class="btn btn-block btn-primary" type="submit" style="margin-top: 20px;">提交</button>
      </form>
	</div>
</div>

<!-- 进程 状态 -->
<div id="stats_sun" class="badge" style="margin-top:-3%;" >日光:&nbsp白天</div>

<!-- 圆桌&&Incepter -->
<div class="container-fluid" style="margin-top: 40px;">
	<div class="row-fluid">
		<div class="span12">
				{{.data.Name}}
		</div>
	</div>
</div>

<!-- 功能行使 -->


<!-- 聊天框 -->

<!-- 记录器 -->

<!-- 骑士信息 -->


<!-- 脚本.加人 -->

{{end}}