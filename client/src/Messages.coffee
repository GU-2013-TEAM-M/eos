class Messages
	# Templates of all the messages. Used for validation and as a hint.
	messages =
		out:
			loginCheck:
				data: ["session_id"],
			login:
				data: ["username", "password"]
			logout:
				data: []
			daemons:
				data: []
			control:
				data: ["daemon_id", "operation"]
			history:
				data: ["daemon_id", "start", "end", "param"]
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
				data: ["daemon_id", "daemon_name", "daemon_state", "daemon_address", "daemon_port", "daemon_platform", "daemon_all_parameters", "daemon_monitored_parameters"],
				processCallback: (data) ->
					processDaemons data
			control:
				data: ["daemon_id", "status", "operation"],
				processCallback: (data) ->
					processControl data
			monitoring:
				data: ["daemon_id", "data"],
				processCallback: (data) ->
					processMonitoring data
			history:
				data: ["daemon_id", "start", "end", "param", "point_distance", "points"]
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
		# status = data.status.toLowerCase()

		# test
		status = "ok"

		switch status
			when "ok"
				console.log "You are logged in"
				
				# message = MessageProcessor.createMessage "loginCheck", session_id: appState.get("session_id")
				# 	if message
				# 		serverSocket.sendMessage message
				loginCheckSuccessful()

				router.navigate("app", {trigger: true})
			when "unauthorized"
				console.log "You are not logged in"
				# loginCheckUnsuccessful()
				router.navigate("login", {trigger: true})


	# Processes the login
	# Params:	data - response data, that contain session id field. If session id is empty/0/null/etc the used is not logged in
	processLogin = (data) ->
		session_id = data.session_id
		if session_id
			console.log "You have successfully logged in (" + session_id + ")"
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
				# logoutSuccessful()
			when "not_ok"
				console.log "Sorry, an error has occured"
				# logoutError()

	# Processes daemons response
	# Params:	data - response data, that contain daemon id, name and state for every available daemon
	processDaemons = (data) ->
		for daemon in data
			daemon_id = daemon.daemon_id
			daemon_name = daemon.daemon_name
			daemon_state = daemon.daemon_state
			daemon_address = daemon.daemon_address
			daemon_port = daemon.daemon_port
			daemon_platform = daemon.daemon_platform
			daemon_all_parameters = daemon.daemon_all_parameters
			daemon_monitored_parameters = daemon.daemon_monitored_parameters		

			str = ""
			for own key, value of daemon_platform
				str += key + ": " + value + "; "

			str += " ALL PARAMS: "
			for param in daemon_all_parameters
				str += param + ", "

			str += " MON PARAMS: "
			for param in daemon_monitored_parameters
				str += param + ", "

			console.log "ID " + daemon_id + "; Name " + daemon_name + "; State " + daemon_state + "; address " + daemon_address + "; port " + daemon_port + ". " + str

		
			# UPDATING
			existing = daemons.find (model) =>
				model.get("daemon_id") == daemon_id
			if existing
				existing.setDaemonProperties {"daemon_name": daemon_name, "daemon_state": daemon_state, "daemon_address": daemon_address, "daemon_port": daemon_port, "daemon_platform": daemon_platform, "daemon_all_parameters": daemon_all_parameters, "daemon_monitored_parameters": daemon_monitored_parameters}
			else 
				daemons.add(new Daemon({"daemon_id": daemon_id, "daemon_name": daemon_name, "daemon_state": daemon_state, "daemon_address": daemon_address, "daemon_port": daemon_port, "daemon_platform": daemon_platform, "daemon_all_parameters": daemon_all_parameters, "daemon_monitored_parameters": daemon_monitored_parameters}))
			if !appState.get("current_daemon")
				appState.set("current_daemon", daemons.models[0])
		
		# updateDaemons(data)
		


	# Processes control response
	# Params:	data - response data, that contain daemon in and status
	processControl = (data) ->
		# What was controlled??????????
		daemon_id = data.daemon_id
		status = data.status
		console.log "Control for daemon " + daemon_id + " was " + status
		# controlStatus(data)


	# Processes daemon monitoring data
	# Params:	data
	processMonitoring = (data) ->
		daemon_id = data.daemon_id
		mon = data.data
		str = ""
		for own key, value of mon
			str += key + ": " + value + "; " 
		console.log "Monitoring for " + daemon_id + ". " + str
		

		daemon = daemons.find (model) ->
			model.get('daemon_id') == daemon_id

		if daemon
			daemon.processMonitoring data.data
		else 
			console.log "Can't find a daemon with id " + daemon_id
		# monitoringData data


	processHistory = (data) ->
		daemon_id = data.daemon_id
		start = data.start
		end = data.end
		param = data.param
		console.log "history for " + daemon_id + " param " + param + " for period from " + start + " to " + end
		# processHistory data


	# Processes not implemented message
	# Params:	data - the inital data sent
	processNotImplemented = (data) ->
		console.log "Not implemented: " + JSON.stringify(data)
		# notImplemented(data)


	# Processes an error message
	# Params:	data - log?
	processError = (data) ->
		console.log "Not implemented: " + JSON.stringify(data)
		# error(data)