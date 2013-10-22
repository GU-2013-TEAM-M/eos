isUserLoggedIn = false

cookieExpiryDays = 3 #Why 3 days? Maybe it should be specified in response?

session_id = getCookie("session_id")

daemons = []

# Callback for successfull login check. All the UI stuff should be done here
loginCheckSuccessful = () ->
	isUserLoggedIn = true

# Callback for unsuccessfull login check. All the UI stuff should be done here
loginCheckUnsuccessful = () ->
	isUserLoggedIn = false

# Callback for successfull login. All the UI stuff should be done here
loginSuccessful = (session_id) ->
	isUserLoggedIn = true;
	setCookie "session_id", session_id, cookieExpiryDays

# Callback for unsuccessfull login. All the UI stuff should be done here
loginUnsuccessful = () ->
	isUserLoggedIn = false;
	setCookie "session_id", null, cookieExpiryDays

# Callback for successfull logout. All the UI stuff should be done here
logoutSuccessful = () ->
	isUserLoggedIn = false;
	setCookie "session_id", null, cookieExpiryDays		
	daemons = []

logoutError = () ->
# TODO:

createDaemons = (data) ->
	daemons = []
	for daemon in data
		daemon_id = daemon.daemon_id
		daemon_name = daemon.daemon_name
		daemon_state = daemon.daemon_state
		daemons.push new Daemon daemon_id, daemon_name, daemon_state

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
		console.error "Sorry, you have tried to update a daemon that does not exist"

controlStatus = (data) ->
# TODO:

monitoringData = (data) ->
# TODO:

notImplemented = (data) ->
# TODO:

error = (data) ->
# TODO: