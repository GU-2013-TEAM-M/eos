class MessageProcessor

	# Creates a new message
	# Params:	type - type of the message
	#			data - message data
	#			direction - in or out
	# Return: a new message	
	@createMessage: (type, data, direction) ->
		message = 
			type: type
			data: data
		if @checkMessage message, direction
			message
		else
			null


	@process: (messageEvent) ->
		data = messageEvent.data
		target = messageEvent.target

		console.log "Incoming message: " + data
		
		message = JSON.parse data
		if @checkMessage message, "in"
			processCallback = Messages.getMessageCallback("in", message.type)
			processCallback message.data


	@checkMessage: (message, direction) ->
		# todo
		true