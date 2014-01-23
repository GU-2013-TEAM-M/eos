AlertInstancesView = Backbone.Marionette.CompositeView.extend {
	template: "#alert-instance-list-template",
	itemView: AlertInstance,
	tagName: "ul"
}