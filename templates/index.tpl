<html>
<head>
	<title>Blade Symphony Server Browser</title>
	<link href='http://fonts.googleapis.com/css?family=Lato:400,700' rel='stylesheet' type='text/css'>
	<link rel="stylesheet" href="/static/style.css">
	<body>
		<div class="contents">
			<p class="header">Blade Symphony Server Browser</p>
			{{range .}}
			<div class="server">
			<span class="server-players">{{.PlayerCount}}/{{.Capacity}}</span>
			<a class="server-name" href="steam://connect/{{.Address}}:{{.Port}}">{{.Name}}</a><br>
			<span class="server-map">{{.Map}}</span>
			</div>
			{{end}}
		</div>
	</body>
</html>