HistoryContentView = Backbone.Marionette.ItemView.extend {
	template: "#history-content-template",	

	getTemplate: () ->
		if !@model
			"#history-content-template-empty"
		else
			"#history-content-template"	

	onBeforeRender: () ->
		@model = appState.get("current_daemon")
		if !@model
			@model = daemons.models[0]
			appState.set("current_daemon", @model)

	onRender: () ->
		if @model
			$("#dp_start", @el).datepicker()
			$("#dp_end", @el).datepicker()


			$("#requestHistory", @el).on('click', (event) =>
				dateStart = $("#dp_start", @el)
				dateEnd = $("#dp_end", @el)
				unixDateStart = new Date(dateStart[0].value).getTime()/1000
				unixDateEnd = new Date(dateEnd[0].value).getTime()/1000
				param = $("select", @el).val()
				console.log param
				daemon_id = @model.get("daemon_id")

				message = MessageProcessor.createMessage "monitoring", {daemon_id: daemon_id, from: unixDateStart, to: unixDateEnd, parameter: param}
				if message
					serverSocket.sendMessage message				
			)
}