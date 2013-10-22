simulate = (timing) ->
	new Graph("simulation", "karma");
	new Graph("simulation", "cpu", {count: 4});
	new Graph("simulation", "ram", {total: 2048});

	setInterval () ->
		karmaMessage = 
			type: "monitoring"
			data: {"daemon_id": "simulation", "data": {"karma": Math.floor((Math.random()*100)+1), "cpu": [Math.floor((Math.random()*100)+1), Math.floor((Math.random()*100)+1), Math.floor((Math.random()*100)+1), Math.floor((Math.random()*100)+1)], "ram": Math.floor((Math.random()*2048)+1)}}
		trySendMessage karmaMessage
	, timing
	