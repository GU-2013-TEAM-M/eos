WelcomePageLayout = Backbone.Marionette.Layout.extend {
	template: "#welcome-page-layout-template",

	regions: {
		header: "header",
		content: "#content"
		footer: "footer"
	},

	onRender: () ->
		@header.show views.welcomePageHeaderView
		@content.show views.welcomePageContentView
		@footer.show views.welcomePageFooterView
}