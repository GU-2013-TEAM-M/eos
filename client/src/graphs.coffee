graphs = []
graph_id = 0
class Graph
	constructor: (daemon_id, type, params, location) ->
		graphs.push(this)
		@daemon_id = daemon_id
		@type = type
		@graph_id = "gpaph_" + graph_id
		graph_id++
		@canvas = "<canvas id='" + @graph_id + "' width='400' height='400'></canvas>"
		if !location
			location = $("#graphs")
		location.append(@canvas);
		@ctx = ($("#" + @graph_id).get 0).getContext "2d"
		@data = []
		switch type
			when "cpu"
				@cpuCount = params.count
				@options = {animation : false, scaleOverride : true, scaleSteps : 10, scaleStepWidth : 10, scaleStartValue : 0}
				@getGraph = (data) ->
					new Chart(@ctx).Bar(data, @options)

				labels = []
				for i in [0...@cpuCount]
					labels.push "CPU " + (i+1)

				@data = 
					labels: labels
					datasets: [
						{
							fillColor : "rgba(220,220,220,0.5)",
							strokeColor : "rgba(220,220,220,1)",
							data : []
						}

					]
				# data - an array of numbers					
				@setData = (data) ->
					@data.datasets[0].data = data
			
			when "ram"					
				@totalRam = params.total
				@options = {animation : false, scaleOverride : true, scaleSteps : @totalRam/256, scaleStepWidth : 256, scaleStartValue : 0}
				@getGraph = (data) ->
					new Chart(@ctx).Bar(data, @options)

				labels = ["RAM"]

				@data = 
					labels: labels
					datasets: [
						{
							fillColor : "rgba(220,220,220,0.5)",
							strokeColor : "rgba(220,220,220,1)",
							data : []
						}

					]
				# data - a number
				@setData = (data) ->
					@data.datasets[0].data = [data]

			when "karma"
				@pointNumber = 10 # change to param
				@lastPoints = new Array(@pointNumber)
				for i in [0...@pointNumber]
					@lastPoints[i] = 0;
				@options = {animation : false, scaleOverride : true, scaleSteps : 10, scaleStepWidth : 10, scaleStartValue : 0}
				@getGraph = (data) ->
					new Chart(@ctx).Line(data, @options)

				labels = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

				@data = 
					labels: labels
					datasets: [
						{
							fillColor : "rgba(220,220,220,0.5)",
							strokeColor : "rgba(220,220,220,1)",
							data : []
						}

					]
				# data - a number
				@setData = (data) ->
					for i in [1...@pointNumber]
						@lastPoints[i-1] = @lastPoints[i]
					@lastPoints[@pointNumber-1] = data
					@data.datasets[0].data = @lastPoints

	create: (data) ->
		@setData(data)
		@getGraph(@data)

	show: () ->
		$("#" + @graph_id).show();

	hide: () ->
		$("#" + @graph_id).hide();

	update: (data) ->
		@create(data);

getGraph = (daemon_id, type) ->
	for graph in graphs
		if graph.daemon_id == daemon_id && graph.type == type
			return graph