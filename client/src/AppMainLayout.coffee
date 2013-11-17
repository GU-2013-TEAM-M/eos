AppMainLayout = Backbone.Marionette.Layout.extend {
	template: "#app-page-content-template",

	regions: {
		tabSelector: "#tab-selector"
		tab: "#tab"
	}

	initialize: () ->

	onRender: () ->
		@tabSelector.show views.tabSelectorView
		if appState.get("current_tab") == ""
			appState.set("current_tab", tabs.models[0].get("tab"))
		else 
			appState.trigger "change:current_tab"
}