{{ define "error.tpl" }}
{{template "head.tpl"}}
<h8>-- 脑补背景 --</h8>


<h1>
       [ERROR-PAGE]
</h1>


<div>
        DATA:
        <div>
               {{range $k, $v := .data}}
                       {{$k}}-{{$v}}</br>
               {{end}}
		</div>
		<dt> code: {{.code}} </dt>
   		<dt> msg: {{.msg}} </dt>
</div>

<h8>-- 脑补背景 --</h8>

{{ end }}