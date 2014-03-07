AlertsView = Backbone.Marionette.CompositeView.extend {
	template: "#alert-list-template",
	itemView: AlertView,
	itemViewContainer: "ul",
	
	template: "#alert-list-template",
	
	###
	
	getTemplate: () ->
		if (@currentAlerts.length() == 0)
			"#alert-list-template-empty"
		else
			"#alert-list-template"
	
	###

	onBeforeRender: () ->
		alertTrigger = appState.get("current_alert_trigger")
		
		if !alertTrigger
			alertTrigger = alertTriggers.models[0]
			appState.set("current_alert_trigger", alertTrigger)
			
		id = alertTrigger.get("trigger_id")
		currentAlerts = new Alerts
		
		for alert in alerts.models
			if alert.get("trigger_id") == id
				currentAlerts.add(alert)
				
		@model = currentAlerts
		views.alertsView = new AlertsView({collection: currentAlerts})
			
		
		console.log @model
		###
		@model = appState.get("current_alert_trigger")
		if !@model
			@model = alertTriggers.models[0]
			appState.set("current_alert_trigger", @model)
		@currentAlerts = alerts.where(trigger_id: @model.get("trigger_id"))
		###
		
}