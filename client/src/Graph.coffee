Graph = Backbone.Model.extend {
	defaults: {
		daemon_id: null
		canvas: null
		data: null
		pointNumber: 100
		width: 650
		height: 350		
	}


	initialize: (params) ->
		if params.pointNumber
			@set("pointNumber", params.pointNumber)

		if params.width
			@set("width", params.width)

		if params.height
			@set("height", params.height)
		canvas = document.createElement "canvas"
		canvas.width = @get("width")
		canvas.height = @get("height")
		canvas.id = @cid
		
		@set("canvas", canvas)


	getDataSet: () ->
		data = @get("data")
		data = data.slice(data.length-@get("pointNumber"), data.length)
		labels = []
		values = []
		for el, index in data
			label = ""
			if index % 11 == 0
				label = el.label
			labels.push @getLabel(label)
			values.push @getValue(el.value)

		dataSet = 
			labels: labels
			datasets: [
				{
					fillColor : "rgba(66,139,202,0.7)",
					strokeColor : "rgba(220,220,220,1)",
					data : values
				}
			]

		dataSet

	getLabel: (value) ->
		if value.length > 0
			(new Date(value*1000)).toLocaleTimeString()
		else 
			value

	getValue: (value) ->
		value

	setData: (data) ->
		@set("data", data)

	appendData: (data) ->
		@get("data").push data
		

	createGraph: () ->
		id = @get("canvas").id
		if $("#" + id)[0]
			new Chart($("#" + id)[0].getContext "2d").Line(@getDataSet(), @graphOptions)		
}