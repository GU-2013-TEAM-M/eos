GraphCPU = Graph.extend {
	defaults: {
		type: "cpu"
		cpuCount: null
	}

	initialize: () ->
		Graph.prototype.initialize.apply(@, arguments)

		cpuCount = @get("options").cpuCount
		@set("cpuCount", cpuCount)

		@graphOptions = {animation : false, scaleOverride : true, scaleSteps : 10, scaleStepWidth : 10, scaleStartValue : 0}

		labels = []
		for i in [0...cpuCount]
			labels.push "CPU " + (i+1)

		data = 
			labels: labels
			datasets: [
				{
					fillColor : "rgba(153,153,153,0.5)",
					strokeColor : "rgba(220,220,220,1)",
					data : []
				}

			]
		@set("data", data)

	createGraph: () ->
		new Chart(($("#graph_" + @cid, @get("context")).get 0).getContext "2d").Bar(@get("data"), @graphOptions)

	# data - an array of numbers					
	setData: (data) ->
		graphData = @get("data")
		graphData.datasets[0].data = data

	update: (data) ->
		@setData([data])
		@createGraph()

	reset: () ->
		data = []
		for i in [0...@get("cpuCount")]
			data.push 0
		@update(data)

}