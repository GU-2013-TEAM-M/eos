class Daemon
	constructor: (daemon_id, daemon_name, daemon_state) ->
		@daemon_id = daemon_id
		@daemon_name = daemon_name
		@daemon_state = daemon_state
		@daemon_address = null
		@daemon_port = null
		@daemon_platform = null
		@daemon_all_parameters = null
		@daemon_monitored_parameters = null

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