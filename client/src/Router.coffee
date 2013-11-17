Controller = Backbone.Router.extend {
    routes: {
        "": "welcome",
        "welcome": "welcome",
        "login": "login",
        "app": "app"
    },

    welcome: () ->
        appState.set("current_state_layout", layouts.welcomePageLayout)
    ,

    login: () ->
        appState.set("current_state_layout", layouts.loginPageLayout)
    ,

    app: () ->
        appState.set("current_state_layout", layouts.appPageLayout)
}