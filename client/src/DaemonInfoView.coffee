DaemonInfoView = Backbone.Marionette.ItemView.extend {

	# events: {
	# 	'click li': 'paramClicked',
	# },

	getTemplate: () ->
		if !@model
			"#daemon-info-template-empty"
		else
			"#daemon-info-template"

	onBeforeRender: () ->
		@model = appState.get("current_daemon")
		if !@model
			@model = daemons.models[0]
			appState.set("current_daemon", @model)
	
	onRender: () ->
		if @model
			monitoredParams = @model.get("daemon_monitored_parameters")
			for param in monitoredParams
				graph = graphs.find (model) =>
					model.get("daemon_id") == @model.get("daemon_id") && model.get("type") == param
				if !graph
					switch param
						when "cpu"
							graph = new GraphCPU({daemon_id: @model.get("daemon_id"), options: {cpuCount: 1} })
						when "ram"
							graph = new GraphRAM_2({daemon_id: @model.get("daemon_id"), options: {totalRam: 32768} })
						when "hdd"
							graph = new GraphHDD({daemon_id: @model.get("daemon_id"), options: {totalHdd: 4096} })
						when "net"
							graph = new GraphNET_2({daemon_id: @model.get("daemon_id"), options: {maxNet: 100000} })							
					if graph
						graphs.add graph

				$("#graphs", @el).append(graph.get("canvas"))
				graph.set("context", $("#graphs", @el))
				graph.reset()

			$("li", @el).on('click', (event) =>
				el = event.target
				param = $(el).text().trim()
				@model.toggleMonitor(param)
				@render()
			)

	# paramClicked: (event) ->
	# 	console.log "lololo"
	# 	el = event.target
	# 	param = $(el).text().trim()
	# 	@model.toggleMonitor(param)
	# 	@render()
		
}