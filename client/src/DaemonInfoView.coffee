DaemonInfoView = Backbone.Marionette.ItemView.extend {

	# events: {
	# 	'click li': 'paramClicked',
	# },

	getTemplate: () ->
		if !@model
			"#daemon-info-template-empty"
		else
			"#daemon-info-template"

	onBeforeRender: () ->
		@model = appState.get("current_daemon")
		if !@model
			@model = daemons.models[0]
			appState.set("current_daemon", @model)
	
	onRender: () ->
		$("li", @el).on('click', (event) =>
			el = event.target
			param = $(el).text().trim()
			@model.toggleMonitor(param)
			@render()
		)
}