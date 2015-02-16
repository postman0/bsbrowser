<html>
<head>
	<title>Blade Symphony Server Browser</title>
	<link href='http://fonts.googleapis.com/css?family=Lato:400,700' rel='stylesheet' type='text/css'>
	<link rel="stylesheet" href="/static/style.css">
	<script src="http://underscorejs.org/underscore-min.js"></script>
	<script src="/static/shit.js"></script>
	<body>
		<div class="contents">
			<p class="header">Blade Symphony Server Browser</p>
			{{range .}}
			<div class="server" id="{{.Address}}:{{.Port}}">
				<div class="server-players">
					<span class="server-players-count">{{.PlayerCount}}/{{.Capacity}}</span><br>
					<a class="server-players-button" onclick="showServerPlayers('{{.Address}}:{{.Port}}', true);" href="#">Show players</a>
				</div>
				<div class="server-shit">
					<a class="server-name" href="steam://connect/{{.Address}}:{{.Port}}">{{.Name}}</a><br>
					<span class="server-map">{{.Map}}</span><br>
				</div>
				<div class="clear"></div>
			</div>
			{{end}}
		</div>
	</body>
</html>