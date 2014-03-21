Graph = Backbone.Model.extend {
	defaults: {
		daemon_id: null
		data: null
		canvas: null
	}

	initialize: () ->
		w = arguments[0].options.width || 200
		h = arguments[0].options.height || 200
		canvas = "<canvas id='graph_" + @cid + "' width='"+w+"' height='"+h+"'></canvas>"
		@set("canvas", canvas)

	show: () ->

	hide: () ->


}