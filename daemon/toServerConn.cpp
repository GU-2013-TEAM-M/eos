#include "toServerConn.h"

outClient::outClient() {
	ready = false;
	
	// set up access channels to only log interesting things
	m_client.clear_access_channels(websocketpp::log::alevel::all);
	//m_client.set_access_channels(websocketpp::log::alevel::devel);
	m_client.set_access_channels(websocketpp::log::alevel::connect);
	m_client.set_access_channels(websocketpp::log::alevel::disconnect);
	m_client.set_access_channels(websocketpp::log::alevel::app);

	// Initialize the Asio transport policy
	m_client.init_asio();

	// Bind the handlers we are using
	m_client.set_open_handler(websocketpp::lib::bind(&outClient::on_open,this,websocketpp::lib::placeholders::_1));
	m_client.set_close_handler(websocketpp::lib::bind(&outClient::on_close,this,websocketpp::lib::placeholders::_1));
	m_client.set_fail_handler(websocketpp::lib::bind(&outClient::on_fail,this,websocketpp::lib::placeholders::_1));
	m_client.set_message_handler(websocketpp::lib::bind(&outClient::on_message,this,websocketpp::lib::placeholders::_1,websocketpp::lib::placeholders::_2));
}

void outClient::init( const std::string & uri) {
	// Create a new connection to the given URI
	websocketpp::lib::error_code ec;
	client::connection_ptr con = m_client.get_connection(uri, ec);
	if (ec) {
		m_client.get_alog().write(websocketpp::log::alevel::app,
			"Get Connection Error: "+ec.message());
		return;
	}
	con->replace_header("Origin","shacron.twilightparadox.com:8080");
	// Grab a handle for this connection so we can talk to it in a thread
	// safe manor after the event loop starts.
	m_hdl = con->get_handle();

	// Queue the connection. No DNS queries or network connections will be
	// made until the io_service event loop is run.

	m_client.connect(con);
}

void outClient::run() {
	// Create a thread to run the ASIO io_service event loop
	websocketpp::lib::thread asio_thread(&client::run, &m_client);

	asio_thread.join();
	//telemetry_thread.join();
}

void outClient::bindMsg(std::function<void(std::string)> msgF) {
	msgRecv = msgF;
}

void outClient::send(std::string payload) {
	websocketpp::lib::error_code ec;

	m_client.get_alog().write(websocketpp::log::alevel::app, payload);
	m_client.send(m_hdl,payload,websocketpp::frame::opcode::text,ec);

	// The most likely error that we will get is that the connection is
	// not in the right state. Usually this means we tried to send a
	// message to a connection that was closed or in the process of
	// closing.
	if (ec) {
		m_client.get_alog().write(websocketpp::log::alevel::app,"Send Error: "+ec.message());
	}
}

void outClient::on_message( websocketpp::connection_hdl hdl, websocketpp::client<websocketpp::config::asio_client>::message_ptr msg) {
	std::cout << "on_message called with hdl: " << hdl.lock().get()	<< " and message: " << msg->get_payload() << std::endl;
	msgRecv(msg->get_payload());
}

void outClient::on_open( websocketpp::connection_hdl ) {
	m_client.get_alog().write(websocketpp::log::alevel::app,
		"Connection opened");
	ready = true;
	scoped_lock guard(m_lock);
	m_open = true;
}

void outClient::on_close( websocketpp::connection_hdl ) {
	m_client.get_alog().write(websocketpp::log::alevel::app,
		"Connection closed");
	ready = false;
	scoped_lock guard(m_lock);
	m_done = true;
}

void outClient::on_fail( websocketpp::connection_hdl ) {
	m_client.get_alog().write(websocketpp::log::alevel::app,
		"Connection failed");

	scoped_lock guard(m_lock);
	m_done = true;
}
