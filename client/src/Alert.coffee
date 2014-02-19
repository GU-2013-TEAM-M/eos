# An instance of one of a daemon's triggers

Alert = Backbone.Model.extend {
	
	defaults: {
		trigger: null
		time: null
		value: null
	}
	
}