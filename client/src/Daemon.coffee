Daemon = Backbone.Model.extend {
	defaults: {
		daemon_id: ""
		daemon_name: ""
		daemon_state: ""
		daemon_address: ""
		daemon_port: ""
		daemon_platform: ""
		daemon_all_parameters: ""
		daemon_monitored_parameters: ""
	}

	# The data already guaranteed to be correct
	setDaemonProperties: (properties) ->
		for own key, value of properties
			if key != "daemon_id"
				this.set(key, value);

	stop: () ->
		sendControl this.get("daemon_id"), "stop"

	start: () ->
		sendControl this.get("daemon_id"), "start"

	monitor: (parameter) ->
		# What if the parameter is already being monitored?
		operation = 
			"start": [parameter]
		sendControl this.get("daemon_id"), operation

	unmonitor: (parameter) ->
		# What if the parameter is not being monitored?
		operation = 
			"stop": [parameter]
		sendControl this.get("daemon_id"), operation
}