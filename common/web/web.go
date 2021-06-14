package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

// 1、注册接口，找到服务下面的所有接口
// 2、开启服务，监听服务
// 3、根据请求路由，调用注册的接口

type HandlerFunc func(ctx context.Context, in interface{})

// 定义一个 WebRoute 结构体用于存放单个路由
type WebRoute struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

type WebServer struct {
	port       string
	isBigCamel bool
	prefix     string //前缀
	service    string
	webRoutes  []WebRoute
}

func NewWebServe() *WebServer {
	return &WebServer{
		port:       "8080",
		prefix:     "api",
		isBigCamel: true,
		webRoutes:  make([]WebRoute, 0, 4),
	}
}

// 注册接口，找到服务下面的所有接口
func (ws *WebServer) Register(methods ...interface{}) {
	for _, method := range methods {
		typ := reflect.TypeOf(method)
		val := reflect.ValueOf(method)
		t := reflect.Indirect(val).Type()
		objName := t.Name()

		for i := 0; i < typ.NumMethod(); i++ {
			method := typ.Method(i)

			path, m := ws.getDefaultComment(objName, method.Name, method.Type.NumIn())
			handler := ws.getHandlerFunc(method.Name, method.Func, val)
			ws.webRoutes = append(ws.webRoutes, WebRoute{
				method.Name,
				m,
				path,
				handler,
			})
		}
	}
}

func (ws *WebServer) NewRouter() *mux.Router {
	// 创建 mux.Router 路由器示例
	router := mux.NewRouter().StrictSlash(true)

	// 遍历 web.go 中定义的所有 webRoutes
	for _, route := range ws.webRoutes {
		// 将每个 web 路由应用到路由器
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.Handler)
	}

	return router
}

func (ws *WebServer) WebStart(port string) {
	r := ws.NewRouter()
	http.Handle("/", r)
	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here
	if err != nil {
		fmt.Println("An error occured starting HTTP listener at port " + port)
		fmt.Println("Error: " + err.Error())
	}
}

func (ws *WebServer) getDefaultComment(objName, objFunc string, num int) (routerPath string, method string) {
	method = "ANY"
	if num == 2 { // parm 2 , post default
		method = "POST"
	}

	if ws.isBigCamel { // big camel style.大驼峰
		routerPath = ws.prefix + "/" + ws.service + "." + objName + "/" + objFunc
	} else {
		routerPath = ws.prefix + "/" + ws.service + "." + Ucfirst(objName) + "/" + Ucfirst(objFunc)
	}

	return
}

func (ws *WebServer) getHandlerFunc(methodName string, tvl, obj reflect.Value) http.HandlerFunc {
	typ := tvl.Type()
	if typ.NumIn() < 2 {
		panic(fmt.Sprintf("invalid method args of %s", methodName))
	}
	if typ.NumIn() == 2 {
		ctxType := typ.In(1)
		if ctxType != reflect.TypeOf(context.Background()) {
			panic(fmt.Sprintf("invalid method args context of %s", methodName))
		}
	}
	reqType, respType := typ.In(2), typ.Out(2)

	// 获取请求参数
	// 将请求参数转化为结构体
	return func(w http.ResponseWriter, r *http.Request) {
		// 解析请求参数，转换为第二个参数
		req, _ := reflect.New(reqType), reflect.New(respType)

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(req.Interface())
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		vals := make([]reflect.Value, 2)
		vals[0] = reflect.ValueOf(context.Background())
		vals[1] = reflect.ValueOf(req)

		rsv := tvl.Call(vals)
		if err := rsv[1].Interface().(error); err != nil {

		}
		resp := rsv[0].Interface()
		json.NewEncoder(w).Encode(&resp)

	}
}
