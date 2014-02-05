ServerSocket = Backbone.Model.extend {
	defaults: {
		ws: null,
		reconnectTimer: null,
		reconnectTiming: 5 * 1000,
	}	

	initialize: () ->
		@set("ws", @createServerSocket())
	
	# Raw web socket initializer	
	createServerSocket: () ->
		ws =  new WebSocket appState.get "server_address"

		ws.onopen = () =>
			console.log "Connection to the server is established."
			appState.set "serverConnected", true

			message = MessageProcessor.createMessage "loginCheck", session_id: appState.get("session_id")
			if message
				@sendMessage message

		ws.onclose = (event) =>
			wasClean = event.wasClean
			if (!wasClean)
				console.log "Connection to the server (" + ws.url + ") was lost unexpectedly. Re-connecting (" + @get("reconnectTiming")/1000 + ")"

				setTimeout () =>
					ws = @createServerSocket()
				, @get("reconnectTiming")

			else 
				console.log "Connection to the server was successfully closed."

			appState.set "serverConnected", false

		ws.onerror = () ->
			console.log "An error occured while talking to the server (" + @.url + ")"

		ws.onmessage = (messageEvent) ->
			#console.log messageEvent
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
