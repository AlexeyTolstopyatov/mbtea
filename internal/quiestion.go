package internal

type Question struct {
	TargetId                  string `json:"target_id"`
	TargetPosition            int32  `json:"target_position"`
	Text                      string `json:"text"`
	Answer                    int32  `json:"answer"`
	Count                     int32  `json:"count"`
	Scope                     int32  `json:"scope"`
	UseAnswerSelectionCount   bool   `json:"use_answer_selection_count"`
	UseAnswerSelectionScope   bool   `json:"use_answer_selection_scope"`
	UseQuestionSelectionCount bool   `json:"use_question_selection_count"`
}
