class Messages
	# Templates of all the messages. Used for validation and as a hint.
	messages =
		out:
			loginCheck:
				data: ["session_id"]
			login:
				data: ["email", "password"]
			logout:
				data: []
			daemons:
				data: []
			daemon:
				data: ["daemon_id"]
			control:
				data: ["daemon_id", "operation"]
			monitoring:
				data: ["daemon_id", "from", "to", "parameter"]
		in:
			loginCheck: 
				data: ["status"],
				processCallback: (data) ->
					processLoginCheck data
			login:
				data: ["session_id"],
				processCallback: (data) ->
					processLogin data
			logout:
				data: ["status"],
				processCallback: (data) ->
					processLogout data
			daemons:
				data: ["daemon_id", "daemon_name", "daemon_state"],
				processCallback: (data) ->
					processDaemons data
			daemon:
				data: ["daemon_id", "daemon_address", "daemon_platform", "daemon_all_parameters", "daemon_monitored_parameters"],
				processCallback: (data) ->
					processDaemon data
			control:
				data: ["daemon_id", "status", "operation"],
				processCallback: (data) ->
					processControl data
			monitoring:
				data: ["daemon_id", "data"],
				processCallback: (data) ->
					processMonitoring data
			history:
				data: ["daemon_id", "parameter", "values"]
				processCallback: (data) ->
					processHistory data
			not_implemented:
				data: [],
				processCallback: (data) ->
					processNotImplemented data
			error:
				data: [],
				processCallback: (data) ->
					processError data
	
	
	@getMessageCallback = (direction, type) ->
		return messages[direction][type].processCallback


	# Processes the login check
	# Params:	data - response data, that contain login check status
	processLoginCheck = (data) ->
		status = data.status.toLowerCase()
		switch status
			when "ok"
				console.log "You are logged in"
				loginCheckSuccessful()

			when "unauthorized"
				console.log "You are not logged in"
				loginCheckUnsuccessful()


	# Processes the login
	# Params:	data - response data, that contain session id field. If session id is empty/0/null/etc the used is not logged in
	processLogin = (data) ->
		session_id = data.session_id
		if session_id
			console.log "You have successfully logged in."
			loginSuccessful session_id
		else
			console.log "Username or password are incorrect. Please log in"
			loginUnsuccessful session_id


	# Processes the logout
	# Params:	data - response data, that contain logout status
	processLogout = (data) ->
		status = data.status.toLowerCase()
		switch status
			when "ok"
				console.log "You are no longer logged in"
				logoutSuccessful()
			when "not_ok"
				console.log "Sorry, an error has occured"

	# Processes daemons response
	# Params:	data - response data, that contain daemon id, name and state for every available daemon
	processDaemons = (data) ->
		for daemon in data.list
			daemon_id = daemon.daemon_id
			daemon_name = daemon.daemon_name
			daemon_state = daemon.daemon_state

			daemons.add(new Daemon({"daemon_id": daemon_id, "daemon_name": daemon_name, "daemon_state": daemon_state}))
			
			message = MessageProcessor.createMessage "daemon", daemon_id: daemon_id
			if message
				serverSocket.sendMessage message
		
	processDaemon = (data) ->
		daemon = data 
		daemon_id = daemon.daemon_id
		# daemon_address = daemon.daemon_address
		daemon_address = "ws://localhost:9005"
		daemon_platform = daemon.daemon_platform
		daemon_all_parameters = daemon.daemon_all_parameters
		daemon_monitored_parameters = daemon.daemon_monitored_parameters		

		existing = daemons.find (model) =>
			model.get("daemon_id") == daemon_id

		if existing
			existing.setDaemonProperties {"daemon_address": daemon_address, "daemon_platform": daemon_platform, "daemon_all_parameters": daemon_all_parameters, "daemon_monitored_parameters": daemon_monitored_parameters}
		
		# existing.createSocket()
		for x in daemon_monitored_parameters
			start = new Date()
			end = new Date()
			start.setDate(start.getDate()-1)
			start = start.getTime()/1000
			end = end.getTime()/1000
			message = MessageProcessor.createMessage "monitoring", {daemon_id: daemon_id, from: start, to: end, parameter: x}
			if message
				serverSocket.sendMessage message

			existing.addMonChart x
			existing.addHisChart x



		if !appState.get("current_daemon")
			appState.set("current_daemon", daemons.models[0])

	# Processes control response
	# Params:	data - response data, that contain daemon in and status
	processControl = (data) ->
		# controlStatus(data)


	# Processes daemon monitoring data
	# Params:	data
	processMonitoring = (data) ->
		if data.values
			processHistory data
			return

		daemon_id = data.daemon_id

		daemon = daemons.find (model) ->
			model.get('daemon_id') == daemon_id

		if daemon
			daemon.processMonitoring data.data
		else 
			console.log "Can't find a daemon with id " + daemon_id


	processHistory = (data) ->
		# daemon_id = data.daemon_id
		# start = data.start
		# end = data.end
		# param = data.param
		# console.log "history for " + daemon_id + " param " + param + " for period from " + start + " to " + end
		historyData data


	# Processes not implemented message
	# Params:	data - the inital data sent
	processNotImplemented = (data) ->
		console.log "Not implemented: " + JSON.stringify(data)
		# notImplemented(data)


	# Processes an error message
	# Params:	data - log?
	processError = (data) ->
		console.log "error"
		switch data.handler
			when "loginCheck"
				loginCheckUnsuccessful()
		# 	when "control"
