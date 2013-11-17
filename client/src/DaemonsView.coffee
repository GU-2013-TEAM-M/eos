DaemonsView = Backbone.Marionette.CompositeView.extend {
	id: "daemons"
	template: "#daemons-list-template",
	itemView: DaemonView,

	appendHtml: (collectionView, itemView) -> 
		collectionView.$("ul").append itemView.el	
}