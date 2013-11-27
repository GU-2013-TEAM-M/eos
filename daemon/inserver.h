#ifndef inserver_h__
#define inserver_h__

#include <websocketpp/config/asio_no_tls.hpp>
#include <websocketpp/server.hpp>

class serveToClient {

private:
	websocketpp::server<websocketpp::config::asio> s;
	websocketpp::connection_hdl handler;
	void on_message(websocketpp::connection_hdl, websocketpp::server<websocketpp::config::asio>::message_ptr);
	void on_open(websocketpp::connection_hdl);
	void on_close(websocketpp::connection_hdl);

public:
	serveToClient();
	~serveToClient();
	void send(std::string);
	void run();
	bool open;
};

#endif // inserver_h__