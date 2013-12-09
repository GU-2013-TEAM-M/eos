DaemonsView = Backbone.Marionette.CompositeView.extend {
	template: "#daemons-list-template"
	itemView: DaemonView
	itemViewContainer: "ul"

	# onRender: () ->
	# 	if daemons.length == 0 
	# 		$(@el).append("wait")

	onRender: () ->
		currentDaemon = appState.get("current_daemon")
		if currentDaemon
			el = views.daemonsView.children.findByModel(currentDaemon).el
			$(".activeDaemon").removeClass("activeDaemon")
			$(el).addClass("activeDaemon")


}