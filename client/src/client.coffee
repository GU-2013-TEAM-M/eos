tabs = []

daemons = new Daemons [
	# {"daemon_id": "123", "daemon_name": "foo", "daemon_state": "RUNNING", "daemon_address": "ws://127.0.0.1:8080/ws", "daemon_port": "666", "daemon_platform": {"OS": "Linux", "Architecture": "64 bit"}, "daemon_all_parameters": ["CPU", "RAM", "HDD"], "daemon_monitored_parameters": ["CPU"]},
	# {"daemon_id": "234", "daemon_name": "bar", "daemon_state": "STOPPED", "daemon_address": "123.456.789.123", "daemon_port": "777", "daemon_platform": {"OS": "Linux", "Architecture": "32 bit"}, "daemon_all_parameters": ["CPU", "RAM", "HDD"], "daemon_monitored_parameters": ["CPU", "RAM"]},
	# {"daemon_id": "345", "daemon_name": "foobar", "daemon_state": "NOT_KNOWN", "daemon_address": "456.789.123.456", "daemon_port": "111", "daemon_platform": {"OS": "Linux", "Architecture": "128 bit"}, "daemon_all_parameters": ["CPU", "RAM", "HDD"], "daemon_monitored_parameters": ["CPU", "HDD"]},
	# {"daemon_id": "456", "daemon_name": "Bob", "daemon_state": "EATING A PIZZA", "daemon_address": "789.123.456.789", "daemon_port": "222", "daemon_platform": {"OS": "Windows 7", "Architecture": "16 bit"}, "daemon_all_parameters": ["CPU", "RAM", "HDD"], "daemon_monitored_parameters": ["HDD"]},
]

graphs = new Graphs()

views = {}
layouts = {}

serverSocket = {}

appState = new AppState({server_address: "ws://127.0.0.1:8080/ws"});

router = new Router()

MyApp = new Backbone.Marionette.Application()

MyApp.addRegions {
	mainRegion: "#mainRegion"
}

MyApp.addInitializer (options) ->
	tabs = new Tabs [	
	    { name: "Home", tab: new HomeTabLayout() },
	    { name: "History", tab: new HistoryTabLayout() },
	    { name: "Alerts", tab: new AlertsTabLayout() },
	    { name: "User", tab: new UserTabLayout() },
	]


	layouts = {
		welcomePageLayout: new WelcomePageLayout(),
		loginPageLayout: new LoginPageLayout(),
		appPageLayout: new AppPageLayout(),
		appMainLayout: new AppMainLayout()
	}

	views = {
		tabSelectorView: new TabSelectorView({collection: tabs}),
		appPageHeaderView: new AppPageHeaderView(),
		appPageFooterView: new AppPageFooterView(),
		loginPageHeaderView: new LoginPageHeaderView(),
		loginPageContentView: new LoginPageContentView(),
		loginPageFooterView: new LoginPageFooterView(),
		welcomePageHeaderView: new WelcomePageHeaderView(),
		welcomePageContentView: new WelcomePageContentView(),
		welcomePageFooterView: new WelcomePageFooterView(),

		daemonsView: new DaemonsView({collection: daemons}),
		daemonInfoView: new DaemonInfoView(),
	}


MyApp.addInitializer (options) ->
	serverSocket = new ServerSocket()



# MyApp.addRegions {
# 	tabsRegion: "#tabs"
# 	daemonsListRegion: "#daemonsList"
# }

# MyApp.addInitializer (options) ->
# 	tabsView = new TabsView {
# 		collection: options.tabs
# 	}
# 	MyApp.tabsRegion.show(tabsView)

# MyApp.addInitializer (options) ->
# 	daemonsView = new DaemonsView {
# 		collection: options.daemons
# 	}
# 	MyApp.daemonsListRegion.show(daemonsView)


MyApp.addInitializer (options) ->
	Backbone.history.start()

$(document).ready () ->
	MyApp.start()
	router.navigate("welcome", {trigger: true})