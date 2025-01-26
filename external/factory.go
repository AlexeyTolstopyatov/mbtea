package external

// (C) Bilbo Backends
// Schizo notes

type Mbtea struct {
	Type string `json:"type"`
	// Function suggestions based on answers for every
	// function's positions [Dom, Aux, Tert, Inf]
	FuncSuggestions struct {
		Ti [4]float64 `json:"Ti"`
		Fi [4]float64 `json:"Fi"`
		Si [4]float64 `json:"Si"`
		Ni [4]float64 `json:"Ni"`
		Te [4]float64 `json:"Te"`
		Se [4]float64 `json:"Se"`
		Fe [4]float64 `json:"Fe"`
		Ne [4]float64 `json:"Ne"`
	} `json:"func_suggestions"`
	// Axes suggestions based on answers too
	// if result's offset moves to right function position
	// it will be negative.
	// (e.g. NiSe offset -0.0003 means offset to Extroverted Sensing)
	AxesSuggestions struct {
		NiSe int32 `json:"NiSe"`
		NeSi int32 `json:"NeSi"`
		TiFe int32 `json:"TiFe"`
		TeFi int32 `json:"TeFi"`
	} `json:"axes_suggestions"`

	// This is what I've seen while sleeping.
	// Usually, questions in test may be very abstract.
	// (means: answers scope is 1 to 5) but they require concrete
	// answer, that not fully describes us. User's tries to select
	// mostly correct answer, I suggest, his emotional deviation
	// says about it. Count and Interval of user's selections can tell
	// about his problems. Once exception: Question may be read incorrect.
	Details struct {
		// If client selects 1 ("No") but next time this answer changes to
		// 4 ("May be Yes"), his deviation processes like |1 - 4| = 3, and
		// memorizes like inverval of free-moves.
		// It recognizes like 2 ways: or Middle of two/three answers
		// or simply array segment of choises, what needs for next static analysis.
		AnswerSelectionScope float64 `json:"answer_selection_scope"`
		// If client's count of selection more than 2, it may be
		// "red flag" and tells about some troubles.
		AnswerSelectionCount int32 `json:"answer_selection_count"`
		// Client's middle value of free-moves.
		AnswerSelectionMid float64 `json:"answer_selection_mid"`
		// When client tries to read question and compare it,
		// may be situations of UI reflection, when user selects text of
		// question, or clicks to the UI block with question.
		// I'll make sence on it.
		QuestionSelectionCount int32 `json:"question_selection_count"`
	} `json:"details"`
}

// New calls internal modules for make structure
// and fill results
func New() {

}
