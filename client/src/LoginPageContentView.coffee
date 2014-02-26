LoginPageContentView = Backbone.Marionette.ItemView.extend {
	template: "#login-page-content-template",
	className: ".login-page-content",

	onRender: () ->
		$("button", @el).on("click", () ->
			console.log "Login pressed"
			login()
		)

}