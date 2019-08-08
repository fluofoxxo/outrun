package responses

import "github.com/fluofoxxo/outrun/objects"

type InformationResponse struct {
	BaseResponse
	Informations      []objects.Information  `json:"informations"`
	OperatorEachInfos []objects.OperatorInfo `json:"operatorEachInfos"`
	NumOperatorInfo   int64                  `json:"numOperatorInfo"`
}

func NewInformationResponse(base BaseInfo, informations []objects.Information, opInfo []objects.OperatorInfo) InformationResponse {
	numOperatorInfo := int64(len(opInfo))
	br := NewBaseResponse(base)
	ir := InformationResponse{
		br,
		informations,
		opInfo,
		numOperatorInfo,
	}
	return ir
}

// TODO: insert into consts
var DEFAULT_INFORMATIONS = []objects.Information{
	objects.NewInformation(531, 1, 1465808400, 1466413199, "3__90000001_14"),
	objects.NewInformation(6001070, 1, 1464336180, 1580608922, "1_NOTICEThis is the open beta version of the game. This means, that attempting to access some features may result in an error!Thank you for your attention.\r\n_10600001_0"),
	objects.NewInformation(1000230, 3, 1465981200, 1466413199, "1__90000001_1"),
	objects.NewInformation(1000157, 600, 1448614800, 1609459199, "1__90000002_1"),
}

func DefaultInformationResponse(base BaseInfo) InformationResponse {
	informations := DEFAULT_INFORMATIONS
	opInfo := []objects.OperatorInfo{}
	ir := NewInformationResponse(
		base,
		informations,
		opInfo,
	)
	return ir
}
