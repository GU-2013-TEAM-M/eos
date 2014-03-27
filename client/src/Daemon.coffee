Daemon = Backbone.Model.extend {
	defaults: {
		daemon_id: null
		daemon_name: null
		daemon_state: null
		daemon_address: null
		daemon_port: null
		daemon_platform: null
		daemon_all_parameters: null
		daemon_monitored_parameters: null
		socket: null
		mon_charts: new Graphs()
		his_charts: new Graphs()
	}

	addMonChart: (type) ->
		switch type
			when "cpu"
				@get("mon_charts").add new GraphCPU_2({daemon_id: @get("daemon_id")})
			when "ram"
				@get("mon_charts").add new GraphRAM_2({daemon_id: @get("daemon_id")})
			when "net"
				@get("mon_charts").add new GraphNET_2({daemon_id: @get("daemon_id")})

	addHisChart: (type) ->
		switch type
			when "cpu"
				@get("his_charts").add new GraphCPU_2({daemon_id: @get("daemon_id")})
			when "ram"
				@get("his_charts").add new GraphRAM_2({daemon_id: @get("daemon_id")})
			when "net"
				@get("his_charts").add new GraphNET_2({daemon_id: @get("daemon_id")})

	getMonChart: (type) ->
		chart = mon_charts.find (model) =>
			model.get("type") == type
		return chart

	getHisChart: (type) ->
		chart = his_charts.find (model) =>
			model.get("type") == type
		return chart			

	createSocket: () ->
		new DaemonSocket(@)

	# The data already guaranteed to be correct
	setDaemonProperties: (properties) ->
		for own key, value of properties
			if key != "daemon_id"
				@set(key, value);

	stop: () ->
		message = MessageProcessor.createMessage "control", daemon_id: @get("daemon_id"), operation: "stop"
		if message
			# @get("socket").sendMessage message
			serverSocket.sendMessage message

	start: () ->
		message = MessageProcessor.createMessage "control", daemon_id: @get("daemon_id"), operation: "start"
		if message
			# @get("socket").sendMessage message
			serverSocket.sendMessage message	

	monitor: (parameter) ->
		operation = 
			"start": [parameter]

		message = MessageProcessor.createMessage "control", daemon_id: @get("daemon_id"), operation: operation
		if message
			serverSocket.sendMessage message


		monitored = @get "daemon_monitored_parameters"
		if monitored.indexOf(parameter) == -1 
			monitored.push parameter

		chart = @get("mon_charts").find (model) =>
			model.get("type") == parameter

		$(chart.get("canvas")).show()			

	unmonitor: (parameter) ->
		operation = 
			"stop": [parameter]

		message = MessageProcessor.createMessage "control", daemon_id: @get("daemon_id"), operation: operation
		if message
			serverSocket.sendMessage message

		monitored = @get "daemon_monitored_parameters"
		index = monitored.indexOf(parameter)
		monitored.splice index, 1

		chart = @get("mon_charts").find (model) =>
			model.get("type") == parameter

		$(chart.get("canvas")).hide()

	toggleMonitor: (parameter) ->
		monitored = @get "daemon_monitored_parameters"
		if monitored.indexOf(parameter) > -1 
			@unmonitor parameter
		else 
			@monitor parameter

	processMonitoring: (data) ->
		for own key, value of data
			graph = @get("mon_charts").find (model) =>
				model.get("type") == key
			if graph
				val = parseFloat value
				graph.appendData({"label": (Math.floor((new Date()).getTime()/1000)).toString(), "value": val})
				graph.createGraph()
			else
				console.log "No graph found"
}