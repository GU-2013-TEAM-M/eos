EoS ProtoDaemon v0.02
=======================

Provides performance watching and reporting capabilities.

Changelog
----------
Version v0.02
* Refactored everything to better organize files and classes in order to make adding things easier
* Added a websockets client with the ability to send strings
* Added a windows CPU watcher

Version v0.01
* Initial version, provides a simple linux CPU watcher that parses proc/stat to find usage

Compilation
-----------
A makefile is provided to easily build under unix systems. The project is developed using Visual Studio, and the solution and project files are included for building under windows.

Requires boost 1.50.0 or higher (tested with 1.54.0)
Requires websocketpp 0.3.x or higher (http://www.zaphoyd.com/websocketpp)

The provided makefile assumes websocketpp is located under /usr/include/c++/websocketpp