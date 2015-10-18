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
				<div class="server-info">
					<div class="server-map">
						<div class="server-map-no-image">
							no image :^)
						</div>
						<div style="background-image: url(/static/maps/{{.Map}}.jpg);" class="server-map-image"></div>
						<p class="server-map-name">{{.Map}}</p>
					</div>
					<div class="server-shit">
						<a class="server-name" href="steam://connect/{{.Address}}:{{.Port}}">{{.Name}}</a>
					</div>
					<div class="server-players">
						<p class="server-players-count">{{.PlayerCount}}/{{.Capacity}}</p>
						<a class="server-players-button" data-address='{{.Address}}:{{.Port}}' href="#">Show players</a>
					</div>
				</div>
			</div>
			{{end}}
		</div>
	</body>
</html>