AlertsView = Backbone.Marionette.CompositeView.extend {
	template: "#alert-list-template",
	itemView: AlertView,
	itemViewContainer: "ul"
}