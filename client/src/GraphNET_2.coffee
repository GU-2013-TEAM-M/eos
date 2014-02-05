GraphNET_2 = Graph.extend {
	defaults: {
		type: "net"
		maxNet: null
		pointNumber: 10
		lastPoints: null
	}

	initialize: () ->
		Graph.prototype.initialize.apply(@, arguments)


		maxNet = @get("options").maxNet
		@set("maxNet", maxNet)

		pointNumber = @get("pointNumber")

		lastPoints = new Array(pointNumber)
		for i in [0...pointNumber]
			lastPoints[i] = 0;		

		@set("lastPoints", lastPoints)

		@graphOptions = {animation : false, scaleOverride : true, scaleSteps : maxNet/2048, scaleStepWidth : 2048, scaleStartValue : 0}

		labels = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

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
		new Chart(($("#graph_" + @cid, @get("context")).get 0).getContext "2d").Line(@get("data"), @graphOptions)

	# data - a number
	setData: (data) ->
		pointNumber = @get("pointNumber")
		lastPoints = @get("lastPoints")		
		if (data)
			for i in [1...pointNumber]
				lastPoints[i-1] = lastPoints[i]
			lastPoints[pointNumber-1] = data
			graphData = @get("data")
			graphData.datasets[0].data = lastPoints
		else 
			graphData = @get("data")
			graphData.datasets[0].data = lastPoints			

	update: (data) ->
		@setData(data)
		@createGraph()

	reset: () ->
		@update()
}