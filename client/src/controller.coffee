isUserLoggedIn = false
cookieExpiryDays = 3
daemons = []

# Callback for successfull login check. All the UI stuff should be done here
loginCheckSuccessful = () ->
	appState.set "is_user_logged_in", true
	router.navigate("app", {trigger: true})
	message = MessageProcessor.createMessage "daemons";
	if message
		serverSocket.sendMessage message


# Callback for unsuccessfull login check. All the UI stuff should be done here
loginCheckUnsuccessful = () ->
	router.navigate("login", {trigger: true})
	appState.set "is_user_logged_in", false

# Callback for successfull login. All the UI stuff should be done here
loginSuccessful = (session_id) ->
	appState.set "is_user_logged_in", true
	Service.setCookie "session_id", session_id, cookieExpiryDays

	message = MessageProcessor.createMessage "daemons";
	if message
		serverSocket.sendMessage message

# Callback for unsuccessfull login. All the UI stuff should be done here
loginUnsuccessful = () ->
	appState.set "is_user_logged_in", false
	Service.setCookie "session_id", null, cookieExpiryDays

# Callback for successfull logout. All the UI stuff should be done here
logoutSuccessful = () ->
	appState.set "is_user_logged_in", false
	Service.setCookie "session_id", null, cookieExpiryDays		

logoutError = () ->
# TODO:

getDaemon = (daemon_id) ->
	for daemon in daemons
		if daemon.daemon_id == daemon_id
			return daemon

updateDaemons = (data) ->
	daemon_id = data.daemon_id
	daemon = getDaemon(daemon_id)
	if daemon
		daemon.setDaemonProperties data
	else
		daemons.push new Daemon data

controlStatus = (data) ->
# TODO:

monitoringData = (data) ->
	daemon_id = data.daemon_id

	data = data.data
	for own key, value of data
		graph = getGraph(daemon_id, key)
		if (graph)
			graph.update(value)
		else
			console.error "There is no graph associated with daemon: " + daemon_id + " for attribute: " + key

historyData = (data) ->
	daemon_id = data.daemon_id
	parameter = data.parameter
	values = data.values

	daemon = daemons.find (model) =>
		model.get("daemon_id") == daemon_id

	if daemon
		chart = null

		if appState.attributes.current_tab.el.id.length > 0
			chart = daemon.get("mon_charts").find (model) =>
				model.get("type") == parameter

			data = []
			if chart
				for own key, value of values
					data.push {"label": key, "value": value}

				chart.setData data

				chart.createGraph()

			if !daemon.get("socket")
				daemon.createSocket()
		else 
			chart = daemon.get("his_charts").find (model) =>
				model.get("type") == parameter

			data = []
			if chart
				for own key, value of values
					data.push {"label": key, "value": value}

				chart.setData data

				$("#history_graph").empty()
				$("#history_graph").append(chart.get("canvas"))

				chart.createGraph()

notImplemented = (data) ->
# TODO:

error = (data) ->
# TODO:

initUI = () ->
