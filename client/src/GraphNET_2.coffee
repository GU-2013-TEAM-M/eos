GraphNET_2 = Graph.extend {
	defaults: _.extend({}, Graph.prototype.defaults,
		{
			type: "net"
			maxNet: null
		}
	)

	initialize: () ->
		Graph.prototype.initialize.apply(@, arguments)

		if arguments.maxNet
			@set("maxNet", arguments.maxNet)

		data = []
		for i in [0...@get("pointNumber")]
			data[i] = {"label": "", "value": 0}

		@set("data", data)

		@graphOptions = {animation : false}

	appendData: (data) ->
		@get("data").push {"label": (Math.floor((new Date()).getTime()/1000)).toString(), "value": Math.random()*10000}

}