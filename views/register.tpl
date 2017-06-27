<!DOCTYPE html>
<html>
<head>
<link  rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css">
<script src="/static/js/jquery-3.0.0.min.js"></script>
<script src="/static/js/bootstrap.min.js"></script>
</head>
<body>
	<div class="container" >
		<h2>请填写注册信息</h2>
	</div>
	<form class="form-horizontal" role="form" action="/hello/register" method="post">
	<div class="form-group">
		<label for="name" class="col-sm-2 control-label">名称</label>
		<div class="col-sm-2">
			<input type="text" class="form-control" name="name" 
				   placeholder="请输入名字">
		</div>
			<P class="text-danger">{{.flash.NameErr}}</P>
	</div>
	<div class="form-group">
		<label for="password" class="col-sm-2 control-label">密码</label>
		<div class="col-sm-2">
			<input type="password" class="form-control" name="password" 
				   placeholder="请输入密码">
		</div>
		<P class="text-danger">{{.flash.PasswordErr}}</P>
	</div>
	<div class="form-group">
		<label for="email" class="col-sm-2 control-label">邮箱</label>
		<div class="col-sm-2">
			<input type="text" class="form-control" name="email" 
				   placeholder="请输入邮箱">
		</div>
		<P class="text-danger">{{.flash.EmailErr}}</P>
	</div>
	<div class="form-group">
		<label for="phone" class="col-sm-2 control-label">手机号</label>
		<div class="col-sm-2">
			<input type="text" class="form-control" name="phone" 
				   placeholder="请输入手机号">
		</div>
		<P class="text-danger">{{.flash.PhoneErr}}</P>
	</div>
	
	<div class="form-group">
		<div class="col-sm-offset-2 col-sm-10">
			<button type="submit" class="btn btn-default">注册</button>
		</div>
	</div>
</form>
</body>
</html>