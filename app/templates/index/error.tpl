{{ define "error.tpl" }}
<html>
<h1>
       [ERROR] {{.data.code}}
</h1>


<div>
       INFO:
       <div>
               {{range $k, $v := .data.data}}
                       -{{$v}}</br>
               {{end}}
       </div>

       <td>{{.data.msg}}</td>
</div>


{{ end }}