HistoryTabLayout = Backbone.Marionette.Layout.extend {
	template: "#history-tab-layout",

	regions: {
		daemonList: "#daemonList"
		historyContent: "#historyContent"
	}

	onRender: () ->
		@daemonList.show views.daemonsView
		@historyContent.show views.historyContentView
}