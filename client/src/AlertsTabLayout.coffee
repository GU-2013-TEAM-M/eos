AlertsTabLayout = Backbone.Marionette.Layout.extend {
	id: "AlertsTabWrapper"
	template: "#alerts-tab-layout",

	regions: {
		daemonList: "#daemonList"
		alertList: "#alertList"
		instanceList: "#instanceList"
	}

	onRender: () ->
		@daemonList.show views.daemonsView
		@alertList.show views.alertTriggersView
		@instanceList.show views.alertsView
}