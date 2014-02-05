GraphView = Backbone.Marionette.ItemView.extend {
	template: "#graph-item-template",

	render: () ->
		console.log @model
}