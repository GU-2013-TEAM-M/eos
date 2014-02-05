DaemonInfoView = Backbone.Marionette.ItemView.extend {
	# template: "#daemon-info-template",

	events: {
		'click li': 'paramClicked',
	},

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
						when "CPU"
							graph = new GraphCPU({daemon_id: @model.get("daemon_id"), options: {cpuCount: 4} })
						when "RAM"
							graph = new GraphRAM({daemon_id: @model.get("daemon_id"), options: {totalRam: 2048} })
					graphs.add graph

				$("#graphs", @el).append(graph.get("canvas"))
				graph.set("context", $("#graphs", @el))
				graph.reset()

	paramClicked: (event) ->
		el = event.target
		param = $(el).text().trim()
		@model.toggleMonitor(param)
		@render()
		
}