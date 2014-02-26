isUserLoggedIn = false

cookieExpiryDays = 3 #Why 3 days? Maybe it should be specified in response?

daemons = []

# Callback for successfull login check. All the UI stuff should be done here
loginCheckSuccessful = () ->
	appState.set "is_user_logged_in", true

	message = MessageProcessor.createMessage "daemons";
	if message
		serverSocket.sendMessage message


# Callback for unsuccessfull login check. All the UI stuff should be done here
loginCheckUnsuccessful = () ->
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
	# daemons = []

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

processHistory = (data) ->
# TODO:

notImplemented = (data) ->
# TODO:

error = (data) ->
# TODO:

initUI = () ->
