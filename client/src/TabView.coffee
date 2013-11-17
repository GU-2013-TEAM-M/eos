TabView = Backbone.Marionette.ItemView.extend {
	template: "#tab-template",
	className: "tabItem",
	tagName: "li"

	events: {
		'click': 'tabClicked',
	},

	tabClicked: () ->
		appState.set("current_tab", @model.get("tab"))
}