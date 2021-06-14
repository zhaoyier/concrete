package spike

// import (
// 	"context"
// 	"reflect"
// )

// // type WebRoutes []WebRoute
// type HandleFunc func(ctx context.Context, in interface{})

// // 定义一个 WebRoute 结构体用于存放单个路由
// type WebRoute struct {
// 	Name        string
// 	Method      string
// 	Pattern     string
// 	HandlerFunc HandleFunc
// }

// type WebServer struct {
// 	port       string
// 	isBigCamel bool
// 	prefix     string //前缀
// 	service    string
// 	webRoutes  []WebRoute
// }

// func NewWebServe() *WebServer {
// 	return &WebServer{
// 		prefix:     "api",
// 		isBigCamel: true,
// 		webRoutes:  make([]WebRoute, 0, 4),
// 	}
// }

// func (ws *WebServer) Register(methods ...interface{}) {
// 	for _, method := range methods {
// 		typ := reflect.TypeOf(method)
// 		val := reflect.ValueOf(method)
// 		t := reflect.Indirect(val).Type()
// 		objName := t.Name()

// 		for i := 0; i < typ.NumMethod(); i++ {
// 			method := typ.Method(i)
// 			// num, _b := b.checkHandlerFunc(method.Type /*.Interface()*/, true)
// 			path, m := ws.getDefaultComment(objName, method.Name, method.Type.NumIn())
// 			// wb.registerHandlerObj(router, methods, routerPath, method.Name, method.Func, refVal)
// 			ws.webRoutes = append(ws.webRoutes, WebRoute{
// 				"Home",
// 				"GET",
// 				path,
// 				nil,
// 			})
// 		}
// 	}
// }

// func (ws *WebServer) getDefaultComment(objName, objFunc string, num int) (routerPath string, method string) {
// 	method = "ANY"
// 	if num == 2 { // parm 2 , post default
// 		method = "POST"
// 	}

// 	if ws.isBigCamel { // big camel style.大驼峰
// 		routerPath = ws.prefix + "/" + ws.service + "." + objName + "/" + objFunc
// 	} else {
// 		routerPath = ws.prefix + "/" + ws.service + "." + Ucfirst(objName) + "/" + Ucfirst(objFunc)
// 	}

// 	return
// }

// func (ws *WebServer) registerHandlerObj() {

// }

// func (ws *WebServer) WebStart(port string) {

// }

// func New() *Spike {
// 	return new(Spike)
// }

// func (s *Spike) Port() {

// }

// func (s *Spike) AddMiddleware() {

// }

// func (s *Spike) RegisterServer() {

// }

// func (s *Spike) Serve() {

// }
