package obj

// TODO: create a param object that makes it easier to create informations

type OperatorInformation struct {
    ID      int64  `json:"id"`
    Content string `json:"content"`
}

func NewOperatorInformation(id int64, content string) OperatorInformation {
    return OperatorInformation{
        id,
        content,
    }
}
