// From github.com/gin-contrib/sessions
//
// 配置文件格式：
//
//  tcgo:
//    session:
//      name:
//      key:
//      maxage: 0
//      redis:
//    	  addr: localhost:6379
//    	  password: xxxx
//    	  dbname: 0
//    	  mode: single or sentinel or cluster
//    	  addrs: ["localhost:6379"]
//    	  master: master
package sessions
