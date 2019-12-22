package controller

type TodoIF interface {
	list()
	get()
	create()
	update()
	delete()
}
