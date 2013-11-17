TabsView = Backbone.Marionette.CompositeView.extend {
	id: "tabs",
	template: "#tabs-template",
	itemView: TabView,

	appendHtml: (collectionView, itemView) -> 
		collectionView.$("#tabsList").append itemView.el
}