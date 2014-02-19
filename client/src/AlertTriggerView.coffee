AlertTriggerView = Backbone.Marionette.ItemView.extend {
	template: "#alert-trigger-item-template",
	className: "alertItem",
	tagName: "li",
	
	events: {
		'click': 'triggerClicked',
	},
	
	triggerClicked: () ->
		appState.set("current_alert_trigger", @model)
	
}