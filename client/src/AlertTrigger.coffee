AlertTrigger = Backbone.Model.extend {
	
	idAttribute: "trigger_id",
	
	defaults: {
		daemon_id: null
		trigger_id: null
		trigger_name: null
		trigger_parameter: null
		trigger_min: null
		trigger_max: null
	}

}