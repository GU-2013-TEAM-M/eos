HomeTabLayout = Backbone.Marionette.Layout.extend {
	id: "homeTabWrapper"
	template: "#home-tab-layout",

	regions: {
		daemonList: "#daemonList"
		daemonInfo: "#daemonInfo"
	}

	onRender: () ->
		@daemonList.show views.daemonsView
		@daemonInfo.show views.daemonInfoView
}