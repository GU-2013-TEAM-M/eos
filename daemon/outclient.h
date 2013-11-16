#ifndef outclient_h__
#define outclient_h__

#include <websocketpp/config/asio_no_tls_client.hpp>

//Include this when devel-level debug logs are required
//#include "websocketpp/config/debug_asio_no_tls.hpp"
#include <websocketpp/client.hpp>

// This header pulls in the WebSocket++ abstracted thread support that will
// select between boost::thread and std::thread based on how the build system
// is configured.
#include <websocketpp/common/thread.hpp>

class outClient {

private:
	typedef websocketpp::client<websocketpp::config::asio_client> client;
	typedef websocketpp::lib::lock_guard<websocketpp::lib::mutex> scoped_lock;
	client m_client;
	websocketpp::connection_hdl m_hdl;
	websocketpp::lib::mutex m_lock;
	bool m_open;
	bool m_done;

public:
	outClient();
	void init(const std::string &);

	void send(std::string payload);

	// The open handler will signal that we are ready to start sending telemetry
	void on_open(websocketpp::connection_hdl);

	// The close handler will signal that we should stop sending telemetry
	void on_close(websocketpp::connection_hdl);

	// The fail handler will signal that we should stop sending telemetry
	void on_fail(websocketpp::connection_hdl);

};


#endif // outclient_h__