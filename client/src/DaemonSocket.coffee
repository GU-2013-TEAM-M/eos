DaemonSocket = Backbone.Model.extend {
	defaults: {
		ws: null
		reconnectTimer: null
		reconnectTiming: 5 * 1000
		daemon: null
	}	

	initialize: (daemon) ->	
		@set("daemon", daemon)	
		daemon.set("socket", @)
		@set("ws", @createDaemonSocket())
	
	# Raw web socket initializer	
	createDaemonSocket: () ->
		daemon = @get("daemon")
		ws =  new WebSocket daemon.get("daemon_address")

		ws.onopen = () ->
			console.log "Connection to the daemon is established."
			# appState.set "serverConnected", true

		ws.onclose = (event) =>
			wasClean = event.wasClean
			if (!wasClean)
				console.log "Connection to the daemon (" + ws.url + ") was lost unexpectedly. Re-connecting (" + @get("reconnectTiming")/1000 + ")"

				setTimeout () =>
					ws = @createDaemonSocket()
				, @get("reconnectTiming")

			else 
				console.log "Connection to the daemon was successfully closed."

			# appState.set "serverConnected", false

		ws.onerror = () ->
			console.log "An error occured while talking to the daemon (" + @.url + ")"

		ws.onmessage = (messageEvent) ->
			MessageProcessor.process messageEvent

		ws


	# Sends a message
	# Params:	message
	# Return: true/false depending on success
	sendMessage: (message) ->
		ws = @get("ws")
		if ws.readyState == 1
			stringMessage = JSON.stringify message
			ws.send stringMessage
			console.log message.type + " message was sent " + stringMessage
			return true    	
		else
			console.log "message was  NOT sent"
			return false

}