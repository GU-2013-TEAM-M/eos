AlertTriggersView = Backbone.Marionette.CompositeView.extend {
	template: "#alert-trigger-list-template",
	itemView: AlertTriggerView,
	itemViewContainer: "ul",
	
	onBeforeRender: () ->
		@model = appState.get("current_alert_trigger")
		if !@model
			@model = alerts.models[0]
			appState.set("current_alert_trigger", @model)
			
			###
	
	onRender: () ->
		currentAlert = appState.get("current_alert")
		if currentAlert
			el = views.alertsView.children.findByModel(currentAlert).el
			$(".activeAlert").removeClass("activeAlert")
			$(el).addClass("activeAlert")
			
			###
}