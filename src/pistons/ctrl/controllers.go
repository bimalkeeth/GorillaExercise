package ctrl

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/reverse"
	"html/template"
	"net/http"
	"net/url"
)

var (
	login    *loginController    = new(loginController)
	parts    *partController     = new(partController)
	checkout *checkoutController = new(checkoutController)
	admin    *adminController    = new(adminController)
)

var (
	partModelsRegexp *reverse.Regexp
	TemplateFunc     template.FuncMap
)

func init() {
	partModelsRegexp, _ = reverse.CompileRegexp(`/parts/makes/(?p<makeId>.+)/models`)
	TemplateFunc = template.FuncMap{"partModels": reverter(partModelsRegexp)}
}

func reverter(regexp *reverse.Regexp) func(...string) (string, error) {
	return func(params ...string) (string, error) {
		values := url.Values{}
		for i := 0; i < len(params); i += 2 {
			values[params[i]] = []string{params[i+1]}
		}
		return regexp.Revert(values)
	}
}

func Setup(tc *template.Template) {
	SetTemplateCache(tc)
	createResourceServer()

	r := mux.NewRouter()

	r.HandleFunc("/", login.GetLogin)
	r.HandleFunc("/parts/makes", parts.GetMake)
	r.HandleFunc("/parts/models", parts.PostModel)
	r.HandleFunc("/parts/years", parts.PostYear)
	r.HandleFunc("/parts/engines", parts.PostEngine)
	r.HandleFunc("/parts/searchresults", parts.PostSearch)
	r.HandleFunc("/parts", parts.GetPartSearchPartial)
	r.HandleFunc("/parts/detail", parts.GetPart)
	r.HandleFunc("/checkout", checkout.HandleCheckout)
	r.HandleFunc("/admin", admin.HandleLogin)
	r.HandleFunc("/admin/menu", admin.GetMenu)
	r.HandleFunc("/admin/employees/new", admin.HandleCreateEmp)
	r.HandleFunc("/admin/employee", admin.GetEmployeeView)

	r.HandleFunc("/api/makes", parts.AutocompleteMake)
	r.HandleFunc("/api/models", parts.AutocompleteModel)
}

func createResourceServer() {
	http.Handle("/res/lib/", http.StripPrefix("/res/lib", http.FileServer(http.Dir("node_modules"))))
	http.Handle("/res/", http.StripPrefix("/res", http.FileServer(http.Dir("res"))))
}

func SetTemplateCache(tc *template.Template) {
	login.loginTemplate = tc.Lookup("login.html")

	parts.autoMakeTemplate = tc.Lookup("make.html")
	parts.autoModelTemplate = tc.Lookup("model.html")
	parts.autoYearTemplate = tc.Lookup("year.html")
	parts.autoEngineTemplate = tc.Lookup("engine.html")
	parts.searchResultTemplate = tc.Lookup("search_results.html")
	parts.partTemplate = tc.Lookup("part.html")
	parts.searchResultPartialTemplate = tc.Lookup("_result.html")

	checkout.template = tc.Lookup("checkout.html")

	admin.loginTemplate = tc.Lookup("admin_login.html")
	admin.menuTemplate = tc.Lookup("admin_menu.html")
	admin.createEmpTemplate = tc.Lookup("admin_create_emp.html")
	admin.viewEmpTemplate = tc.Lookup("admin_employee.html")
}
