
var playersTpl = _.template('<div class="server-players-list">\
	<table>\
	<tr><td>Name</td><td>Score</td><td>Time on Server</td></tr>\
	<% _.map(players, function(player) { %>\
		<tr><td><%- player.Name %></td><td><%- player.Score %></td><td><%- Math.round(player.Duration / 60) %> minutes</td></tr>\
	<% }) %>\
	</table>\
	</div>');

function showServerPlayers(ip, enable) {
	if (enable) {
		var request = new XMLHttpRequest();
		request.open('GET', '/players?address='+ip, true);

		request.onload = function() {
			if (this.status >= 200 && this.status < 400) {
				var data = JSON.parse(this.response);
				var srvdiv = document.getElementById(ip);
				srvdiv.insertAdjacentHTML("beforeend", playersTpl({players: data}));

				var link = document.getElementById(ip).querySelector(".server-players-button");
				link.onclick = function() {
					showServerPlayers(ip, false);
				}
			}
		};
		request.send();

	} else {
		var tablediv = document.getElementById(ip).querySelector(".server-players-list");
		tablediv.parentNode.removeChild(tablediv);
		var link = document.getElementById(ip).querySelector(".server-players-button");
		link.onclick = function() {
			showServerPlayers(ip, true);
		}
	}

}