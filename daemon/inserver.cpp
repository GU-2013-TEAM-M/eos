#include "inserver.h"

serveToClient::serveToClient() {
	open = false;
}

void serveToClient::run() {
	try {
		// Set logging settings
		s.set_access_channels(websocketpp::log::alevel::all);
		s.clear_access_channels(websocketpp::log::alevel::frame_payload);

		// Initialize ASIO
		s.init_asio();

		// Register our message handler
		s.set_message_handler(bind(&serveToClient::on_message,this,::_1,::_2));
		s.set_open_handler(bind(&serveToClient::on_open,this,::_1));
		s.set_close_handler(bind(&serveToClient::on_close,this,::_1));

		// Listen on port 9005
		s.listen(9005);

		// Start the server accept loop
		s.start_accept();
		std::cout<<"listening"<<std::endl;
		// Start the ASIO io_service run loop
		s.run();
	} catch (const std::exception & e) {
		std::cout << e.what() << std::endl;
	} catch (websocketpp::lib::error_code e) {
		std::cout << e.message() << std::endl;
	} catch (...) {
		std::cout << "other exception" << std::endl;
	}
}

void serveToClient::on_message( websocketpp::connection_hdl hdl, websocketpp::server<websocketpp::config::asio>::message_ptr msg) {
	std::cout << "on_message called with hdl: " << hdl.lock().get()	<< " and message: " << msg->get_payload() << std::endl;
}

void serveToClient::on_open( websocketpp::connection_hdl hdl) {
	handler = hdl;
	std::cout<<"set handler"<<std::endl;
	open=true;
}

void serveToClient::on_close( websocketpp::connection_hdl hdl) {
	std::cout<<"conn handler to client lost"<<std::endl;
	open=false;
}

void serveToClient::send(std::string payload) {
	websocketpp::lib::error_code ec;
	try {
		s.send(handler, payload, websocketpp::frame::opcode::text);
	} catch (const websocketpp::lib::error_code& e) {
		std::cout << "Sending failed: " << e << "(" << e.message() << ")" << std::endl;
	}
}


