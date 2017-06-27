<!DOCTYPE html>
<html>
<head>
<link  rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css">
<script src="/static/js/jquery-3.0.0.min.js"></script>
<script src="/static/js/bootstrap.min.js"></script>
 <style type="text/css">
	 .row-margin-top { margin-top: 10px; } 
</style>
</head>
<body>
	<div class="container">
		 <form role="form" action="/hello/login" method="post">
	        <h2>请登录</h2>
			<div class="row">
			<div class="row-margin-top col-lg-4">
	        <label for="inputEmail" class="sr-only">请输入邮箱</label>
	        <input type="email" name="inputEmail" class="form-control" placeholder="请邮箱地址" required autofocus>
			</div>
			<P class="text-danger">{{.flash.error}}</P>
			</div>
			<div class="row">
			<div class="row-margin-top col-lg-4">
	        <label for="inputPassword" class="sr-only">密码</label>
	        <input type="password" name="inputPassword" class="form-control" placeholder="请输入密码" required autofocus>
		
	        <button class="btn btn-lg btn-primary btn-block" type="submit" id="login_btn">登录</button>
			<button class="btn btn-lg btn-primary btn-block" type="button" id="register_btn">注册</button>
			</div>
			</div>
      </form>
	</div>
	<script type="text/javascript">
	$("#register_btn").click(function(){
		location.href = "/hello/register"	
	}	
	);
	</script>
</body>
</html>