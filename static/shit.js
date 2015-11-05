

function scoreToLeague(score) {
	return ["Oak", "Iron", "Steel", "Diamond", "Master"][score]
}


var playersTpl = _.template('<div class="server-players-list">\
	<table>\
	<tr><td>Name</td><td>League</td><td>Time on Server</td></tr>\
	<% _.map(players, function(player) { %>\
		<tr><td><%- player.Name %></td><td><%- scoreToLeague(player.Score) %></td><td><%- Math.round(player.Duration / 60) %> minutes</td></tr>\
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
				link.onclick = function(evt) {
					showServerPlayers(ip, false);
					evt.preventDefault();
				}
			}
		};
		request.send();

	} else {
		var tablediv = document.getElementById(ip).querySelector(".server-players-list");
		tablediv.parentNode.removeChild(tablediv);
		var link = document.getElementById(ip).querySelector(".server-players-button");
		link.onclick = function(evt) {
			showServerPlayers(ip, true);
			evt.preventDefault();
		}
	}

}

document.addEventListener("DOMContentLoaded", function(evt) {
	var elems = document.querySelectorAll(".server-players-button");
	for(var i = 0; i < elems.length; i++) {
		elems[i].onclick = function(evt) {
			showServerPlayers(evt.target.dataset.address, true);
			evt.preventDefault();
		};
	}
});
