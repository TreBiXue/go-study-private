package main

import "go-studying/cmd"

//	@title			Trial EC入荷 API
//	@version		1.0
//	@description	This is 「Trial EC入荷 API」 server.

// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	server := cmd.InitServer()
	server.Start()

}
