GraphRAM_2 = Graph.extend {
	defaults: _.extend({}, Graph.prototype.defaults,
		{
			type: "ram"
			totalRam: null
		}
	)

	initialize: () ->
		Graph.prototype.initialize.apply(@, arguments)

		if arguments.totalRam
			@set("totalRam", arguments.totalRam)

		data = []
		for i in [0...@get("pointNumber")]
			data[i] = {"label": "", "value": 0}

		@set("data", data)

		@graphOptions = {animation : false}

	getValue: (value) ->
		value/1000

}