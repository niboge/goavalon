{{ define "index.tpl" }}
<html>
<h1>
	[主页]
</h1>
<h2>
	【桌游法官】
</h2>


<div>
	轮播图:
	<div>
		{{range $k, $v := .data.game_slogn}}
			-{{$v}}</br>
		{{end}}
	</div>

	我的信息:
	<tr>{{.data.info}}</tr>
</div>

</html>

<h6>{{.code}}:{{.msg}}</h6>

{{ end }}