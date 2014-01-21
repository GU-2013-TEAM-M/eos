AlertsView = Backbone.Maroinette.CompositeView.extend {
	template: "#alerts-list-template",
	itemView: AlertView,
	tagname: "ul"
}