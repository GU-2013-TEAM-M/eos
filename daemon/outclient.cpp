#include "outclient.h"

outClient::outClient() {
	// set up access channels to only log interesting things
	m_client.clear_access_channels(websocketpp::log::alevel::all);
	m_client.set_access_channels(websocketpp::log::alevel::connect);
	m_client.set_access_channels(websocketpp::log::alevel::disconnect);
	m_client.set_access_channels(websocketpp::log::alevel::app);

	// Initialize the Asio transport policy
	m_client.init_asio();

	// Bind the handlers we are using
	using websocketpp::lib::placeholders::_1;
	using websocketpp::lib::bind;
	m_client.set_open_handler(bind(&outClient::on_open,this,::_1));
	m_client.set_close_handler(bind(&outClient::on_close,this,::_1));
	m_client.set_fail_handler(bind(&outClient::on_fail,this,::_1));
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

	// Grab a handle for this connection so we can talk to it in a thread
	// safe manor after the event loop starts.
	m_hdl = con->get_handle();

	// Queue the connection. No DNS queries or network connections will be
	// made until the io_service event loop is run.
	m_client.connect(con);

	// Create a thread to run the ASIO io_service event loop
	websocketpp::lib::thread asio_thread(&client::run, &m_client);

	// Create a thread to run the telemetry loop
	//websocketpp::lib::thread telemetry_thread(&outClient::telemetry_loop,this);

	asio_thread.join();
	//telemetry_thread.join();
}

void outClient::send(std::string payload) {
	websocketpp::lib::error_code ec;

	m_client.get_alog().write(websocketpp::log::alevel::app, payload);
	m_client.send(m_hdl,payload,websocketpp::frame::opcode::text,ec);

	// The most likely error that we will get is that the connection is
	// not in the right state. Usually this means we tried to send a
	// message to a connection that was closed or in the process of
	// closing. While many errors here can be easily recovered from,
	// in this simple example, we'll stop the telemetry loop.
	if (ec) {
		m_client.get_alog().write(websocketpp::log::alevel::app,"Send Error: "+ec.message());
	}
}

void outClient::on_open( websocketpp::connection_hdl ) {
	m_client.get_alog().write(websocketpp::log::alevel::app,
		"Connection opened");

	scoped_lock guard(m_lock);
	m_open = true;
}

void outClient::on_close( websocketpp::connection_hdl ) {
	m_client.get_alog().write(websocketpp::log::alevel::app,
		"Connection closed");

	scoped_lock guard(m_lock);
	m_done = true;
}

void outClient::on_fail( websocketpp::connection_hdl ) {
	m_client.get_alog().write(websocketpp::log::alevel::app,
		"Connection failed");

	scoped_lock guard(m_lock);
	m_done = true;
}
