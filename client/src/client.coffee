# root = exports ? this

# The main server address
# My local echo server
serverAddress = "ws://192.168.1.100:8080/ws"
serverws = null
daemonws = null

testing = false;
testOUT = false;
testIN = false;
# Ready
$(document).ready ->
	init()

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
			data: ["daemon_id", "daemon_address", "daemon_port", "daemon_platform", "daemon_all_parameters", "daemon_monitored_parameters"],
			processCallback: (data) ->
				processDaemon data
		control:
			data: ["daemon_id", "status"],
			processCallback: (data) ->
				processControl data
		monitoring:
			data: ["daemon_id", "data"],
			processCallback: (data) ->
				processMonitoring data
		not_implemented:
			data: [],
			processCallback: (data) ->
				processNotImplemented data
		error:
			data: [],
			processCallback: (data) ->
				processError data

# Processes an incoming message if the data is well-formatted
# Params:	msg - message
processIncomingMessage = (msg, from, messageEvent) ->
	console.log "Incoming message: " + msg
	message = JSON.parse msg
	type = message.type
	if messages.in[type]
		# console.log "Received a message of type " + type
		data = message.data
		if checkData type, data, "in"
			messages["in"][type].processCallback data
	else
		console.error "Received a message of unknown type"

# Checks if the data format relates to the type. This works both for incoming and outcoming data. For incoming the function returns true if the check is passed. False if failed. For outcoming - a message object if passed, false if failed.
# Params:	type - type of data
#			data - data
#			direction - "in" or "out". Is the data incoming or outcoming?
# Return:	true/false/object(JSON)
checkData = (type, data, direction) ->
	dataTemplate = messages[direction][type].data
	
	for own key of dataTemplate
		keyname = dataTemplate[key]
		
		if type == "daemons" && direction = "in"
			for daemon in data 
				if !daemon.hasOwnProperty keyname
					console.error "Wrong data format"
					return false
		else
			if !data.hasOwnProperty keyname
				console.error "Wrong data format"
				return false
	if direction == "out"
		message = 
			type: type
			data: data
		return message
	else if direction == "in"
		return true;
	else
		console.error "wrong direction"
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
processLogin = (data) ->
	session_id = data.session_id
	#Why 3 days? Maybe it should be specified it response?
	setCookie "session_id", session_id, 3
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

	str = ""
	for own key, value of daemon_platform
		str += key + ": " + value + "; "

	str += " ALL PARAMS: "
	for param in daemon_all_parameters
		str += param + ", "

	str += " MON PARAMS: "
	for param in daemon_monitored_parameters
		str += param + ", "

	console.log "ID " + daemon_id + "; address " + daemon_address + "; port " + daemon_port + ". " + str

# Processes control response
# Params:	data - response data, that contain daemon in and status
processControl = (data) ->
	# What was controlled??????????
	daemon_id = data.daemon_id
	status = data.status
	console.log "Control for daemon " + daemon_id + " was " + status

# Processes daemon monitoring data
# Params:	data
processMonitoring = (data) ->
	daemon_id = data.daemon_id;
	mon = data.data;
	str = ""
	for own key, value of mon
		str += key + ": " + value + "; " 
	console.log "Monitoring for " + daemon_id + ". " + str

# Processes not implemented message
# Params:	data - the inital data sent
processNotImplemented = (data) ->
	console.log "Not implemented: " + JSON.stringify(data)

# Processes an error message
# Params:	data - log?
processError = (data) ->
	console.log "Not implemented: " + JSON.stringify(data)

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
trySendMessage = (msg) ->
	try
		message = JSON.stringify msg
		serverws.send message
		console.log msg.type + " message was sent"
	catch err
		console.error err
		return false
	finally
		return true    	

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
	serverws.onopen = wsServerOnOpenHandler
	serverws.onclose = wsServerOnCloseHandler
	serverws.onerror = wsServerOnErrorHandler
	serverws.onmessage = wsServerOnMessageHandler
	return serverws

# Creates and initializes a web socket connection with the daemon
# Params:	address - server address
# Return:	web socket object
createDaemonWebSocket = (address) ->
	daemonws = new WebSocket address
	daemonws.onopen = wsDaemonOnOpenHandler
	daemonws.onclose = wsDaemonOnCloseHandler
	daemonws.onerror = wsDaemonOnErrorHandler
	daemonws.onmessage = wsDaemonOnMessageHandler
	return daemonws

wsServerOnOpenHandler = () ->
	if (!testing)
		sendLoginCheck()
	else
		test()
		
wsServerOnCloseHandler = (event) ->
	ws = event.target;
	console.log "Connection lost. IMPLEMENT RETRYING!"

wsServerOnErrorHandler = (error) ->
	console.error error

wsServerOnMessageHandler = (messageEvent) ->
	processIncomingMessage messageEvent.data, messageEvent.target, messageEvent

wsDaemonOnOpenHandler = () ->
	sendLoginCheck()

wsDaemonOnCloseHandler = (event) ->
	ws = event.target;
	console.log "Connection lost. IMPLEMENT RETRYING!"

wsDaemonOnErrorHandler = (error) ->
	console.error error

wsDaemonOnMessageHandler = (messageEvent) ->
	processIncomingMessage messageEvent.data, messageEvent.target, messageEvent

test = () ->
	if testOUT
		# THESE TEST IS DESIGNED FOR ECHO SERVER ONLY!!!
		# OUTGOING TEST. DON'T PAY ATTENTION TO THE ERRORS - ECHO SERVER GIVES BACK BAD DATA (NOT BAD, BUT THE SAME)
		console.log "OUTGOING TEST. DON'T PAY ATTENTION TO THE ERRORS - ECHO SERVER GIVES BACK BAD DATA (NOT BAD, BUT THE SAME)"
		sendLoginCheck()
		sendLogin "foo", "bar"
		sendLogout()
		sendDaemons()
		sendDaemon("123123123")
		sendControl("123123123", "DIE")	

	if testIN
		# INCOMING TEST
		console.log "INCOMING TEST"
		loginCheckMessageOK = 
			type: "loginCheck"
			data: "status": "OK"
		trySendMessage loginCheckMessageOK

		loginCheckMessageUNAUTHORIZED = 
			type: "loginCheck"
			data: "status": "UNAUTHORIZED"
		trySendMessage loginCheckMessageUNAUTHORIZED

		loginMessageNOTNULL = 
			type: "login"
			data: "session_id": "123123123"
		trySendMessage loginMessageNOTNULL
		
		loginMessageNULL = 
			type: "login"
			data: "session_id": null
		trySendMessage loginMessageNULL

		logoutMessageOK = 
			type: "logout"
			data: "status": "OK"
		trySendMessage logoutMessageOK
		
		logoutMessageNOTOK = 
			type: "logout"
			data: "status": "NOT_OK"
		trySendMessage logoutMessageNOTOK

		daemonsMessage = 
			type: "daemons"
			data: [
				{"daemon_id": "123", "daemon_name": "foo", "daemon_state": "RUNNING"},
				{"daemon_id": "234", "daemon_name": "bar", "daemon_state": "STOPPED"},
				{"daemon_id": "345", "daemon_name": "foobar", "daemon_state": "NOT_KNOWN"},
				{"daemon_id": "456", "daemon_name": "Bob", "daemon_state": "EATING A PIZZA"},
			]
		trySendMessage daemonsMessage

		daemonMessage = 
			type: "daemon"
			data: {"daemon_id": "123", "daemon_address": "123.123.123.123", "daemon_port": "666", "daemon_platform": {"OS": "Linux", "Architecture": "64 bit"}, "daemon_all_parameters": ["CPU", "RAM", "HDD"], "daemon_monitored_parameters": ["CPU"]}
		trySendMessage daemonMessage

		controlMessageOK = 
			type: "control"
			data: {"daemon_id": "123", "status": "OK"}
		trySendMessage controlMessageOK

		controlMessageNOTOK = 
			type: "control"
			data: {"daemon_id": "123", "status": "NOT_OK"}
		trySendMessage controlMessageNOTOK

		monitoringMessage1 = 
			type: "monitoring"
			data: {"daemon_id": "123", "data": {"CPU": "100", "RAM": "111"}}
		trySendMessage monitoringMessage1

		monitoringMessage2 = 
			type: "monitoring"
			data: {"daemon_id": "234", "data": {"CPU": "50", "RAM": "222"}}
		trySendMessage monitoringMessage2

		monitoringMessage3 = 
			type: "monitoring"
			data: {"daemon_id": "345", "data": {"CPU": "25", "RAM": "333"}}
		trySendMessage monitoringMessage3


		not_implementedMessage = 
			type: "not_implemented"
			data: {"this": "is a wrong message"}
		trySendMessage not_implementedMessage

		errorMessage = 
			type: "error"
			data: {"error": "error info"}
		trySendMessage errorMessage	