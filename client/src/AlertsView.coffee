AlertsView = Backbone.Maroinette.CompositeView.extend {
	template: "#alert-list-template",
	itemView: AlertView,
	tagname: "ul"
}