package main

type civilian struct {
	name  string
	email string
	phone int64
}

type nineNine struct {
	civilian
	ltk bool
}

func (civil civilian) speak() string {
	return "Civilian" + civil.name
}
func (nN nineNine) speak() string {
	return "Civilian" + nN.name
}

type human interface {
	speak()
}

func main() {
	hitchcock := civilian{"hitchcock", "cock@gm", 456}
	jake := nineNine{civilian: civilian{"jake", "noise@gm", 99}, ltk: true}
	hitchcock.speak()
	jake.speak()
}
