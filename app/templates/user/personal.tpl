{{ define "personal.tpl" }}

{{template "nav.tpl"}}

<h3>用户信息</h3>
<h5>账户名: {{.data.Account}}</h5>
<h5>昵称: {{.data.NickName}}</h5>
<h5>积分: {{.data.Score}}</h5>
<h5>胜场: {{.data.Win}}</h5>
<h5>负场: {{.data.Lose}}</h5>
<h5>胜率: {{.data.WinRate}}</h5>

<h4>底部</h4>

{{ end }}