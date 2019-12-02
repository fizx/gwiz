package app

type App struct {
}

func (*App) Execute() {

}

func NewApp() *App {
	return &App{}
}
