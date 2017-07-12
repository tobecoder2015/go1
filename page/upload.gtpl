<html>
<head>
<title></title>
</head>
<body>
<form action="/upload" method="post" enctype="multipart/form-data">
	文件:<input type="file" name="uploadfile">
	<input type="hidden" name="token" value="{{.}}">
	<input type="submit" value="提交">
</form>
</body>
</html>