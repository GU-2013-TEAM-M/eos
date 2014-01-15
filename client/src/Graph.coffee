Graph = Backbone.Model.extend {
	defaults: {
		daemon_id: null
		data: null
		canvas: null
	}

	initialize: () ->
		canvas = "<canvas id='graph_" + @cid + "' width='400' height='400'></canvas>"
		@set("canvas", canvas)

	show: () ->

	hide: () ->


}