# An instance of one of a daemon's alerts being triggered

AlertInstance = Backbone.Model.extend {
	
	defaults: {
		alert: null
		time: null
		value: null
	}
	
}