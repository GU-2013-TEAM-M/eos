AlertsView = Backbone.Marionette.CompositeView.extend {
	template: "#alerts-list-template",
	itemView: AlertView,
	tagname: "ul"
}