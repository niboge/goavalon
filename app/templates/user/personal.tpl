{{ define "personal.tpl" }}

{{template "nav.tpl"}}


<h3>用户信息</h3>

<h5>账户名: {{.data.account}}</h5>
<h5>昵称: {{.data.nick}}</h5>
<h5>积分: {{.data.score}}</h5>
<h5>胜场: {{.data.win}}</h5>
<h5>负场: {{.data.lose}}</h5>
<h5>胜率: {{.data.winrate}}</h5>

<h4>底部</h4>

{{ end }}