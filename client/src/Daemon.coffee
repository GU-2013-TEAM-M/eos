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
	}

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
			# @get("socket").sendMessage message
			serverSocket.sendMessage message


		monitored = @get "daemon_monitored_parameters"
		if monitored.indexOf(parameter) == -1 
			monitored.push parameter

	unmonitor: (parameter) ->
		operation = 
			"stop": [parameter]

		message = MessageProcessor.createMessage "control", daemon_id: @get("daemon_id"), operation: operation
		if message
			# @get("socket").sendMessage message
			serverSocket.sendMessage message

		monitored = @get "daemon_monitored_parameters"
		index = monitored.indexOf(parameter)
		monitored.splice index, 1

	toggleMonitor: (parameter) ->
		monitored = @get "daemon_monitored_parameters"
		if monitored.indexOf(parameter) > -1 
			@unmonitor parameter
		else 
			@monitor parameter

	processMonitoring: (data) ->
		for own key, value of data
			graph = graphs.find (model) =>
				model.get("daemon_id") == @get("daemon_id") && model.get("type") == key
			if graph
				graph.update(parseFloat value)
			else
				console.log "No graph found"			
			


}