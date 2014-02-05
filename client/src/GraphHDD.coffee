GraphHDD = Graph.extend {
	defaults: {
		type: "hdd"
		totalHdd: null
	}

	initialize: () ->
		Graph.prototype.initialize.apply(@, arguments)

		totalHdd = @get("options").totalHdd
		@set("totalHdd", totalHdd)

		@graphOptions = {animation : false, scaleOverride : true, scaleSteps : totalHdd/256, scaleStepWidth : 256, scaleStartValue : 0}

		labels = ["HDD"]

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

	# data - a number
	setData: (data) ->
		console.log data
		graphData = @get("data")
		graphData.datasets[0].data = [data]

	update: (data) ->
		@setData(data)
		@createGraph()

	reset: () ->
		@update([0])
}