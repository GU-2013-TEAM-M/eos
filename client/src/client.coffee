# root = exports ? this

# The main server address
serverAddress = "ws://localhost/ws"
serverws = null
daemonws = null

# Ready
$(document).ready ->
	init

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
		daemon:
			data: ["daemon_id"]
		control:
			data: ["daemon_id", "operation"]
	in:
		loginCheck: 
			data: ["status"],
			process: processLoginCheck
		login:
			data: ["session_id"],
			process: processLogin
		logout:
			data: ["status"],
			process: processLogout
		daemons:
			data: [daemon : ["daemon_id", "daemon_name", "daemon_state"]],
			process: processDaemons
		daemon:
			data: ["daemon_id", "daemon_address", "daemon_port", daemon_platform: ["platform", "bla1", "bla2"], daemon_all_parameters: [], daemon_monitored_parameters: []],
			process: processDaemon
		control:
			data: ["daemon_id", "status"],
			process: processControl
		monitoring:
			data: ["daemon_id", data: []],
			process: processMonitoring
		not_implemented:
			data: [],
			process: processNotImplemented
		error:
			data:	[],
			process: processError

# Processes an incoming message if the data is well-formatted
# Params:	msg - message
processIncomingMessage = (msg) ->
	message = JSON.parse msg
	type = message.type
	if messages.in[type]
		console.log "Received a message of type " + type
		data = message.data
		if checkData type, data, "in"
			messages[type].process data
	else
		console.err "Received a message of unknown type"

# Checks if the data format relates to the type. This works both for incoming and outcoming data. For incoming the function returns true if the check is passed. False if failed. For outcoming - a message object if passed, false if failed.
# Params:	type - type of data
#			data - data
#			direction - "in" or "out". Is the data incoming or outcoming?
# Return:	true/false/object
checkData = (type, data, direction) ->
	dataTemplate = messages[direction][type].data
	
	for own key of dataTemplate
		if !data.key
			console.err "Wrong data format"
			return false
	if direction == "out"
		message = 
			type: type
			data: JSON.stringify(data);
		return message
	else if direction == "in"
		return true;
	else
		console.err "wrong direction"
# Processes the login check
# Params:	data - response data, that contain login check status
processLoginCheck = (data) ->
	status = data.status
	switch status
		when "OK"
			console.log "You are logged in"
		when "UNAUTHORIZED"
			console.log "You are not logged in"

# Processes the login
# Params:	data - response data, that contain session id field. If session id is empty/0/null/etc the used is not logged in
procesLogin = (data) ->
	session_id = data.session_id
	if session_id
		console.log "You have successfully logged in"
	else
		console.log "Username or password are incorrect. Please log in"

# Processes the logout
# Params:	data - response data, that contain logout status
processLogout = (data) ->
	status = data.status  
	switch status
		when "OK"
			console.log "You are no longer logged in"
		when "NOT_OK"
			console.log "Sorry, an error has occured"

# Processes daemons response
# Params:	data - response data, that contain daemon id, name and state for every available daemon
processDaemons = (data) ->
	for daemon in data
		daemon_id = daemon.daemon_id
		daemon_name = daemon.daemon_name
		daemon_state = daemon.daemon_state
		console.log "ID " + daemon_id + "; Name " + daemon_name + "; State " + daemon_state + ";"

# Processes daemon response
# Params:	data - response data, that contain daemon id, address, port, platform, all parameters, monitores parameters for a daemon
processDaemon = (data) ->
	daemon_id = data.daemon_id
	daemon_address = data.daemon_address
	daemon_port = data.daemon_port
	daemon_platform = data.daemon_platform
	daemon_all_parameters = data.daemon_all_parameters
	daemon_monitored_parameters = data.daemon_monitored_parameters
	console.log "data for " + daemon_id + " loaded"

# Processes daemon monitoring data
# Params:	data
processMonitoring = (data) ->
	console.log data

# Processes not implemented message
# Params:	data - the inital data sent
processNotImplemented = (data) ->
	console.log data

# Processes an error message
# Params:	data - log?
processError = (data) ->
	console.log data

# Creates a new message
# Params:	type - type of the message
#			data - message data
# Return: a new message
createMessage = (type, data) ->
	message = checkData type, data, "out"
	return message

# Tries to send a message
# Params:	message
# Return: true/false depending on success
trySendMessage = (message) ->
	if message
		serverws.send(message)
		console.log "Login check message was sent"
		return true;
	else 
		console.err "Login check message was aborted"	
		return false;

# Sends a login check message
sendLoginCheck = () ->
	session_id = getCookie("session_id");
	message = createMessage "loginCheck", (session_id: session_id)
	return trySendMessage message

# Sends a login message
# Params:	username
#			password
# Return: true/false depending on success
sendLogin = (username, password) ->
	message = createMessage "login", (username: username, password: password)
	return trySendMessage message

# Sends a logout message
# Return: true/false depending on success
sendLogout = () ->
	message = createMessage "logout"
	return trySendMessage message

# Sends a daemons message
# Return: true/false depending on success
sendDaemons = () ->
	message = createMessage "daemons"
	return trySendMessage message

# Sends a daemon message
# Params:	daemon_id
# Return: true/false depending on success
sendDaemon = (daemon_id) ->
	message = createMessage "daemon", daemon_id: daemon_id
	return trySendMessage message

# Sends a control message
# Params:	daemon_id
#			operation
# Return: true/false depending on success
sendControl = (daemon_id, operation) ->
	message = createMessage "control", daemon_id: daemon_id, operation: operation
	trySendMessage message

# Sets a cookie
# Params:	name - cookie name
#			value - cookie value
#			days - days until a cookie should expire
setCookie = (name, value, days) ->
	if days
		date = new Date
		date.setTime(date.getTime() + (days*24*60*60*1000))
		expires = "; expires=" + date.toGMTString()
	else
		expires = ""
	document.cookie = name + "=" + value + expires + "; path=/"

# Gets a cookie value
# Params:	key - cookie name
# Return:	cookie value or null if the cookie was not found
getCookie = (key) ->
	key = key + "="
	for c in document.cookie.split(';')
		c.substring(1, c.length) while c.charAt(0) is ' '
		return c.substring(key.length, c.length) if c.indexOf(key) == 0
	return null

# Initialization
init = () ->
	serverws = createServerWebSocket(serverAddress)

# Creates and initializes a web socket connection with the server
# Params:	address - server address
# Return:	web socket object
createServerWebSocket = (address) ->
	serverws = new WebSocket serverAddress

	serverws.onopen ->
		sendLoginCheck

	serverws.onclose ->

	serverws.onerror (event) ->
		console.err event

	serverws.onmessage (msg) ->
		processIncomingMessage msg		

	return serverws

# Creates and initializes a web socket connection with the daemon
# Params:	address - server address
# Return:	web socket object
createDaemonWebSocket = (address) ->
	daemonws = new WebSocket address

	daemonws.onopen ->
		sendLoginCheck # Is it neccessary?

	daemonws.onclose ->

	daemonws.onerror (event) ->
		console.err event

	daemonws.onmessage (msg) ->
		processIncomingMessage msg

	return daemonws