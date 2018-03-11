{{define "room/main.tpl"}}
{{template "nav.tpl" }}
<script type="text/javascript"> var roomName = {{.data.Name}}  </script>
<script type="text/javascript"  src="/public/room/room.js"> </script>
{{end}}