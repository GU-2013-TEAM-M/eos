AppState = Backbone.Model.extend {
    defaults: {
        user_id: "",
        user_name: "",
        state: "",
        current_state_layout: "",
        current_tab: "",
        server_address: ""
    }

    # triggers:
    	
    initialize: () ->
    	@listenTo this, "change:current_state_layout", () ->
            MyApp.mainRegion.show @get "current_state_layout"

        @listenTo this, "change:current_tab", () ->
            layouts.appMainLayout.tab.show @get "current_tab"
}