AppState = Backbone.Model.extend {
    defaults: {
        user_id: "",
        user_name: "",
        state: "",
        current_state_layout: "",
        current_tab: "",
        server_address: null,
        current_daemon: null, current_alert_trigger: null,
        serverConnected: false,
        session_id: Service.getCookie("session_id"),
        is_user_logged_in: false,
    }

    initialize: () ->
    	@listenTo this, "change:current_state_layout", () ->
            MyApp.mainRegion.show @get "current_state_layout"

        @listenTo this, "change:current_tab", () ->
            layouts.appMainLayout.tab.show @get "current_tab"

        @listenTo this, "change:current_daemon", () ->
            views.daemonsView.render()
            views.daemonInfoView.render()
			
        @listenTo this, "change:current_alert_trigger", () ->
            views.alertTriggersView.render(); views.alertsView.render(); console.log "Hi"
		
}