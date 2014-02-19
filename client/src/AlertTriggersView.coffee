AlertTriggersView = Backbone.Marionette.CompositeView.extend {
	template: "#alert-trigger-list-template",
	itemView: AlertTriggerView,
	itemViewContainer: "ul",
	
	onRender: () ->
		currentTrigger = appState.get("current_alert_trigger")
		if currentTrigger 
			el = views.alertTriggersView.children.findByModel(currentTrigger).el
			$(".activeTrigger").removeClass("activeTrigger")
			$(el).addClass("activeTrigger")
			

}