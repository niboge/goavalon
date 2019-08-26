{{ define "index.tpl" }}

{{template "nav.tpl"}}

<html>
<h3>
       [主页]
</h3>
<h4>
       【桌游法官】
</h4>


<div>
       <div>
               {{range $k, $v := .data.game_slogn}}
                       <li>{{$v}}</li>
               {{end}}
       </div>
</div>

</html>


<div class="view">
    <div class="carousel slide" id="carousel-52109">
        <ol class="carousel-indicators">
            <li data-slide-to="0" data-target="#carousel-52109" class="active"></li>
            <li data-slide-to="1" data-target="#carousel-52109" ></li>
            <li data-slide-to="2" data-target="#carousel-52109"></li>
        </ol>
        <div class="carousel-inner">
          <div class="item active">
              <img alt="" src="../../public/bg1.jpg">
              <div class="carousel-caption">
                  <h4>一</h4>
                  <p>
                    《the garden of sinners》
                  </p>
              </div>
          </div>
          <div class="item">
              <img alt="" src="../../public/bg2.jpg">
              <div class="carousel-caption">
                  <h4>二</h4>
                  <p>
                      《晚自习、路灯、大雪》
                  </p>
              </div>
          </div>
          <div class="item">
              <img alt="" src="../../public/bg3.jpg">
              <div class="carousel-caption">
                  <h4>三</h4>
                  <p>
                      《压抑和未知》
                  </p>
              </div>
          </div>
        </div>

        <a class="left carousel-control" href="#carousel-52109" data-slide="prev">
            <span class="glyphicon glyphicon-chevron-left"></span>
        </a>
        <a class="left carousel-control" href="#carousel-52109" data-slide="prev">
            <span class="glyphicon glyphicon-chevron-left"></span>
        </a>
        <a class="right carousel-control" href="#carousel-52109" data-slide="next">
            <span class="glyphicon glyphicon-chevron-right"></span>
        </a>
    </div>
</div>


{{ end }}