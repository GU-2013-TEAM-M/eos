LoginPageLayout = Backbone.Marionette.Layout.extend {
	template: "#login-page-layout-template",

	regions: {
		header: "header",
		content: "#content"
		footer: "footer"
	}

	onRender: () ->
		@header.show views.loginPageHeaderView
		@content.show views.loginPageContentView
		@footer.show views.loginPageFooterView
}