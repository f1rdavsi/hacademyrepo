package student

type Subjects struct {
	Math       int
	Physics    int
	Chemistry  int
	Literature int
}

type Candidate struct {
	Name     string
	Age      int
	Score    int
	Subjects Subjects
}
