GraphRAM_2 = Graph.extend {
	defaults: {
		type: "ram"
		totalRam: null
		pointNumber: 20
		lastPoints: null
	}

	initialize: () ->
		Graph.prototype.initialize.apply(@, arguments)

		totalRam = @get("options").totalRam
		@set("totalRam", totalRam)

		pointNumber = @get("pointNumber")

		lastPoints = new Array(pointNumber)
		for i in [0...pointNumber]
			lastPoints[i] = 0;		

		@set("lastPoints", lastPoints)

		# @graphOptions = {animation : false, scaleOverride : true, scaleSteps : totalRam/2048, scaleStepWidth : 2048, scaleStartValue : 0}
		@graphOptions = {animation : false}

		labels = []
		for  i in [0...pointNumber-1]
			labels[i] = i+1

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
		@setData(data/1048576)
		@createGraph()

	reset: () ->
		@update()
}