package core

import (
	"html/template"
	"log"
	"net/http"

	"muidea.com/magicCommon/agent"
	"muidea.com/magicCommon/model"
	engine "muidea.com/magicEngine"
)

type route struct {
	pattern string
	method  string
	handler interface{}
}

func (s *route) Pattern() string {
	return s.pattern
}

func (s *route) Method() string {
	return s.method
}

func (s *route) Handler() interface{} {
	return s.handler
}

func newRoute(pattern, method string, handler interface{}) engine.Route {
	return &route{pattern: pattern, method: method, handler: handler}
}

// New 新建Protal
func New(centerServer, name, endpointID, authToken string) (Protal, bool) {
	protal := Protal{}

	agent := agent.New()
	authToken, sessionID, ok := agent.Start(centerServer, endpointID, authToken)
	if !ok {
		return protal, false
	}
	protalCatalog, ok := agent.FetchSummary(name, model.CATALOG, authToken, sessionID)
	if !ok {
		log.Print("fetch protal root ctalog failed.")
		return protal, false
	}

	protalContent := agent.QuerySummaryContent(protalCatalog.ID, model.CATALOG, authToken, sessionID)

	protal.centerAgent = agent
	protal.protalInfo = protalCatalog
	protal.protalContent = protalContent
	protal.endpointID = endpointID
	protal.authToken = authToken
	protal.sessionID = sessionID

	return protal, true
}

// Protal Protal对象
type Protal struct {
	centerAgent   agent.Agent
	protalInfo    model.SummaryView
	protalContent []model.SummaryView
	endpointID    string
	authToken     string
	sessionID     string
}

// Startup 启动
func (s *Protal) Startup(router engine.Router) {
	defaultRoute := newRoute("/", "GET", s.mainPage)
	router.AddRoute(defaultRoute)

	indexRoute := newRoute("/index.html", "GET", s.mainPage)
	router.AddRoute(indexRoute)

	productRoute := newRoute("/product.html", "GET", s.productPage)
	router.AddRoute(productRoute)

	blogRoute := newRoute("/blog.html", "GET", s.blogPage)
	router.AddRoute(blogRoute)

	aboutRoute := newRoute("/about.html", "GET", s.aboutPage)
	router.AddRoute(aboutRoute)

	contactRoute := newRoute("/contact.html", "GET", s.contactPage)
	router.AddRoute(contactRoute)

	noFoundRoute := newRoute("/404.html", "GET", s.noFoundPage)
	router.AddRoute(noFoundRoute)
}

// Teardown 销毁
func (s *Protal) Teardown() {
	if s.centerAgent != nil {
		s.centerAgent.Stop()
	}
}

func (s *Protal) getIndexView() (model.SummaryView, bool) {
	for _, v := range s.protalContent {
		if v.Name == "Index" && v.Type == model.CATALOG {
			return v, true
		}
	}

	return model.SummaryView{}, false
}

func (s *Protal) getProductView() (model.SummaryView, bool) {
	for _, v := range s.protalContent {
		if v.Name == "Product" && v.Type == model.CATALOG {
			return v, true
		}
	}

	return model.SummaryView{}, false
}

func (s *Protal) getBlogView() (model.SummaryView, bool) {
	for _, v := range s.protalContent {
		if v.Name == "Blog" && v.Type == model.CATALOG {
			return v, true
		}
	}

	return model.SummaryView{}, false
}

func (s *Protal) getAboutView() (model.SummaryView, bool) {
	for _, v := range s.protalContent {
		if v.Name == "About" && v.Type == model.ARTICLE {
			return v, true
		}
	}

	return model.SummaryView{}, false
}

func (s *Protal) getContactView() (model.SummaryView, bool) {
	for _, v := range s.protalContent {
		if v.Name == "Contact" && v.Type == model.ARTICLE {
			return v, true
		}
	}

	return model.SummaryView{}, false
}

func (s *Protal) get404View() (model.SummaryView, bool) {
	for _, v := range s.protalContent {
		if v.Name == "404" && v.Type == model.ARTICLE {
			return v, true
		}
	}

	return model.SummaryView{}, false
}

func (s *Protal) mainPage(res http.ResponseWriter, req *http.Request) {
	log.Print("mainPage")

	pageFile := "static/default/index.html"
	_, ok := s.getIndexView()
	if ok {
		pageFile = "static/template/index.html"
	}
	t, err := template.ParseFiles(pageFile)
	if err != nil {
		log.Printf("parseFiles exception, err:%s", err.Error())

		http.Redirect(res, req, "/404.html", http.StatusNotFound)
		return
	}

	t.Execute(res, nil)
}

func (s *Protal) productPage(res http.ResponseWriter, req *http.Request) {
	log.Print("productPage")

	pageFile := "static/default/product.html"
	_, ok := s.getProductView()
	if ok {
		pageFile = "static/template/product.html"
	}
	t, err := template.ParseFiles(pageFile)
	if err != nil {
		log.Printf("parseFiles exception, err:%s", err.Error())

		http.Redirect(res, req, "/404.html", http.StatusNotFound)
		return
	}

	t.Execute(res, nil)
}

func (s *Protal) blogPage(res http.ResponseWriter, req *http.Request) {
	log.Print("blogPage")

	pageFile := "static/default/blog.html"
	_, ok := s.getProductView()
	if ok {
		pageFile = "static/template/blog.html"
	}
	t, err := template.ParseFiles(pageFile)
	if err != nil {
		log.Printf("parseFiles exception, err:%s", err.Error())

		http.Redirect(res, req, "/404.html", http.StatusNotFound)
		return
	}

	t.Execute(res, nil)
}

func (s *Protal) aboutPage(res http.ResponseWriter, req *http.Request) {
	log.Print("aboutPage")

	pageFile := "static/default/about.html"
	_, ok := s.getProductView()
	if ok {
		pageFile = "static/template/about.html"
	}
	t, err := template.ParseFiles(pageFile)
	if err != nil {
		log.Printf("parseFiles exception, err:%s", err.Error())

		http.Redirect(res, req, "/404.html", http.StatusNotFound)
		return
	}

	t.Execute(res, nil)
}

func (s *Protal) contactPage(res http.ResponseWriter, req *http.Request) {
	log.Print("contactPage")

	pageFile := "static/default/contact.html"
	_, ok := s.getProductView()
	if ok {
		pageFile = "static/template/contact.html"
	}
	t, err := template.ParseFiles(pageFile)
	if err != nil {
		log.Printf("parseFiles exception, err:%s", err.Error())

		http.Redirect(res, req, "/404.html", http.StatusNotFound)
		return
	}

	t.Execute(res, nil)
}

func (s *Protal) noFoundPage(res http.ResponseWriter, req *http.Request) {
	log.Print("noFoundPage")

	pageFile := "static/default/404.html"
	_, ok := s.getProductView()
	if ok {
		pageFile = "static/template/404.html"
	}
	t, err := template.ParseFiles(pageFile)
	if err != nil {
		log.Printf("parseFiles exception, err:%s", err.Error())

		http.Redirect(res, req, "/404.html", http.StatusNotFound)
		return
	}

	t.Execute(res, nil)
}
