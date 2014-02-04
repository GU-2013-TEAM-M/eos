AppPageLayout = Backbone.Marionette.Layout.extend {
	template: "#app-page-layout-template",

	regions: {
		header: "header",
		content: "#content"
		footer: "footer"
	}

	onRender: () ->
		@header.show views.appPageHeaderView
		@content.show layouts.appMainLayout
		@footer.show views.appPageFooterView
}