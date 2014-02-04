AlertsTabLayout = Backbone.Marionette.Layout.extend {
	id: "AlertsTabWrapper"
	template: "#alerts-tab-layout",

	regions: {
		daemonList: "#daemonList"
		alertList: "#alertList"
	}

	onRender: () ->
		@daemonList.show views.daemonsView
}