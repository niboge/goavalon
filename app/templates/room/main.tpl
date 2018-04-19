{{define "room/main.tpl"}}
{{template "nav.tpl" }}
<script type="text/javascript"> var roomName = {{.data.Name}}  </script>
<script type="text/javascript"  src="/public/room/room.js"> </script>
  
<div class="container-fluid" style="background-color: #101010;">
<div class="row-fluid" style="color: #FFF; margin-top: 10px; margin-bottom: 10px;">
  <table width=100%>
  <tr>
  <td width=10%>
      <a href="/login">
  <i class="glyphicon glyphicon-home" style="font-size:23px; color:#FFF;"></i>
  </a>
  </td>
  <td class="text-center" style="color: #FFF; font-size:23px" >
     {{.data.Name}}
  </td>
  <td width=10%>
  </td>
  </tr>
  </table>
  
</div>
</div>


<div class="container-fluid" style="margin-top: 40px;">
	<div class="row-fluid">
		<div class="span12">

      <form class="form-signin" method="post" action="/enter_room_post">
        <!-- <h2 class="form-signin-heading">Please sign in</h2> -->
        <!-- <label for="username" class="sr-only">用户名</label> -->

        游戏类型
        <select name="type" class="form-control">
        	<option value=1>kill</option><option value=2>avalon</option>
        </select>
        公告
        <input name="notice" class="form-control" required autofocus value={{.data.Notice}}>
        职阶
        <input name="username" class="form-control"
        placeholder="填写姓名可以方便大家交流哦">
        
        
<!--         <label for="inputPassword" class="sr-only">Password</label>
        <input type="password" id="inputPassword" class="form-control" placeholder="Password" required>
 -->        

        <button class="btn btn-block btn-success" type="submit" style="margin-top: 20px;">进入房间</button>
      </form>



		</div>
	</div>

</div>


<div class="container-fluid" style="margin-top: 40px;">
	<div class="row-fluid">
		<div class="span12">
				{{.data.Notice}}
		</div>
	</div>
</div>



{{end}}