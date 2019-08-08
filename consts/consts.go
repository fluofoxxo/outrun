package consts

import (
	"strconv"
)

type EMess string

func (s EMess) MarshalJSON() ([]byte, error) {
	return []byte(strconv.QuoteToASCII(string(s))), nil
}

// player states
const STATE_REGISTER = "0"
const STATE_LOGIN_STAGE_1 = "1"
const STATE_LOGIN_STAGE_2 = "2"

// error messages
var EM_REQUEST_PASSWORD = "\u30d1\u30b9\u30ef\u30fc\u30c9\u304c\u4e0d\u6b63"
var EM_OK = "\u6210\u529f"

// status codes
const SC_SERVER_SECURITY_ERROR = -19001
const SC_VERSION_DIFFERENCE = -19002
const SC_DECRYPTION_FAILURE = -19003
const SC_PARAM_HASH_DIFFERENCE = -19004
const SC_SERVER_NEXT_VERSION = -19990
const SC_SERVER_MAINTENANCE = -19997
const SC_SERVER_BUSY_ERROR = -19998
const SC_SERVER_SYSTEM_ERROR = -19999
const SC_REQUEST_PARAM_ERROR = -10100
const SC_NOT_AVAILABLE_PLAYER = -10101
const SC_MISSING_PLAYER = -10102
const SC_EXPIRATION_SESSION = -10103
const SC_INVALID_PASSWORD = -10104
const SC_INVALID_SERIAL_CODE = -10105
const SC_USED_SERIAL_CODE = -10106
const SC_HSP_WEB_API_ERROR = -10110
const SC_APOLLO_WEB_API_ERROR = -10115
const SC_DATA_MISMATCH = -30120
const SC_MASTER_DATA_MISMATCH = -10121
const SC_NOT_ENOUGH_RED_STAR_RINGS = -21030
const SC_NOT_ENOUGH_RINGS = -20131
const SC_NOT_ENOUGH_ENERGY = -20132
const SC_ROULETTE_USE_LIMIT = -30401
const SC_ROULETTE_BOARD_RESET = -30411
const SC_CHARACTER_LEVEL_LIMIT = -20601
const SC_ALL_CHAO_LEVEL_LIMIT = -20602
const SC_ALREADY_INVITED_FRIEND = -30801
const SC_ALREADY_REQUESTED_ENERGY = -30901
const SC_ALREADY_SENT_ENERGY = -30902
const SC_RECEIVE_FAILURE_MESSAGE = -30910
const SC_ALREADY_EXISTED_PRE_PURCHASE = -11001
const SC_ALREADY_REMOVED_PRE_PURCHASE = -11002
const SC_INVALID_RECEIPT_DATA = -11003
const SC_ALREADY_PROCESSED_RECEIPT = -11004
const SC_ENERGY_LIMIT_PURCHASE_TRIGGER = -21010
const SC_NOT_START_EVENT = -10201
const SC_ALREADY_END_EVENT = -10202
const SC_VERSION_FOR_APPLICATION = -999002
const SC_TIMEOUT = -7
const SC_OTHER_ERROR = -8
const SC_NOT_REACHABLE = -10
const SC_INVALID_RESPONSE = -20
const SC_CLIENT_ERROR = -400
const SC_INTERNAL_SERVER_ERROR = -500
const SC_HSP_PURCHASE_ERROR = -600
const SC_SERVER_BUSY = -700 // why is this different from the other busy?

// wheel consts
var WHEEL_ITEMS = []string{"200000", "120000", "120001", "120002", "200000", "900000", "120003", "120004"}
var WHEEL_ITEM = []int64{1, 2, 2, 2, 1, 3, 2, 2}
var WHEEL_ITEMWEIGHT = []int64{1250, 1250, 1250, 1250, 1250, 1250, 1250, 1250}

const WHEEL_ITEMWON = 5
const WHEEL_NEXTFREESPIN = 9999999999 // should be end of UTC day (the game automatically balances the time for its own region). ideally, this constant is never used
const WHEEL_SPINCOST = 100
const WHEEL_ROULETTERANK = 0
const WHEEL_NUMROULETTETOKEN = 0
const WHEEL_NUMJACKPOTRING = 50000
const WHEEL_NUMREMAININGROULETTE = 0

// challenge constants
const CHAL_DAILY_INCENTIVELISTCONT = 7
const CHAL_DAILY_NUMDAILYCHALCONT = 0
const CHAL_DAILY_NUMDAILYCHALDAY = 2
const CHAL_DAILY_MAXDAILYCHALDAY = 10

// misc consts
var CHAR_IDS = []string{"300000", "300001", "300002", "300003", "300004", "300005", "300006", "300007", "300008", "300009", "300010", "300011", "300012", "300013", "300014", "300015", "300016", "300017", "300018", "300019", "300020"}
var CHAO_IDS = []string{"400000", "400001", "400002", "400003", "400004", "400005", "400006", "400007", "400008", "400009", "400010", "400011", "400012", "400013", "400014", "400015", "400016", "400017", "400018", "400019", "400020", "400021", "400022", "400023", "400024", "400025", "401000", "401001", "401002", "401003", "401004", "401005", "401006", "401007", "401008", "401009", "401010", "401011", "401012", "401013", "401014", "401015", "401016", "401017", "401018", "401019", "401020", "401021", "401022", "401023", "401024", "401025", "401026", "401027", "401028", "401029", "401030", "401031", "401032", "401033", "401034", "401035", "401036", "401037", "401038", "401039", "401040", "401041", "401042", "401043", "401044", "401045", "401046", "401047", "402000", "402001", "402002", "402003", "402004", "402005", "402006", "402007", "402008", "402009", "402010", "402011", "402012", "402013", "402014", "402015", "402016", "402017", "402018", "402019", "402020", "402021", "402022", "402023", "402024", "402025", "402026", "402027", "402028", "402029", "402030", "402031", "402032", "402033", "402034"}

// Note about the Chao: The third digit indicates their rarity value.

const CHAR_NUM_ABILITIES = 11

// new user consts
const USR_DEFAULT_CHARACTERSTATE_CHARACTERID = "300000"
const USR_DEFAULT_CHARACTERSTATE_LEVEL = 1
const USR_DEFAULT_CHARACTERSTATE_NUMRINGS = 0
const USR_DEFAULT_CHARACTERSTATE_NUMREDRINGS = 0

var USR_DEFAULT_CHARACTERSTATE_ABILITYLEVEL = []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}
var USR_DEFAULT_CHARACTERSTATE_ABILITYNUMRINGS = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const USR_DEFAULT_CHARACTERSTATE_EXP = 0
const USR_DEFAULT_CHARACTERSTATE_STAR = 0
const USR_DEFAULT_CHARACTERSTATE_STARMAX = 10
const USR_DEFAULT_CHARACTERSTATE_LOCKCONDITION = 0
const USR_DEFAULT_CHARACTERSTATE_PRICENUMRINGS = 0
const USR_DEFAULT_CHARACTERSTATE_PRICENUMREDRINGS = 0
const USR_DEFAULT_CHARACTERSTATE_STATUS = 1

var USR_DEFAULT_CHARACTERSTATE_CAMPAIGNLIST = []string{}

const USR_DEFAULT_PLAYERSTATE_NUMRINGS = 0
const USR_DEFAULT_PLAYERSTATE_NUMBUYRINGS = 0
const USR_DEFAULT_PLAYERSTATE_NUMREDRINGS = 0
const USR_DEFAULT_PLAYERSTATE_NUMBUYREDRINGS = 0
const USR_DEFAULT_PLAYERSTATE_ENERGY = 5
const USR_DEFAULT_PLAYERSTATE_ENERGYBUY = 0
const USR_DEFAULT_PLAYERSTATE_ENERGYRENEWSAT = 0

//ITEMS
const USR_DEFAULT_PLAYERSTATE_MUMMESSAGES = 0
const USR_DEFAULT_PLAYERSTATE_RANKINGLEAGUE = 0
const USR_DEFAULT_PLAYERSTATE_QUICKRANKINGLEAGUE = 0
const USR_DEFAULT_PLAYERSTATE_NUMROULETTETICKET = 0
const USR_DEFAULT_PLAYERSTATE_TOTALHIGHSCORE = 0
const USR_DEFAULT_PLAYERSTATE_TOTALDISTANCE = 0
const USR_DEFAULT_PLAYERSTATE_MAXIMUMDISTANCE = 0
const USR_DEFAULT_PLAYERSTATE_DAILYMISSIONID = 68
const USR_DEFAULT_PLAYERSTATE_DAILYMISSIONENDTIME = 9999999999 // REPLACE WITH END OF DAY
const USR_DEFAULT_PLAYERSTATE_DAILYCHALLENGEVALUE = 0
const USR_DEFAULT_PLAYERSTATE_DAILYCHALLENGECOMPLETE = 0
const USR_DEFAULT_PLAYERSTATE_NUMDAILYCHALCONT = 0
const USR_DEFAULT_PLAYERSTATE_MAINCHARAID = "300000"
const USR_DEFAULT_PLAYERSTATE_SUBCHARAID = "300001"
const USR_DEFAULT_PLAYERSTATE_MAINCHAOID = "400000"
const USR_DEFAULT_PLAYERSTATE_SUBCHAOID = "400004"
const USR_DEFAULT_PLAYERSTATE_NUMPLAYING = 0
const USR_DEFAULT_PLAYERSTATE_NUMANIMALS = 0
const USR_DEFAULT_PLAYERSTATE_NUMRANK = 0

var USR_DEFAULT_PLAYERSTATE_EQUIPITEMLIST = []string{"-1", "-1", "-1"}

const USR_DEFAULT_PLAYERSTATE_QUICKTOTALHIGHSCORE = 0

const USR_DEFAULT_CHAO_STATUS = 0
const USR_DEFAULT_CHAO_LEVEL = 0
const USR_DEFAULT_CHAO_ACQUIRED = 0
const USR_DEFAULT_CHAO_RARITY = 0
const USR_DEFAULT_CHAO_HIDDEN = 0

// mileage consts
const MILE_EPISODE = 1
const MILE_CHAPTER = 1
const MILE_POINT = 0
const MILE_MAPDISTANCE = 0
const MILE_NUMBOSSATTACK = 0
const MILE_STAGEDISTANCE = 0
const MILE_STAGETOTALSCORE = 0
const MILE_STAGEMAXSCORE = 0
const MILE_CHAPTERSTARTTIME = 1469718000 // Is there significance to this magic number?

// chao wheel consts
var CHAOWHEEL_RARITY = []int64{2, 1, 100, 1, 2, 1, 100, 1}
var CHAOWHEEL_ITEMWEIGHT = []int64{6, 17, 5, 17, 16, 17, 5, 17}

const CHAOWHEEL_SPINCOST = 0
const CHAOWHEEL_CHAOROULETTETYPE = 0
const CHAOWHEEL_NUMSPECIALEGG = 0
const CHAOWHEEL_ROULETTEAVAILABLE = 1

// const CHAOWHEEL_CAMPAIGNLIST  // TODO: resolve cyclical dependencies to allow
const CHAOWHEEL_NUMCHAOROULETTE = 0
const CHAOWHEEL_NUMCHAOROULETTETOKEN = 0
const CHAOWHEEL_STARTTIME = 1564822800 // this const should never be used. it should be current day at 9 AM UTC
const CHAOWHEEL_ENDTIME = 1564909199   // this const should never be used. it should be next day at 8:59 AM UTC

// json consts
const JSON_DEFAULT_PRIZECHAOWHEEL_PRIZELIST = `[{"chao_id":"300000","rarity":"100"},{"chao_id":"300001","rarity":"100"},{"chao_id":"300002","rarity":"100"},{"chao_id":"401000","rarity":"1"},{"chao_id":"401001","rarity":"1"},{"chao_id":"401002","rarity":"1"},{"chao_id":"401010","rarity":"1"},{"chao_id":"401015","rarity":"1"},{"chao_id":"401005","rarity":"1"},{"chao_id":"401003","rarity":"1"},{"chao_id":"401004","rarity":"1"},{"chao_id":"401022","rarity":"1"},{"chao_id":"401016","rarity":"1"},{"chao_id":"401029","rarity":"1"},{"chao_id":"401020","rarity":"1"},{"chao_id":"401031","rarity":"1"},{"chao_id":"401033","rarity":"1"},{"chao_id":"401026","rarity":"1"},{"chao_id":"401027","rarity":"1"},{"chao_id":"401009","rarity":"1"},{"chao_id":"401017","rarity":"1"},{"chao_id":"401030","rarity":"1"},{"chao_id":"401038","rarity":"1"},{"chao_id":"401037","rarity":"1"},{"chao_id":"401034","rarity":"1"},{"chao_id":"401035","rarity":"1"},{"chao_id":"401036","rarity":"1"},{"chao_id":"401028","rarity":"1"},{"chao_id":"402001","rarity":"2"},{"chao_id":"402003","rarity":"2"},{"chao_id":"402005","rarity":"2"},{"chao_id":"402008","rarity":"2"},{"chao_id":"402015","rarity":"2"},{"chao_id":"402007","rarity":"2"},{"chao_id":"402021","rarity":"2"},{"chao_id":"402014","rarity":"2"},{"chao_id":"402024","rarity":"2"},{"chao_id":"402000","rarity":"2"},{"chao_id":"402002","rarity":"2"},{"chao_id":"402020","rarity":"2"},{"chao_id":"402016","rarity":"2"},{"chao_id":"402017","rarity":"2"},{"chao_id":"402018","rarity":"2"},{"chao_id":"402006","rarity":"2"},{"chao_id":"300005","rarity":"100"},{"chao_id":"300014","rarity":"100"},{"chao_id":"300015","rarity":"100"},{"chao_id":"300018","rarity":"100"},{"chao_id":"300020","rarity":"100"},{"chao_id":"300016","rarity":"100"},{"chao_id":"300017","rarity":"100"},{"chao_id":"300019","rarity":"100"},{"chao_id":"401011","rarity":"1"},{"chao_id":"401023","rarity":"1"},{"chao_id":"401024","rarity":"1"},{"chao_id":"401025","rarity":"1"},{"chao_id":"401046","rarity":"1"},{"chao_id":"401047","rarity":"1"},{"chao_id":"401040","rarity":"1"},{"chao_id":"401041","rarity":"1"},{"chao_id":"402034","rarity":"2"},{"chao_id":"402030","rarity":"2"}]`
