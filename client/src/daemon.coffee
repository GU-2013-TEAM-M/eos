class Daemon
	constructor: (params) ->
		@daemon_id = params.daemon_id
		@daemon_name = params.daemon_name
		@daemon_state = params.daemon_state
		@daemon_address = params.daemon_address
		@daemon_port = params.daemon_port
		@daemon_platform = params.daemon_platform
		@daemon_all_parameters = params.daemon_all_parameters
		@daemon_monitored_parameters = params.daemon_monitored_parameters


	# The data already guaranteed to be correct
	setDaemonProperties: (properties) ->
		for own key, value of properties
			if key != "daemon_id"
				this[key] = value

	stop: () ->
		sendControl @daemon_id, "stop"
		return

	start: () ->
		sendControl @daemon_id, "start"

	monitor: (parameter) ->
		# What if the parameter is already being monitored?
		operation = 
			"start": [parameter]
		sendControl @daemon_id, operation

	unmonitor: (parameter) ->
		# What if the parameter is not being monitored?
		operation = 
			"stop": [parameter]
		sendControl @daemon_id, operation