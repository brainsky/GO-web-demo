<!DOCTYPE html>
<html>
<head>
<link  rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css">
<script src="/static/js/jquery-3.0.0.min.js"></script>
<script src="/static/js/bootstrap.min.js"></script>

</head>
<body>
	<div class="main_frame">
		<nav class="navbar navbar-default navbar-static-top" role="navigation">
		<div class="navbar-header">
				<button type="button" class="navbar-toggle" data-toggle="collapse"
					data-target="#mycollapse">
					<span class="sr-only">切换导航</span> <span class="icon-bar"></span> <span
						class="icon-bar"></span> <span class="icon-bar"></span> <span
						class="icon-bar"></span>
				</button>
				<a href="" class="navbar-brand">LKD WEB应用</a>
		</div>
		<div class="collapse navbar-collapse" id="mycollapse">
			<ul class="nav navbar-nav navbar-left" id="nav_collapse">
					<li><a href="#"><span class="glyphicon glyphicon-home"></span>
							首页</a></li>
					<li><a href="/hello/login"><span class="glyphicon glyphicon-user"></span>
							用户登录</a></li>
			</ul>
		</div>
		<div class="jumbotron">
			<div class="container">
				 <h1>欢迎{{.name}}登陆</h1>
				<p>{{.Website}}</p>
				<p>{{.Email}}</p>
			</div>
		</div>
	</div>
</body>
</html>