DaemonView = Backbone.Marionette.ItemView.extend {
	template: "#daemon-item-template",
	className: "daemonItem"
	tagName: "li"

	events: {
		'click': 'daemonClicked',
	},

	modelEvents: {
		'change': 'fieldsChanged'
	},

	fieldsChanged: () ->
		@render()


	daemonClicked: () ->
		appState.set("current_daemon", @model)
}