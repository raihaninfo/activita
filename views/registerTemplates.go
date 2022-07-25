package views

var (
	HomeView *View = NewView("views/ui/home.gohtml")
	AddView *View = NewView("views/ui/activity.gohtml")
	LoginView *View = NewView("views/ui/login.gohtml")
	RegisterView *View = NewView("views/ui/register.gohtml")
	ForgotView *View = NewView("views/ui/request-forgot.gohtml")
	ProfileView *View = NewView("views/ui/profile.gohtml")
	VerifyView *View = NewView("views/ui/verify.gohtml")
	VerifyReqView *View = NewView("views/ui/verify-req.gohtml")
	ResetPassView *View = NewView("views/ui/re-password.gohtml")
	UpdateView *View = NewView("views/ui/update-activity.gohtml")
	NotFoundView *View = NewView("views/ui/404.gohtml")
)